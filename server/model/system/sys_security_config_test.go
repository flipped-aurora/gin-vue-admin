package system

import (
	"testing"
	"time"
)

func TestSysSecurityConfigTableName(t *testing.T) {
	if got := (SysSecurityConfig{}).TableName(); got != "sys_security_config" {
		t.Fatalf("TableName() = %q, want %q", got, "sys_security_config")
	}
}

func TestCaptchaTimeoutDuration(t *testing.T) {
	c := SysSecurityConfig{CaptchaTimeout: 3600}
	if got := c.CaptchaTimeoutDuration(); got != time.Hour {
		t.Fatalf("CaptchaTimeoutDuration() = %v, want %v", got, time.Hour)
	}
}

func TestDefaultSecurityConfig(t *testing.T) {
	cfg := DefaultSecurityConfig()
	if cfg.CaptchaOpen != 0 || cfg.CaptchaTimeout != 3600 {
		t.Fatalf("captcha open/timeout defaults wrong: %+v", cfg)
	}
	if cfg.KeyLong != 6 || cfg.ImgWidth != 240 || cfg.ImgHeight != 80 {
		t.Fatalf("captcha dimension defaults wrong: %+v", cfg)
	}
	if cfg.PwdMinLength != 8 {
		t.Fatalf("PwdMinLength default = %d, want 8", cfg.PwdMinLength)
	}
	if cfg.LockThreshold != 5 || cfg.LockDuration != 30 {
		t.Fatalf("lock defaults wrong: %+v", cfg)
	}
	if cfg.PwdExpireDays != 90 {
		t.Fatalf("PwdExpireDays default = %d, want 90", cfg.PwdExpireDays)
	}
	if cfg.ForceNewUserChangePassword {
		t.Fatal("ForceNewUserChangePassword default = true, want false")
	}
}
