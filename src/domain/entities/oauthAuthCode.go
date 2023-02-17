package entities

import (
	"github.com/gofrs/uuid"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/valueObjects"
	"time"
)

type OauthAuthCode struct {
	id         uuid.UUID `json:"id"`
	userId     uuid.UUID `json:"user_id"`
	clientId   uuid.UUID `json:"client_id"`
	scopes     []string  `json:"scopes"`
	expiresAt  time.Time `json:"expires_at"`
	timestamps valueObjects.TimestampValueObject
}

func NewOauthAuthCode(
	id uuid.UUID,
	userId uuid.UUID,
	clientId uuid.UUID,
	scopes []string,
	expiresAt time.Time,
	timestamps valueObjects.TimestampValueObject,
) *OauthAuthCode {
	return &OauthAuthCode{
		id:         id,
		userId:     userId,
		clientId:   clientId,
		scopes:     scopes,
		expiresAt:  expiresAt,
		timestamps: timestamps,
	}
}

func CreateOauthCode(
	userId uuid.UUID,
	clientId uuid.UUID,
	scopes []string,
	expiresAt time.Time,
) (*OauthAuthCode, domainErrors.DomainError) {
	if userId.IsNil() {
		return nil, domainErrors.ThrowFieldsRequired("user_id")
	}

	if clientId.IsNil() {
		return nil, domainErrors.ThrowFieldsRequired("client_id")
	}

	if expiresAt.IsZero() {
		return nil, domainErrors.ThrowFieldsRequired("expires_at")
	}

	timestamps := valueObjects.NewTimestampValueObject(time.Now(), time.Now(), nil)
	return &OauthAuthCode{
		userId:     userId,
		clientId:   clientId,
		scopes:     scopes,
		expiresAt:  expiresAt,
		timestamps: timestamps,
	}, nil
}
