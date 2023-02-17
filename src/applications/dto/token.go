package dto

type IssueToken struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	Scope        []string
	RedirectUri  string
}

func NewIssueToken(grantType string, clientId string, clientSecret string, scope []string) *IssueToken {
	return &IssueToken{GrantType: grantType, ClientId: clientId, ClientSecret: clientSecret, Scope: scope}
}
