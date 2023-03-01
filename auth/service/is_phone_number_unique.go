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

func (s *service) IsPhoneNumberUnique(ctx context.Context, in *auth_service.IsPhoneNumberUniqueRequest) (*auth_service.IsPhoneNumberUniqueResponse, error) {
	res := &auth_service.IsPhoneNumberUniqueResponse{}
	err := s.Panic(func() *auth_service.Error {
		query := repositories.SelectCount(models.UserName, map[string]any{
			"phone_number": in.PhoneNumber,
		})
		count := -1
		err := g.DB.QueryRowContext(ctx, query).Scan(&count)
		if err != nil {
			return &auth_service.Error{
				Code:    int32(errors.UnexpectedStatus),
				Action:  int32(errors.Report),
				Message: "PhoneNumberDuplicationCheckFailed",
			}
		}
		res.OK = count == 0

		return nil
	})()
	res.Error = err

	if err != nil {
		if g.CFG.Debug {
			log.Println(err)
		}
		return res, nil
	}
	return res, nil
}
