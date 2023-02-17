package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type RecordNotFoundError struct{}

func (*RecordNotFoundError) GetStatusCode() int {
	return http.StatusNotFound
}

func (r *RecordNotFoundError) GetCode() uint {
	return constants.DBQueryNotFoundCode
}

func (r *RecordNotFoundError) Error() string {
	return constants.RecordNotFound
}

func (r *RecordNotFoundError) GetTitle() string {
	return constants.DataInvalid
}

func ThrowRecordNotFoundError() DomainError {
	return &RecordNotFoundError{}
}
