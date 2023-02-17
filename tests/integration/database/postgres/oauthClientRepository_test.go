package postgres

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/tests/integration/factories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type OauthClientRepository struct {
	db *gorm.DB
	suite.Suite
}

func (p *PostgresDatabaseSuiteTest) TestOauthClientRepository_FindByClientIdAndClientSecret() {
	ctx := context.Background()
	model := factories.NewOauthClientFactory()
	repository := p.app.GetRepository().OauthClientRepository
	//gormDb := database.ProvideDatabaseGorm(p.db.DB(), p.app.GetLogger())

	err := repository.Create(ctx, &model)
	p.Assert().NoError(err)
	data, err := repository.FindByClientIdAndClientSecret(ctx, model.ID.String(), model.Secret)
	p.Assert().NoError(err)

	p.Assert().Equal(model.ID, data.ID)
}

func (p *PostgresDatabaseSuiteTest) TestOauthClientRepository_FindByClientIdAndClientSecret_ShouldReturnDomainNotFound() {
	repository := p.app.GetRepository().OauthClientRepository
	ctx := context.Background()
	clientId, clientSecret := gofakeit.UUID(), gofakeit.LetterN(40)
	data, err := repository.FindByClientIdAndClientSecret(ctx, clientId, clientSecret)
	assert.Nil(p.T(), data)
	assert.Equal(p.T(), err, oauth.ThrowClientIdAndSecretNotFound(clientId, clientSecret))
}
