package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils/upload"
	"mime/multipart"
	"strings"
)

// @title    Upload
// @description   创建文件上传记录
// @param     file            model.ExaFileUploadAndDownload
// @auth                     （2020/04/05  20:22）
// @return                    error

func Upload(file model.ExaFileUploadAndDownload) error {
	err := global.GVA_DB.Create(&file).Error
	return err
}

// @title    FindFile
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error

func FindFile(id uint) (error, model.ExaFileUploadAndDownload) {
	var file model.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}

// @title    DeleteFile
// @description   删除文件记录
// @auth                     （2020/04/05  20:22）
// @param     file            model.ExaFileUploadAndDownload
// @return                    error

func DeleteFile(file model.ExaFileUploadAndDownload) (err error) {
	var fileFromDb model.ExaFileUploadAndDownload
	err, fileFromDb = FindFile(file.ID)
	if err = upload.Oss.DeleteFile(fileFromDb.Key); err != nil{
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(file).Error
	return err
}

// @title    GetFileRecordInfoList
// @description   分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return    err             error
// @return    list            error
// @return    total           error

func GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var fileLists []model.ExaFileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

// @title    UploadFile
// @description   根据配置文件判断是文件上传到本地或者七牛云
// @auth                     （2020/04/05  20:22）
// @param     header          *multipart.FileHeader
// @param     noSave          string
// @return    err             error
// @return    file            file model.ExaFileUploadAndDownload

func UploadFile(header *multipart.FileHeader, noSave string) (err error, file model.ExaFileUploadAndDownload) {
	filePath, key, uploadErr := upload.Oss.Upload(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return Upload(f), f
	}
	return
}
