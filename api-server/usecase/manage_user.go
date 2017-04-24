package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jtreu/sleeplog-app/api-server/domain"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

var vanityUsernames map[string]bool

func NewUserInteractor() *UserInteractor {
	// Retrieve vanity names
	vanityUsernamesBytes, err := ioutil.ReadFile("usecase/vanity_usernames.json")
	if err != nil {
		log.Fatal(fmt.Errorf("Error loading vanity_usernames: %v", err.Error()))
	}

	err = json.Unmarshal(vanityUsernamesBytes, &vanityUsernames)
	if err != nil {
		log.Fatal(fmt.Errorf("Error parsing vanity usernames: %v", err.Error()))
	}

	return &UserInteractor{}
}

func (interactor *UserInteractor) Get(uid string) (domain.User, error) {
	return interactor.UserRepository.Get(uid)
}

// GetByUsername gets a User by username
func (interactor *UserInteractor) GetByUsername(username string) (domain.User, error) {
	return interactor.UserRepository.FindByUsername(username)
}

func (interactor *UserInteractor) ListUsers(field string, query string, lastID string, limit int, sort string) domain.Users {
	return interactor.UserRepository.Filter(field, query, lastID, limit, sort)
}

// Create creates a new User
func (interactor *UserInteractor) Create(email, username, fullname, password string) (domain.User, error) {
	if !isValidEmail(email) {
		return domain.User{}, fmt.Errorf("Invalid Email")
	}

	if uMinLength, uMaxLength, uAlphaNum, uNonVanity := isValidUsername(username); !(uMinLength && uMaxLength && uAlphaNum && uNonVanity) {
		if !uMinLength {
			return domain.User{}, fmt.Errorf("Username must have at least 1 character")
		} else if !uMaxLength {
			return domain.User{}, fmt.Errorf("Username must have less than 16 characters")
		} else if !uAlphaNum {
			return domain.User{}, fmt.Errorf("Username must contain alphanumerical characters or _")
		} else if !uNonVanity {
			return domain.User{}, fmt.Errorf("Invalid Username")
		}

		return domain.User{}, fmt.Errorf("Invalid Username")
	}

	if aMinLength, aMaxLength := isValidAlias(fullname); !(aMinLength && aMaxLength) {
		if !aMinLength {
			return domain.User{}, fmt.Errorf("Full name must have at least 1 character")
		} else if !aMaxLength {
			return domain.User{}, fmt.Errorf("Full name must have less than 20 characters")
		}

		return domain.User{}, fmt.Errorf("Invalid Full Name")
	}

	if pLength, pLetter, pNumber, pUppercase, pSpecial := isValidPassword(password); !(pLength && pLetter && pNumber && pUppercase && pSpecial) {
		if !pLength {
			return domain.User{}, fmt.Errorf("Password must contain at least 7 characters")
		} else if !pLetter {
			return domain.User{}, fmt.Errorf("Password must contain at least 1 letter")
		} else if !pUppercase {
			return domain.User{}, fmt.Errorf("Password must contain at least 1 uppercase character")
		} else if !pSpecial {
			return domain.User{}, fmt.Errorf("Password must contain at least 1 special character")
		}

		return domain.User{}, fmt.Errorf("Invalid Password")
	}

	if !interactor.isAvailableEmail(email) {
		return domain.User{}, fmt.Errorf("Email is taken")
	}

	if !interactor.isAvailableUsername(username) {
		return domain.User{}, fmt.Errorf("Username is taken")
	}

	media := domain.UserMedia{}
	media.Icon.UID = fmt.Sprintf("%v-%v", "default-icon", string([]rune(fullname)[0]))
	media.Icon.Source = generateDefaultUserIcon(string([]rune(fullname)[0]))
	media.Banner.UID = fmt.Sprintf("%v-%v", "default-banner", fullname)
	media.Banner.Source = generateDefaultUserBanner(fullname)

	hashedPassword, _ := hashPassword(password)
	emailConfirmationCord := generateEmailConfirmationCode()

	user := domain.User{
		Email:                 strings.TrimSpace(email),
		Username:              strings.TrimSpace(username),
		Alias:                 strings.TrimSpace(fullname),
		Media:                 media,
		HashedPassword:        hashedPassword,
		EmailConfirmationCode: emailConfirmationCord,
	}

	return interactor.UserRepository.Create(user)
}

