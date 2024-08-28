package biz_apphub

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/compress"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/oss"
	"os"
	"path/filepath"
	"strings"
)

type FileSrc struct {
	SrcPath     string `json:"srcPath"`     //文件中的原始地址
	OssFullPath string `json:"OssFullPath"` //替换到上传oss后的地址
	OssPath     string `json:"ossPath"`     //上传oss的地址
	LocalPath   string `json:"localPath"`   //本地所处的本地地址
}

func FileGetAll(path string) (files []string, err error) {
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // 打印错误信息
			return err       // 可以选择返回错误，但这里我们选择继续遍历
		}
		if !info.IsDir() {
			files = append(files, path) // 将文件路径添加到切片中
		}
		return nil
	})

	fmt.Println(files, err)
	return files, nil
}

func (bizAppHubService *BizAppHubService) ParseFileSrc(unZipPath string) (fileList []FileSrc, err error) {
	fmt.Println(unZipPath)
	unZipPath = strings.ReplaceAll(unZipPath, "\\", "/")
	refPath := unZipPath
	unZipPath = "./" + unZipPath //./soft/beiluo/json-tool/json_conv/v1.0/dist
	files, err := FileGetAll(unZipPath)
	if err != nil {
		return nil, err
	}
	var filesSrc []FileSrc
	for _, file := range files {
		filePath := strings.ReplaceAll(file, "\\", "/")
		srcPath := strings.ReplaceAll(filePath, refPath, "")
		filesSrc = append(filesSrc, FileSrc{
			SrcPath:     srcPath,
			LocalPath:   "./" + filePath,
			OssFullPath: "http://cdn.geeleo.com/" + filePath,
			OssPath:     filePath,
		})
	}

	return filesSrc, nil
}

// FileSrcReplaceAndUpload 替换文件中的地址用oss的地址
func (bizAppHubService *BizAppHubService) FileSrcReplaceAndUpload(fileList []FileSrc) error {
	fileReplace := make(map[string]string)
	for _, file := range fileList {
		fileReplace[file.SrcPath] = file.OssFullPath
	}

	for _, src := range fileList {
		fileBytes, err := os.ReadFile(src.LocalPath)
		if err != nil {
			return err
		}
		fileContent := string(fileBytes)
		for k, v := range fileReplace { //把文件中的本地地址替换成oss地址
			fileContent = strings.ReplaceAll(fileContent, k, v)
		}

		err = os.Remove(src.LocalPath)
		if err != nil {
			return err
		}
		create, err := os.Create(src.LocalPath)
		if err != nil {
			return err
		}
		_, err = create.Write([]byte(fileContent))
		if err != nil {
			return err
		}
		create.Close()
	}

	return nil
}

func (bizAppHubService *BizAppHubService) Deploy(req biz_apphub.BizAppHub) (index string, err error) {

	absPath := "./soft"
	absPath = strings.Join([]string{absPath, req.OperateUser, req.AppCode, req.Version}, "/")
	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		return "", err
	}
	url := "http://cdn.geeleo.com/" + req.OssPath

	path := strings.Split(req.OssPath, "/")
	fileName := path[len(path)-1]
	err = httpx.DownloadFile(url, absPath+"/"+fileName)
	if err != nil {
		return "", err
	}
	unZipPath, err := compress.DeCompress(filepath.Join(absPath, fileName), absPath)
	if err != nil {
		return "", err
	}

	fileList, err := bizAppHubService.ParseFileSrc(unZipPath)
	if err != nil {
		return "", err
	}

	fmt.Println(fileList)
	err = bizAppHubService.FileSrcReplaceAndUpload(fileList)
	if err != nil {
		return "", err
	}

	store := oss.NewDefaultQiNiu()
	for _, file := range fileList {
		if strings.HasSuffix(file.OssPath, "/dist/index.html") { //上传入口文件
			replace := fmt.Sprintf("/%s/dist", req.Version)
			indexFile := strings.ReplaceAll(file.OssPath, replace, "")
			indexFile = strings.ReplaceAll(indexFile, "/index.html", "/index")
			ossIndexFile := strings.ReplaceAll(file.OssFullPath, replace, "")
			ossIndexFile = strings.ReplaceAll(ossIndexFile, "/index.html", "/index")

			index = ossIndexFile
			err = store.DeleteFile(indexFile)
			if err != nil {
				fmt.Println(err)
			}
			_, err := store.UploadLocalFile(file.LocalPath, indexFile)
			if err != nil {
				return "", err
			}
		}
		_, err := store.UploadLocalFile(file.LocalPath, file.OssPath)
		if err != nil {
			return "", err
		}
	}

	return index, nil
}

// GetDeployList 发布历史
func (bizAppHubService *BizAppHubService) GetDeployList(req request.GetDeployList) (res []biz_apphub.BizAppHubRecord, totalCount int64, err error) {
	db := global.GVA_DB.Model(&biz_apphub.BizAppHubRecord{}).Where("app_id = ?", req.AppId)
	err = db.Offset(req.GetOffset()).Limit(req.GetLimit()).Order("created_at desc").Find(&res).Error
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

// RollbackVersion 版本回滚
func (bizAppHubService *BizAppHubService) RollbackVersion(req request.RollbackVersion) (err error) {
	var record biz_apphub.BizAppHubRecord
	var app biz_apphub.BizAppHub
	err = global.GVA_DB.Model(&biz_apphub.BizAppHubRecord{}).Where("id =? ", req.RecordID).First(&record).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Model(&biz_apphub.BizAppHub{}).Where("id =? ", req.Appid).First(&app).Error
	if err != nil {
		return err
	}

	tempPath := fmt.Sprintf("./temp/%s/%s/", req.OperateUser, app.AppCode)
	err = os.MkdirAll(tempPath, 0755)
	if err != nil {
		return err
	}
	file := tempPath + "index.html"
	//下载index文件
	url := "http://cdn.geeleo.com/" + record.GetOssStorePath()
	err = httpx.DownloadFile(url, file)
	if err != nil {
		return err
	}
	defer os.ReadFile(file)

	store := oss.NewDefaultQiNiu()
	deleteOss := app.GetUseOssPathV2()
	err = store.DeleteFile(deleteOss)
	if err != nil {
		fmt.Println(err)
	}

	_, err = store.UploadLocalFile(file, deleteOss) //版本回滚后，更新app版本
	if err != nil {
		return err
	}

	ups := map[string]interface{}{
		"app_name":     record.AppName,
		"title":        record.Title,
		"desc":         record.Desc,
		"classify":     record.Classify,
		"version":      record.Version,
		"mode":         record.Mode,
		"develop_mode": record.DevelopMode,
		"cover":        record.Cover,
		"tags":         record.Tags,
		"video":        record.Video,
	}

	err = global.GVA_DB.Model(&biz_apphub.BizAppHub{}).Where("id = ?", req.Appid).Updates(ups).Error
	if err != nil {
		return err
	}
	return nil
}
