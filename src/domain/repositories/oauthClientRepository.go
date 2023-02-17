package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	errors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
)

type IOauthClientRepository interface {
	ITransactionRepository
	FindByClientIdAndClientSecret(ctx context.Context, clientId, clientSecret string) (*entities.OauthClient, errors.DomainError)
}
