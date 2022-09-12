package errors

import (
	"fmt"
	"strings"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ErrorCode uint

type ValidationErrors map[string]string

type CommonError struct {
	ClientMessage    string           `json:"message"`
	SystemMessage    interface{}      `json:"errorMessages"`
	ValidationErrors ValidationErrors `json:"validationErrors,omitempty"`
	ErrorCode        ErrorCode        `json:"code"`
	ErrorMessage     *string          `json:"-"`
	ErrorTrace       *string          `json:"-"`
}
func (err CommonError) Error() string {
	return fmt.Sprintf("CommonError: %+v. Trace: %+v", err.ErrorMessage, err.ErrorTrace)
}

func buildValidationError(err error) ValidationErrors {
	var errors ValidationErrors = map[string]string{}

	errValidate := strings.Split(err.Error(), ";")
	for _, err := range errValidate {
		errPerField := strings.Split(err, ":")
		if len(errPerField[0]) <= 1 {
			errors["error"] = errPerField[0]
		} else {
			errors[strings.TrimSpace(errPerField[0])] = strings.TrimSpace(errPerField[1])
		}
	}

	return errors
}

func NewError(errCode ErrorCode, err error) *CommonError {
	if _err, ok := err.(*CommonError); ok {
		return _err
	}

	var errMsg *string
	var errTrace *string
	var clientMessage string = "Unknown error."
	var systemMessage interface{} = "Unknown error."
	var commonError = errorCodes[errCode]

	if err != nil {
		s := err.Error()
		errMsg = &s

		ss := fmt.Sprintf("%+v", err)
		errTrace = &ss

		if errCode == UNKNOWN_ERROR {
			systemMessage = ss
		}
	}

	if commonError == nil {
		return &CommonError{
			ClientMessage: clientMessage,
			SystemMessage: systemMessage,
			ErrorCode:     errCode,
			ErrorTrace:    errTrace,
			ErrorMessage:  errMsg,
		}
	}

	return &CommonError{
		ClientMessage: commonError.ClientMessage,
		SystemMessage: commonError.SystemMessage,
		ErrorCode:     errCode,
		ErrorTrace:    errTrace,
		ErrorMessage:  errMsg,
	}
}

func (err *CommonError) SetClientMessage(message string) {
	err.ClientMessage = message
}

func (err *CommonError) SetSystemMessage(message interface{}) {
	err.SystemMessage = message
}

func (err *CommonError) SetValidationMessage(message interface{}) {
	if _err, ok := message.(validation.Errors); ok {
		err.ValidationErrors = buildValidationError(_err)
	}
}
