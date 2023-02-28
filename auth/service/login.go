package service

import (
	"context"
	"errors"
	g "service/auth/global"
	"service/auth/service_definition"
)

func (s *service) Login(ctx context.Context, in *service_definition.LoginRequest) (*service_definition.LoginResponse, error) {
	res := &service_definition.LoginResponse{}
	err := s.Panic(func() *service_definition.Error {

		user, err := s.GetUserByEmailOrPhone(g.DB, ctx, in.Email, in.PhoneNumber)
		if err != nil {
			return err
		}
		res.User = &service_definition.User{
			ID:          int32(user.ID),
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
		}
		return nil
	})()

	if err != nil {
		res.Error = err
		return nil, errors.New(err.Message)
	}
	return res, nil
}
