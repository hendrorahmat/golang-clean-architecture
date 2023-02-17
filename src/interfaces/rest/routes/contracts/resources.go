package contracts

import "github.com/gin-gonic/gin"

type IResourcesHandler interface {
	Index(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Show(ctx *gin.Context)
}
