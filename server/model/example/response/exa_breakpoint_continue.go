package response

import "github.com/flipped-aurora/gin-vue-admin/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.ExaFile `json:"file"`
}
