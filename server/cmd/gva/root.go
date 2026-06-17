package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	configPath := defaultConfigPath()
	cfg, err := loadConfig(configPath)
	if err != nil {
		cfg = CliConfig{}
	}

	var manifestPath string

	path := cliManifestPathFromArgs(os.Args[1:])
	if path == "" {
		path = cfg.ManifestPath
	}

	programName := "gva"
	var loadedManifest *Manifest
	manifest, embedded, manifestOK := loadActiveManifest(path)
	if manifestOK {
		loadedManifest = &manifest
		if name := sanitizeProgramName(manifest.Name); name != "" {
			programName = name
		}
		if cfg.BaseURL == "" {
			cfg.BaseURL = manifest.Server.BaseURL
		}
		if cfg.AuthHeader == "" {
			cfg.AuthHeader = manifest.Server.AuthHeader
		}
		if !embedded {
			cfg.ManifestPath = path
		}
	}

	root := &cobra.Command{
		Use:   programName,
		Short: "Gin-Vue-Admin dynamic CLI",
	}
	root.PersistentFlags().StringVar(&manifestPath, "manifest", "", "Manifest 文件路径")
	root.PersistentFlags().StringVar(&cfg.BaseURL, "base-url", cfg.BaseURL, "服务地址")
	root.PersistentFlags().StringVar(&cfg.AuthHeader, "auth-header", cfg.AuthHeader, "认证头名称")
	root.AddCommand(newLoginCmd(&cfg, configPath))
	root.AddCommand(newSetBaseURLCmd(&cfg, configPath))
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "显示版本",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = cmd.OutOrStdout().Write([]byte(programName + "\n"))
		},
	})

	if loadedManifest != nil {
		attachDynamicCommands(root, *loadedManifest, &cfg)
	}
	return root
}

// loadActiveManifest 优先使用编译期内嵌的 manifest；
// 没有内嵌时再从 path 加载。第三个返回值表示是否成功取得 manifest。
func loadActiveManifest(path string) (Manifest, bool, bool) {
	if len(embeddedManifest) > 0 {
		var m Manifest
		if err := json.Unmarshal(embeddedManifest, &m); err != nil {
			return Manifest{}, true, false
		}
		return m, true, true
	}
	if strings.TrimSpace(path) == "" {
		return Manifest{}, false, false
	}
	m, err := loadManifest(path)
	if err != nil {
		return Manifest{}, false, false
	}
	return m, false, true
}

func cliManifestPathFromArgs(args []string) string {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--manifest" && i+1 < len(args) {
			return strings.TrimSpace(args[i+1])
		}
		if strings.HasPrefix(arg, "--manifest=") {
			return strings.TrimSpace(strings.TrimPrefix(arg, "--manifest="))
		}
	}
	return ""
}

// sanitizeProgramName 把 manifest.Name 规整为合法的命令行程序名：
// 转小写，非 [a-z0-9] 字符合并为单个 "-"，并去掉首尾 "-"。
func sanitizeProgramName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return ""
	}
	var b strings.Builder
	prevDash := false
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevDash = false
			continue
		}
		if !prevDash && b.Len() > 0 {
			b.WriteRune('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}


func execute() int {
	if err := newRootCmd().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