func (interactor *UserInteractor) UpdateAlias(alias string, user domain.User) (domain.User, error) {
	if alias == user.Alias {
		return domain.User{}, nil // No need to update with the same alias
	}

	if aMinLength, aMaxLength := isValidAlias(alias); !(aMinLength && aMaxLength) {
		if !aMinLength {
			return domain.User{}, fmt.Errorf("Full name must have at least 1 character")
		} else if !aMaxLength {
			return domain.User{}, fmt.Errorf("Full name must have less than 20 characters")
		}

		return domain.User{}, fmt.Errorf("Invalid Full Name")
	}

	user.Alias = alias

	return interactor.UserRepository.UpdateAlias(user)
}

func (interactor *UserInteractor) UpdateUsername(username string, user domain.User) (domain.User, error) {
	if uMinLength, uMaxLength, uAlphaNum, uNonVanity := isValidUsername(username); !(uMinLength && uMaxLength && uAlphaNum && uNonVanity) {
		if !uMinLength {
			return domain.User{}, fmt.Errorf("Username must have at least 1 character")
		} else if !uMaxLength {
			return domain.User{}, fmt.Errorf("Username must have less than 16 characters")
		} else if !uAlphaNum {
			return domain.User{}, fmt.Errorf("Username must contain alphanumerical characters or _")
		} else if !uNonVanity {
			return domain.User{}, fmt.Errorf("Invalid Username")
		}

		return domain.User{}, fmt.Errorf("Invalid Username")
	}

	if !interactor.isAvailableUsername(username) {
		return domain.User{}, fmt.Errorf("Username is taken")
	}

	user.Username = username

	return interactor.UserRepository.UpdateUsername(user)
}

func (interactor *UserInteractor) UpdateEmail(email string, user domain.User) (domain.User, error) {
	if email == user.Email {
		return domain.User{}, nil // No need to update with the same email
	}

	if !isValidEmail(email) {
		return domain.User{}, fmt.Errorf("Invalid Email")
	}

	if !interactor.isAvailableEmail(email) {
		return domain.User{}, fmt.Errorf("Email is taken")
	}

	user.Email = email
	user.EmailConfirmationCode = generateEmailConfirmationCode()

	return interactor.UserRepository.UpdateEmail(user)
}

func (interactor *UserInteractor) UpdateDescription(description string, user domain.User) (domain.User, error) {
	if description == user.Description {
		return domain.User{}, nil // No need to update with the same description
	}

	if !isValidUserDescription(description) {
		return domain.User{}, fmt.Errorf("Invalid Description")
	}

	user.Description = description

	return interactor.UserRepository.UpdateDescription(user)
}

func (interactor *UserInteractor) UpdatePassword(oldPassword string, newPassword string, user domain.User) error {
	if !hasValidPassword(user.HashedPassword, oldPassword) {
		return fmt.Errorf("Old password is invalid")
	}

	if pLength, pLetter, pNumber, pUppercase, pSpecial := isValidPassword(newPassword); !(pLength && pLetter && pNumber && pUppercase && pSpecial) {
		if !pLength {
			return fmt.Errorf("New Password must contain at least 7 characters")
		} else if !pLetter {
			return fmt.Errorf("New Password must contain at least 1 letter")
		} else if !pUppercase {
			return fmt.Errorf("New Password must contain at least 1 uppercase character")
		} else if !pSpecial {
			return fmt.Errorf("New Password must contain at least 1 special character")
		}

		return fmt.Errorf("New Password is invalid")
	}

	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword

	return interactor.UserRepository.UpdatePassword(user)
}

func (interactor *UserInteractor) UpdateEntries(entries domain.UserEntries, user domain.User) error {
	if user.Entries == nil {
		user.Entries = make(domain.UserEntries)
	}

	for date := range entries {
		user.Entries[date] = entries[date]
	}

	return interactor.UserRepository.UpdateEntries(user)
}

