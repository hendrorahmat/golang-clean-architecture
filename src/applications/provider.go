//go:build wireinject
// +build wireinject

package applications

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database"
	"github.com/sirupsen/logrus"
)

//var OauthUsecaseSet = wire.NewSet(wire.Struct(new(usecases.OauthUsecase), "*"))

func ProvideOauthUsecase(repository *database.Repository, logger *logrus.Logger) *usecases.OauthUsecase {
	return &usecases.OauthUsecase{
		OauthClientRepository:      repository.OauthClientRepository,
		OauthAccessTokenRepository: repository.OauthAccessTokenRepository,
		Logger:                     logger,
	}
}

var (
	ProviderUsecaseSet wire.ProviderSet = wire.NewSet(
		ProvideOauthUsecase,
		wire.Struct(new(usecases.Usecase), "*"),
		wire.Bind(new(usecases.IOauthUsecase), new(*usecases.OauthUsecase)),
	)
)

func InjectUsecase(repository *database.Repository, logger *logrus.Logger, defaultJoins ...string) *usecases.Usecase {
	panic(wire.Build(ProviderUsecaseSet))
}
