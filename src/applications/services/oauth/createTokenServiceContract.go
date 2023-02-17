package oauth

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
)

type ICreateTokenService interface {
	Handle(ctx context.Context) (*entities.Token, domainErrors.DomainError)
}
