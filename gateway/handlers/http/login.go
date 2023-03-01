package httpHandlers

import (
	"encoding/json"
	"net/http"
	"service/auth/auth_service"
	"service/auth/models"
	"service/gateway/calls"
	"service/gateway/dto"
	g "service/gateway/global"
	"service/gateway/handlers/utils"
	"service/pkg/translator"
	"strings"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	req := &dto.LoginRequest{}
	ctx := r.Context()
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	utils.ParseBody(r.Body, req)
	utils.ValidateBody(req, dto.LoginValidator, translate)

	email := ""
	phone_number := ""
	if strings.Contains(req.Username, "@") {
		email = req.Username
	} else {
		phone_number = req.Username
	}

	s := calls.NewAuthService()
	var res *dto.LoginResponse = nil
	s.CallAuth(func(auth auth_service.AuthClient) {
		resGrpc, err := auth.Login(ctx, &auth_service.LoginRequest{
			PhoneNumber: phone_number,
			Email:       email,
			Password:    req.Password,
		})
		if resGrpc != nil {
			s.Check(resGrpc.Error, err)
		} else {
			s.Check(nil, err)
		}

		res = &dto.LoginResponse{
			AccessToken:  resGrpc.AccessToken,
			RefreshToken: resGrpc.RefreshToken,
			User: models.User{
				ID:                   int64(resGrpc.User.ID),
				PhoneNumber:          resGrpc.User.PhoneNumber,
				Email:                resGrpc.User.Email,
				PhoneNumberConfirmed: resGrpc.User.PhoneNumberConfirmed,
				EmailConfirmed:       resGrpc.User.EmailConfirmed,
				Role:                 resGrpc.User.Role,
				JoinedDate:           time.Unix(resGrpc.User.JoinedDate, 0),
			},
		}
	})

	data, _ := json.Marshal(res)
	w.Write(data)
}

var Login = g.Handler{
	Handler: login,
}
