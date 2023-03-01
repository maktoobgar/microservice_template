package g

import (
	_ "embed"
	"net/http"

	"service/config"

	"service/pkg/logging"
	"service/pkg/translator"
)

//go:embed version
var Version string

//go:embed name
var Name string

// Context type and keys
type contextKey string

var (
	TranslateContext contextKey = "translate"
)

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
