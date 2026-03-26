package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/spf13/cobra"
)

const (
	configFlagShort = "-c"
	configFlagLong  = "--config"
)

func runCli() error {
	var configPath string
	var resetPasswordUsername string

	rootCmd := &cobra.Command{
		Use:           filepath.Base(os.Args[0]),
		Short:         "gva",
		Example:       fmt.Sprintf("%s -c config.yaml --reset-password admin", filepath.Base(os.Args[0])),
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if resetPasswordUsername != "" {
				return resetPasswordByUsername(resetPasswordUsername)
			}
			return cmd.Help()
		},
	}

	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "choose config file.")
	rootCmd.Flags().StringVar(&resetPasswordUsername, "reset-password", "", "reset the specified user's password and print the new password")

	return rootCmd.Execute()
}

func resetPasswordByUsername(username string) error {
	username = strings.TrimSpace(username)
	if username == "" {
		return errors.New("用户名不能为空")
	}
	if global.GVA_DB == nil {
		return errors.New("db not init")
	}

	password, err := utils.RandomPassword(16)
	if err != nil {
		return err
	}
	if err = system.UserServiceApp.ResetPasswordByUsername(username, password); err != nil {
		return err
	}
	fmt.Printf("++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("the new password is: %s\n", password)
	fmt.Printf("++++++++++++++++++++++++++++++++++\n")
	return nil
}

func hasCommandArgs(args []string) bool {
	return len(filterCommonArgs(args)) > 0
}

func filterCommonArgs(args []string) []string {
	filtered := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == configFlagShort || arg == configFlagLong:
			if i+1 < len(args) {
				i++
			}
		case strings.HasPrefix(arg, configFlagShort+"="), strings.HasPrefix(arg, configFlagLong+"="):
			continue
		default:
			filtered = append(filtered, arg)
		}
	}
	return filtered
}

func lookupConfigPathArg(args []string) (string, bool) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == configFlagShort || arg == configFlagLong:
			if i+1 < len(args) {
				return args[i+1], true
			}
			return "", false
		case strings.HasPrefix(arg, configFlagShort+"="):
			return strings.TrimPrefix(arg, configFlagShort+"="), true
		case strings.HasPrefix(arg, configFlagLong+"="):
			return strings.TrimPrefix(arg, configFlagLong+"="), true
		}
	}
	return "", false
}
