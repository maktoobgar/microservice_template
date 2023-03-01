package g

import (
	"net/http"

	"service/config"

	"service/pkg/logging"
	"service/pkg/translator"
)

type contextKey string

// Handling section
type Handler struct {
	Handler func(w http.ResponseWriter, r *http.Request)
}

// Function that gets executed to host a url
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler(w, r)
}

// Config
var CFG *config.Config = nil

// Utilities
var Logger logging.Logger = nil
var Translator translator.Translator = nil

// Microservices
var AuthMic *config.Microservice = nil

// App
var Server *http.Server = nil

// Context Key Maker
func ContextKey(key string) contextKey {
	return contextKey(key)
}
