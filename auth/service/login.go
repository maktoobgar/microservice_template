package service

import (
	"context"
	"log"
	"service/auth/auth_service"
	g "service/auth/global"
	cerrors "service/pkg/errors"
)

func (s *service) Login(ctx context.Context, in *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	res := &auth_service.LoginResponse{}
	err := s.Panic(func() *auth_service.Error {
		user, err := s.GetUserByEmailOrPhone(g.DB, ctx, in.Email, in.PhoneNumber)
		if err != nil {
			return err
		}

		if !s.CheckPasswordHash(in.Password, user.Password) {
			return &auth_service.Error{
				Code:    int32(cerrors.UnauthorizedStatus),
				Action:  int32(cerrors.ReSignIn),
				Message: "PasswordDoesNotMatch",
			}
		}

		accessToken, err := s.CreateAccessToken(user.ID)
		if err != nil {
			return err
		}
		refreshToken, err := s.CreateRefreshToken(user.ID)
		if err != nil {
			return err
		}

		user.AccessToken = accessToken
		user.RefreshToken = refreshToken

		err = s.UpdateAccessRefreshToken(g.DB, ctx, user)
		if err != nil {
			return err
		}

		res.User = &auth_service.User{
			ID:                   int32(user.ID),
			PhoneNumber:          user.PhoneNumber,
			Email:                user.Email,
			PhoneNumberConfirmed: user.PhoneNumberConfirmed,
			EmailConfirmed:       user.EmailConfirmed,
			JoinedDate:           user.JoinedDate.Unix(),
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
