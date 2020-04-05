package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
)

type ExaFileUploadAndDownload struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Tag  string `json:"tag"`
	Key  string `json:"key"`
}

// @title    Upload
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (f *ExaFileUploadAndDownload) Upload() error {
	err := global.GVA_DB.Create(f).Error
	return err
}

// @title    DeleteFile
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (f *ExaFileUploadAndDownload) DeleteFile() error {
	err := global.GVA_DB.Where("id = ?", f.ID).Unscoped().Delete(f).Error
	return err
}

// @title    FindFile
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (f *ExaFileUploadAndDownload) FindFile() (error, ExaFileUploadAndDownload) {
	var file ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", f.ID).First(&file).Error
	return err, file
}

// @title    GetInfoList
// @description   分页获取数据
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func (f *ExaFileUploadAndDownload) GetInfoList(info PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	if err != nil {
		return
	} else {
		var fileLists []ExaFileUploadAndDownload
		err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
		return err, fileLists, total
	}
}
