package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	"strings"
)

type CmdSoft struct {
	OssPath string `json:"oss_path"`
	User    string `json:"user"`
	Name    string `json:"name"`
}

func (s *CmdSoft) Install() error {
	absPath := "./soft_cmd"
	absPath = strings.Join([]string{absPath, s.User, s.Name}, "/")

	path := strings.Split(s.OssPath, "/")
	fileName := path[len(path)-1]
	err := httpx.DownloadFile("http://cdn.geeleo.com"+s.OssPath, absPath+"/"+fileName)
	if err != nil {
		return err
	}

	return nil
	//unZipPath, err := compress.DeCompress(filepath.Join(absPath, fileName), absPath)
	//if err != nil {
	//	return err
	//}
	//exec.Command("chmod")
}

func (b *BizToolCmdSrvApiService) Install() {

}
