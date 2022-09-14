package simkah_app

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/v1/simkah_app/handler"
)

func RouteSimkahAppV1(routeGroup *gin.RouterGroup, handler handler.IBankHandler) {
	routeGroup.GET("/banks", handler.Index)
}
