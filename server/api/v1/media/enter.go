package media

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AttachmentCategoryApi
	FileUploadAndDownloadApi
	MediaUploadApi
}

var (
	attachmentCategoryService    = service.ServiceGroupApp.MediaServiceGroup.AttachmentCategoryService
	fileUploadAndDownloadService = service.ServiceGroupApp.MediaServiceGroup.FileUploadAndDownloadService
	mediaUploadService           = service.ServiceGroupApp.MediaServiceGroup.MediaUploadService
)
