package service

import (
	"context"
	"database/sql"
	"net/http"
	"runtime/debug"
	g "service/auth/global"
	"service/auth/models"
	"service/auth/service_definition"
	"service/pkg/errors"
	"service/pkg/repositories"
)

type service struct {
	service_definition.UnimplementedAuthServer
}

var serverInstance = service{}

func New() service_definition.AuthServer {
	return &serverInstance
}

func (s *service) Panic(next func() *service_definition.Error) func() *service_definition.Error {
	return func() (e *service_definition.Error) {
		defer func() {
			errInterface := recover()
			if errInterface == nil {
				return
			}
			stack := string(debug.Stack())
			g.Logger.PanicMicroservice(errInterface, g.Name, stack)

			e = &service_definition.Error{
				Message: "InternalServerError",
				Action:  int32(errors.Report),
				Code:    http.StatusInternalServerError,
			}
		}()
		e = next()
		return
	}
}

func (s *service) GetUserByID(db *sql.DB, ctx context.Context, id string) (*models.User, *service_definition.Error) {
	user := &models.User{}
	query := repositories.Select(user.Name(), user, map[string]any{
		"id": id,
	})
	err := db.QueryRowContext(ctx, query).Scan(user)
	if err != nil {
		err := &service_definition.Error{
			Code:    int32(errors.NotFoundStatus),
			Action:  int32(errors.Resend),
			Message: "UserNotFound",
		}

		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByEmailOrPhone(db *sql.DB, ctx context.Context, email, phone string) (*models.User, *service_definition.Error) {
	user := &models.User{}
	query := ""
	if email != "" {
		query = repositories.Select(user.Name(), user, map[string]any{
			"email": email,
		})
	} else if phone != "" {
		query = repositories.Select(user.Name(), user, map[string]any{
			"email": phone,
		})
	} else {
		return nil, &service_definition.Error{
			Code:    int32(errors.InvalidStatus),
			Action:  int32(errors.Resend),
			Message: "EmailOrPhoneRequired",
		}
	}

	err := db.QueryRowContext(ctx, query).Scan(user)
	if err != nil {
		return nil, &service_definition.Error{
			Code:    int32(errors.NotFoundStatus),
			Action:  int32(errors.Resend),
			Message: "UserNotFound",
		}
	}

	return user, nil
}
