package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// @title         FindOrCreateFile
// @description   Check your file if it does not exist, or return current slice of the file
// 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
// @auth                     （2020/04/05  20:22）
// @param     fileMd5         string
// @param     fileName        string
// @param     chunkTotal      int
// @return    err             error
// @return    file            ExaFile

func FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (err error, file model.ExaFile) {
	var cfile model.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal
	notHaveSameMd5Finish := global.GVA_DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).RecordNotFound()
	if notHaveSameMd5Finish {
		err = global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
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
// @auth                       （2020/04/05  20:22）
// @param     id                unit
// @param     fileChunkPath     string
// @param     fileChunkNumber   int
// @return                      error

func CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk model.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileId = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.GVA_DB.Create(&chunk).Error
	return err
}

// @title         FileCreateComplete
// @description   file creation, 文件合成完成
// @auth                     （2020/04/05  20:22）
// @param     fileMd5         string
// @param     fileName        string
// @param     filePath        string
// @return                    error

func FileCreateComplete(fileMd5 string, fileName string, filePath string) error {
	var file model.ExaFile
	upDateFile := make(map[string]interface{})
	upDateFile["FilePath"] = filePath
	upDateFile["IsFinish"] = true
	err := global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Updates(upDateFile).Error
	return err
}

// @title    DeleteFileChunk
// @description   delete a chuck of the file, 删除文件切片记录
// @auth                     （2020/04/05  20:22）
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
