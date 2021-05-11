package response

import "github.com/eyotang/game-api-admin/server/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
