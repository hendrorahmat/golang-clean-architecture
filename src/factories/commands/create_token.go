package commands

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/commands/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/dto"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
)

func CreateTokenCommandFactory(dto *dto.IssueToken) (oauth.IIssueTokenCommand, domainErrors.DomainError) {
	switch dto.GrantType {
	case constants.ClientCredentialsGrantType:
		return oauth.NewCreateTokenClientCredentialCommand(dto.ClientId, dto.ClientSecret, dto.Scope), nil
	default:
		return nil, domainErrors.ThrowInvalidData("Grant type invalid.")
	}
}
