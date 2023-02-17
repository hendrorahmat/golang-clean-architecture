package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type Errorlists []string
type ValidationErrors map[string]Errorlists

type ValidationError struct {
	Message          string           `json:"message"`
	ValidationErrors ValidationErrors `json:"errors"`
	ErrorCode        uint             `json:"code"`
	StatusCode       int              `json:"-"`
}

func NewError(errCode uint) *ValidationError {
	return &ValidationError{
		Message:          constants.ValidationError,
		ErrorCode:        errCode,
		StatusCode:       http.StatusUnprocessableEntity,
		ValidationErrors: make(ValidationErrors),
	}
}
