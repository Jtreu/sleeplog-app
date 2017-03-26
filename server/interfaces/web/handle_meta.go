package web

import "net/http"

type GetIndexResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func (handler *WebHandler) GetIndex(res http.ResponseWriter, req *http.Request) {
	handler.util.renderer.Render(res, req, http.StatusOK, GetIndexResponseV0{
		Message: "Welcome to api.enoplay.com",
		Success: true,
	})
}

func (handler *WebHandler) GetCertBotKey(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("place_certbot_key_here"))
}
