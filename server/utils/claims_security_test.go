package utils

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

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
	parsed, err := NewJWT().ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken err = %v", err)
	}
	if !parsed.MustChangePwd {
		t.Fatalf("parsed.MustChangePwd = false, want true")
	}
}
