package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type InternalServerError struct {
	message string
}

func (*InternalServerError) GetStatusCode() int {
	return http.StatusInternalServerError
}

func (i *InternalServerError) Error() string {
	return i.message
}

func (i *InternalServerError) GetTitle() string {
	return constants.InternalServerError
}

func (i *InternalServerError) GetCode() uint {
	return constants.InternalCodeError
}

func ThrowInternalServerError(msg string) DomainError {
	return &InternalServerError{message: msg}
}
