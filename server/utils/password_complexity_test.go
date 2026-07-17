package utils

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func TestValidatePasswordComplexity(t *testing.T) {
	base := system.SysSecurityConfig{PwdMinLength: 8}
	all := system.SysSecurityConfig{
		PwdMinLength:      8,
		PwdRequireUpper:   true,
		PwdRequireLower:   true,
		PwdRequireDigit:   true,
		PwdRequireSpecial: true,
	}
	cases := []struct {
		name    string
		pwd     string
		cfg     system.SysSecurityConfig
		wantErr bool
	}{
		{"too short", "Ab1!", base, true},
		{"min length ok", "abcdefgh", base, false},
		{"all rules ok", "Abcdef1!", all, false},
		{"missing upper", "abcdef1!", all, true},
		{"missing lower", "ABCDEF1!", all, true},
		{"missing digit", "Abcdefg!", all, true},
		{"missing special", "Abcdefg1", all, true},
		{"empty pwd fails min", "", base, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePasswordComplexity(tc.pwd, tc.cfg)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("expected nil, got %v", err)
			}
		})
	}
}
