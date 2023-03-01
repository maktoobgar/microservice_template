package utils

import (
	"service/pkg/errors"
	"service/pkg/translator"

	"github.com/golodash/galidator"
)

func ValidateBody(data any, validator galidator.Validator, translate translator.TranslatorFunc) {
	if errs := validator.Validate(data, galidator.Translator(translate)); errs != nil {
		panic(errors.New(errors.InvalidStatus, errors.Resend, "BodyNotProvidedProperly", errs))
	}
}
