package upload

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (url, key string, err error)
	DeleteFile(ctx context.Context, key string) error
	// DeleteFiles 批量删除，返回逐个失败项；整体 err 仅用于鉴权失败等致命错误。
	DeleteFiles(ctx context.Context, keys []string) (failed []DeleteFailure, err error)
	// ListFiles 按前缀列举存储对象，cursor 为不透明分页游标（实现内部映射到各 SDK 的 marker/token）。
	ListFiles(ctx context.Context, prefix, cursor string, limit int) (files []FileInfo, nextCursor string, hasMore bool, err error)
	// Exists 检查对象是否存在，"不存在"统一降级为 (false, nil)。
	Exists(ctx context.Context, key string) (bool, error)
}

// FileInfo 跨 SDK 统一的文件元信息
type FileInfo struct {
	Key          string    // 存储对象 key
	Size         int64     // 字节大小
	LastModified time.Time // 最后修改时间
	ContentType  string    // MIME 类型（部分实现可能为空）
}

// DeleteFailure 批量删除中单个失败项
type DeleteFailure struct {
	Key string
	Err error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	case "cloudflare-r2":
		return &CloudflareR2{}
	case "minio":
		return &Minio{}
	default:
		return &Local{}
	}
}
