package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/middleware"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/oauth2"
)

func NewRoute(
	ctx context.Context,
	handler *Handler,
	config *config.Config,
) *gin.Engine {

	if config.App.Environment == constants.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(func(ginContext *gin.Context) {
		ginContext.Request = ginContext.Request.WithContext(ctx)
		ginContext.Next()
	})
	router.Use(middleware.TimeoutHandler(config.Http))

	router.GET("/health", HealthGET)

	oauth2Group := router.Group("/oauth2")
	{
		oauth2.RouteOauth2Client(oauth2Group, handler.Oauth2Handler)
	}
	return router
}

func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":       "UP",
		"service_name": "sdf",
	})
}
