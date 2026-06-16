package system

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

const (
	cliSourceDir            = "cmd/gva"
	cliBuildDir             = ".gva-build"
	cliBuildTimeout         = 5 * time.Minute
	cliEmbeddedGoName       = "embedded.go"
	cliEmbeddedManifestName = "manifest.embed.json"
)

var (
	cliBuildGOOSAllowList   = map[string]struct{}{"windows": {}, "linux": {}, "darwin": {}}
	cliBuildGOARCHAllowList = map[string]struct{}{"amd64": {}, "arm64": {}}
)

// normalizeBuildTarget 校验并归一化目标平台。空值回退到当前运行平台。
func normalizeBuildTarget(goos, goarch string) (string, string, error) {
	goos = strings.ToLower(strings.TrimSpace(goos))
	goarch = strings.ToLower(strings.TrimSpace(goarch))
	if goos == "" {
		goos = runtime.GOOS
	}
	if goarch == "" {
		goarch = runtime.GOARCH
	}
	if _, ok := cliBuildGOOSAllowList[goos]; !ok {
		return "", "", fmt.Errorf("不支持的目标系统: %s", goos)
	}
	if _, ok := cliBuildGOARCHAllowList[goarch]; !ok {
		return "", "", fmt.Errorf("不支持的目标架构: %s", goarch)
	}
	return goos, goarch, nil
}

// cliBinaryExt 返回目标平台对应的二进制后缀。
func cliBinaryExt(goos string) string {
	if goos == "windows" {
		return ".exe"
	}
	return ""
}

// BuildCliBinary 把指定 CLI 的 manifest 内嵌进 gva 源码副本，交叉编译出一个开箱即用的二进制。
func (s *cliService) BuildCliBinary(req systemReq.BuildSysCliBinaryRequest) (string, []byte, error) {
	goos, goarch, err := normalizeBuildTarget(req.GOOS, req.GOARCH)
	if err != nil {
		return "", nil, err
	}
	cli, bindings, err := s.getCliAndBindings(req.CliID)
	if err != nil {
		return "", nil, err
	}
	manifest, err := buildSysCliManifest(cli, bindings)
	if err != nil {
		return "", nil, err
	}
	manifestBytes, err := marshalSysCliManifest(manifest)
	if err != nil {
		return "", nil, err
	}
	if _, err := exec.LookPath("go"); err != nil {
		return "", nil, fmt.Errorf("服务器未安装 Go 工具链，无法编译: %w", err)
	}
	absBuildRoot, err := filepath.Abs(cliBuildDir)
	if err != nil {
		return "", nil, fmt.Errorf("解析编译目录失败: %w", err)
	}
	if err := os.MkdirAll(absBuildRoot, 0o755); err != nil {
		return "", nil, fmt.Errorf("创建编译根目录失败: %w", err)
	}

	buildDir, err := os.MkdirTemp(absBuildRoot, "cli-")
	if err != nil {
		return "", nil, fmt.Errorf("创建编译目录失败: %w", err)
	}
	defer os.RemoveAll(buildDir)

	if err := copyCliSources(cliSourceDir, buildDir); err != nil {
		return "", nil, err
	}
	if err := writeEmbeddedManifest(buildDir, manifestBytes); err != nil {
		return "", nil, err
	}

	binaryName := sanitizeSingleSegmentSlug(cli.Command)
	if binaryName == "" {
		binaryName = "cli"
	}
	binaryName += cliBinaryExt(goos)
	binaryPath := filepath.Join(buildDir, binaryName)

	if err := runGoBuild(buildDir, binaryPath, goos, goarch); err != nil {
		return "", nil, err
	}
	binary, err := os.ReadFile(binaryPath)
	if err != nil {
		return "", nil, err
	}
	return binaryName, binary, nil
}

// copyCliSources 复制 gva 源码到编译目录，跳过测试文件和默认的 embedded.go。
func copyCliSources(srcDir, destDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("读取 CLI 源码目录失败(%s): %w", srcDir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		if name == cliEmbeddedGoName {
			continue
		}
		data, err := os.ReadFile(filepath.Join(srcDir, name))
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(destDir, name), data, 0o644); err != nil {
			return err
		}
	}
	return nil
}

// writeEmbeddedManifest 写入内嵌版的 embedded.go 与 manifest.embed.json。
func writeEmbeddedManifest(destDir string, manifestBytes []byte) error {
	if err := os.WriteFile(filepath.Join(destDir, cliEmbeddedManifestName), manifestBytes, 0o644); err != nil {
		return err
	}
	embedded := "package main\n\nimport _ \"embed\"\n\n//go:embed " + cliEmbeddedManifestName + "\nvar embeddedManifest []byte\n"
	return os.WriteFile(filepath.Join(destDir, cliEmbeddedGoName), []byte(embedded), 0o644)
}

// runGoBuild 在编译目录里交叉编译 gva。
func runGoBuild(buildDir, binaryPath, goos, goarch string) error {
	ctx, cancel := context.WithTimeout(context.Background(), cliBuildTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binaryPath, ".")
	cmd.Dir = buildDir
	cmd.Env = append(os.Environ(),
		"GOOS="+goos,
		"GOARCH="+goarch,
		"CGO_ENABLED=0",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go build 失败: %w\n%s", err, strings.TrimSpace(string(output)))
	}
	return nil
}
