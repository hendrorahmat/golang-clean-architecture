package oauth

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/commands/oauth"
	oauthService "github.com/hendrorahmat/golang-clean-architecture/src/applications/services/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	domainErrorsOauth "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/models"
	repositoriesMock "github.com/hendrorahmat/golang-clean-architecture/tests/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	command := &oauth.CreateTokenClientCredentialCommand{}
	oauthClientRepository := new(repositoriesMock.OauthClientRepositoryMock)
	oauthAccessTokenRepository := new(repositoriesMock.OauthAccessTokenMock)
	entityOauthClient := &entities.OauthClient{}
	oauthAccessTokenModel := &models.OauthAccessToken{}

	gofakeit.Struct(command)
	gofakeit.Struct(entityOauthClient)
	gofakeit.Struct(oauthAccessTokenModel)

	oauthClientRepository.
		On("FindByClientIdAndClientSecret", context.Background(), command.ClientId, command.ClientSecret).
		Return(
			entityOauthClient,
			nil,
		)
	oauthAccessTokenRepository.
		On("Create", context.Background(), mock.AnythingOfType("*models.OauthAccessToken")).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.OauthAccessToken)
			arg.ID = oauthAccessTokenModel.ID
			arg.CreatedAt = oauthAccessTokenModel.CreatedAt
			arg.UpdatedAt = oauthAccessTokenModel.UpdatedAt
			arg.UserId = oauthAccessTokenModel.UserId
			arg.GrantType = oauthAccessTokenModel.GrantType
			arg.Scopes = oauthAccessTokenModel.Scopes
			arg.ExpiresAt = oauthAccessTokenModel.ExpiresAt
			arg.DeletedAt = oauthAccessTokenModel.DeletedAt
			arg.ClientId = oauthAccessTokenModel.ClientId
		}).
		Return(nil)
	service := oauthService.NewCreateTokenClientCredentialService(command, oauthClientRepository, oauthAccessTokenRepository)
	token, domainError := service.Handle(context.Background())
	oneDay := time.Now().UTC().Add(24 * time.Hour).Format(constants.SQLTimestampFormat)

	assert.Nil(t, domainError)
	assert.NotEmpty(t, token)
	assert.Equal(t, token.ExpiresIn().Format(constants.SQLTimestampFormat), oneDay)
	assert.Equal(t, token.TokenType(), constants.TokenTypeBearer)
	oauthClientRepository.AssertExpectations(t)
	oauthAccessTokenRepository.AssertExpectations(t)
}

func TestShouldReturnDomainErrorWhenClientIdAndSecreteNotFound(t *testing.T) {
	command := &oauth.CreateTokenClientCredentialCommand{}
	oauthClientRepository := new(repositoriesMock.OauthClientRepositoryMock)
	entityOauthClient := &entities.OauthClient{}
	oauthAccessTokenRepository := new(repositoriesMock.OauthAccessTokenMock)

	gofakeit.Struct(command)
	gofakeit.Struct(entityOauthClient)

	expectedError := domainErrorsOauth.ThrowClientIdAndSecretNotFound(command.ClientId, command.ClientSecret)
	oauthClientRepository.
		On("FindByClientIdAndClientSecret", context.Background(), command.ClientId, command.ClientSecret).
		Return(
			nil,
			expectedError,
		)
	service := oauthService.NewCreateTokenClientCredentialService(command, oauthClientRepository, oauthAccessTokenRepository)
	data, domainError := service.Handle(context.Background())

	assert.Equal(t, expectedError, domainError)
	assert.Nil(t, data)
	oauthClientRepository.AssertExpectations(t)
}
