package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type DatabaseCommandError struct {
	message string
}

func (*DatabaseCommandError) GetStatusCode() int {
	return http.StatusInternalServerError
}

func (d *DatabaseCommandError) Error() string {
	return d.message
}

func (d *DatabaseCommandError) GetTitle() string {
	return constants.DatabaseCommandError
}

func (d *DatabaseCommandError) GetCode() uint {
	return 0_01_002_001
}

func ThrowDatabaseCommandError(message string) DomainError {
	return &DatabaseCommandError{message: message}
}
