package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/dto"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	factoryCommand "github.com/hendrorahmat/golang-clean-architecture/src/factories/commands"
	"github.com/hendrorahmat/golang-clean-architecture/src/factories/services"
	"github.com/sirupsen/logrus"
)

type IOauthUsecase interface {
	IssueToken(ctx *gin.Context, token *dto.IssueToken) (*entities.Token, domainErrors.DomainError)
}

type OauthUsecase struct {
	OauthClientRepository      repositories.IOauthClientRepository
	OauthAccessTokenRepository repositories.IOauthAccessTokenRepository
	Logger                     *logrus.Logger
}

func (o *OauthUsecase) IssueToken(ctx *gin.Context, dtoIssueToken *dto.IssueToken) (*entities.Token, domainErrors.DomainError) {
	commandFactory, err := factoryCommand.CreateTokenCommandFactory(dtoIssueToken)
	if err != nil {
		return nil, err
	}

	service, err := services.CreateTokenServiceFactory(commandFactory, o.OauthClientRepository, o.OauthAccessTokenRepository)
	if err != nil {
		return nil, err
	}

	tokenEntity, err := service.Handle(ctx)
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
}

var _ IOauthUsecase = &OauthUsecase{}
