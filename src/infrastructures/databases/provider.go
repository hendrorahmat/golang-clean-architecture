//go:build wireinject
// +build wireinject

package databases

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories/gorm_types"
	repositories2 "github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

var repoOnce sync.Once

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

var GormBankRepositorySet = wire.NewSet(wire.Struct(new(repositories2.GormBankRepository), "*"))

var (
	ProviderRepositorySet wire.ProviderSet = wire.NewSet(
		ProvideDatabaseRepository,
		GormBankRepositorySet,
		wire.Struct(new(Repository), "*"),
		wire.Bind(new(gorm_types.TransactionRepository), new(*gormRepository)),
		wire.Bind(new(gorm_types.Repository), new(*gormRepository)),
		wire.Bind(new(repositories.IBankRepository), new(*repositories2.GormBankRepository)),
	)
)

func InjectRepository(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *Repository {
	panic(wire.Build(ProviderRepositorySet))
}
