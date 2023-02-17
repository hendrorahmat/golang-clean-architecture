package entities

import (
	"github.com/gofrs/uuid"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"time"
)

type OauthClient struct {
	ID               uuid.UUID  `json:"id"`
	Name             string     `json:"name"`
	EnabledGrantType []string   `json:"enabled_grant_type"`
	Secret           string     `json:"secret"`
	Redirect         string     `json:"redirect"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

func MakeOauthClientEntity(
	ID uuid.UUID,
	name string,
	enabledGrantType []string,
	secret string,
	redirect string,
	createdAt *time.Time,
	updatedAt *time.Time,
	deletedAt *time.Time,
) (*OauthClient, errors.DomainError) {
	return &OauthClient{
		ID:               ID,
		Name:             name,
		EnabledGrantType: enabledGrantType,
		Secret:           secret,
		Redirect:         redirect,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        deletedAt,
	}, nil
}
