package media

import "github.com/flipped-aurora/gin-vue-admin/server/global"

const (
	UploadStatusUploading = "uploading"
	UploadStatusMerging   = "merging"
	UploadStatusCompleted = "completed"
	UploadStatusFailed    = "failed"
)

// MediaUpload 大文件上传会话
type MediaUpload struct {
	global.GVA_MODEL
	UserID     uint   `json:"userId" gorm:"index;column:user_id;comment:上传者"`
	FileName   string `json:"fileName" gorm:"column:file_name;comment:文件名"`
	FileHash   string `json:"fileHash" gorm:"index;column:file_hash;comment:整文件MD5"`
	FileSize   int64  `json:"fileSize" gorm:"column:file_size;comment:总字节"`
	ChunkSize  int64  `json:"chunkSize" gorm:"column:chunk_size;comment:分片字节"`
	ChunkTotal int    `json:"chunkTotal" gorm:"column:chunk_total;comment:分片总数"`
	Status     string `json:"status" gorm:"column:status;default:uploading;comment:状态"`
	StorageKey string `json:"storageKey" gorm:"column:storage_key;comment:最终对象key"`
	MediaID    uint   `json:"mediaId" gorm:"column:media_id;comment:媒体库记录ID"`
}

func (MediaUpload) TableName() string { return "media_uploads" }

// MediaUploadChunk 分片收讫记录
type MediaUploadChunk struct {
	global.GVA_MODEL
	UploadID   uint   `json:"uploadId" gorm:"uniqueIndex:idx_upload_chunk;column:upload_id"`
	ChunkIndex int    `json:"chunkIndex" gorm:"uniqueIndex:idx_upload_chunk;column:chunk_index"`
	ChunkHash  string `json:"chunkHash" gorm:"column:chunk_hash"`
	Size       int64  `json:"size" gorm:"column:size"`
}

func (MediaUploadChunk) TableName() string { return "media_upload_chunks" }
