package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string  `json:"fileName"`
	FileMd5      string	 `json:"fileMd5"`
	FilePath     string	 `json:"filePath"`
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
