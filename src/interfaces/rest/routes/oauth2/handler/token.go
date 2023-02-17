package oauth2_handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/dto"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/form_request"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/contracts"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type IOauthTokenHandler interface {
	contracts.IResourcesHandler
}

type OauthTokenHandler struct {
	Usecase usecases.IOauthUsecase
	Logger  *logrus.Logger
}

func (o OauthTokenHandler) Index(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OauthTokenHandler) Store(ctx *gin.Context) {
	var request form_request.Oauth2Request
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	errV := request.Validate(ctx)
	if errV != nil {
		pd := errors.NewProblemDetails(422, errV, "", "")
		messages, _ := json.Marshal(pd)
		o.Logger.Error(string(messages))
		ctx.JSON(422, pd)
		return
	}
	scopes := strings.Split(request.Scope, ",")
	issueTokenDto := dto.NewIssueToken(request.GrantType, request.ClientId, request.ClientSecret, scopes)

	tokenEntity, errUsecase := o.Usecase.IssueToken(ctx, issueTokenDto)
	if errUsecase != nil {
		pd := errors.NewProblemDetails(errUsecase.GetStatusCode(), errUsecase, "", "")
		messages, _ := json.Marshal(pd)
		o.Logger.Warnf(string(messages))
		ctx.JSON(errUsecase.GetStatusCode(), pd)
		return
	}
	ctx.Header("cache-control", "no-store, private")
	ctx.Header("connection", "keep-alive")
	ctx.JSON(http.StatusOK, tokenEntity)
}

func (o OauthTokenHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OauthTokenHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OauthTokenHandler) Show(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

var _ IOauthClientHandler = OauthTokenHandler{}
