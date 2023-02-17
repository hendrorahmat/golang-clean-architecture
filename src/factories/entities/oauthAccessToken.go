package entities

import (
	"github.com/gofrs/uuid"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/dto"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"time"
)

func OauthAccessTokenFactory(token *dto.IssueToken) (*entities.OauthAccessToken, domainErrors.DomainError) {
	switch token.GrantType {
	case constants.ClientCredentialsGrantType:
		oauthAccessTokenEntity, err := entities.CreateGrantTypeClientCredentials(
			uuid.FromStringOrNil(token.ClientId),
			token.Scope,
			time.Now().Add(config.LoginExpirationDuration),
		)
		if err != nil {
			return nil, err
		}
		return oauthAccessTokenEntity, nil
	default:
		return nil, domainErrors.ThrowRecordFieldsNotFound("grant_type")
	}
}
