package httpHandlers

import (
	"net/http"

	g "service/gateway/global"

	"service/pkg/errors"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	panic(errors.New(errors.NotFoundStatus, errors.DoNothing, "PageNotFound"))
}

var NotFound = g.Handler{
	Handler: notFound,
}
