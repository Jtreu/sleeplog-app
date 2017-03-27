package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-zoo/bone"
	"github.com/jtreu/sleeplog-app/api-server/infrastructure"
	"github.com/jtreu/sleeplog-app/api-server/interfaces/repository"
	"github.com/jtreu/sleeplog-app/api-server/interfaces/web"
	"github.com/jtreu/sleeplog-app/api-server/usecase"
	"github.com/twinj/uuid"
	"github.com/urfave/negroni"
)

func main() {
	// Determine Environment
	var environment string

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		environment = "dev"
	} else {
		environment = "prod"
	}

	// Initialize UUID generator
	uuid.Init()

	// Initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// Retrieve Signing Keys for token authority
	privateSigningKeyBytes, err := ioutil.ReadFile("config/keys/domain.key")
	if err != nil {
		panic(fmt.Errorf("Error loading private signing key: %v", err.Error()))
	}
	publicSigningKeyBytes, err := ioutil.ReadFile("config/keys/domain.key.pub")
	if err != nil {
		panic(fmt.Errorf("Error loading public signing key: %v", err.Error()))
	}
	privateSigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateSigningKeyBytes)
	if err != nil {
		log.Fatal("Error parsing public key")
		return
	}
	publicSigningKey, err := jwt.ParseRSAPublicKeyFromPEM(publicSigningKeyBytes)
	if err != nil {
		log.Fatal("Error parsing public key")
		return
	}

	// Retrieve Database config settings
	var dbFilePath string
	if environment == "dev" {
		dbFilePath = "config/db.dev.json"
	} else if environment == "prod" {
		dbFilePath = "config/db.prod.json"
	}

	dbFile, err := ioutil.ReadFile(dbFilePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading database info: %v", err.Error()))
	}
	var dbConfig struct {
		Database string `json:"database"`
		FullURI  string `json:"fullURI"`
	}

	err = json.Unmarshal(dbFile, &dbConfig)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error parsing database info: %v", err.Error()))
	}

	// Setup Database options
	dbOptions := &infrastructure.MongoOptions{
		ServerName:   fmt.Sprintf(dbConfig.FullURI),
		DatabaseName: dbConfig.Database,
	}

	// Setup Database
	dbHandler := infrastructure.NewMongoHandler(dbOptions)
	dbHandler.NewSession()

	handlers := make(map[string]repository.DbHandler)
	handlers["DbSessionTokenRepo"] = dbHandler
	handlers["DbUserRepo"] = dbHandler

	// Setup Mailgun - Email Serivce
	mailHandlerOptions := infrastructure.MailHandlerOptions{
		Domain:        "mg.sleeplog-app.com",
		PrivateApiKey: "key-d915e528e573a92b003eb765373c8593",
		PublicApiKey:  "pubkey-98cc18ae956af012719065834959c94a",
	}

	mailHandler := infrastructure.NewMailHandler(mailHandlerOptions)

	sessionTokenInterator := usecase.NewSessionTokenInteractor()
	userInteractor := usecase.NewUserInteractor()

	tokenAuthorityOptions := repository.NewTokenAuthorityOptions(privateSigningKey, publicSigningKey)
	sessionTokenInterator.SessionTokenRepository = repository.NewDbSessionTokenRepo(handlers, tokenAuthorityOptions)
	userInteractor.UserRepository = repository.NewDbUserRepo(handlers, mailHandler)

	sessionTokenInterator.UserRepository = userInteractor.UserRepository

	// Setup web handler
	w := web.NewWebHandler()
	w.SessionTokenInteractor = sessionTokenInterator
	w.UserInteractor = userInteractor

	// Setup route handling

	router := bone.New()
	n := negroni.New()
	n.UseFunc(w.ForceHTTPS)
	n.UseFunc(w.CORS)

	router.Get("/", http.HandlerFunc(w.GetIndex))
	router.Get("/sessions", w.EnsureAuth(w.GetSessionTokenUser))   // Get account information
	router.Post("/sessions", w.CheckAuth(w.CreateSessionToken))    // Login
	router.Delete("/sessions", w.EnsureAuth(w.DeleteSessionToken)) // Logout
	// router.Post("/sessions", w.EnsureAuth(w.UpdateSessionToken))

	router.Get("/users", http.HandlerFunc(w.ListUsers))
	router.Post("/users", http.HandlerFunc(w.CreateUser))                                 // Register
	router.Post("/users/email_confirmation", http.HandlerFunc(w.ConfirmUserEmail))        // Confirm Email request
	router.Post("/users/password_reset", http.HandlerFunc(w.ResetUserPassword))           // Reset Password request
	router.Post("/users/password_update", http.HandlerFunc(w.UpdateUserPasswordWithCode)) // Update Password w/ code

	router.Get("/users/:username", http.HandlerFunc(w.GetUser))
	router.Post("/users/:username/alias", w.EnsureAuth(w.UpdateUserAlias))
	router.Post("/users/:username/username", w.EnsureAuth(w.UpdateUsername))
	router.Post("/users/:username/email", w.EnsureAuth(w.UpdateUserEmail))
	router.Post("/users/:username/description", w.EnsureAuth(w.UpdateUserDescription))
	router.Post("/users/:username/password", w.EnsureAuth(w.UpdateUserPassword))

	router.Get("/.well-known/acme-challenge/osvGIOUEm8cHMDa30G3Hpbhg_oI7ncvypeVEtdTarGs", http.HandlerFunc(w.GetCertBotKey)) // Certbot - Certificate verification

	n.UseHandler(router)
	fmt.Printf("Running on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), n))
}
