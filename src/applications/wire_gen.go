// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package applications

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/postgres/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Injectors from provider.go:

func InjectUsecase(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *usecases.Usecase {
	gormRepository := databases.ProvideDatabaseRepository(db, logger, defaultJoins...)
	gormBankRepository := &repositories_postgres.GormBankRepository{
		TransactionRepository: gormRepository,
	}
	bankUsecase := &usecases.BankUsecase{
		Repository: gormBankRepository,
	}
	usecase := &usecases.Usecase{
		BankUsecase: bankUsecase,
	}
	return usecase
}

// provider.go:

var BankUsecaseSet = wire.NewSet(wire.Struct(new(usecases.BankUsecase), "*"))

var (
	ProviderUsecaseSet wire.ProviderSet = wire.NewSet(databases.ProviderRepositorySet, BankUsecaseSet, wire.Struct(new(usecases.Usecase), "*"), wire.Bind(new(usecases.IBankUsecase), new(*usecases.BankUsecase)))
)
