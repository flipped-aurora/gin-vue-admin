package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
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

func DeleteFile(file model.ExaFileUploadAndDownload) error {
	err := global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(file).Error
	return err
}

// @title    GetFileRecordInfoList
// @description   分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return    err             error
// @return    list            error
// @return    total           error

func GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var fileLists []model.ExaFileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}
