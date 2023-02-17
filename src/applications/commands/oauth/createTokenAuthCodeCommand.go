package oauth

type CreateTokenAuthCodeCommand struct {
	ClientId     string
	ClientSecret string
	Scopes       []string
	RedirectUri  string
}

func NewCreateTokenAuthCodeCommand(
	clientId string,
	clientSecret string,
	scopes []string,
	redirectUri string,
) IIssueTokenCommand {
	return &CreateTokenAuthCodeCommand{ClientId: clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		RedirectUri:  redirectUri,
	}
}
