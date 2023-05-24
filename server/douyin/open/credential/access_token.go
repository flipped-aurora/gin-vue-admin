package credential

// AccessTokenHandle AccessToken 接口
type AccessTokenHandle interface {
	GetAccessToken(openID string) (accessToken string, err error)
	SetAccessToken(accessToken *AccessToken) (err error)
	GetClientToken() (clientToken *ClientToken, err error)
}
