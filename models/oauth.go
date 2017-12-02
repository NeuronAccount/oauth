package models

type OauthClient struct {
	ClientId     string
	AccountId    string
	PasswordHash string
	RedirectUri  string
}

type AuthorizeParams struct {
	Jwt          string
	ResponseType string
	ClientID     string
	Scope        string
	RedirectURI  string
	State        string
}

type AuthorizationCode struct {
	Code          string
	ClientId      string
	AccountId     string
	Scope         string
	RedirectUri   string
	ExpireSeconds int64
}

type AccessToken struct {
	AccessToken  string
	TokenType    string
	ClientId     string
	AccountId    string
	Scope        string
	ExpiresIn    int64
	RefreshToken string
}
