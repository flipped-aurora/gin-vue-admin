package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/media"

type FileUploadAndDownloadResponse struct {
	File media.FileUploadAndDownload `json:"file"`
}
