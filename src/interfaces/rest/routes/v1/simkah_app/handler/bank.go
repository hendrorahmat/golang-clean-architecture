package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
)

type IBankHandler interface {
	Index(ctx *gin.Context)
}

type bankHandler struct {
	Usecase usecases.IBankUsecase
}

func (b bankHandler) Index(ctx *gin.Context) {
	b.Usecase.GetListBank(ctx.Request.Context())
	//ctx.Header("accept", "application/json")
	ctx.JSON(200, gin.H{
		"status":  "posted",
		"message": "message",
		"nick":    "nick",
	})
}

func NewBankHandler(usecase usecases.IBankUsecase) IBankHandler {
	return bankHandler{
		Usecase: usecase,
	}
}
