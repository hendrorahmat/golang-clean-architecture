package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type InvalidHeader struct {
	message string
}

func (*InvalidHeader) GetStatusCode() int {
	return http.StatusBadRequest
}

func (i *InvalidHeader) Error() string {
	return i.message
}

func (i *InvalidHeader) GetTitle() string {
	return constants.HeaderInvalid
}

func (i *InvalidHeader) GetCode() uint {
	return constants.InvalidHeaderCode
}

func ThrowInvalidHeader(message string) DomainError {
	return &InvalidHeader{message: message}
}
