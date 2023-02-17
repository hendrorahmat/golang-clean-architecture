package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type QueryParamInvalid struct {
	fields string
}

func (*QueryParamInvalid) GetStatusCode() int {
	return http.StatusBadRequest
}

func (q *QueryParamInvalid) Error() string {
	return q.fields
}

func (q *QueryParamInvalid) GetTitle() string {
	return constants.QueryParamInvalid
}

func (q *QueryParamInvalid) GetCode() uint {
	return constants.QueryParamDataInvalidCode
}

func ThrowQueryParamInvalid(fields string) DomainError {
	return &QueryParamInvalid{fields: fields}
}
