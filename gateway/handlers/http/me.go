package httpHandlers

import (
	"encoding/json"
	"net/http"
	"service/auth/auth_service"
	"service/auth/models"
	"service/gateway/calls"
	"service/gateway/dto"
	g "service/gateway/global"
	"service/pkg/errors"
	"time"
)

func me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := r.Header.Get("Token")
	if token == "" {
		panic(errors.New(errors.InvalidStatus, errors.ReSignIn, "TokenNotSpecified"))
	}

	s := calls.NewAuthService()
	var res *dto.MeResponse = nil
	s.CallAuth(func(auth auth_service.AuthClient) {
		resGrpc, err := auth.Me(ctx, &auth_service.MeRequest{AccessToken: token})
		if resGrpc != nil {
			s.Check(resGrpc.Error, err)
		} else {
			s.Check(nil, err)
		}

		res = &dto.MeResponse{
			User: models.User{
				ID:                   int64(resGrpc.User.ID),
				PhoneNumber:          resGrpc.User.PhoneNumber,
				Email:                resGrpc.User.Email,
				PhoneNumberConfirmed: resGrpc.User.PhoneNumberConfirmed,
				EmailConfirmed:       resGrpc.User.EmailConfirmed,
				JoinedDate:           time.Unix(resGrpc.User.JoinedDate, 0),
			},
		}
	})

	data, _ := json.Marshal(res)
	w.Write(data)
}

var Me = g.Handler{
	Handler: me,
}
