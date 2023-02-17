package repositories

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
)

type OauthAccessTokenRepository struct {
	repositories.ITransactionRepository
}

var _ repositories.IOauthAccessTokenRepository = &OauthAccessTokenRepository{}
