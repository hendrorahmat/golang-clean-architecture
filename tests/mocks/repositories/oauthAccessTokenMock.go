package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/stretchr/testify/mock"
)

type OauthAccessTokenMock struct {
	mock.Mock
	repositories.ITransactionRepository
}

func (mockRepository *OauthAccessTokenMock) Create(ctx context.Context, oauthAccessTokenModel interface{}) errors.DomainError {
	args := mockRepository.Called(ctx, oauthAccessTokenModel)
	if args[0] == nil {
		return nil
	}
	return args[0].(errors.DomainError)
}
