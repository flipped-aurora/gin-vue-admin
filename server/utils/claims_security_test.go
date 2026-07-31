package utils

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

func tokenCookie(t *testing.T, recorder *httptest.ResponseRecorder) *http.Cookie {
	t.Helper()
	cookies := recorder.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("Set-Cookie count = %d, want 1", len(cookies))
	}
	return cookies[0]
}

func tokenContext(target string) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, target, nil)
	c.Request.Host = "admin.example.com:8443"
	return c, recorder
}

func TestTokenCookieSecurityAttributes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name      string
		configure func(*http.Request)
		secure    bool
	}{
		{name: "http"},
		{name: "tls", configure: func(r *http.Request) { r.TLS = &tls.ConnectionState{} }, secure: true},
		{name: "forwarded_https", configure: func(r *http.Request) { r.Header.Set("X-Forwarded-Proto", "https") }, secure: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, recorder := tokenContext("http://admin.example.com/api")
			if tt.configure != nil {
				tt.configure(c.Request)
			}
			SetToken(c, "signed-token", 600)
			cookie := tokenCookie(t, recorder)
			if cookie.Name != "x-token" || cookie.Value != "signed-token" || cookie.Path != "/" {
				t.Fatalf("unexpected token cookie: %+v", cookie)
			}
			if !cookie.HttpOnly {
				t.Error("token cookie HttpOnly = false, want true")
			}
			if cookie.SameSite != http.SameSiteStrictMode {
				t.Errorf("token cookie SameSite = %v, want Strict", cookie.SameSite)
			}
			if cookie.Secure != tt.secure {
				t.Errorf("token cookie Secure = %v, want %v", cookie.Secure, tt.secure)
			}
			if cookie.Domain != "" {
				t.Errorf("token cookie Domain = %q, want host-only", cookie.Domain)
			}
		})
	}
}

func TestClearTokenUsesSecurityAttributes(t *testing.T) {
	c, recorder := tokenContext("https://admin.example.com/api")
	c.Request.TLS = &tls.ConnectionState{}
	ClearToken(c)
	cookie := tokenCookie(t, recorder)
	if cookie.MaxAge >= 0 {
		t.Errorf("cleared token MaxAge = %d, want negative", cookie.MaxAge)
	}
	if !cookie.HttpOnly || !cookie.Secure || cookie.SameSite != http.SameSiteStrictMode || cookie.Domain != "" {
		t.Errorf("cleared token cookie is not hardened: %+v", cookie)
	}
}

func TestGetTokenPrefersHeaderAndFallsBackToCookie(t *testing.T) {
	t.Run("header_priority", func(t *testing.T) {
		c, _ := tokenContext("http://admin.example.com/api")
		c.Request.Header.Set("x-token", "header-token")
		c.Request.AddCookie(&http.Cookie{Name: "x-token", Value: "cookie-token"})
		if got := GetToken(c); got != "header-token" {
			t.Fatalf("GetToken() = %q, want header token", got)
		}
	})

	t.Run("cookie_fallback", func(t *testing.T) {
		c, _ := tokenContext("http://admin.example.com/api")
		c.Request.AddCookie(&http.Cookie{Name: "x-token", Value: "cookie-token"})
		if got := GetToken(c); got != "cookie-token" {
			t.Fatalf("GetToken() = %q, want cookie token", got)
		}
	})
}

func TestLoginTokenWithExpireSetsMustChangePwd(t *testing.T) {
	global.GVA_CONFIG.JWT = config.JWT{
		SigningKey:  "test-signing-key",
		ExpiresTime: "7d",
		BufferTime:  "1d",
		Issuer:      "GVA",
	}
	user := &system.SysUser{Username: "tester"}
	token, claims, err := LoginTokenWithExpire(user, true)
	if err != nil {
		t.Fatalf("LoginTokenWithExpire err = %v", err)
	}
	if !claims.MustChangePwd {
		t.Fatalf("claims.MustChangePwd = false, want true")
	}
	if claims.UserType != system.UserTypeAdmin {
		t.Fatalf("claims.UserType = %q, want %q", claims.UserType, system.UserTypeAdmin)
	}
	parsed, err := NewJWT().ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken err = %v", err)
	}
	if !parsed.MustChangePwd {
		t.Fatalf("parsed.MustChangePwd = false, want true")
	}
	if parsed.UserType != system.UserTypeAdmin {
		t.Fatalf("parsed.UserType = %q, want %q", parsed.UserType, system.UserTypeAdmin)
	}
}
