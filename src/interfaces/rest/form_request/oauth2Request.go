package form_request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/errors"
)

type Oauth2Request struct {
	ClientId     string `form:"client_id" json:"client_id"`
	ClientSecret string `form:"client_secret" json:"client_secret"`
	GrantType    string `form:"grant_type" json:"grant_type"`
	RedirectUri  string `form:"redirect_uri" json:"redirect_uri"`
	Scope        string `form:"scope" json:"scope"`
	Code         string `form:"code" json:"code"`
	State        string `form:"state" json:"state"`
}

func (oauthRequest *Oauth2Request) Validate(ctx *gin.Context) *errors.ValidationError {
	request := NewRequestValidation()
	request.AddParam("client_id", &oauthRequest.ClientId, validation.Required, validation.Length(1, 36), is.UUID)
	request.AddParam("client_secret", &oauthRequest.ClientSecret, validation.When(
		oauthRequest.GrantType == constants.ClientCredentialsGrantType,
		validation.Required,
	))
	request.AddParam("grant_type", &oauthRequest.GrantType,
		validation.Required,
		validation.Length(1, len(constants.AuthorizationGrantType)),
		validation.In(
			constants.AuthorizationGrantType,
			constants.ClientCredentialsGrantType,
			constants.RefreshTokenGrantType,
			constants.DeviceCodeGrantType,
		),
	)
	request.AddParam("redirect_uri", &oauthRequest.RedirectUri, validation.When(
		oauthRequest.GrantType == constants.AuthorizationGrantType,
		validation.Required,
		is.URL,
	))
	request.AddParam("code", &oauthRequest.Code, validation.When(
		oauthRequest.GrantType != "" &&
			&oauthRequest.GrantType != nil &&
			oauthRequest.GrantType == constants.AuthorizationGrantType,
		validation.Required,
	))
	request.AddParam("scope", &oauthRequest.Scope, validation.Required)

	return request.Validate(ctx)
}
