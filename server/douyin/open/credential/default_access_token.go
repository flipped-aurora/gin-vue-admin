package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

const (
	accessTokenURL       = "https://open.douyin.com/oauth/access_token?client_key=%s&client_secret=%s&code=%s&grant_type=authorization_code"
	refreshTokenURL      = "https://open.douyin.com/oauth/oauth/refresh_token?client_key=%s&grant_type=refresh_token&refresh_token=%s"
	renewRefreshTokenURL = "https://open.douyin.com/oauth/oauth/renew_refresh_token?client_key=%s&refresh_token=%s"
	clientTokenURL       = "https://open.douyin.com/oauth/oauth/client_token?client_key=%s&client_secret=%s&grant_type=client_credential"
	// CacheKeyPrefix 抖音open cache key前缀
	CacheKeyPrefix = "douyin_open"
)

// DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	ClientKey       string
	ClientSecret    string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(clientKey, clientSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is need")
	}
	return &DefaultAccessToken{
		ClientKey:       clientKey,
		ClientSecret:    clientSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// AccessToken struct
type AccessToken struct {
	util.CommonError

	AccessToken    string `json:"access_token"`
	ExpiresIn      int64  `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	RefreshTokenIn int64  `json:"refresh_expires_in"`
	OpenID         string `json:"openid"`
	Scope          string `json:"scope"`
}

// GetAccessToken 获取access_token,先从cache中获取，没有则从服务端获取
func (ak *DefaultAccessToken) GetAccessToken(openID string) (accessToken string, err error) {
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从抖音服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, openID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	// 刷新AccessToken
	refreshTokenCacheKey := fmt.Sprintf("%s_refresh_token_%s", ak.cacheKeyPrefix, openID)
	refreshToken := ak.cache.Get(refreshTokenCacheKey)
	if refreshToken == nil {
		err = fmt.Errorf("user need auth")
		return
	}

	var resAccessToken *AccessToken
	resAccessToken, err = ak.RefreshAccessToken(refreshToken.(string))
	if err != nil {
		return
	}

	// 缓存AccessToken
	err = ak.SetAccessToken(resAccessToken)
	if err != nil {
		return
	}

	accessToken = resAccessToken.AccessToken
	return
}

// SetAccessToken 设置access_token
func (ak *DefaultAccessToken) SetAccessToken(accessToken *AccessToken) (err error) {
	// access token cache
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, accessToken.OpenID)
	expires := accessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, accessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}

	// refresh access token cache
	refreshAccessTokenCacheKey := fmt.Sprintf("%s_refresh_token_%s", ak.cacheKeyPrefix, accessToken.OpenID)
	refreshTokenExpires := accessToken.RefreshTokenIn - 1500
	err = ak.cache.Set(refreshAccessTokenCacheKey, accessToken.RefreshToken, time.Duration(refreshTokenExpires)*time.Second)
	if err != nil {
		return
	}

	return
}

type accessTokenRes struct {
	Message string                `json:"message"`
	Extra   util.CommonErrorExtra `json:"extra"`

	Data AccessToken `json:"data"`
}

// RefreshAccessToken 刷新AccessToken.
// 当access_token过期（过期时间15天）后，可以通过该接口使用refresh_token（过期时间30天）进行刷新
func (ak *DefaultAccessToken) RefreshAccessToken(refreshToken string) (accessToken *AccessToken, err error) {
	uri := fmt.Sprintf(refreshTokenURL, ak.ClientKey, refreshToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result accessTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}

	err = ak.SetAccessToken(&result.Data)
	if err != nil {
		return
	}
	accessToken = &result.Data
	return
}

// RefreshToken .
type RefreshToken struct {
	util.CommonError

	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type refreshTokenRes struct {
	Message string                `json:"message"`
	Extra   util.CommonErrorExtra `json:"extra"`
	Data    RefreshToken          `json:"data"`
}

// RenewRefreshToken 刷新refresh_token.
// 前提： client_key需要具备renew_refresh_token这个权限
// 接口说明： 可以通过旧的refresh_token获取新的refresh_token,调用后旧refresh_token会失效，新refresh_token有30天有效期。最多只能获取5次新的refresh_token，5次过后需要用户重新授权。
func (ak *DefaultAccessToken) RenewRefreshToken(refreshToken string) (refreshTokenData *RefreshToken, err error) {
	uri := fmt.Sprintf(renewRefreshTokenURL, ak.ClientKey, refreshToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result refreshTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("RenewRefreshToken error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	refreshTokenData = &result.Data
	return
}

// ClientToken struct.
type ClientToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type clientTokenRes struct {
	Message string      `json:"message"`
	Data    ClientToken `json:"data"`
}

// GetClientToken 该接口用于获取接口调用的凭证client_access_token，主要用于调用不需要用户授权就可以调用的接口.
func (ak *DefaultAccessToken) GetClientToken() (clientToken *ClientToken, err error) {
	uri := fmt.Sprintf(clientTokenURL, ak.ClientKey, ak.ClientSecret)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result clientTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	clientToken = &result.Data
	return
}
