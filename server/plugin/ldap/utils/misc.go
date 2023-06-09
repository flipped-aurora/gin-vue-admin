package utils

import (
	"fmt"
	systemModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/go-ldap/ldap/v3"
	"net/url"
	"strings"
)

// 类型转换
func InterfaceConvertInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func FillSysUser(u *systemModel.SysUser, entry *ldap.Entry, fm string) error {
	if len(fm) == 0 {
		return nil
	}
	if !strings.HasPrefix(fm, "?") {
		fm = fmt.Sprintf("?%s", fm)
	}
	up, err := url.Parse(fm)
	if err != nil {
		return err
	}
	queries := up.Query()
	for key := range queries {
		val := queries.Get(key)
		attr := entry.GetAttributeValue(val)
		switch key {
		case "phone":
			u.Phone = attr
		case "email":
			u.Email = attr
		}
	}
	return nil
}
