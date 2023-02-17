package oauth2_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/contracts"
	"github.com/sirupsen/logrus"
)

type IOauthClientHandler interface {
	contracts.IResourcesHandler
}

type OauthClientHandler struct {
	Usecase usecases.IOauthUsecase
	Logger  *logrus.Logger
}

func (o *OauthClientHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o *OauthClientHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o *OauthClientHandler) Show(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o *OauthClientHandler) Index(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o *OauthClientHandler) Store(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

var _ IOauthClientHandler = &OauthClientHandler{}
