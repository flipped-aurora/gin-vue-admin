package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string `gorm:"size:256;"`
	FileMd5      string `gorm:"size:256;"`
	FilePath     string `gorm:"size:256;"`
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string `gorm:"size:256;"`
}
