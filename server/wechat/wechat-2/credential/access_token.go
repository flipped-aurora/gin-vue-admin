package credential

import "context"

// AccessTokenHandle AccessToken 接口
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}

// AccessTokenContextHandle AccessToken 接口
type AccessTokenContextHandle interface {
	AccessTokenHandle
	GetAccessTokenContext(ctx context.Context) (accessToken string, err error)
}
