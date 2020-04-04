package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
)

//文件结构体
type ExaFile struct {
	gorm.Model
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

//切片结构体
type ExaFileChunk struct {
	gorm.Model
	ExaFileId       uint
	FileChunkNumber int
	FileChunkPath   string
}

//文件合成完成
func (f *ExaFile) FileCreateComplete(FileMd5 string, FileName string, FilePath string) error {
	var file ExaFile
	upDateFile := make(map[string]interface{})
	upDateFile["FilePath"] = FilePath
	upDateFile["IsFinish"] = true
	err := global.GVA_DB.Where("file_md5 = ? AND file_name = ?", FileMd5, FileName).First(&file).Updates(upDateFile).Error
	return err
}

//第一次上传或者断点续传时候检测当前文件属性，没有则创建，有则返回文件的当前切片
func (f *ExaFile) FindOrCreateFile(FileMd5 string, FileName string, ChunkTotal int) (err error, file ExaFile) {
	var cfile ExaFile
	cfile.FileMd5 = FileMd5
	cfile.FileName = FileName
	cfile.ChunkTotal = ChunkTotal
	notHaveSameMd5Finish := global.GVA_DB.Where("file_md5 = ? AND is_finish = ?", FileMd5, true).First(&file).RecordNotFound()
	if notHaveSameMd5Finish {
		err = global.GVA_DB.Where("file_md5 = ? AND file_name = ?", FileMd5, FileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return err, file
	} else {
		cfile.IsFinish = true
		cfile.FilePath = file.FilePath
		err = global.GVA_DB.Create(&cfile).Error
		return err, cfile
	}
}

// 创建文件切片记录
func (f *ExaFile) CreateFileChunk(FileChunkPath string, FileChunkNumber int) error {
	var chunk ExaFileChunk
	chunk.FileChunkPath = FileChunkPath
	chunk.ExaFileId = f.ID
	chunk.FileChunkNumber = FileChunkNumber
	err := global.GVA_DB.Create(&chunk).Error
	return err
}

// 删除文件切片记录
func (f *ExaFile) DeleteFileChunk(fileMd5 string, fileName string, filePath string) error {
	var chunks []ExaFileChunk
	var file ExaFile
	err := global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Update("IsFinish", true).Update("file_path", filePath).Error
	err = global.GVA_DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
