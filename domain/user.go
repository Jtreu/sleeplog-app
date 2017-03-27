package domain

import "time"

type UserRepository interface {
	Create(user User) (User, error)
	// Update(uid string, user User) (User, error)
	// Delete(uid string) error
	Get(userUID string) (User, error)
	GetAll() Users

	UpdateAlias(user User) (User, error)
	UpdateUsername(user User) (User, error)
	UpdateEmail(user User) (User, error)
	UpdateDescription(user User) (User, error)
	UpdatePassword(user User) error

	ConfirmEmail(user User) error
	ResetPassword(user User) error
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	FindByEmailConfirmationCode(code string) (User, error)
	FindByPasswordResetCode(code string) (User, error)

	ExistsByEmail(email string) bool
	ExistsByUsername(username string) bool
	Filter(field, query, lastID string, limit int, sort string) Users
}

type User struct {
	UID      string `json:"uid" bson:"uid,omitempty"`
	Username string `json:"username" bson:"username"`
	Alias    string `json:"alias" bson:"alias"`

	Description string    `json:"description" bson:"description"`
	Media       UserMedia `json:"media" bson:"media"`

	DateCreated time.Time `json:"dateCreated,omitempty" bson:"dateCreated"`
	LastOnline  time.Time `json:"lastOnline,omitempty" bson:"lastOnline"`

	// fields not exposed to JSON
	Status                string `json:"-" bson:"status"`
	Email                 string `json:"-" bson:"email"`
	EmailConfirmationCode string `json:"-" bson:"emailConfirmationCode"`
	PasswordResetCode     string `json:"-" bson:"passwordResetCode"`
	HashedPassword        string `json:"-" bson:"hashedPassword"`
}

// Users is a list of User
type Users []User

type UserMedia struct {
	Icon struct {
		UID    string `json:"uid" bson:"uid"`
		Source string `json:"source" bson:"source"`
	} `json:"icon,omitempty"`
	Banner struct {
		UID    string `json:"uid" bson:"uid"`
		Source string `json:"source" bson:"source"`
	} `json:"banner,omitempty" bson:"banner"`
}

const (
	UserStatusActive    = "active"    // email confirmed
	UserStatusPending   = "pending"   // email confirmation pending
	UserStatusSuspended = "suspended" // suspended
	UserStatusDeleted   = "deleted"   // deleted
)
