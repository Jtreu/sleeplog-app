package web

import "net/http"

type GetIndexResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func (handler *WebHandler) GetIndex(res http.ResponseWriter, req *http.Request) {
	handler.util.renderer.Render(res, req, http.StatusOK, GetIndexResponseV0{
		Message: "Welcome to api.sleeplog-app.com",
		Success: true,
	})
}

func (handler *WebHandler) GetCertBotKey(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("osvGIOUEm8cHMDa30G3Hpbhg_oI7ncvypeVEtdTarGs.DH62LGmvcUoyR7ZuXidRyIjLGdZEdDnaVwK-q3pFy7s"))
}
