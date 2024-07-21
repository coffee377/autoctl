package unilink

type AppConnect interface {
	Init()
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

type AccountData struct {
	Username string
	Password string
	Code     string
}

type SSO interface {
	AppConnect
	AccountLogin(AccountData) TokenResponse
	OAuth2Login() TokenResponse
	Logout() bool
}

type App struct {
	Id   string
	Name string
	Icon string
}

type AppInstance interface {
	GetId() uint64
	GetBaseUrl() string
	GeTHomePageUri() string
	GetConfig() map[string]any
}

//type
