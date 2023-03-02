package service

import (
	"context"
	"log"
	"service/auth/auth_service"
	g "service/auth/global"
)

func (s *service) RefreshTokens(ctx context.Context, in *auth_service.RefreshTokensRequest) (*auth_service.RefreshTokensResponse, error) {
	res := &auth_service.RefreshTokensResponse{}
	err := s.Panic(func() *auth_service.Error {
		claims, err := s.GetClaimsFromRefreshToken(in.RefreshToken)
		if err != nil {
			return err
		}

		accessToken, err := s.CreateAccessToken(claims.UserID)
		if err != nil {
			return err
		}
		refreshToken, err := s.CreateRefreshToken(claims.UserID)
		if err != nil {
			return err
		}

		res.AccessToken = accessToken
		res.RefreshToken = refreshToken

		return nil
	})()

	res.Error = err
	if g.CFG.Debug && err != nil {
		log.Println(err)
	}
	return res, nil
}
