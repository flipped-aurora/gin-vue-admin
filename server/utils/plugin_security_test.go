package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidatePluginName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "announcement"},
		{name: "my_plugin2"},
		{name: "", wantErr: true},
		{name: "../escape", wantErr: true},
		{name: `a\b`, wantErr: true},
		{name: "a/b", wantErr: true},
		{name: "APlugin", wantErr: true},
		{name: "for", wantErr: true},
		{name: "plugin.name", wantErr: true},
		{name: "_private", wantErr: true},
		{name: "插件", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePluginName(tt.name)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidatePluginName(%q) error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestJoinWithinRoot(t *testing.T) {
	root := t.TempDir()
	got, err := JoinWithinRoot(root, "announcement", "initialize", "api.go")
	if err != nil {
		t.Fatalf("JoinWithinRoot(valid) error = %v", err)
	}
	rel, err := filepath.Rel(root, got)
	if err != nil || rel != filepath.Join("announcement", "initialize", "api.go") {
		t.Fatalf("JoinWithinRoot(valid) = %q, rel = %q, err = %v", got, rel, err)
	}

	outside := filepath.Join(filepath.Dir(root), "outside")
	tests := [][]string{
		{"..", "outside"},
		{outside},
		{string(os.PathSeparator)},
	}
	for _, elems := range tests {
		if got, err = JoinWithinRoot(root, elems...); err == nil {
			t.Fatalf("JoinWithinRoot(%q) = %q, want error", elems, got)
		}
	}
}
