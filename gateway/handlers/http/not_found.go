package httpHandlers

import (
	"net/http"

	g "service/gateway/global"

	"service/pkg/errors"
	"service/pkg/translator"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	panic(errors.New(errors.NotFoundStatus, errors.DoNothing, translate("PageNotFound")))
}

var NotFound = g.Handler{
	Handler: notFound,
}
