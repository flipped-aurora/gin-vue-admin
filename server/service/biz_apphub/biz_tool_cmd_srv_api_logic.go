package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/compress"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type CmdSoft struct {
	OssPath string `json:"oss_path"`
	User    string `json:"user"`
	Name    string `json:"name"`
}

type InstallInfo struct {
	InstallDir string `json:"install_dir"`
	SoftName   string `json:"soft_name"`
}

func (s *CmdSoft) Install() (installInfo *InstallInfo, err error) {
	absPath := "./soft_cmd"
	//absPath = strings.Join([]string{absPath, s.User, s.Name}, "/")
	absPath = strings.Join([]string{absPath, s.User}, "/")

	path := strings.Split(s.OssPath, "/")
	fileName := path[len(path)-1]

	appName := ""
	if runtime.GOOS == "windows" {
		//fileName  =soft.zip
		appName = strings.Split(fileName, ".")[0] + ".exe"
	} else {
		appName = strings.Split(fileName, ".")[0]
	}
	out := absPath + "/" + fileName
	defer os.ReadFile(out)
	os.MkdirAll(absPath, os.ModePerm)
	url := "http://cdn.geeleo.com/" + s.OssPath
	err = httpx.DownloadFile(url, out)
	if err != nil {
		return nil, err
	}

	unZipPath, err := compress.DeCompress(filepath.Join(absPath, fileName), absPath)
	if err != nil {
		return nil, err
	}
	//todo 设置权限
	//exec.Command("chmod")

	//判断是否存在该软件
	return &InstallInfo{InstallDir: unZipPath, SoftName: appName}, nil
}

// Install 工具安装
func (b *BizToolCmdSrvApiService) Install(req biz_apphub.BizToolCmdSrvApi) (installInfo *InstallInfo, err error) {
	//todo
	app := CmdSoft{OssPath: req.OssPath, User: req.OperateUser, Name: req.AppCode}
	return app.Install()
}