func (interactor *UserInteractor) ConfirmEmail(code string) error {
	if !isValidEmailConfirmationCode(code) {
		return fmt.Errorf("Invalid confirmation code")
	}
	user, err := interactor.UserRepository.FindByEmailConfirmationCode(code)
	if err != nil {
		return fmt.Errorf("Invalid confirmation code")
	}

	if user.EmailConfirmationCode != code {
		return fmt.Errorf("Invalid confirmation code")
	}

	return interactor.UserRepository.ConfirmEmail(user)
}

func (interactor *UserInteractor) ResetPassword(email string) error {
	if !isValidEmail(email) {
		return fmt.Errorf("Invalid email")
	}

	user, err := interactor.UserRepository.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("Invalid email: %v", err.Error())
	}

	user.PasswordResetCode = generatePasswordResetCode()

	return interactor.UserRepository.ResetPassword(user)
}

func (interactor *UserInteractor) UpdatePasswordWithCode(code, password string) error {
	if !isValidPasswordResetCode(code) {
		return fmt.Errorf("Invalid reset code")
	}
	user, err := interactor.UserRepository.FindByPasswordResetCode(code)
	if err != nil {
		return fmt.Errorf("Invalid reset code")
	}

	if user.PasswordResetCode != code {
		return fmt.Errorf("Invalid reset code")
	}

	if pLength, pLetter, pNumber, pUppercase, pSpecial := isValidPassword(password); !(pLength && pLetter && pNumber && pUppercase && pSpecial) {
		if !pLength {
			return fmt.Errorf("Password must contain at least 7 characters")
		} else if !pLetter {
			return fmt.Errorf("Password must contain at least 1 letter")
		} else if !pUppercase {
			return fmt.Errorf("Password must contain at least 1 uppercase character")
		} else if !pSpecial {
			return fmt.Errorf("Password must contain at least 1 special character")
		}

		return fmt.Errorf("Invalid Password")
	}

	hashedPassword, _ := hashPassword(password)
	user.HashedPassword = hashedPassword

	return interactor.UserRepository.UpdatePassword(user)
}

var defaultUserIconColors = []string{"FF6344", "3D4E5E", "22C19A"}

func generateDefaultUserIcon(letter string) string {
	color := defaultUserIconColors[rand.Intn(len(defaultUserIconColors))]
	return fmt.Sprintf("https://www.dummyimage.com/128x128/%v/F7F7F8.png&text=%v", color, letter)
}

func generateDefaultUserBanner(name string) string {
	return fmt.Sprintf("https://www.dummyimage.com/2120x352/000/ffffff&text=%v", name)
}

func generateEmailConfirmationCode() string {
	return uuid.NewV4().String()
}

func generatePasswordResetCode() string {
	return uuid.NewV4().String()
}

func hashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (interactor *UserInteractor) isAvailableEmail(email string) bool {
	return !interactor.UserRepository.ExistsByEmail(email)
}

func (interactor *UserInteractor) isAvailableUsername(username string) bool {
	return !interactor.UserRepository.ExistsByUsername(username)
}

func isValidEmailConfirmationCode(code string) bool {
	return utf8.RuneCountInString(code) > 5
}

func isValidPasswordResetCode(code string) bool {
	return utf8.RuneCountInString(code) > 5
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

func isValidUsername(username string) (minLength, maxLength, alphaNum, nonVanity bool) {
	username = strings.TrimSpace(username)
	length := utf8.RuneCountInString(username)
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return length > 0, length < 16, re.MatchString(username), !vanityUsernames[username]
}

func isValidAlias(alias string) (minLength, maxLength bool) {
	alias = strings.TrimSpace(alias)
	length := utf8.RuneCountInString(alias)
	return length > 0, length < 21
}

func isValidUserDescription(description string) bool {
	return true
}

func isValidPassword(password string) (length, letter, number, uppercase, special bool) {
	length = utf8.RuneCountInString(password) >= 5
	special = true
	uppercase = true
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
			/*
				case unicode.IsUpper(c):
					uppercase = true
				case unicode.IsPunct(c) || unicode.IsSymbol(c):
					special = true
			*/
		case unicode.IsLetter(c) || c == ' ':
			letter = true
		default:
			return
		}
	}
	return
}
