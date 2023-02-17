package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	domainErrorsOauth "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors/oauth"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/models"
)

type OauthClientRepository struct {
	repositories.ITransactionRepository
}

func (db *OauthClientRepository) FindByClientIdAndClientSecret(
	ctx context.Context,
	clientId,
	clientSecret string,
) (*entities.OauthClient, domainErrors.DomainError) {
	var oauthClientModel = new(models.OauthClient)
	err := db.FindOneByFields(ctx, oauthClientModel, map[string]interface{}{
		"id":     clientId,
		"secret": clientSecret,
	})

	if err != nil {
		switch err.(type) {
		case *domainErrors.RecordNotFoundError:
			return nil, domainErrorsOauth.ThrowClientIdAndSecretNotFound(clientId, clientSecret)
		default:
			return nil, err
		}
	}
	entity, err := oauthClientModel.ToEntity()
	if err != nil {
		return nil, err
	}
	return entity, nil
}

var _ repositories.IOauthClientRepository = &OauthClientRepository{}
