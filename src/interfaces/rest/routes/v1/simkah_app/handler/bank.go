package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/form_request"
	"github.com/sirupsen/logrus"
)

type IBankHandler interface {
	Index(ctx *gin.Context)
}

type BankHandler struct {
	Usecase usecases.IBankUsecase
	Logger  *logrus.Logger
}

func (b *BankHandler) Index(ctx *gin.Context) {
	var bankRequest form_request.BankRequest
	ctx.ShouldBind(&bankRequest)
	ctx.Handler()
	errV := bankRequest.Validate(ctx)

	if errV != nil {
		messages, _ := json.Marshal(errV)

		b.Logger.Warnf(string(messages))
		ctx.JSON(422, errors.NewProblemDetails(422, errV, "", ""))
		return
	}

	listBank, err := b.Usecase.GetListBank(ctx.Request.Context())
	if err != nil {
		return
	}

	resp := utils.NewResponse[[]entities.Bank, string](listBank)
	resp.MakeMetaData("halo", nil, nil, nil, nil)
	ctx.JSON(200, resp)
}
