package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

const UserUIDContextKey = "userUID"
const SessionTokenContextKey = "sessionToken"

func (handler *WebHandler) EnsureAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearerToken := req.Header.Get("Bearer")
		if bearerToken == "" {
			handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires Authentication: Invalid Bearer Header.")
			return
		}

		userUID, err := handler.SessionTokenInteractor.VerifySessionToken(bearerToken)
		if err != nil {
			handler.util.renderer.Error(res, req, http.StatusUnauthorized, "Requires Authentication: Invalid Bearer Header.")
			return
		}

		userCtx := context.WithValue(req.Context(), UserUIDContextKey, userUID)
		req = req.WithContext(userCtx)

		sessionCtx := context.WithValue(req.Context(), SessionTokenContextKey, bearerToken)
		req = req.WithContext(sessionCtx)

		h.ServeHTTP(res, req)
	})
}

func (handler *WebHandler) CheckAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearerToken := req.Header.Get("Bearer")

		userUID, _ := handler.SessionTokenInteractor.VerifySessionToken(bearerToken)

		userCtx := context.WithValue(req.Context(), UserUIDContextKey, userUID)
		req = req.WithContext(userCtx)

		sessionCtx := context.WithValue(req.Context(), SessionTokenContextKey, bearerToken)
		req = req.WithContext(sessionCtx)

		h.ServeHTTP(res, req)
	})
}

func (handler *WebHandler) ForceHTTPS(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if req.Header.Get("X-Forwarded-Proto") == "http" {
		urlStr := fmt.Sprintf("https://%v%v", req.Host, req.RequestURI)
		http.Redirect(res, req, urlStr, http.StatusMovedPermanently)
		return
	}
	next(res, req)
}

func (handler *WebHandler) CORS(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if req.Method == "OPTIONS" {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Bearer")
		res.Write([]byte{1})
		return
	}

	if req.Method == "DELETE" {
		res.Header().Set("Access-Control-Allow-Origin", "*")
	}

	cors.Default().ServeHTTP(res, req, next)
}
