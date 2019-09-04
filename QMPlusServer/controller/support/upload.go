package support

import (
	"bytes"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"mime/multipart"
)

var accessKey string = "25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ" // 你在七牛云的accessKey
var secretKey string = "pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY" // 你在七牛云的secretKey
var bucket string = "a303176530"                                  // 你七牛云标准空间的名字
func Upload(file multipart.File) (err error, path string) {

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
	data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
