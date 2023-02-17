package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type InvalidData struct {
	message string
}

func (*InvalidData) GetStatusCode() int {
	return http.StatusBadRequest
}

func (d *InvalidData) Error() string {
	return d.message
}

func (d *InvalidData) GetTitle() string {
	return constants.DataInvalid
}

func (d *InvalidData) GetCode() uint {
	return constants.DataPayloadInvalidCode
}

func ThrowInvalidData(msg string) DomainError {
	return &InvalidData{message: msg}
}
