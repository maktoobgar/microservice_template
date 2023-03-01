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
		res, err := auth.IsPhoneNumberUnique(context.TODO(), &auth_service.IsPhoneNumberUniqueRequest{
			PhoneNumber: input.(string),
		})
		s.Check(res.Error, err)

		ok = res.OK
	})
	return ok
}
