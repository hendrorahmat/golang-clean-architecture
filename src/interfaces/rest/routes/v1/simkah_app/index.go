package simkah_app

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/v1/simkah_app/handler"
)

func RouteSimkahAppV1(routeGroup *gin.RouterGroup, usecase usecases.UsecaseList) {
	bankController := handler.NewBankHandler(usecase.BankUsecase)

	routeGroup.GET("/banks", bankController.Index)
}