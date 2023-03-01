package service

import (
	"context"
	"database/sql"
	"net/http"
	"runtime/debug"
	"service/auth/auth_service"
	g "service/auth/global"
	"service/auth/models"
	"service/pkg/errors"
	"service/pkg/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	auth_service.UnimplementedAuthServer
}

var serverInstance = service{}

func New() auth_service.AuthServer {
	return &serverInstance
}

func (s *service) Panic(next func() *auth_service.Error) func() *auth_service.Error {
	return func() (e *auth_service.Error) {
		defer func() {
			errInterface := recover()
			if errInterface == nil {
				return
			}
			stack := string(debug.Stack())
			g.Logger.PanicMicroservice(errInterface, g.Name, stack)

			e = &auth_service.Error{
				Message: "InternalServerError",
				Action:  int32(errors.Report),
				Code:    http.StatusInternalServerError,
			}
		}()
		e = next()
		return
	}
}

func (s *service) CreateUser(db *sql.DB, ctx context.Context, phone_number string, password string) *auth_service.Error {
	user := &models.User{
		PhoneNumber: phone_number,
		JoinedDate:  time.Now(),
		Password:    s.HashPassword(password),
	}

	query := repositories.InsertInto(user.Name(), user)
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		return &auth_service.Error{
			Code:    int32(errors.InvalidStatus),
			Action:  int32(errors.Resend),
			Message: "CreateUserFailure",
		}
	}
	user.ID, _ = result.LastInsertId()

	return nil
}

func (s *service) GetUserByID(db *sql.DB, ctx context.Context, id string) (*models.User, *auth_service.Error) {
	user := &models.User{}
	query := repositories.Select(user.Name(), user, map[string]any{
		"id": id,
	})
	err := db.QueryRowContext(ctx, query).Scan(user)
	if err != nil {
		err := &auth_service.Error{
			Code:    int32(errors.NotFoundStatus),
			Action:  int32(errors.Resend),
			Message: "UserNotFound",
		}

		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByEmailOrPhone(db *sql.DB, ctx context.Context, email, phone string) (*models.User, *auth_service.Error) {
	user := &models.User{}
	query := ""
	if email != "" {
		query = repositories.Select(user.Name(), user, map[string]any{
			"email": email,
		})
	} else if phone != "" {
		query = repositories.Select(user.Name(), user, map[string]any{
			"phone_number": phone,
		})
	} else {
		return nil, &auth_service.Error{
			Code:    int32(errors.InvalidStatus),
			Action:  int32(errors.Resend),
			Message: "EmailOrPhoneRequired",
		}
	}

	// id, phone_number, email, password, phone_number_confirmed, email_confirmed, role, joined_date FROM users
	err := db.QueryRowContext(ctx, query).Scan(&user.ID, &user.PhoneNumber, &user.Email, &user.Password, &user.PhoneNumberConfirmed, &user.EmailConfirmed, &user.Role, &user.JoinedDate)
	if err != nil {
		return nil, &auth_service.Error{
			Code:    int32(errors.NotFoundStatus),
			Action:  int32(errors.Resend),
			Message: "UserNotFound",
		}
	}

	return user, nil
}

func (s *service) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (s *service) CheckPasswordHash(password, hash string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *service) CreateAccessToken(user *models.User) (string, *auth_service.Error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		ID:   user.ID,
		Type: models.AccessTokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(g.SecretKey)
	if err != nil {
		return "", &auth_service.Error{
			Code:    int32(errors.UnexpectedStatus),
			Action:  int32(errors.Report),
			Message: "TokenGenerationFailed",
		}
	}

	return tokenString, nil
}

func (s *service) CreateRefreshToken(user *models.User) (string, *auth_service.Error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &models.Claims{
		ID:   user.ID,
		Type: models.RefreshTokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(g.SecretKey)
	if err != nil {
		return "", &auth_service.Error{
			Code:    int32(errors.UnexpectedStatus),
			Action:  int32(errors.Report),
			Message: "TokenGenerationFailed",
		}
	}

	return tokenString, nil
}
