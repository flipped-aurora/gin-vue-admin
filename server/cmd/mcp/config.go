package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gopkg.in/yaml.v3"
)

type standaloneConfig struct {
	MCP      config.MCP      `yaml:"mcp"`
	AutoCode config.Autocode `yaml:"autocode"`
}

func loadStandaloneConfig() (string, error) {
	configPath, err := resolveConfigPath()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("读取 MCP 配置失败: %w", err)
	}

	var cfg standaloneConfig
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return "", fmt.Errorf("解析 MCP 配置失败: %w", err)
	}

	applyStandaloneDefaults(configPath, &cfg)

	global.GVA_CONFIG.MCP = cfg.MCP
	global.GVA_CONFIG.AutoCode = cfg.AutoCode

	return configPath, nil
}

func resolveConfigPath() (string, error) {
	explicit, err := parseConfigFlag(os.Args[1:])
	if err != nil {
		return "", err
	}
	if explicit != "" {
		return filepath.Abs(explicit)
	}

	if envPath := strings.TrimSpace(os.Getenv("GVA_MCP_CONFIG")); envPath != "" {
		return filepath.Abs(envPath)
	}

	wd, _ := os.Getwd()
	exe, _ := os.Executable()
	exeDir := filepath.Dir(exe)

	candidates := []string{
		filepath.Join(wd, "config.yaml"),
		filepath.Join(wd, "mcp.yaml"),
		filepath.Join(wd, "cmd", "mcp", "config.yaml"),
		filepath.Join(wd, "server", "cmd", "mcp", "config.yaml"),
		filepath.Join(exeDir, "config.yaml"),
		filepath.Join(exeDir, "mcp.yaml"),
	}

	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return filepath.Abs(candidate)
		}
	}

	return "", errors.New("未找到 MCP 独立配置文件，请在当前目录、cmd/mcp 目录或通过 -config / GVA_MCP_CONFIG 指定 config.yaml")
}

func parseConfigFlag(args []string) (string, error) {
	fs := flag.NewFlagSet("gva-mcp", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var configPath string
	fs.StringVar(&configPath, "c", "", "MCP config file path")
	fs.StringVar(&configPath, "config", "", "MCP config file path")
	if err := fs.Parse(args); err != nil {
		return "", err
	}

	return strings.TrimSpace(configPath), nil
}

func applyStandaloneDefaults(configPath string, cfg *standaloneConfig) {
	if cfg.MCP.Name == "" {
		cfg.MCP.Name = "GVA_MCP"
	}
	if cfg.MCP.Version == "" {
		cfg.MCP.Version = "v1.0.0"
	}
	if cfg.MCP.Path == "" {
		cfg.MCP.Path = "/mcp"
	}
	if cfg.MCP.Addr == 0 {
		cfg.MCP.Addr = 8889
	}
	if cfg.MCP.AuthHeader == "" {
		cfg.MCP.AuthHeader = "x-token"
	}
	if cfg.MCP.RequestTimeout <= 0 {
		cfg.MCP.RequestTimeout = 15
	}
	if cfg.MCP.UpstreamBaseURL == "" {
		cfg.MCP.UpstreamBaseURL = "http://127.0.0.1:8888"
	}
	if cfg.MCP.BaseURL == "" {
		cfg.MCP.BaseURL = fmt.Sprintf("http://127.0.0.1:%d%s", cfg.MCP.Addr, cfg.MCP.Path)
	}

	configDir := filepath.Dir(configPath)
	if cfg.AutoCode.Server == "" {
		cfg.AutoCode.Server = "server"
	}
	if cfg.AutoCode.Web == "" {
		cfg.AutoCode.Web = "web/src"
	}
	if cfg.AutoCode.Root == "" {
		if root, err := detectProjectRoot(configDir); err == nil {
			cfg.AutoCode.Root = root
		}
	} else if !filepath.IsAbs(cfg.AutoCode.Root) {
		cfg.AutoCode.Root = filepath.Clean(filepath.Join(configDir, cfg.AutoCode.Root))
	}

	if cfg.AutoCode.Module == "" && cfg.AutoCode.Root != "" {
		goModPath := filepath.Join(cfg.AutoCode.Root, cfg.AutoCode.Server, "go.mod")
		if module, err := detectGoModule(goModPath); err == nil {
			cfg.AutoCode.Module = module
		}
	}
}

func detectProjectRoot(startDir string) (string, error) {
	dir := startDir
	for {
		serverDir := filepath.Join(dir, "server")
		webDir := filepath.Join(dir, "web")
		if isDir(serverDir) && isDir(webDir) {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", errors.New("未能自动识别项目根目录，请在 MCP 配置中设置 autocode.root")
}

func detectGoModule(goModPath string) (string, error) {
	file, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New("go.mod 中未找到 module 定义")
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
