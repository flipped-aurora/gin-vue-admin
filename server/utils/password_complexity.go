package utils

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// ValidatePasswordComplexity 按安全配置校验密码复杂度 不满足返回可读错误
func ValidatePasswordComplexity(pwd string, cfg system.SysSecurityConfig) error {
	if cfg.PwdMinLength > 0 && utf8.RuneCountInString(pwd) < cfg.PwdMinLength {
		return fmt.Errorf("密码长度不能少于 %d 位", cfg.PwdMinLength)
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range pwd {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}
	var missing []string
	if cfg.PwdRequireUpper && !hasUpper {
		missing = append(missing, "大写字母")
	}
	if cfg.PwdRequireLower && !hasLower {
		missing = append(missing, "小写字母")
	}
	if cfg.PwdRequireDigit && !hasDigit {
		missing = append(missing, "数字")
	}
	if cfg.PwdRequireSpecial && !hasSpecial {
		missing = append(missing, "特殊字符")
	}
	if len(missing) > 0 {
		return fmt.Errorf("密码必须包含%s", strings.Join(missing, "、"))
	}
	return nil
}
