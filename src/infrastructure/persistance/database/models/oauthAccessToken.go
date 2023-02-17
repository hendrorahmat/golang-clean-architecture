package models

import (
	"github.com/gofrs/uuid"
	domainModelEntities "github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type OauthAccessToken struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ClientId  uuid.UUID      `json:"client_id" gorm:"type:uuid;primaryKey;column:client_id"`
	UserId    *uuid.UUID     `json:"user_id" gorm:"type:uuid;column:user_id"`
	GrantType string         `json:"grant_type" gorm:"column:grant_type"`
	Scopes    pq.StringArray `json:"scopes" gorm:"type:varchar(100)[];column:scopes"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	ExpiresAt time.Time      `json:"expires_at" gorm:"column:expires_at"`
	DeletedAt *time.Time     `json:"revoked_at" gorm:"column:deleted_at"`
}

func (o *OauthAccessToken) BeforeCreate(db *gorm.DB) (err error) {
	for {
		// UUID version 4
		o.ID, _ = uuid.NewV4()
		var model OauthAccessToken
		db.Raw("SELECT id where id = ?", o.ID).Scan(&model)
		if reflect.ValueOf(model).IsZero() {
			break
		}
	}
	return
}

func CreateModelFromEntityOauthAccessToken(entity domainModelEntities.OauthAccessToken) *OauthAccessToken {
	oauthAccessTokenModel := &OauthAccessToken{}
	oauthAccessTokenModel.ClientId = entity.ClientId()
	oauthAccessTokenModel.UserId = entity.UserId()
	oauthAccessTokenModel.Scopes = entity.Scopes()
	oauthAccessTokenModel.GrantType = entity.GrantType()
	oauthAccessTokenModel.DeletedAt = entity.RevokedAt()
	oauthAccessTokenModel.CreatedAt = entity.CreatedAt()
	oauthAccessTokenModel.UpdatedAt = entity.UpdatedAt()
	oauthAccessTokenModel.ExpiresAt = entity.ExpiresAt()
	return oauthAccessTokenModel
}

func (o *OauthAccessToken) ToEntity() (*domainModelEntities.OauthAccessToken, errors.DomainError) {
	entity, err := domainModelEntities.NewOauthAccessTokenEntity(
		o.ID,
		o.UserId,
		o.GrantType,
		o.ClientId,
		o.Scopes,
		o.ExpiresAt,
		o.CreatedAt,
		o.UpdatedAt,
		o.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

//func (oauth *OauthAccessToken) BeforeCreate(tx *gorm.DBGorm) (err error) {
//	field := tx.Statement.Schema.LookUpField("ID")
//	fmt.Println(field)
//}
