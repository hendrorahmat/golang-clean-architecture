package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type DatabaseQueryError struct {
	message string
}

func (*DatabaseQueryError) GetStatusCode() int {
	return http.StatusInternalServerError
}

func (d *DatabaseQueryError) Error() string {
	return d.message
}

func (d *DatabaseQueryError) GetTitle() string {
	return constants.DatabaseQueryError
}

func (d *DatabaseQueryError) GetCode() uint {
	return constants.DBQueryErrorCode
}

func ThrowDatabaseQueryError(message string) DomainError {
	return &DatabaseQueryError{message: message}
}
