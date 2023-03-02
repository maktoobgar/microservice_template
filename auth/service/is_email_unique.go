package service

import (
	"context"
	"log"
	"service/auth/auth_service"
	g "service/auth/global"
	"service/auth/models"
	"service/pkg/errors"
	"service/pkg/repositories"
)

func (s *service) IsEmailUnique(ctx context.Context, in *auth_service.IsEmailUniqueRequest) (*auth_service.IsEmailUniqueResponse, error) {
	res := &auth_service.IsEmailUniqueResponse{}
	err := s.Panic(func() *auth_service.Error {
		query := repositories.SelectCount(models.UserName, map[string]any{
			"email": in.Email,
		})
		count := -1
		err := g.DB.QueryRowContext(ctx, query).Scan(&count)
		if err != nil {
			return &auth_service.Error{
				Code:    int32(errors.UnexpectedStatus),
				Action:  int32(errors.Report),
				Message: "EmailDuplicated",
			}
		}
		res.OK = count == 0

		return nil
	})()

	res.Error = err
	if g.CFG.Debug && err != nil {
		log.Println(err)
	}
	return res, nil
}
