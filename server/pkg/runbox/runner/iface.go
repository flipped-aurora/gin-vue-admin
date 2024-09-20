package runner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/dto/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/dto/response"
)

type Runner interface {
	Install() (installInfo *InstallInfo, err error)
	UnInstall() (unInstallInfo *UnInstallInfo, err error)
	Call(req *request.Call) (*response.Call, error)
}
