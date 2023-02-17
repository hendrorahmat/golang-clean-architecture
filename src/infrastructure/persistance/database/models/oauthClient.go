package models

import (
	"github.com/gofrs/uuid"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type OauthClient struct {
	ID               uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;column:id;default:uuid_generate_v4()"`
	Name             string         `json:"name" gorm:"column:name"`
	EnabledGrantType pq.StringArray `json:"enabled_grant_type" gorm:"type:varchar(100)[];column:enabled_grant_type"`
	Secret           string         `json:"secret" gorm:"column:secret"`
	Redirect         string         `json:"redirect" gorm:"column:redirect"`
	CreatedAt        *time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        *time.Time     `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (oauthClient *OauthClient) ToEntity() (*entities.OauthClient, errors.DomainError) {
	entity, err := entities.MakeOauthClientEntity(
		oauthClient.ID,
		oauthClient.Name,
		oauthClient.EnabledGrantType,
		oauthClient.Secret,
		oauthClient.Redirect,
		oauthClient.CreatedAt,
		oauthClient.UpdatedAt,
		&oauthClient.DeletedAt.Time,
	)
	return entity, err
}
