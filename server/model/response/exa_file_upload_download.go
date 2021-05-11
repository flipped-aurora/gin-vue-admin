package response

import "github.com/eyotang/gin-vue-admin/server/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
