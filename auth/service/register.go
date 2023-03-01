package service

import (
	"context"
	"log"
	"service/auth/auth_service"
	g "service/auth/global"
)

func (s *service) Register(ctx context.Context, in *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {
	res := &auth_service.RegisterResponse{}
	err := s.Panic(func() *auth_service.Error {
		err := s.CreateUser(g.DB, ctx, in.PhoneNumber, in.Password)
		if err != nil {
			return err
		}

		return nil
	})()

	if err != nil {
		res.Error = err
		if g.CFG.Debug {
			log.Println(err)
		}
		return res, nil
	}
	return res, nil
}
