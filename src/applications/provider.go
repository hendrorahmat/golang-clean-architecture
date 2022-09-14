//go:build wireinject
// +build wireinject

package applications

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var BankUsecaseSet = wire.NewSet(wire.Struct(new(usecases.BankUsecase), "*"))

var (
	ProviderUsecaseSet wire.ProviderSet = wire.NewSet(
		databases.ProviderRepositorySet,
		BankUsecaseSet,
		wire.Struct(new(usecases.Usecase), "*"),
		wire.Bind(new(usecases.IBankUsecase), new(*usecases.BankUsecase)),
	)
)

func InjectUsecase(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *usecases.Usecase {
	panic(wire.Build(ProviderUsecaseSet))
}
