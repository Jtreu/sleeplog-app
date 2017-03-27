package web

import (
	"encoding/json"
	"net/http"
)

type WebHandler struct {
	SessionTokenInteractor SessionTokenInteractor
	UserInteractor         UserInteractor
	util                   *Utility
}

type Utility struct {
	renderer *Renderer
}

func NewWebHandler() *WebHandler {
	util := Utility{}
	util.renderer = NewRenderer(RendererOptions{
		IndentJSON: true,
	}, JSON)
	webHandler := &WebHandler{}
	webHandler.util = &util
	return webHandler
}

func (util *Utility) DecodeRequestBody(res http.ResponseWriter, req *http.Request, target interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		return err
	}
	return nil
}
