package dto

import (
	"service/auth/models"

	"github.com/golodash/galidator"
)

type LoginRequest struct {
	Username string `json:"username" g:"or=phone|email,required"`
	Password string `json:"password" g:"max=32,required"`
}

var LoginValidator = gen.Validator(LoginRequest{}, galidator.Messages{
	"or": "NotAPhoneNumberOrEmailAddress",
})

type LoginResponse struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
