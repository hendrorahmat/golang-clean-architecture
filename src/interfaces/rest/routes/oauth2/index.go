package oauth2

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/oauth2/handler"
)

func RouteOauth2Client(routeGroup *gin.RouterGroup, handler oauth2_handler.Oauth2Handler) {
	routeGroup.GET("/client", handler.OauthClientHandler.Index)
	routeGroup.POST("/client", handler.OauthClientHandler.Store)
	routeGroup.POST("/token", handler.OauthTokenHandler.Store)
}
