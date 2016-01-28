package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/couchbaselabs/logg"
)

type handler struct {
	responseWriter http.ResponseWriter
	request        *http.Request
}

type handlerMethod func(*handler) error

func makeHandler(method handlerMethod) http.Handler {
	return http.HandlerFunc(func(r http.ResponseWriter, rq *http.Request) {
		h := newHandler(r, rq)
		err := h.invoke(method)
		if err != nil {
			logg.LogPanic("Error creating http handler: %v", err)
		}
	})
}

func newHandler(w http.ResponseWriter, rq *http.Request) *handler {
	return &handler{
		responseWriter: w,
		request:        rq,
	}
}

func (h *handler) invoke(method handlerMethod) error {
	return method(h)
}

func (h *handler) writeJSON(i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	io.WriteString(h.responseWriter, string(b))

	return nil
}
