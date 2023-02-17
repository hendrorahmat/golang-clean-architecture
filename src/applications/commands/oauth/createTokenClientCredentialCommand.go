package oauth

type CreateTokenClientCredentialCommand struct {
	ClientId     string   `fake:"{uuid}"`
	ClientSecret string   `fake:"{lettern:40}"`
	Scopes       []string `fakesize:"2"`
}

func NewCreateTokenClientCredentialCommand(clientId, clientSecret string, scopes []string) IIssueTokenCommand {
	return &CreateTokenClientCredentialCommand{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
	}
}
