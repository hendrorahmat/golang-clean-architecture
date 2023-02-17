package entities

import (
	"github.com/gofrs/uuid"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"time"
)

type OauthAccessToken struct {
	id        uuid.UUID  `json:"id"`
	userId    *uuid.UUID `json:"user_id"`
	grantType string     `json:"grant_type"`
	clientId  uuid.UUID  `json:"client_id"`
	scopes    []string   `json:"scopes"`
	createdAt time.Time  `json:"created_at"`
	updatedAt time.Time  `json:"updated_at"`
	deletedAt *time.Time `json:"updated_at"`
	expiresAt time.Time  `json:"expires_at"`
}

func NewOauthAccessTokenEntity(
	id uuid.UUID,
	userId *uuid.UUID,
	grantType string,
	clientId uuid.UUID,
	scopes []string,
	expiresAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*OauthAccessToken, errors.DomainError) {
	if grantType == "" {
		return nil, errors.ThrowFieldsRequired("grant_type")
	}

	if clientId.IsNil() {
		return nil, errors.ThrowFieldsRequired("client_id")
	}

	if expiresAt.IsZero() {
		return nil, errors.ThrowFieldsRequired("expires_at")
	}

	return &OauthAccessToken{
		id:        id,
		userId:    userId,
		grantType: grantType,
		clientId:  clientId,
		scopes:    scopes,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
		expiresAt: expiresAt,
	}, nil
}

func (oauth *OauthAccessToken) ClientId() uuid.UUID {
	return oauth.clientId
}

func CreateOauthAccessToken(
	userId *uuid.UUID,
	clientId uuid.UUID,
	grantType string,
	scopes []string,
	revokedAt *time.Time,
	expiresAt time.Time,
) (*OauthAccessToken, errors.DomainError) {
	if grantType == "" {
		return nil, errors.ThrowFieldsRequired("grant_type")
	}

	if clientId.IsNil() {
		return nil, errors.ThrowFieldsRequired("client_id")
	}

	if expiresAt.IsZero() {
		return nil, errors.ThrowFieldsRequired("expires_at")
	}

	return &OauthAccessToken{
		userId:    userId,
		grantType: grantType,
		clientId:  clientId,
		scopes:    scopes,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: revokedAt,
		expiresAt: expiresAt,
	}, nil
}

func CreateGrantTypeClientCredentials(
	clientId uuid.UUID,
	scopes []string,
	expiresAt time.Time,
) (*OauthAccessToken, errors.DomainError) {
	data, err := CreateOauthAccessToken(nil, clientId, constants.ClientCredentialsGrantType, scopes, nil, expiresAt)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (oauth *OauthAccessToken) Id() *uuid.UUID {
	return &oauth.id
}

func (oauth *OauthAccessToken) UserId() *uuid.UUID {
	return oauth.userId
}

func (oauth *OauthAccessToken) GrantType() string {
	return oauth.grantType
}

func (oauth *OauthAccessToken) Scopes() []string {
	return oauth.scopes
}

func (oauth *OauthAccessToken) CreatedAt() time.Time {
	return oauth.createdAt
}

func (oauth *OauthAccessToken) UpdatedAt() time.Time {
	return oauth.updatedAt
}

func (oauth *OauthAccessToken) RevokedAt() *time.Time {
	return oauth.deletedAt
}

func (oauth *OauthAccessToken) ExpiresAt() time.Time {
	return oauth.expiresAt
}
