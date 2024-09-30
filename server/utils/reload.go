package utils

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New(global.Translate("utils.sysNoSupport"))
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
