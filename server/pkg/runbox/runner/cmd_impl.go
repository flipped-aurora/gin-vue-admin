package runner

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/dto/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/dto/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/compress"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Cmd struct {
	InstallInfo
}

func (c *Cmd) Install() (*InstallInfo, error) {
	absPath := "./soft_cmd"
	//absPath = strings.Join([]string{absPath, s.User, s.Name}, "/")
	absPath = strings.Join([]string{absPath, c.User}, "/")
	unZipOut := absPath + "/" + c.Name
	path := strings.Split(c.DownloadPath, "/")
	fileName := path[len(path)-1]

	appName := ""
	if runtime.GOOS == "windows" {
		//fileName  =soft.zip
		appName = strings.Split(fileName, ".")[0] + ".exe"
	} else {
		appName = strings.Split(fileName, ".")[0]
	}
	c.FullName = appName
	out := absPath + "/" + fileName
	defer os.Remove(out)
	os.MkdirAll(absPath, os.ModePerm)
	url := "http://cdn.geeleo.com/" + c.DownloadPath
	err := httpx.DownloadFile(url, out)
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
		p := unZipOut + "/" + c.Name
		cmd := exec.Command("chmod", "+x", unZipOut+"/"+c.Name)
		// 执行命令
		err = cmd.Run()
		if err != nil {
			fmt.Printf("cmd.Run() failed with p:%s err:%s\n", p, err)
			return nil, err
		}
	}
	c.InstallPath = unZipPath
	//判断是否存在该软件
	return &c.InstallInfo, nil
}

func (c *Cmd) UnInstall() (*UnInstallInfo, error) {
	return nil, nil
}

func (c *Cmd) Call(r *request.Call) (*response.Call, error) {
	return nil, nil
}
