package httpHandlers

import (
	"encoding/json"
	"net/http"
	"service/auth/auth_service"
	"service/gateway/calls"
	"service/gateway/dto"
	g "service/gateway/global"
	"service/pkg/errors"
)

func refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := r.Header.Get("Token")
	if token == "" {
		panic(errors.New(errors.InvalidStatus, errors.ReSignIn, "TokenNotSpecified"))
	}

	s := calls.NewAuthService()
	var res *dto.RefreshResponse = nil
	s.CallAuth(func(auth auth_service.AuthClient) {
		resGrpc, err := auth.RefreshTokens(ctx, &auth_service.RefreshTokensRequest{RefreshToken: token})
		if resGrpc != nil {
			s.Check(resGrpc.Error, err)
		} else {
			s.Check(nil, err)
		}

		res = &dto.RefreshResponse{
			AccessToken:  resGrpc.AccessToken,
			RefreshToken: resGrpc.RefreshToken,
		}
	})

	data, _ := json.Marshal(res)
	w.Write(data)
}

var Refresh = g.Handler{
	Handler: refresh,
}
