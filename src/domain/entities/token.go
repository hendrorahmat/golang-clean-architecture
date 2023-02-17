package entities

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"time"
)

type Token struct {
	id           uuid.UUID
	tokenType    string    `json:"token_type"`
	expiresIn    time.Time `json:"expires_in"`
	accessToken  string    `json:"access_token"`
	refreshToken *string   `json:"refresh_token,omitempty"`
}

func (token *Token) Id() uuid.UUID {
	return token.id
}

func (token *Token) TokenType() string {
	return token.tokenType
}

func (token *Token) ExpiresIn() time.Time {
	return token.expiresIn
}

func (token *Token) AccessToken() string {
	return token.accessToken
}

func (token *Token) RefreshToken() *string {
	return token.refreshToken
}

func NewToken(
	id uuid.UUID,
	tokenType string,
	expiresIn time.Time,
	accessToken string,
	refreshToken *string,
) (*Token, domainErrors.DomainError) {
	return &Token{
		id:           id,
		tokenType:    tokenType,
		expiresIn:    expiresIn,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}, nil
}

func (token *Token) MarshalJSON() ([]byte, error) {
	response, err := json.Marshal(struct {
		TokenType    string  `json:"token_type"`
		ExpiresIn    int64   `json:"expires_in"`
		AccessToken  string  `json:"access_token"`
		RefreshToken *string `json:"refresh_token,omitempty"`
	}{
		TokenType:    token.tokenType,
		ExpiresIn:    token.expiresIn.Unix(),
		AccessToken:  token.accessToken,
		RefreshToken: token.refreshToken,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
