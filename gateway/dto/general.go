package dto

import (
	"context"
	"service/auth/auth_service"
	"service/gateway/calls"

	"github.com/golodash/galidator"
)

var gen = galidator.G().CustomMessages(galidator.Messages{
	"phone":               "Phone",
	"email":               "Email",
	"required":            "Required",
	"max":                 "Max",
	"unique_phone_number": "PhoneNumberDuplicated",
}).CustomValidators(galidator.Validators{
	"unique_phone_number": uniquePhoneNumber,
})

func uniquePhoneNumber(input interface{}) bool {
	s := calls.NewAuthService()
	ok := false
	s.CallAuth(func(auth auth_service.AuthClient) {
		resGrpc, err := auth.IsPhoneNumberUnique(context.TODO(), &auth_service.IsPhoneNumberUniqueRequest{
			PhoneNumber: input.(string),
		})
		if resGrpc != nil {
			s.Check(resGrpc.Error, err)
		} else {
			s.Check(nil, err)
		}

		ok = resGrpc.OK
	})
	return ok
}
