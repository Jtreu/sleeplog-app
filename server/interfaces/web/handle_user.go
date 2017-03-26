package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"

	"github.com/jtreu/sleeplog-app/server/domain"
)

type ListUsersResposeV0 struct {
	Users domain.Users `json:"users"`
}

type CreateUserRequestV0 struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type GetUserResponseV0 struct {
	User    domain.User `json:"user"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type CreateUserResponseV0 struct {
	User    domain.User `json:"user"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type ListUsersResponseV0 struct {
	Users   domain.Users `json:"users"`
	Message string       `json:"message,omitempty"`
	Success bool         `json:"success,omitempty"`
}

type UpdateUserAliasRequestV0 struct {
	Alias string `json:"alias"`
}

type UpdateUserAliasResponseV0 struct {
	User    domain.User `json:"user"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type UpdateUsernameRequestV0 struct {
	Username string `json:"username"`
}

type UpdateUsernameResponseV0 struct {
	User    domain.User `json:"user"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type UpdateUserEmailRequestV0 struct {
	Email string `json:"email"`
}

type UpdateUserEmailResponseV0 struct {
	User struct {
		Email string `json:"email,omitempty"`
	} `json:"user"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type UpdateUserDescriptionRequestV0 struct {
	Description string `json:"description"`
}

type UpdateUserDescriptionResponseV0 struct {
	User    domain.User `json:"user"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type UpdateUserPasswordRequestV0 struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UpdateUserPasswordResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type VerifyUserGamePlayAccessResponseV0 struct {
	PlayAccessStatus string `json:"playAccessStatus,omitempty"`
	Message          string `json:"message,omitempty"`
	Success          bool   `json:"success,omitempty"`
}

type ConfirmUserEmailRequestV0 struct {
	ConfirmationCode string `json:"confirmationCode,omitempty"`
}

type ConfirmUserEmailResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type ResetUserPasswordRequestV0 struct {
	Email string `json:"email,omitempty"`
}

type ResetUserPasswordResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type UpdateUserPasswordWithCodeRequestV0 struct {
	ResetCode string `json:"resetCode,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UpdateUserPasswordWithCodeResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type UserInteractor interface {
	Create(email, username, fullname, password string) (domain.User, error)
	ListUsers(field, query, lastID string, limit int, sort string) domain.Users
	GetByUsername(username string) (domain.User, error)
	Get(uid string) (domain.User, error)

	UpdateAlias(alias string, user domain.User) (domain.User, error)
	UpdateUsername(username string, user domain.User) (domain.User, error)
	UpdateEmail(email string, user domain.User) (domain.User, error)
	UpdateDescription(description string, user domain.User) (domain.User, error)
	UpdatePassword(oldPassword string, newPassword string, user domain.User) error

	ConfirmEmail(code string) error
	ResetPassword(email string) error
	UpdatePasswordWithCode(code, password string) error
}

func (handler *WebHandler) GetUser(res http.ResponseWriter, req *http.Request) {
	username := bone.GetValue(req, "username")

	user, err := handler.UserInteractor.GetByUsername(username)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusNotFound, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusOK, GetUserResponseV0{
		User:    user,
		Message: "User retrieved",
		Success: true,
	})
}

func (handler *WebHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var body CreateUserRequestV0
	err := handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	user, err := handler.UserInteractor.Create(body.Email, body.Username, body.Fullname, body.Password)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusCreated, CreateUserResponseV0{
		User:    user,
		Message: "User created",
		Success: true,
	})
}

func (handler *WebHandler) ListUsers(res http.ResponseWriter, req *http.Request) {
	field := req.FormValue("field")
	query := req.FormValue("q")
	lastID := req.FormValue("last_id")
	perPageStr := req.FormValue("per_page")
	sort := req.FormValue("sort")

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = 20
	}

	users := handler.UserInteractor.ListUsers(field, query, lastID, perPage, sort)

	handler.util.renderer.Render(res, req, http.StatusOK, ListUsersResponseV0{
		Users:   users,
		Message: "User list retrieved",
		Success: true,
	})
}

func (handler *WebHandler) UpdateUserAlias(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	username := bone.GetValue(req, "username")
	user, err := handler.UserInteractor.GetByUsername(username)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	if user.UID != userUID {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, fmt.Sprintf("Invalid identity"))
		return
	}

	var body UpdateUserAliasRequestV0
	err = handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	user, err = handler.UserInteractor.UpdateAlias(body.Alias, user)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusOK, UpdateUserAliasResponseV0{
		User:    user,
		Message: "Full name updated",
		Success: true,
	})
}

func (handler *WebHandler) UpdateUsername(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	username := bone.GetValue(req, "username")
	user, err := handler.UserInteractor.GetByUsername(username)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	if user.UID != userUID {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, fmt.Sprintf("Invalid identity"))
		return
	}

	var body UpdateUsernameRequestV0
	err = handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	user, err = handler.UserInteractor.UpdateUsername(body.Username, user)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusOK, UpdateUsernameResponseV0{
		User:    user,
		Message: "Username updated",
		Success: true,
	})
}

func (handler *WebHandler) UpdateUserEmail(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	username := bone.GetValue(req, "username")
	user, err := handler.UserInteractor.GetByUsername(username)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	if user.UID != userUID {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, fmt.Sprintf("Invalid identity"))
		return
	}

	var body UpdateUserEmailRequestV0
	err = handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	user, err = handler.UserInteractor.UpdateEmail(body.Email, user)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	data := UpdateUserEmailResponseV0{
		Message: "Email updated",
		Success: true,
	}

	data.User.Email = user.Email

	handler.util.renderer.Render(res, req, http.StatusOK, data)
}

func (handler *WebHandler) UpdateUserDescription(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	username := bone.GetValue(req, "username")
	user, err := handler.UserInteractor.GetByUsername(username)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	if user.UID != userUID {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, fmt.Sprintf("Invalid identity"))
		return
	}

	var body UpdateUserDescriptionRequestV0
	err = handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	user, err = handler.UserInteractor.UpdateDescription(body.Description, user)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusOK, UpdateUserDescriptionResponseV0{
		User:    user,
		Message: "Description updated",
		Success: true,
	})
}

func (handler *WebHandler) UpdateUserPassword(res http.ResponseWriter, req *http.Request) {
	userUID := req.Context().Value(UserUIDContextKey)
	if userUID == nil {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires authentication.")
		return
	}

	username := bone.GetValue(req, "username")
	user, err := handler.UserInteractor.GetByUsername(username)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, err.Error())
		return
	}

	if user.UID != userUID {
		handler.util.renderer.Error(res, req, http.StatusUnauthorized, fmt.Sprintf("Invalid identity"))
		return
	}

	var body UpdateUserPasswordRequestV0
	err = handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	err = handler.UserInteractor.UpdatePassword(body.OldPassword, body.NewPassword, user)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusOK, UpdateUserPasswordResponseV0{
		Message: "Password updated",
		Success: true,
	})
}

func (handler *WebHandler) VerifyUserGamePlayAccess(res http.ResponseWriter, req *http.Request) {
	handler.util.renderer.Render(res, req, http.StatusOK, VerifyUserGamePlayAccessResponseV0{
		Message: "Retrieved play access status",
		Success: true,
	})
}

func (handler *WebHandler) ConfirmUserEmail(res http.ResponseWriter, req *http.Request) {
	var body ConfirmUserEmailRequestV0
	err := handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	err = handler.UserInteractor.ConfirmEmail(body.ConfirmationCode)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusCreated, ConfirmUserEmailResponseV0{
		Message: "Email confirmed",
		Success: true,
	})
}

func (handler *WebHandler) UpdateUserPasswordWithCode(res http.ResponseWriter, req *http.Request) {
	var body UpdateUserPasswordWithCodeRequestV0
	err := handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	err = handler.UserInteractor.UpdatePasswordWithCode(body.ResetCode, body.Password)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusCreated, UpdateUserPasswordWithCodeResponseV0{
		Message: "Password updated!",
		Success: true,
	})
}

func (handler *WebHandler) ResetUserPassword(res http.ResponseWriter, req *http.Request) {
	var body ResetUserPasswordRequestV0
	err := handler.util.DecodeRequestBody(res, req, &body)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return
	}

	err = handler.UserInteractor.ResetPassword(body.Email)
	if err != nil {
		handler.util.renderer.Error(res, req, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.util.renderer.Render(res, req, http.StatusCreated, ResetUserPasswordResponseV0{
		Message: "Password reset code sent",
		Success: true,
	})
}
