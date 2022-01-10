package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string `json:"fileName" example:"test"`
	FileMd5      string `json:"fileMd5" example:"D45547AD2A4695DA2871D5C73FBA6555"`
	FilePath     string `json:"filePath" example:"xxx/xxx"`
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int  `example:"10"`
	IsFinish     bool `example:"true"`
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint   `example:"1"`
	FileChunkNumber int    `example:"2"`
	FileChunkPath   string `example:"xxx/xxx"`
}
