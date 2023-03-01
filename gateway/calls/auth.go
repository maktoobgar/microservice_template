package calls

import (
	"fmt"
	"service/auth/auth_service"
	g "service/gateway/global"
	"service/pkg/errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authService struct{}

var authServiceInstance = &authService{}

func NewAuthService() *authService {
	return authServiceInstance
}

func (a *authService) CallAuth(do func(auth auth_service.AuthClient)) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", g.AuthMic.IP, g.AuthMic.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(errors.New(errors.ServiceUnavailable, errors.TryLater, "AuthServiceUnavailable"))
	}
	defer conn.Close()

	auth := auth_service.NewAuthClient(conn)

	do(auth)
}

func (a *authService) Check(err *auth_service.Error, errCarrier error) {
	if err != nil {
		panic(errors.New(int(err.Code), int(err.Action), err.Message))
	}
	if errCarrier != nil && strings.Contains(errCarrier.Error(), "connection refused") {
		panic(errors.New(errors.ServiceUnavailable, errors.TryLater, "AuthServiceUnavailable"))
	}
	if errCarrier != nil {
		panic(errors.New(errors.UnexpectedStatus, errors.Report, "ServiceDidNotFinishProperly"))
	}
}
