package oauth2_handler

type Oauth2Handler struct {
	OauthClientHandler IOauthClientHandler
	OauthTokenHandler  IOauthTokenHandler
}
