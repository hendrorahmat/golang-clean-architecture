package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
	"strings"
)

// RecordWithFieldsNotFound should only be used where get data from db and return nil
type RecordWithFieldsNotFound struct {
	fields string
}

func (f *RecordWithFieldsNotFound) Error() string {
	return strings.Replace(constants.RecordFieldsNotFound, "{fields}", f.fields, 1)
}

func (f *RecordWithFieldsNotFound) GetTitle() string {
	return constants.DataInvalid
}

func (f *RecordWithFieldsNotFound) GetCode() uint {
	return constants.DBQueryNotFoundCode
}

func (f *RecordWithFieldsNotFound) GetStatusCode() int {
	return http.StatusNotFound
}

func ThrowRecordFieldsNotFound(fields ...string) DomainError {
	dataFields := strings.Join(fields, ",")
	return &RecordWithFieldsNotFound{fields: dataFields}
}
