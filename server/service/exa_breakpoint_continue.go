package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// @title         FindOrCreateFile
// @description   Check your file if it does not exist, or return current slice of the file
// 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     ChunkTotal      int
// @return    err             error
// @return    file            ExaFile
func  FindOrCreateFile(FileMd5 string, FileName string, ChunkTotal int) (err error, file model.ExaFile) {
	var cfile model.ExaFile
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


// @title    CreateFileChunk
// @description   create a chunk of the file, 创建文件切片记录
// @auth                       （2020/04/05  20:22 ）
// @param     FileChunkPath     string
// @param     FileChunkNumber   int
// @return                      error
func CreateFileChunk(id uint, FileChunkPath string, FileChunkNumber int) error {
	var chunk model.ExaFileChunk
	chunk.FileChunkPath = FileChunkPath
	chunk.ExaFileId = id
	chunk.FileChunkNumber = FileChunkNumber
	err := global.GVA_DB.Create(&chunk).Error
	return err
}

// @title         FileCreateComplete
// @description   file creation, 文件合成完成
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func FileCreateComplete(FileMd5 string, FileName string, FilePath string) error {
	var file model.ExaFile
	upDateFile := make(map[string]interface{})
	upDateFile["FilePath"] = FilePath
	upDateFile["IsFinish"] = true
	err := global.GVA_DB.Where("file_md5 = ? AND file_name = ?", FileMd5, FileName).First(&file).Updates(upDateFile).Error
	return err
}

// @title    DeleteFileChunk
// @description   delete a chuck of the file, 删除文件切片记录
// @auth                     （2020/04/05  20:22 ）
// @param     FileMd5         string
// @param     FileName        string
// @param     FilePath        string
// @return                    error
func DeleteFileChunk(fileMd5 string, fileName string, filePath string) error {
	var chunks []model.ExaFileChunk
	var file model.ExaFile
	err := global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Update("IsFinish", true).Update("file_path", filePath).Error
	err = global.GVA_DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
