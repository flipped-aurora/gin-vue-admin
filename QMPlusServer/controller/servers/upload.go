package servers

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"mime/multipart"
	"time"
)

var accessKey string = "25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ" // 你在七牛云的accessKey  这里是我个人测试号的key 仅供测试使用 恳请大家不要乱传东西
var secretKey string = "pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY" // 你在七牛云的secretKey  这里是我个人测试号的key 仅供测试使用 恳请大家不要乱传东西

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func Upload(file *multipart.FileHeader, bucket string, urlPath string) (err error, path string) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return e, ""
	}
	dataLen := file.Size
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	err = formUploader.Put(context.Background(), &ret, upToken, fileKey, f, dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		//qmlog.QMLog.Info(err)
		return err, ""
	}
	return err, urlPath + "/" + ret.Key
}
