package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/middleware"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/v1/simkah_app"
)

func NewRoute(
	handler *Handler,
	config *config.Config,
) *gin.Engine {

	if config.App.Environment == constants.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/health", HealthGET)
	router.Use(middleware.TimeoutHandler(config.Http))

	simkahApp := router.Group("/v1/simkah-app")
	{
		simkah_app.RouteSimkahAppV1(simkahApp, handler.BankHandler)
	}
	return router
}

func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":       "UP",
		"service_name": "sdf",
	})
}
