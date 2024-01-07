package stalcraftgo // import "github.com/re-incarnation/stalcraftgo/object"

type OAuth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
