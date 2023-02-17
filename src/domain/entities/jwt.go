package entities

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	domainErrors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/utils"
	"time"
)

type JWTClaims struct {
	registeredClaims jwt.RegisteredClaims
	scopes           []string `json:"scopes" json:"scopes,omitempty"`
}

func NewJwtClaim(id *uuid.UUID, audiences []uuid.UUID, subject string, scopes []string) *JWTClaims {
	config.MakeJwtConfig()
	claims := jwt.ClaimStrings{}

	for _, audience := range audiences {
		claims = append(claims, audience.String())
	}
	return &JWTClaims{
		registeredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   subject,
			Audience:  claims,
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(config.LoginExpirationDuration)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ID:        id.String(),
		},
		scopes: scopes,
	}
}

func (jwtClaim *JWTClaims) GetJwtId() uuid.UUID {
	return uuid.FromStringOrNil(jwtClaim.registeredClaims.ID)
}

func (jwtClaim *JWTClaims) GetExpiresAt() time.Time {
	return jwtClaim.registeredClaims.ExpiresAt.Time
}

func (jwtClaim *JWTClaims) SetId(id uuid.UUID) *JWTClaims {
	jwtClaim.registeredClaims.ID = id.String()
	return jwtClaim
}

func (jwtClaim *JWTClaims) Generate() (*string, domainErrors.DomainError) {
	if jwtClaim.registeredClaims.ID == "" || &jwtClaim.registeredClaims.ID == nil {
		return nil, domainErrors.ThrowFieldsRequired("jwt id")
	}

	if config.JwtSignatureKey == nil || len(config.JwtSignatureKey) <= 0 {
		return nil, domainErrors.ThrowFieldsRequired("public_key", "private_key")
	}

	privateKeyFromPEM, err := jwt.ParseRSAPrivateKeyFromPEM(config.JwtSignatureKey)
	if err != nil {
		return nil, domainErrors.ThrowInternalServerError(err.Error())
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	if jwtClaim.registeredClaims.Subject != "" {
		claims["sub"] = jwtClaim.registeredClaims.Subject
	}

	claims["iss"] = utils.GetEnvWithDefaultValue("APP_URL", "http://localhost:8000")
	claims["aud"] = jwtClaim.registeredClaims.Audience
	claims["exp"] = jwtClaim.registeredClaims.ExpiresAt
	claims["nbf"] = jwtClaim.registeredClaims.NotBefore
	claims["iat"] = jwtClaim.registeredClaims.IssuedAt
	claims["jti"] = jwtClaim.registeredClaims.ID
	claims["scopes"] = jwtClaim.scopes

	signedToken, err := token.SignedString(privateKeyFromPEM)
	if err != nil {
		return nil, domainErrors.ThrowInternalServerError(err.Error())
	}

	return &signedToken, nil
}
