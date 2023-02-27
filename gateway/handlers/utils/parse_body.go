package utils

import (
	"encoding/json"
	"io"

	"service/pkg/errors"
	"service/pkg/translator"
)

func ParseBody(body io.ReadCloser, translator translator.TranslatorFunc, output interface{}) {
	bytes, err1 := io.ReadAll(body)
	err2 := json.Unmarshal(bytes, output)
	if err1 != nil || err2 != nil {
		panic(errors.New(errors.InvalidStatus, errors.Resend, translator("BodyNotProvidedProperly")))
	}
}
