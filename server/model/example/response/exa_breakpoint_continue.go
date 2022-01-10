package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath" comment:"文件路径"`
}

type FileResponse struct {
	File example.ExaFile `json:"file" comment:"文件详情"`
}
