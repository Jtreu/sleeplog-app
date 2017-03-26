package usecase

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	"github.com/jtreu/sleeplog-app/server/domain"
)

type SessionTokenInteractor struct {
	SessionTokenRepository domain.SessionTokenRepository
	UserRepository         domain.UserRepository
}

func NewSessionTokenInteractor() *SessionTokenInteractor {
	return &SessionTokenInteractor{}
}

func (interactor *SessionTokenInteractor) VerifySessionToken(token string) (string, error) {
	tokenClaim, err := interactor.SessionTokenRepository.Verify(token)
	if err != nil {
		return "", err
	}

	return tokenClaim.UserUID, nil
}

func (interactor *SessionTokenInteractor) CreateSessionToken(usernameOrEmail, password string) (string, domain.User, error) {
	var user domain.User

	if isValidEmail(usernameOrEmail) {
		user, _ = interactor.UserRepository.FindByEmail(usernameOrEmail)
	}

	if uMinLength, uMaxLength, uAlphaNum, uNonVanity := isValidUsername(usernameOrEmail); (uMinLength && uMaxLength && uAlphaNum && uNonVanity) && user.UID == "" {
		user, _ = interactor.UserRepository.FindByUsername(usernameOrEmail)
	}

	// Ensure credentials are valid.
	if !hasValidPassword(user.HashedPassword, password) {
		return "", domain.User{}, fmt.Errorf("Invalid Credentials.")
	}

	sessionTokenClaim := domain.NewSessionTokenClaim(user.UID)
	sessionToken, err := interactor.SessionTokenRepository.Create(sessionTokenClaim)
	if err != nil {
		return "", domain.User{}, fmt.Errorf("Token could not be generated: %v", err)
	}

	return sessionToken, user, nil
}

func (interactor *SessionTokenInteractor) UpdateSessionToken(userUID string) (string, error) {
	sessionTokenClaim := domain.NewSessionTokenClaim(userUID)
	return interactor.SessionTokenRepository.Create(sessionTokenClaim)
}

func (interactor *SessionTokenInteractor) DeleteSessionToken(token string) error {
	tokenClaim, err := interactor.SessionTokenRepository.Verify(token)
	if err != nil {
		return err
	}

	revoked := domain.RevokedSessionToken{
		ID:          bson.ObjectIdHex(tokenClaim.JTI),
		DateExpired: tokenClaim.DateExpired,
	}

	return interactor.SessionTokenRepository.Delete(revoked)
}

func hasValidPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
