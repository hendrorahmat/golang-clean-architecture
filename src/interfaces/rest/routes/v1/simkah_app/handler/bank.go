package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
)

type IBankHandler interface {
	Index(ctx *gin.Context)
}

type BankHandler struct {
	Usecase usecases.IBankUsecase
}

func (b *BankHandler) Index(ctx *gin.Context) {
	listBank, err := b.Usecase.GetListBank(ctx.Request.Context())
	if err != nil {
		return
	}
	ctx.JSON(200, listBank)
}
