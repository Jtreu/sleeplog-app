package repository

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/jtreu/sleeplog-app/api-server/domain"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UserCollection = "users"

type DbUserRepo struct {
	DbRepo
	mailHandler MailHandler
}

var emailConfirmationTemplate string
var passwordResetTemplate string

func NewDbUserRepo(dbHandlers map[string]DbHandler, mailHandler MailHandler) *DbUserRepo {
	dbUserRepo := &DbUserRepo{}
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
	dbUserRepo.mailHandler = mailHandler

	// ensure that collection has the right text index
	// refactor building collection index
	err := dbUserRepo.dbHandler.EnsureIndex(UserCollection, mgo.Index{
		Key: []string{
			"$text:uid",
			"$text:username",
			"$text:alias",
		},
		Background: true,
		Sparse:     true,
	})
	if err != nil {
		fmt.Printf("Ensure User Index error: %v\n", err.Error())
	}

	// Retrieve Email Confirmation && Password Reset templates
	emailConfirmationTemplateBytes, err := ioutil.ReadFile("interfaces/repository/email_templates/email_confirmation.html")
	if err != nil {
		log.Fatal(fmt.Errorf("Error loading email confirmation template: %v", err.Error()))
	}
	emailConfirmationTemplate = string(emailConfirmationTemplateBytes)

	passwordResetTemplateBytes, err := ioutil.ReadFile("interfaces/repository/email_templates/password_reset.html")
	if err != nil {
		log.Fatal(fmt.Errorf("Error loading password reset template: %v", err.Error()))
	}
	passwordResetTemplate = string(passwordResetTemplateBytes)

	return dbUserRepo
}

