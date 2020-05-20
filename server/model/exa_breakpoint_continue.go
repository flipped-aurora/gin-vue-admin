package model

import (
	"github.com/jinzhu/gorm"
)

// file struct, 文件结构体
type ExaFile struct {
	gorm.Model
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	gorm.Model
	ExaFileId       uint
	FileChunkNumber int
	FileChunkPath   string
}
