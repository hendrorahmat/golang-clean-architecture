package form_request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/errors"
)

type BankRequest struct {
	Page int `form:"page" json:"page"`
}

func (b BankRequest) Validate(ctx *gin.Context) *errors.ValidationError {
	request := NewRequestValidation()
	request.AddParam("page", &b.Page, validation.Required, validation.Min(1))
	request.SetCustomCode(constants.QueryParamDataInvalidCode)
	return request.Validate(ctx)
}
