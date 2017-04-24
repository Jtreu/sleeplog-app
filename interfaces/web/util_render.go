package web

import (
	"net/http"

	"github.com/unrolled/render"
)

const JSON = "json"
const XML = "xml"
const Data = "octet-stream"
const Text = "text"

type RendererOptions render.Options

type Renderer struct {
	r                 *render.Render
	DefaultRenderType string
}

func NewRenderer(options RendererOptions, defaultRenderType string) *Renderer {
	r := render.New(render.Options(options))
	return &Renderer{r, defaultRenderType}
}

func (renderer *Renderer) Render(w http.ResponseWriter, req *http.Request, status int, v interface{}) {
	acceptHeaders := NewAcceptHeadersFromString(req.Header.Get("accept"))

	renderType := renderer.DefaultRenderType
	for _, h := range acceptHeaders {
		m := h.MediaType
		if m.SubType == JSON || m.Suffix == JSON {
			renderType = JSON
			break
		}
		if m.SubType == Data || m.Suffix == Data {
			renderType = Data
			break
		}
		if m.SubType == Text || m.Suffix == Text {
			renderType = Text
			break
		}
	}
	switch renderType {
	case JSON:
		renderer.JSON(w, status, v)
	case Data:
		renderer.Data(w, status, v.([]byte))
	case Text:
		renderer.Text(w, status, v.([]byte))
	default:
		renderer.JSON(w, status, v)
	}
}

type ErrorResponseV0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

// RenderError renders an error
func (renderer *Renderer) Error(w http.ResponseWriter, req *http.Request, status int, message string) {
	renderer.Render(w, req, status, ErrorResponseV0{
		Message: message,
		Success: false,
	})
}

func (renderer *Renderer) JSON(w http.ResponseWriter, status int, v interface{}) {
	renderer.r.JSON(w, status, v)
}

func (renderer *Renderer) Data(w http.ResponseWriter, status int, v []byte) {
	renderer.r.Data(w, status, v)
}
func (renderer *Renderer) Text(w http.ResponseWriter, status int, v []byte) {
	w.WriteHeader(status)
	w.Write(v)
}
