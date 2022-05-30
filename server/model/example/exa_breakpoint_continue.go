package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExaFile 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// TableName 实例断点续传表
func (e *ExaFile) TableName() string {
	return "exa_files"
}

// ExaFileChunk 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}

// TableName 实例断点续传切片信息表
func (e *ExaFileChunk) TableName() string {
	return "exa_file_chunks"
}
