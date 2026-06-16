package main

import "testing"

func TestSanitizeProgramName(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"gva", "gva"},
		{"opsctl", "opsctl"},
		{"My Cool CLI", "my-cool-cli"},
		{"user_manager", "user-manager"},
		{"OPS Tool!", "ops-tool"},
		{"  opsctl  ", "opsctl"},
		{"api/cli", "api-cli"},
		{"中文工具", ""},
		{"---", ""},
	}
	for _, c := range cases {
		got := sanitizeProgramName(c.input)
		if got != c.want {
			t.Errorf("sanitizeProgramName(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}
