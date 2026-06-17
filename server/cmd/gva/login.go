package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newLoginCmd(cfg *CliConfig, configPath string) *cobra.Command {
	var token string
	cmd := &cobra.Command{
		Use:   "login",
		Short: "保存JWT token到本地配置",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strings.TrimSpace(token) == "" {
				return fmt.Errorf("token不能为空")
			}
			cfg.Token = strings.TrimSpace(token)
			return saveConfig(configPath, *cfg)
		},
	}
	cmd.Flags().StringVar(&token, "token", "", "JWT token")
	return cmd
}

func newSetBaseURLCmd(cfg *CliConfig, configPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "set-base-url <url>",
		Short: "更改后台 API 地址并保存到本地配置",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.BaseURL = strings.TrimRight(strings.TrimSpace(args[0]), "/")
			return saveConfig(configPath, *cfg)
		},
	}
}
