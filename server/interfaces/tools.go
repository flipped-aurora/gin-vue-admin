package interfaces

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
)

type Tools interface {
	Install() error
	UnInstall() error
	Enable() error
	Disable() error
	Call(req request.Call) (*runner.CallResponse, error)
	RootPath() string
	InstallPath() string
}
