package response

import "github.com/gzpz/golf-sales-system/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
