package gorm_types

import (
	"context"
	"gorm.io/gorm"
)

// Repository is a generic DB handler that cares about default error handling
type Repository interface {
	FindAll(ctx context.Context, target interface{}, preloads ...string) error
	FindBatch(ctx *context.Context, target interface{}, limit, offset int, preloads ...string) error

	FindWhere(ctx context.Context, target interface{}, condition string, preloads ...string) error
	FindWhereBatch(ctx context.Context, target interface{}, condition string, limit, offset int, preloads ...string) error

	FindByField(ctx context.Context, target interface{}, field string, value interface{}, preloads ...string) error
	FindByFields(ctx context.Context, target interface{}, filters map[string]interface{}, preloads ...string) error
	FindByFieldBatch(ctx context.Context, target interface{}, field string, value interface{}, limit, offset int, preloads ...string) error
	FindByFieldsBatch(ctx context.Context, target interface{}, filters map[string]interface{}, limit, offset int, preloads ...string) error

	FindOneByField(ctx context.Context, target interface{}, field string, value interface{}, preloads ...string) error
	FindOneByFields(ctx context.Context, target interface{}, filters map[string]interface{}, preloads ...string) error

	// FindOneByID assumes you have a PK column "id" which is a UUID. If this is not the case just ignore the method
	// and add a custom struct with this Repository embedded.
	FindOneByID(ctx context.Context, target interface{}, id string, preloads ...string) error

	Create(ctx context.Context, target interface{}) error
	Save(ctx context.Context, target interface{}) error
	Delete(ctx context.Context, target interface{}) error

	DB() *gorm.DB
	DBWithPreloads(preloads []string) *gorm.DB
	HandleError(res *gorm.DB) error
	HandleOneError(res *gorm.DB) error
}

// TransactionRepository extends Repository with modifier functions that accept a transaction
type TransactionRepository interface {
	Repository
	CreateTx(ctx context.Context, target interface{}, tx *gorm.DB) error
	SaveTx(ctx context.Context, target interface{}, tx *gorm.DB) error
	DeleteTx(ctx context.Context, target interface{}, tx *gorm.DB) error
}
