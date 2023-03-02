package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	g "service/gateway/global"

	"service/pkg/errors"
	"service/pkg/translator"
)

type PanicResponse struct {
	Message string `json:"message"`
	Action  int    `json:"action"`
	Code    int    `json:"code"`
	Errors  any    `json:"errors"`
}

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		translate := r.Context().Value(g.TranslateContext).(translator.TranslatorFunc)
		defer func() {
			errInterface := recover()
			if errInterface == nil {
				return
			}
			if err, ok := errInterface.(error); ok && errors.IsServerError(err) {
				code, action, message, errors := errors.HttpError(err)
				res := PanicResponse{
					Message: translate(message),
					Action:  action,
					Code:    code,
					Errors:  errors,
				}
				resBytes, _ := json.Marshal(res)
				if g.CFG.Debug {
					log.Println(message)
				}
				w.WriteHeader(code)
				w.Write(resBytes)
			} else {
				stack := string(debug.Stack())
				g.Logger.Panic(errInterface, r, stack)
				res := PanicResponse{
					Message: translate("InternalServerError"),
					Action:  errors.Report,
					Code:    http.StatusInternalServerError,
					Errors:  nil,
				}
				resBytes, _ := json.Marshal(res)
				w.WriteHeader(res.Code)
				w.Write(resBytes)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
