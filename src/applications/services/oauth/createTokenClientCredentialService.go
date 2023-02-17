package oauth

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/commands/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/models"
	"time"
)

type CreateTokenClientCredentialService struct {
	command                    *oauth.CreateTokenClientCredentialCommand
	oauthClientRepository      repositories.IOauthClientRepository
	oauthAccessTokenRepository repositories.IOauthAccessTokenRepository
}

func (c *CreateTokenClientCredentialService) Handle(ctx context.Context) (*entities.Token, errors.DomainError) {
	_, err := c.oauthClientRepository.FindByClientIdAndClientSecret(ctx, c.command.ClientId, c.command.ClientSecret)
	if err != nil {
		return nil, err
	}
	var audiences []uuid.UUID
	audiences = append(audiences, uuid.FromStringOrNil(c.command.ClientId))

	oauthAccessTokenModel := new(models.OauthAccessToken)
	oauthAccessTokenEntity, err := entities.CreateGrantTypeClientCredentials(
		uuid.FromStringOrNil(c.command.ClientId),
		c.command.Scopes,
		time.Now().Add(config.LoginExpirationDuration),
	)

	oauthAccessTokenModel = models.CreateModelFromEntityOauthAccessToken(*oauthAccessTokenEntity)
	err = c.oauthAccessTokenRepository.Create(ctx, oauthAccessTokenModel)
	if err != nil {
		return nil, err
	}

	oauthAccessTokenEntity, err = oauthAccessTokenModel.ToEntity()
	if err != nil {
		return nil, err
	}

	var subject = ""
	if oauthAccessTokenEntity.UserId() != nil {
		subject = oauthAccessTokenEntity.UserId().String()
	}

	jwtClaim := entities.NewJwtClaim(oauthAccessTokenEntity.Id(), audiences, subject, oauthAccessTokenEntity.Scopes())

	token, err := jwtClaim.Generate()
	if err != nil {
		return nil, err
	}

	var tokenEntity *entities.Token
	tokenEntity, err = entities.NewToken(jwtClaim.GetJwtId(), constants.TokenTypeBearer, jwtClaim.GetExpiresAt(), *token, nil)
	if err != nil {
		return nil, err
	}
	return tokenEntity, nil
}

func NewCreateTokenClientCredentialService(
	command *oauth.CreateTokenClientCredentialCommand,
	oauthClientRepository repositories.IOauthClientRepository,
	oauthAccessTokenRepository repositories.IOauthAccessTokenRepository,
) ICreateTokenService {
	return &CreateTokenClientCredentialService{
		command:                    command,
		oauthAccessTokenRepository: oauthAccessTokenRepository,
		oauthClientRepository:      oauthClientRepository,
	}
}
