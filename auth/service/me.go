package service

import (
	"context"
	"log"
	"service/auth/auth_service"
	g "service/auth/global"
)

func (s *service) Me(ctx context.Context, in *auth_service.MeRequest) (*auth_service.MeResponse, error) {
	res := &auth_service.MeResponse{}
	err := s.Panic(func() *auth_service.Error {
		claims, err := s.IsAccessTokenValid(in.AccessToken)
		if err != nil {
			return err
		}

		user, err := s.GetUserByID(g.DB, ctx, claims.UserID)
		if err != nil {
			return err
		}

		res.User = &auth_service.User{
			ID:                   int32(user.ID),
			PhoneNumber:          user.PhoneNumber,
			Email:                user.Email,
			PhoneNumberConfirmed: user.PhoneNumberConfirmed,
			EmailConfirmed:       user.EmailConfirmed,
			Role:                 user.Role,
			JoinedDate:           user.JoinedDate.Unix(),
		}

		return nil
	})()

	res.Error = err
	if g.CFG.Debug && err != nil {
		log.Println(err)
	}
	return res, nil
}
