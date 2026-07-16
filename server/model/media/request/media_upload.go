package request

type UploadInitReq struct {
	FileName   string `json:"fileName" binding:"required"`
	FileHash   string `json:"fileHash" binding:"required"`
	FileSize   int64  `json:"fileSize" binding:"required"`
	ChunkSize  int64  `json:"chunkSize" binding:"required"`
	ChunkTotal int    `json:"chunkTotal" binding:"required"`
	Mime       string `json:"mime"`
}

type UploadCompleteReq struct {
	UploadID uint `json:"uploadId" binding:"required"`
}
