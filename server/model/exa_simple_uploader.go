package model

type ExaSimpleUploader struct {
	ChunkNumber      string `json:"chunkNumber" gorm:"comment:当前切片标记"`
	CurrentChunkSize string `json:"currentChunkSize" gorm:"comment:当前切片容量"`
	CurrentChunkPath string `json:"currentChunkPath" gorm:"comment:切片本地路径"`
	TotalSize        string `json:"totalSize" gorm:"comment:总容量"`
	Identifier       string `json:"identifier" gorm:"comment:文件标识（md5）"`
	Filename         string `json:"filename" gorm:"comment:文件名"`
	TotalChunks      string `json:"totalChunks" gorm:"comment:切片总数"`
	IsDone           bool   `json:"isDone" gorm:"comment:是否上传完成"`
	FilePath         string `json:"filePath" gorm:"comment:文件本地路径"`
}
