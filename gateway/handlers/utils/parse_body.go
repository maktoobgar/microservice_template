package utils

import (
	"encoding/json"
	"io"

	"service/pkg/errors"
)

func ParseBody(body io.ReadCloser, output interface{}) {
	bytes, err1 := io.ReadAll(body)
	err2 := json.Unmarshal(bytes, output)
	if err1 != nil || err2 != nil {
		panic(errors.New(errors.InvalidStatus, errors.Resend, "BodyNotProvidedProperly"))
	}
}
