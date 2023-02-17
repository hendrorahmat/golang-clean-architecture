package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
	"strings"
)

// FieldsRequired should only be used when one of the fields is empty
type FieldsRequired struct {
	fields string
}

func (f *FieldsRequired) Error() string {
	return strings.Replace(constants.ErrorFieldNotFound, "{fields}", f.fields, 1)
}

func (f *FieldsRequired) GetTitle() string {
	return constants.DataInvalid
}

func (f *FieldsRequired) GetCode() uint {
	return constants.FieldsRequiredCode
}

func (f *FieldsRequired) GetStatusCode() int {
	return http.StatusBadRequest
}

func ThrowFieldsRequired(fields ...string) DomainError {
	return &FieldsRequired{fields: strings.Join(fields, ",")}
}
