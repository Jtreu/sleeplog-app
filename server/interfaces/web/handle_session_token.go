package web

import (
	"fmt"
	"net/http"

	"github.com/jtreu/sleeplog-app/server/domain"
)

type UserAccount struct {
	UID         string      `json:"uid"`
	Email       string      `json:"email"`
	Username    string      `json:"username"`
	Alias       string      `json:"alias"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Media       interface{} `json:"media"`
}

type GetSessionTokenUserResponseV0 struct {
	User    UserAccount `json:"user"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type CreateSessionTokenRequestV0 struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type CreateSessionTokenResponseV0 struct {
	SessionToken string      `json:"sessionToken"`
	User         UserAccount `json:"user"`
	Message      string      `json:"message,omitempty"`
	Success      bool        `json:"success,omitempty"`
}

type DeleteSessionTokenResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type SessionTokenInteractor interface {
	VerifySessionToken(token string) (string, error)
	CreateSessionToken(usernameOrEmail, password string) (string, domain.User, error)
	UpdateSessionToken(userUID string) (string, error)
	DeleteSessionToken(token string) error
}

func (handler *WebHandler) GetSessionTokenUser(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	user, err := handler.UserInteractor.Get(userUID.(string))
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	data := GetSessionTokenUserResponseV0{
		Message: "User profile retrieved",
		Success: true,
	}

	data.User.UID = user.UID
	data.User.Email = user.Email
	data.User.Username = user.Username
	data.User.Alias = user.Alias
	data.User.Description = user.Description
	data.User.Status = user.Status
	data.User.Media = user.Media

	handler.util.renderer.Render(res, req, http.StatusOK, data)
}

func (handler *WebHandler) CreateSessionToken(res http.ResponseWriter, req *http.Request) {
	oldToken := req.Context().Value(SessionTokenContextKey)

	var body CreateSessionTokenRequestV0
	err := handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	// Revoke old token
	handler.SessionTokenInteractor.DeleteSessionToken(oldToken.(string))

	// Create new token
	token, user, err := handler.SessionTokenInteractor.CreateSessionToken(body.UsernameOrEmail, body.Password)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	data := CreateSessionTokenResponseV0{
		SessionToken: token,
		Message:      "Token Created",
		Success:      true,
	}

	data.User.UID = user.UID
	data.User.Email = user.Email
	data.User.Username = user.Username
	data.User.Alias = user.Alias
	data.User.Description = user.Description
	data.User.Status = user.Status
	data.User.Media = user.Media

	handler.util.renderer.Render(res, req, http.StatusCreated, data)
}

func (handler *WebHandler) UpdateSessionToken(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	sessionToken := req.Context().Value(SessionTokenContextKey)
	if sessionToken == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	// Revoke old token
	err := handler.SessionTokenInteractor.DeleteSessionToken(sessionToken.(string))
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Failed to revoke old token: %v", err.Error()))
		return
	}

	// Create new token
	token, err := handler.SessionTokenInteractor.UpdateSessionToken(userUID.(string))
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusCreated, CreateSessionTokenResponseV0{
		SessionToken: token,
		Message:      "Token Updated",
		Success:      true,
	})
}

func (handler *WebHandler) DeleteSessionToken(res http.ResponseWriter, req *http.Request) {
	sessionToken := req.Context().Value(SessionTokenContextKey)
	if sessionToken == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication")
		return
	}

	err := handler.SessionTokenInteractor.DeleteSessionToken(sessionToken.(string))
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusCreated, DeleteSessionTokenResponseV0{
		Message: "Token Deleted",
		Success: true,
	})
}
