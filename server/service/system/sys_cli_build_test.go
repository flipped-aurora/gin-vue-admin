package system

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestNormalizeBuildTarget(t *testing.T) {
	cases := []struct {
		goos, goarch string
		wantOS       string
		wantArch     string
		wantErr      bool
	}{
		{"windows", "amd64", "windows", "amd64", false},
		{"linux", "arm64", "linux", "arm64", false},
		{"darwin", "amd64", "darwin", "amd64", false},
		{"WINDOWS", "AMD64", "windows", "amd64", false},
		{"", "", "", "", false},
		{"freebsd", "amd64", "", "", true},
		{"windows", "mips", "", "", true},
	}
	for _, c := range cases {
		gos, garc, err := normalizeBuildTarget(c.goos, c.goarch)
		if c.wantErr {
			if err == nil {
				t.Errorf("normalizeBuildTarget(%q,%q) expected error", c.goos, c.goarch)
			}
			continue
		}
		if err != nil {
			t.Errorf("normalizeBuildTarget(%q,%q) unexpected error: %v", c.goos, c.goarch, err)
			continue
		}
		if c.wantOS != "" && gos != c.wantOS {
			t.Errorf("normalizeBuildTarget(%q,%q) goos=%q want %q", c.goos, c.goarch, gos, c.wantOS)
		}
		if c.wantArch != "" && garc != c.wantArch {
			t.Errorf("normalizeBuildTarget(%q,%q) goarch=%q want %q", c.goos, c.goarch, garc, c.wantArch)
		}
		if c.wantOS == "" && (gos == "" || garc == "") {
			t.Errorf("normalizeBuildTarget(%q,%q) empty default goos/goarch", c.goos, c.goarch)
		}
	}
}

func TestCliBinaryExt(t *testing.T) {
	if cliBinaryExt("windows") != ".exe" {
		t.Fatal(`cliBinaryExt("windows") want ".exe"`)
	}
	if cliBinaryExt("linux") != "" {
		t.Fatal(`cliBinaryExt("linux") want ""`)
	}
	if cliBinaryExt("darwin") != "" {
		t.Fatal(`cliBinaryExt("darwin") want ""`)
	}
}

// TestBuildCliBinaryCompilesBinary 跑真实的 go build，验证整条编译链路产出可用二进制。
func TestBuildCliBinaryCompilesBinary(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping build integration test in short mode")
	}
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available")
	}
	_, testFile, _, _ := runtime.Caller(0)
	serverRoot := filepath.Join(filepath.Dir(testFile), "..", "..")
	absRoot := filepath.Join(serverRoot, ".gva-build-test")
	if err := os.MkdirAll(absRoot, 0o755); err != nil {
		t.Fatalf("mkdirall: %v", err)
	}
	defer os.RemoveAll(absRoot)

	buildDir, err := os.MkdirTemp(absRoot, "t-")
	if err != nil {
		t.Fatalf("mkdirtemp: %v", err)
	}
	if err := copyCliSources(filepath.Join(serverRoot, "cmd", "gva"), buildDir); err != nil {
		t.Fatalf("copyCliSources: %v", err)
	}
	manifest := []byte(`{"name":"cpt","version":"v1","server":{"baseURL":"http://127.0.0.1:8888/api","authHeader":"x-token"},"commands":[]}`)
	if err := writeEmbeddedManifest(buildDir, manifest); err != nil {
		t.Fatalf("writeEmbeddedManifest: %v", err)
	}
	binaryPath := filepath.Join(buildDir, "cpt"+cliBinaryExt(runtime.GOOS))
	if err := runGoBuild(buildDir, binaryPath, runtime.GOOS, runtime.GOARCH); err != nil {
		t.Fatalf("runGoBuild: %v", err)
	}
	info, err := os.Stat(binaryPath)
	if err != nil {
		t.Fatalf("binary not found at %s: %v", binaryPath, err)
	}
	if info.Size() == 0 {
		t.Fatalf("binary is empty: %s", binaryPath)
	}
}
