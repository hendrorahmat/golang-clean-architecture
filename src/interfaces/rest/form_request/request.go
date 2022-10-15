package form_request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/errors"
)

type paramValidation struct {
	name  string
	value any
	rules []validation.Rule
}

type requestValidation struct {
	Params     []paramValidation
	Code       *errors.ErrorCode
	StatusCode int
}

func NewRequestValidation() *requestValidation {
	return &requestValidation{}
}

func (rv *requestValidation) AddParam(name string, value any, rules ...validation.Rule) {
	param := paramValidation{
		name:  name,
		value: value,
		rules: rules,
	}
	rv.Params = append(rv.Params, param)
}

func (rv *requestValidation) Validate(ctx *gin.Context) *errors.GeneralError {
	var customErr *errors.GeneralError
	
	if rv.Code == nil {
		customErr = errors.NewError(errors.DataPayloadInvalid)
	} else {
		customErr = errors.NewError(*rv.Code)
	}

	for _, param := range rv.Params {
		errorLists := rv.ValidateParam(&param.value, param.rules...)
		if len(errorLists) > 0 {
			customErr.ValidationErrors[param.name] = errorLists
		}
	}

	if len(customErr.ValidationErrors) <= 0 {
		return nil
	}

	if rv.StatusCode == 0 {
		customErr.SetStatusCode(422)
	} else {
		customErr.SetStatusCode(rv.StatusCode)
	}
	ctx.Header("Content-Type", "application/problem+json")
	return customErr
}

func (rv *requestValidation) ValidateParam(value any, rules ...validation.Rule) errors.Errorlists {
	var errorLists []string
	for _, rule := range rules {
		errorParam := validation.Validate(&value,
			rule,
		)

		if errorParam != nil {

			errorLists = append(errorLists, fmt.Sprintf("%s", errorParam))
		}
	}

	if len(errorLists) <= 0 {
		return nil
	}

	return errorLists
}

func (rv *requestValidation) SetCustomCode(code errors.ErrorCode) {
	rv.Code = &code
}

func (rv *requestValidation) SetStatusCode(statusCode int) {
	rv.StatusCode = statusCode
}
