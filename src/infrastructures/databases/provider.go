//go:build wireinject
// +build wireinject

package databases

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories/gorm_types"
	repositories_postgres "github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/postgres/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	BankRepository        repositories.IBankRepository
	TransactionRepository gorm_types.TransactionRepository
	BaseRepository        gorm_types.Repository
}

func ProvideDatabaseRepository(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *gormRepository {
	return &gormRepository{
		defaultJoins: defaultJoins,
		logger:       logger,
		db:           db,
	}
}

//var GormBankRepositoryPostgresSet = wire.NewSet(new(*repositories2.GormBankRepository))
var GormBankRepositorySet = wire.NewSet(wire.Struct(new(repositories_postgres.GormBankRepository), "*"))

var (
	ProviderRepositorySet wire.ProviderSet = wire.NewSet(
		ProvideDatabaseRepository,
		//GormBankRepositoryPostgresSet,
		GormBankRepositorySet,
		wire.Struct(new(Repository), "*"),
		wire.Bind(new(gorm_types.TransactionRepository), new(*gormRepository)),
		wire.Bind(new(gorm_types.Repository), new(*gormRepository)),
		wire.Bind(new(repositories.IBankRepository), new(*repositories_postgres.GormBankRepository)),
		//wire.Bind(new(repositories.IBankRepository), new(*repositories3.GormBankRepository)),
	)
)

func InjectRepository(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *Repository {
	panic(wire.Build(ProviderRepositorySet))
}
