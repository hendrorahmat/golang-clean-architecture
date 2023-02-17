//go:build wireinject
// +build wireinject

package database

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	repositories_postgres2 "github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/postgres/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	TransactionRepository      repositories.ITransactionRepository
	BaseRepository             repositories.IRepository
	BankRepository             repositories.IBankRepository
	OauthClientRepository      repositories.IOauthClientRepository
	OauthAccessTokenRepository repositories.IOauthAccessTokenRepository
}

func ProvideDatabaseGorm(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *gormRepository {
	return &gormRepository{
		defaultJoins: defaultJoins,
		logger:       logger,
		db:           db,
	}
}

// var GormBankRepositoryPostgresSet = wire.NewSet(new(*repositories2.PostgresBankRepository))
var GormBankRepositorySet = wire.NewSet(wire.Struct(new(repositories_postgres2.PostgresBankRepository), "*"))
var GormOauthClientRepositorySet = wire.NewSet(wire.Struct(new(repositories_postgres2.OauthClientRepository), "*"))
var GormOauthAccessTokenRepositorySet = wire.NewSet(wire.Struct(new(repositories_postgres2.OauthAccessTokenRepository), "*"))

var (
	ProviderRepositorySet wire.ProviderSet = wire.NewSet(
		ProvideDatabaseGorm,
		//GormBankRepositoryPostgresSet,
		GormBankRepositorySet,
		GormOauthClientRepositorySet,
		GormOauthAccessTokenRepositorySet,
		wire.Struct(new(Repository), "*"),
		wire.Bind(new(repositories.ITransactionRepository), new(*gormRepository)),
		wire.Bind(new(repositories.IRepository), new(*gormRepository)),
		wire.Bind(new(repositories.IBankRepository), new(*repositories_postgres2.PostgresBankRepository)),
		wire.Bind(new(repositories.IOauthClientRepository), new(*repositories_postgres2.OauthClientRepository)),
		wire.Bind(new(repositories.IOauthAccessTokenRepository), new(*repositories_postgres2.OauthAccessTokenRepository)),
		//wire.Bind(new(repositories.IBankRepository), new(*repositories3.PostgresBankRepository)),
	)
)

func InjectRepository(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *Repository {
	panic(wire.Build(ProviderRepositorySet))
}
