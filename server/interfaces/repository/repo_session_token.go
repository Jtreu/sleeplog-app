package repository

import (
	"crypto/rsa"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jtreu/sleeplog-app/server/domain"
	"gopkg.in/mgo.v2/bson"
)

const SessionTokenCollection = "sessions"

var rsaSampleSecret = []byte("something-super-secret")

type DbSessionTokenRepo struct {
	DbRepo
	TAOptions *TokenAuthorityOptions
}

type TokenAuthorityOptions struct {
	PrivateSigningKey *rsa.PrivateKey
	PublicSigningKey  *rsa.PublicKey
}

func NewDbSessionTokenRepo(dbHandlers map[string]DbHandler, taOptions *TokenAuthorityOptions) *DbSessionTokenRepo {
	dbSessionTokenRepo := &DbSessionTokenRepo{}
	dbSessionTokenRepo.dbHandlers = dbHandlers
	dbSessionTokenRepo.dbHandler = dbHandlers["DbSessionTokenRepo"]
	dbSessionTokenRepo.TAOptions = taOptions
	return dbSessionTokenRepo
}

func NewTokenAuthorityOptions(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *TokenAuthorityOptions {
	taOptions := &TokenAuthorityOptions{
		PrivateSigningKey: privateKey,
		PublicSigningKey:  publicKey,
	}

	return taOptions
}

func (repo *DbSessionTokenRepo) Verify(tokenString string) (domain.SessionTokenClaim, error) {
	// Ensure token is valid
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return repo.TAOptions.PublicSigningKey, nil
	})
	if err != nil {
		return domain.SessionTokenClaim{}, fmt.Errorf("Error parsing session token: %v", err)
	}

	var claim domain.SessionTokenClaim

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim.UserUID = claims["userUID"].(string)
		claim.DateExpired = time.Unix(int64(claims["dateExpired"].(float64)), 0)
		claim.DateIssued = time.Unix(int64(claims["dateIssued"].(float64)), 0)
		claim.JTI = claims["jti"].(string)
	} else {
		return domain.SessionTokenClaim{}, fmt.Errorf("Invalid session token.")
	}

	// Ensure token hasn't been revoked
	isRevoked := repo.dbHandler.Exists(SessionTokenCollection, Query{"_id": bson.ObjectIdHex(claim.JTI)})
	if isRevoked {
		return domain.SessionTokenClaim{}, fmt.Errorf("Invalid session token.")
	}

	return claim, nil
}

func (repo *DbSessionTokenRepo) Create(claim domain.SessionTokenClaim) (string, error) {
	claims := make(jwt.MapClaims)
	claims["userUID"] = claim.UserUID
	claims["dateExpired"] = time.Now().Add(time.Hour * 72).Unix() // 3 days
	claims["dateIssued"] = time.Now().Unix()
	claims["jti"] = generateJTI()

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	return token.SignedString(repo.TAOptions.PrivateSigningKey)
}

func (repo *DbSessionTokenRepo) Delete(revoked domain.RevokedSessionToken) error {
	revoked.DateRevoked = time.Now()
	err := repo.dbHandler.Insert(SessionTokenCollection, revoked)
	if err != nil {
		err = fmt.Errorf("Error deleting session token.")
	}
	return err
}

func generateJTI() string {
	// We will use mongodb's object id as JTI
	// we then will use this id to blacklist tokens,
	// along with `exp` and `iat` claims.
	// As far as collisions go, ObjectId is guaranteed unique
	// within a collection; and this case our collection is `sessions`
	return bson.NewObjectId().Hex()
}
