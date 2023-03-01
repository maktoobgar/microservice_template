package dto

type RegisterRequest struct {
	// TODO: Add unique for PhoneNumber
	PhoneNumber string `json:"phone_number" g:"phone,unique_phone_number,required"`
	Password    string `json:"password" g:"max=32,required"`
}

var RegisterValidator = gen.Validator(RegisterRequest{})
