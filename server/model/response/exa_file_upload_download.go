package response

import "github.com/eyotang/game-proxy/server/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
