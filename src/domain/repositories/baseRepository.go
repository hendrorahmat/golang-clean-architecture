package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"gorm.io/gorm"
)

// IRepository is a generic DB handler that cares about default error handling
type IRepository interface {
	FindAll(ctx context.Context, target interface{}, preloads ...string) errors.DomainError
	FindBatch(ctx context.Context, target interface{}, limit, offset int, preloads ...string) errors.DomainError

	FindWhere(ctx context.Context, target interface{}, condition string, preloads ...string) errors.DomainError
	FindWhereBatch(ctx context.Context, target interface{}, condition string, limit, offset int, preloads ...string) errors.DomainError

	FindByField(ctx context.Context, target interface{}, field string, value interface{}, preloads ...string) errors.DomainError
	FindByFields(ctx context.Context, target interface{}, filters map[string]interface{}, preloads ...string) errors.DomainError
	FindByFieldBatch(ctx context.Context, target interface{}, field string, value interface{}, limit, offset int, preloads ...string) errors.DomainError
	FindByFieldsBatch(ctx context.Context, target interface{}, filters map[string]interface{}, limit, offset int, preloads ...string) errors.DomainError

	FindOneByField(ctx context.Context, target interface{}, field string, value interface{}, preloads ...string) errors.DomainError
	FindOneByFields(ctx context.Context, target interface{}, filters map[string]interface{}, preloads ...string) errors.DomainError

	// FindOneByID assumes you have a PK column "id" which is a UUID. If this is not the case just ignore the method
	// and add a custom struct with this IRepository embedded.
	FindOneByID(ctx context.Context, target interface{}, id string, preloads ...string) errors.DomainError

	Create(ctx context.Context, target interface{}) errors.DomainError
	Save(ctx context.Context, target interface{}) errors.DomainError
	Delete(ctx context.Context, target interface{}) errors.DomainError

	DBGorm() *gorm.DB
	DBWithPreloads(preloads []string) *gorm.DB
	HandleQueryError(res *gorm.DB) errors.DomainError
	HandleCommandError(res *gorm.DB) errors.DomainError
	HandleOneError(res *gorm.DB) errors.DomainError
}

// ITransactionRepository extends IRepository with modifier functions that accept a transaction
type ITransactionRepository interface {
	IRepository
	CreateTx(ctx context.Context, target interface{}, tx *gorm.DB) errors.DomainError
	SaveTx(ctx context.Context, target interface{}, tx *gorm.DB) errors.DomainError
	DeleteTx(ctx context.Context, target interface{}, tx *gorm.DB) errors.DomainError
}
