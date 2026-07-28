package system

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSysUserMustChangePasswordIsNotSerialized(t *testing.T) {
	data, err := json.Marshal(SysUser{MustChangePassword: true})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "mustChangePassword") {
		t.Fatalf("SysUser JSON exposes mustChangePassword: %s", data)
	}
}