func (repo *DbUserRepo) Create(user domain.User) (domain.User, error) {
	user.UID = bson.NewObjectId().Hex()
	user.DateCreated = time.Now()
	user.Status = domain.UserStatusPending

	err := repo.dbHandler.Insert(UserCollection, user)
	if err != nil {
		return domain.User{}, err
	}

	err = repo.SendEmailConfirmation(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *DbUserRepo) Get(userUID string) (domain.User, error) {
	var user domain.User
	err := repo.dbHandler.FindOne(UserCollection, Query{"uid": userUID}, &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repo *DbUserRepo) GetAll() domain.Users {
	users := domain.Users{}
	err := repo.dbHandler.FindAll(UserCollection, nil, &users, 50, "")
	if err != nil {
		return domain.Users{}
	}
	return users
}

func (repo *DbUserRepo) UpdateAlias(user domain.User) (domain.User, error) {
	update := Query{
		"alias": user.Alias,
	}

	change := Change{
		Update:    Query{"$set": update},
		ReturnNew: true,
	}

	var updatedUser domain.User
	err := repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
	return updatedUser, err
}

func (repo *DbUserRepo) UpdateUsername(user domain.User) (domain.User, error) {
	update := Query{
		"username": user.Username,
	}

	change := Change{
		Update:    Query{"$set": update},
		ReturnNew: true,
	}

	var updatedUser domain.User
	err := repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
	return updatedUser, err
}

func (repo *DbUserRepo) UpdateEmail(user domain.User) (domain.User, error) {
	update := Query{
		"email":                 user.Email,
		"status":                domain.UserStatusPending, // Pending until email is confirmed
		"emailConfirmationCode": user.EmailConfirmationCode,
	}

	change := Change{
		Update:    Query{"$set": update},
		ReturnNew: true,
	}

	var updatedUser domain.User
	err := repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, repo.SendEmailConfirmation(user)
}

func (repo *DbUserRepo) UpdateDescription(user domain.User) (domain.User, error) {
	update := Query{
		"description": user.Description,
	}

	change := Change{
		Update:    Query{"$set": update},
		ReturnNew: true,
	}

	var updatedUser domain.User
	err := repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (repo *DbUserRepo) UpdatePassword(user domain.User) error {
	update := Query{
		"hashedPassword":    user.HashedPassword,
		"passwordResetCode": "", // Reset Code
	}

	change := Change{
		Update: Query{"$set": update},
	}

	var updatedUser domain.User
	return repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
}

func (repo *DbUserRepo) ConfirmEmail(user domain.User) error {
	update := Query{
		"emailConfirmationCode": "",                      // Reset Code
		"status":                domain.UserStatusActive, // Activate User Account
	}

	change := Change{
		Update:    Query{"$set": update},
		ReturnNew: true,
	}

	var updatedUser domain.User
	err := repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
	return err
}

func (repo *DbUserRepo) ResetPassword(user domain.User) error {
	update := Query{
		"passwordResetCode": user.PasswordResetCode,
	}

	change := Change{
		Update: Query{"$set": update},
	}

	var updatedUser domain.User
	err := repo.dbHandler.Update(UserCollection, Query{"uid": user.UID}, change, &updatedUser)
	if err != nil {
		return err
	}

	err = repo.SendPasswordReset(user)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DbUserRepo) ExistsByEmail(email string) bool {
	query := Query{
		"email": Query{
			"$regex":   fmt.Sprintf("^%v$", email),
			"$options": "mi",
		},
	}
	return repo.dbHandler.Exists(UserCollection, query)
}

func (repo *DbUserRepo) ExistsByUsername(username string) bool {
	query := Query{
		"username": Query{
			"$regex":   fmt.Sprintf("^%v$", username),
			"$options": "mi",
		},
	}
	return repo.dbHandler.Exists(UserCollection, query)
}

func (repo *DbUserRepo) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	query := Query{
		"email": Query{
			"$regex":   fmt.Sprintf("^%v$", email),
			"$options": "mi",
		},
	}
	err := repo.dbHandler.FindOne(UserCollection, query, &user)
	return user, err
}

func (repo *DbUserRepo) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	query := Query{
		"username": Query{
			"$regex":   fmt.Sprintf("^%v$", username),
			"$options": "mi",
		},
	}
	err := repo.dbHandler.FindOne(UserCollection, query, &user)
	return user, err
}

func (repo *DbUserRepo) FindByEmailConfirmationCode(code string) (domain.User, error) {
	var user domain.User
	err := repo.dbHandler.FindOne(UserCollection, Query{"emailConfirmationCode": code}, &user)
	return user, err
}

func (repo *DbUserRepo) FindByPasswordResetCode(code string) (domain.User, error) {
	var user domain.User
	err := repo.dbHandler.FindOne(UserCollection, Query{"passwordResetCode": code}, &user)
	return user, err
}

func (repo *DbUserRepo) Filter(field, query, lastID string, limit int, sort string) domain.Users {
	users := domain.Users{}

	// parse sort string
	allowedSortMap := map[string]bool{
		"_id":  true,
		"-_id": true,
	}
	// ensure that sort string is allowed
	// we are basically concerned about sorting on un-indexed keys
	if !allowedSortMap[sort] {
		sort = "-_id" // set it to default sort
	}

	q := Query{}
	if lastID != "" && bson.IsObjectIdHex(lastID) {
		if sort == "_id" {
			q["_id"] = Query{
				"$gt": bson.ObjectIdHex(lastID),
			}
		} else {
			q["_id"] = Query{
				"$lt": bson.ObjectIdHex(lastID),
			}
		}
	}

	if query != "" {
		if field != "" {
			q[field] = Query{
				"$regex":   fmt.Sprintf("^%v.*", query),
				"$options": "i",
			}
		} else {
			// if not field is specified, we do a text search on pre-defined text index
			q["$text"] = Query{
				"$search": query,
			}
		}
	}

	err := repo.dbHandler.FindAll(UserCollection, q, &users, limit, sort)
	if err != nil {
		fmt.Printf("FilterUsers Error: %v\n", err)
		return domain.Users{}
	}

	return users
}

const ConfirmEmailSender = "Neurologist Sleep Log App <noreply@sleeplog-app.com>"
const ConfirmEmailSubject = "Activate your Sleep log account"

func (repo *DbUserRepo) SendEmailConfirmation(user domain.User) error {
	emailTemplate := emailConfirmationTemplate
	emailTemplate = strings.Replace(emailTemplate, "{{ name }}", user.Alias, -1)
	emailTemplate = strings.Replace(emailTemplate, "{{ code }}", user.EmailConfirmationCode, -1)

	return repo.mailHandler.Send(ConfirmEmailSender, user.Email, ConfirmEmailSubject, emailTemplate)
}

const ResetPasswordSender = "Neurologist Sleep Log App <noreply@sleeplog-app.com>"
const ResetPasswordSubject = "Password reset request"

func (repo *DbUserRepo) SendPasswordReset(user domain.User) error {
	passwordTemplate := passwordResetTemplate
	passwordTemplate = strings.Replace(passwordTemplate, "{{ name }}", user.Alias, -1)
	passwordTemplate = strings.Replace(passwordTemplate, "{{ code }}", user.PasswordResetCode, -1)

	return repo.mailHandler.Send(ResetPasswordSender, user.Email, ResetPasswordSubject, passwordTemplate)
}
