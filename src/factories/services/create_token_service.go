package services

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/commands/oauth"
	oauthService "github.com/hendrorahmat/golang-clean-architecture/src/applications/services/oauth"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
)

func CreateTokenServiceFactory(
	command oauth.IIssueTokenCommand,
	oauthClientRepository repositories.IOauthClientRepository,
	oauthAccessTokenRepository repositories.IOauthAccessTokenRepository,
) (oauthService.ICreateTokenService, domainErrors.DomainError) {
	switch commandType := command.(type) {
	case *oauth.CreateTokenClientCredentialCommand:
		service := oauthService.NewCreateTokenClientCredentialService(commandType, oauthClientRepository, oauthAccessTokenRepository)
		return service, nil
	default:
		return nil, domainErrors.ThrowInternalServerError(constants.ErrorCommandTypeNotFound)
	}
}
