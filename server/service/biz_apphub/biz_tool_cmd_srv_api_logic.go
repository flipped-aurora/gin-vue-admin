package biz_apphub

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/osx"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/compress"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	copy2 "github.com/otiai10/copy"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func NewCmdSoft(new *biz_apphub.BizToolCmdSrvApi, old *biz_apphub.BizToolCmdSrvApi) *CmdSoft {
	//global.GVA_DB.Model()
	return &CmdSoft{
		OssPath:        new.OssPath,
		User:           new.TenantUser,
		Name:           new.AppCode,
		NewVersionConf: new,
		OldVersionConf: old,
	}
}

type CmdSoft struct {
	OssPath        string `json:"oss_path"`
	User           string `json:"user"`
	Name           string `json:"name"`
	NewVersionConf *biz_apphub.BizToolCmdSrvApi
	OldVersionConf *biz_apphub.BizToolCmdSrvApi
}

type InstallInfo struct {
	InstallDir string `json:"install_dir"`
	SoftName   string `json:"soft_name"`
}

func (s *CmdSoft) Install() (installInfo *InstallInfo, err error) {
	absPath := "./soft_cmd"
	//absPath = strings.Join([]string{absPath, s.User, s.Name}, "/")
	absPath = strings.Join([]string{absPath, s.User}, "/")
	unZipOut := absPath + "/" + s.Name
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
	defer os.Remove(out)
	os.MkdirAll(absPath, os.ModePerm)
	url := "http://cdn.geeleo.com/" + s.OssPath
	err = httpx.DownloadFile(url, out)
	if err != nil {
		return nil, err
	}

	unZipPath, err := compress.UnZip(filepath.Join(absPath, fileName), unZipOut)
	if err != nil {
		return nil, err
	}
	//todo 设置权限
	//exec.Command("chmod")
	// 创建一个命令来添加执行权限
	if runtime.GOOS != "windows" {
		p := unZipOut + "/" + s.Name
		cmd := exec.Command("chmod", "+x", unZipOut+"/"+s.Name)

		// 执行命令
		err = cmd.Run()
		if err != nil {
			fmt.Printf("cmd.Run() failed with p:%s err:%s\n", p, err)
			return nil, err
		}
	}

	//判断是否存在该软件
	return &InstallInfo{InstallDir: unZipPath, SoftName: appName}, nil
}

func (s *CmdSoft) UpdateVersion() (installInfo *InstallInfo, err error) {
	src := "./soft_cmd"

	//ps: OssPath=  tool/beiluo/1725442391820/helloworld.zip

	src = strings.Join([]string{src, s.NewVersionConf.TenantUser}, "/")

	path := strings.Split(s.NewVersionConf.OssPath, "/")
	fileName := path[len(path)-1] //ps: helloworld.zip

	appDirName := strings.TrimSuffix(fileName, ".zip")                                          //目录名称：helloworld
	backPath := strings.Join([]string{src, ".back", appDirName, s.OldVersionConf.Version}, "/") // ps: ./soft_cmd/beiluo/.back/helloworld/v1.0
	appName := ""                                                                               //
	if runtime.GOOS == "windows" {
		//fileName  =soft.zip
		appName = strings.Split(fileName, ".")[0] + ".exe" //windows ps: helloworld.exe
	} else {
		appName = strings.Split(fileName, ".")[0] //linux ps: helloworld
	}
	//fmt.Println("appName:", appName)
	zipOut := src + "/" + fileName                       //ps: ./soft_cmd/beiluo/helloworld.zip
	currentSoftSrc := strings.TrimSuffix(zipOut, ".zip") //ps: ./soft_cmd/beiluo/helloworld
	url := "http://cdn.geeleo.com/" + s.NewVersionConf.OssPath
	err = httpx.DownloadFile(url, zipOut)
	if err != nil {
		return nil, err
	}
	defer os.Remove(zipOut)

	err = copy2.Copy(currentSoftSrc, backPath) //更换版本前把旧版本备份一份，存档
	if err != nil {
		return nil, err
	}
	copyTempDir := currentSoftSrc + "/.temp/" + appDirName
	defer os.RemoveAll(copyTempDir)
	unZipPath, err := compress.UnZip(zipOut, copyTempDir)
	if err != nil {
		return nil, err
	}
	fmt.Println(unZipPath)
	err = osx.CopyDirectory(copyTempDir, currentSoftSrc) //复制目录
	if err != nil {
		return nil, err
	}
	if runtime.GOOS != "windows" {
		p := currentSoftSrc + "/" + appDirName
		cmd := exec.Command("chmod", "+x", p)

		// 执行命令
		err = cmd.Run()
		if err != nil {
			fmt.Printf("cmd.Run() failed with p:%s err:%s\n", p, err)
			return nil, err
		}
	}

	//复制目录
	return &InstallInfo{SoftName: appName, InstallDir: currentSoftSrc}, nil
}

// Install 工具安装
func (b *BizToolCmdSrvApiService) Install(req biz_apphub.BizToolCmdSrvApi) (installInfo *InstallInfo, err error) {
	//todo
	app := CmdSoft{OssPath: req.OssPath, User: req.OperateUser, Name: req.AppCode}
	global.GVA_LOG.Info(fmt.Sprintf("Install %+v", req))
	return app.Install()
}

// UpdateVersion 工具安装
func (b *BizToolCmdSrvApiService) UpdateVersion(new biz_apphub.BizToolCmdSrvApi, old biz_apphub.BizToolCmdSrvApi) (installInfo *InstallInfo, err error) {
	soft := NewCmdSoft(&new, &old)
	return soft.UpdateVersion()
}
