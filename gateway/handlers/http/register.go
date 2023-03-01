package httpHandlers

import (
	"net/http"
	"service/auth/auth_service"
	"service/gateway/calls"
	"service/gateway/dto"
	g "service/gateway/global"
	"service/gateway/handlers/utils"
	"service/pkg/translator"
)

func register(w http.ResponseWriter, r *http.Request) {
	req := &dto.RegisterRequest{}
	ctx := r.Context()
	translate := ctx.Value(g.TranslateContext).(translator.TranslatorFunc)
	utils.ParseBody(r.Body, req)
	utils.ValidateBody(req, dto.RegisterValidator, translate)

	s := calls.NewAuthService()
	s.CallAuth(func(auth auth_service.AuthClient) {
		resGrpc, err := auth.Register(ctx, &auth_service.RegisterRequest{
			PhoneNumber: req.PhoneNumber,
			Password:    req.Password,
		})
		if resGrpc != nil {
			s.Check(resGrpc.Error, err)
		} else {
			s.Check(nil, err)
		}
	})

	w.WriteHeader(http.StatusCreated)
}

var Register = g.Handler{
	Handler: register,
}
