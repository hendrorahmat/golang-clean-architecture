package oauth

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/commands/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
)

type CreateTokenAuthCodeService struct {
	command *oauth.CreateTokenAuthCodeCommand
}

func NewCreateTokenAuthCodeService(
	command *oauth.CreateTokenClientCredentialCommand,
	oauthClientRepository repositories.IOauthClientRepository,
	oauthAccessTokenRepository repositories.IOauthAccessTokenRepository,
) ICreateTokenService {
	return nil
}
