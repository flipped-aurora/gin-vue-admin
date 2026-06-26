package media

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"gorm.io/gorm"
)

type FileUploadAndDownloadService struct{}

var FileUploadAndDownloadServiceApp = new(FileUploadAndDownloadService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.FileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file media.FileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.FileUploadAndDownload, error

func (e *FileUploadAndDownloadService) FindFile(id uint) (media.FileUploadAndDownload, error) {
	var file media.FileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.FileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(file media.FileUploadAndDownload) (err error) {
	var fileFromDb media.FileUploadAndDownload
	fileFromDb, err = e.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (e *FileUploadAndDownloadService) EditFileName(file media.FileUploadAndDownload) (err error) {
	var fileFromDb media.FileUploadAndDownload
	return global.GVA_DB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.AttachmentCategorySearch
//@return: list interface{}, total int64, err error

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.AttachmentCategorySearch) (list []media.FileUploadAndDownload, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&media.FileUploadAndDownload{})

	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
	}

	if info.ClassId > 0 {
		db = db.Where("class_id = ?", info.ClassId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error
	return list, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: file model.FileUploadAndDownload, err error

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string, classId int) (file media.FileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return file, uploadErr
	}
	s := strings.Split(header.Filename, ".")
	f := media.FileUploadAndDownload{
		Url:     filePath,
		Name:    header.Filename,
		ClassId: classId,
		Tag:     s[len(s)-1],
		Key:     key,
	}
	if noSave == "0" {
		// 检查是否已存在相同key的记录
		var existingFile media.FileUploadAndDownload
		err = global.GVA_DB.Where(&media.FileUploadAndDownload{Key: key}).First(&existingFile).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return f, e.Upload(f)
		}
		return f, err
	}
	return f, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ImportURL
//@description: 导入URL
//@param: file model.FileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) ImportURL(file *[]media.FileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}
