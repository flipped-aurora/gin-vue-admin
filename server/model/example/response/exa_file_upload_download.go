package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
