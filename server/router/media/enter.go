package media

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	AttachmentCategoryRouter
	FileUploadAndDownloadRouter
	MediaUploadRouter
}

var (
	attachmentCategoryApi    = api.ApiGroupApp.MediaApiGroup.AttachmentCategoryApi
	fileUploadAndDownloadApi = api.ApiGroupApp.MediaApiGroup.FileUploadAndDownloadApi
	mediaUploadApi           = api.ApiGroupApp.MediaApiGroup.MediaUploadApi
)
