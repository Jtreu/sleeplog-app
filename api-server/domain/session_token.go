package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	jwt "github.com/dgrijalva/jwt-go"
)

type SessionTokenRepository interface {
	Verify(tokenString string) (SessionTokenClaim, error)
	Create(sessionToken SessionTokenClaim) (string, error)
	Delete(revoked RevokedSessionToken) error
}

type SessionTokenClaim struct {
	UserUID     string    `json:"userUID"`
	DateExpired time.Time `json:"dateExpired"`
	DateIssued  time.Time `json:"dateIssued"`
	JTI         string    `json:"jti"`
}

func NewSessionTokenClaim(userUID string) SessionTokenClaim {
	return SessionTokenClaim{UserUID: userUID}
}

type RevokedSessionToken struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	DateExpired time.Time     `json:"dateExpired" bson:"dateExpired"`
	DateRevoked time.Time     `json:"dateRevoked" bson:"dateRevoked"`
}

type SessionToken struct {
	*jwt.Token
}
