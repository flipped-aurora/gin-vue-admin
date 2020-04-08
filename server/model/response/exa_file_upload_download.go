package response

import "gin-vue-admin/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
