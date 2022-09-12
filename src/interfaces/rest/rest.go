package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/middleware"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/v1/simkah_app"
)

func NewRoute(
	config *config.Config,
	usecase usecases.UsecaseList,
) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.TimeoutHandler(config.Http))

	api := router.Group("/api")
	{
		simkahApp := api.Group("/v1/simkah-app")
		{
			simkah_app.RouteSimkahAppV1(simkahApp, usecase)
		}
	}
	return router
}
