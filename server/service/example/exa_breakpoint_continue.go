package example

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"gorm.io/gorm"
)

type FileUploadAndDownloadService struct{}

var FileUploadAndDownloadServiceApp = new(FileUploadAndDownloadService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindOrCreateFile
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
//@param: fileMd5 string, fileName string, chunkTotal int
//@return: file model.ExaFile, err error

func (e *FileUploadAndDownloadService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file example.ExaFile, err error) {
	var cfile example.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal
	// 检查是否存在已完成上传的文件（is_finish = true）
	// 同时使用 fileMd5 和 fileName 作为查询条件，确保文件唯一性
	// 原因：虽然 MD5 碰撞概率极低，但理论上仍可能发生（不同内容产生相同 MD5）
	// 同时使用 MD5 和文件名可以避免误匹配到错误的文件
	// 使用 errors.Is 判断是否为记录不存在的错误，这是 GORM 的标准做法
	// 好处：可以区分"记录不存在"和"其他数据库错误"，逻辑更清晰
	if errors.Is(global.GVA_DB.Where("file_md5 = ? AND file_name = ? AND is_finish = ?", fileMd5, fileName, true).First(&file).Error, gorm.ErrRecordNotFound) {
		// 不存在已完成的文件，查找或创建未完成的文件记录
		// Preload("ExaFileChunk") 预加载关联的切片记录
		// 好处：一次查询获取文件信息和已上传切片列表，减少数据库查询次数
		// FirstOrCreate 如果找到则返回，找不到则创建，原子操作保证并发安全
		err = global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return file, err
	}

	cfile.IsFinish = true
	cfile.FilePath = file.FilePath
	err = global.GVA_DB.Create(&cfile).Error
	return cfile, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFileChunk
//@description: 创建文件切片记录
//@param: id uint, fileChunkPath string, fileChunkNumber int
//@return: error
·
func (e *FileUploadAndDownloadService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk example.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.GVA_DB.Create(&chunk).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除文件切片记录
//@param: fileMd5 string, fileName string, filePath string
//@return: error

func (e *FileUploadAndDownloadService) DeleteFileChunk(fileMd5 string, filePath string) error {
	var chunks []example.ExaFileChunk
	var file example.ExaFile
	err := global.GVA_DB.Where("file_md5 = ?", fileMd5).First(&file).
		Updates(map[string]interface{}{
			"IsFinish":  true,
			"file_path": filePath,
		}).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
