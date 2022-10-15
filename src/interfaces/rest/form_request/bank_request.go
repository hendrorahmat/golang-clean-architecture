package form_request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/errors"
)

type BankRequest struct {
	Page string `form:"page" json:"page"`
	User string `form:"user" json:"user"`
}

func (b BankRequest) Validate(ctx *gin.Context) *errors.GeneralError {
	request := NewRequestValidation()
	request.AddParam("page", &b.Page, validation.Required, validation.Min(1))
	request.AddParam("limit", nil, validation.Required, validation.Min(1))
	request.SetCustomCode(errors.QueryParamDataInvalid)
	return request.Validate(ctx)
}
