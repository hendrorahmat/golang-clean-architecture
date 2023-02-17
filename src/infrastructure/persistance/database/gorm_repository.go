package database

import (
	"context"
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (r *gormRepository) DBGorm() *gorm.DB {
	return r.DBWithPreloads(nil)
}

type gormRepository struct {
	logger       *log.Logger
	db           *gorm.DB
	defaultJoins []string
}

func (r *gormRepository) FindAll(ctx context.Context, target interface{}, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetAll on %T", target)
	//res := r.db.WithContext(ctx).Find(target)
	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Unscoped().
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindBatch(ctx context.Context, target interface{}, limit, offset int, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetBatch on %T", target)

	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Unscoped().
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindWhere(ctx context.Context, target interface{}, condition string, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetWhere on %T with %v ", target, condition)

	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Where(condition).
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindWhereBatch(ctx context.Context, target interface{}, condition string, limit, offset int, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetWhere on %T with %v ", target, condition)

	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Where(condition).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindByField(ctx context.Context, target interface{}, field string, value interface{}, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetByField on %T with %v = %v", target, field, value)

	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Where(fmt.Sprintf("%v = ?", field), value).
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindByFields(ctx context.Context, target interface{}, filters map[string]interface{}, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetByField on %T with filters = %+v", target, filters)

	db := r.DBWithPreloads(preloads).WithContext(ctx)
	for field, value := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := db.Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindByFieldBatch(ctx context.Context, target interface{}, field string, value interface{}, limit, offset int, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetByField on %T with %v = %v", target, field, value)

	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Where(fmt.Sprintf("%v = ?", field), value).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindByFieldsBatch(ctx context.Context, target interface{}, filters map[string]interface{}, limit, offset int, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetByField on %T with filters = %+v", target, filters)

	db := r.DBWithPreloads(preloads).WithContext(ctx)
	for field, value := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := db.
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleQueryError(res)
}

func (r *gormRepository) FindOneByField(ctx context.Context, target interface{}, field string, value interface{}, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetOneByField on %T with %v = %v", target, field, value)

	res := r.DBWithPreloads(preloads).
		WithContext(ctx).
		Where(fmt.Sprintf("%v = ?", field), value).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneByFields(ctx context.Context, target interface{}, filters map[string]interface{}, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing FindOneByField on %T with filters = %+v", target, filters)

	db := r.DBWithPreloads(preloads).WithContext(ctx)
	for field, value := range filters {
		db = db.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := db.First(target)
	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneByID(ctx context.Context, target interface{}, id string, preloads ...string) errors.DomainError {
	r.logger.Debugf("Executing GetOneByID on %T with id %v", target, id)

	res := r.DBWithPreloads(preloads).WithContext(ctx).
		Where("id = ?", id).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) Create(ctx context.Context, target interface{}) errors.DomainError {
	r.logger.Debugf("Executing Create on %T", target)

	res := r.db.WithContext(ctx).Create(target)
	return r.HandleCommandError(res)
}

func (r *gormRepository) CreateTx(ctx context.Context, target interface{}, tx *gorm.DB) errors.DomainError {
	r.logger.Debugf("Executing Create on %T", target)

	res := tx.WithContext(ctx).Create(target)
	return r.HandleCommandError(res)
}

func (r *gormRepository) Save(ctx context.Context, target interface{}) errors.DomainError {
	r.logger.Debugf("Executing Save on %T", target)

	res := r.db.WithContext(ctx).Save(target)
	return r.HandleCommandError(res)
}

func (r *gormRepository) SaveTx(ctx context.Context, target interface{}, tx *gorm.DB) errors.DomainError {
	r.logger.Debugf("Executing Save on %T", target)

	res := tx.WithContext(ctx).Save(target)
	return r.HandleCommandError(res)
}

func (r *gormRepository) Delete(ctx context.Context, target interface{}) errors.DomainError {
	r.logger.Debugf("Executing Delete on %T", target)

	res := r.db.WithContext(ctx).Delete(target)
	return r.HandleCommandError(res)
}

func (r *gormRepository) DeleteTx(ctx context.Context, target interface{}, tx *gorm.DB) errors.DomainError {
	r.logger.Debugf("Executing Delete on %T", target)

	res := tx.WithContext(ctx).Delete(target)
	return r.HandleCommandError(res)
}

func (r *gormRepository) HandleQueryError(res *gorm.DB) errors.DomainError {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("error: %w", res.Error)
		r.logger.Error(err)
		return errors.ThrowDatabaseQueryError(err.Error())
	}
	return nil
}

func (r *gormRepository) HandleCommandError(res *gorm.DB) errors.DomainError {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("error: %w", res.Error)
		r.logger.Error(err)
		return errors.ThrowDatabaseQueryError(err.Error())
	}
	return nil
}

func (r *gormRepository) HandleOneError(res *gorm.DB) errors.DomainError {
	if err := r.HandleQueryError(res); err != nil {
		return err
	}

	if res.RowsAffected < 1 {
		return errors.ThrowRecordNotFoundError()
	}

	return nil
}

func (r *gormRepository) DBWithPreloads(preloads []string) *gorm.DB {
	dbConn := r.db

	for _, join := range r.defaultJoins {
		dbConn = dbConn.Joins(join)
	}

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}
