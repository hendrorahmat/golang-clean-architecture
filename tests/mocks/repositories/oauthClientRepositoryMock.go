package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/stretchr/testify/mock"
)

type OauthClientRepositoryMock struct {
	mock.Mock
	repositories.IOauthClientRepository
}

func (o *OauthClientRepositoryMock) FindByClientIdAndClientSecret(ctx context.Context, clientId, clientSecret string) (*entities.OauthClient, errors.DomainError) {
	args := o.Called(ctx, clientId, clientSecret)
	var expectedError errors.DomainError
	var expectedData *entities.OauthClient
	if args[1] == nil {
		expectedError = nil
	} else {
		expectedError = args[1].(errors.DomainError)
	}

	if args[0] == nil {
		expectedData = nil
	} else {
		expectedData = args[0].(*entities.OauthClient)
	}
	return expectedData, expectedError
}
