package upload

import (
	"context"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
)

var mu sync.Mutex

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.GVA_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		logger.WithCtx(ctx).Mod("upload").Err(mkdirErr).Error("function os.MkdirAll() failed")
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	filepath := global.GVA_CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		logger.WithCtx(ctx).Mod("upload").Err(openError).Error("function file.Open() failed")
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		logger.WithCtx(ctx).Mod("upload").Err(createErr).Error("function os.Create() failed")

		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		logger.WithCtx(ctx).Mod("upload").Err(copyErr).Error("function io.Copy() failed")
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (l *Local) DeleteFile(ctx context.Context, key string) error {
	p, err := l.localPath(ctx, key)
	if err != nil {
		return err
	}

	// 检查文件是否存在
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return errors.New("文件不存在")
	}

	// 使用文件锁防止并发删除
	mu.Lock()
	defer mu.Unlock()

	err = os.Remove(p)
	if err != nil {
		return errors.New("文件删除失败: " + err.Error())
	}

	return nil
}

// localPath 校验 key 并拼接出本地存储的绝对路径，复用 DeleteFile 中的路径穿越防护逻辑。
func (*Local) localPath(ctx context.Context, key string) (string, error) {
	// 检查 key 是否为空
	if key == "" {
		return "", errors.New("key不能为空")
	}

	// 验证 key 是否包含非法字符或尝试访问存储路径之外的文件
	if strings.Contains(key, "..") || strings.ContainsAny(key, `\/:*?"<>|`) {
		return "", errors.New("非法的key")
	}

	return filepath.Join(global.GVA_CONFIG.Local.StorePath, key), nil
}

// Exists 检查本地文件是否存在，"不存在"统一降级为 (false, nil)。
func (l *Local) Exists(ctx context.Context, key string) (bool, error) {
	p, err := l.localPath(ctx, key)
	if err != nil {
		return false, err
	}

	info, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function os.Stat() failed")
		return false, err
	}
	if info.IsDir() {
		// 仅把普通文件视为存在的对象
		return false, nil
	}
	return true, nil
}

// DeleteFiles 批量删除本地文件，逐个调用 DeleteFile，失败的收集到返回列表。
func (l *Local) DeleteFiles(ctx context.Context, keys []string) ([]DeleteFailure, error) {
	failed := make([]DeleteFailure, 0)
	for _, key := range keys {
		if err := l.DeleteFile(ctx, key); err != nil {
			logger.WithCtx(ctx).Mod("upload").Err(err).Error("function Local.DeleteFile() failed")
			failed = append(failed, DeleteFailure{Key: key, Err: err})
		}
	}
	return failed, nil
}

// ListFiles 按前缀列举本地文件，cursor 为上次返回的最后一条文件名（不透明游标）。
func (*Local) ListFiles(ctx context.Context, prefix, cursor string, limit int) ([]FileInfo, string, bool, error) {
	if limit <= 0 {
		limit = 100
	}

	entries, err := os.ReadDir(global.GVA_CONFIG.Local.StorePath)
	if err != nil {
		logger.WithCtx(ctx).Mod("upload").Err(err).Error("function os.ReadDir() failed")
		return nil, "", false, errors.New("function os.ReadDir() failed, err:" + err.Error())
	}

	// 仅保留普通文件，按文件名过滤
	names := make([]string, 0, len(entries))
	nameToEntry := make(map[string]os.DirEntry, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if prefix != "" && !strings.HasPrefix(name, prefix) {
			continue
		}
		names = append(names, name)
		nameToEntry[name] = entry
	}

	// 按文件名排序，按 cursor 之后开始分页
	start := 0
	if cursor != "" {
		for i, name := range names {
			if name > cursor {
				start = i
				break
			}
			// 全部 <= cursor，则无可返回项
			if i == len(names)-1 {
				start = len(names)
			}
		}
	}

	end := start + limit
	if end > len(names) {
		end = len(names)
	}
	hasMore := end < len(names)

	page := names[start:end]
	files := make([]FileInfo, 0, len(page))
	for _, name := range page {
		entry := nameToEntry[name]
		info, infoErr := entry.Info()
		if infoErr != nil {
			continue
		}
		files = append(files, FileInfo{
			Key:          name,
			Size:         info.Size(),
			LastModified: info.ModTime(),
			ContentType:  mime.TypeByExtension(filepath.Ext(name)),
		})
	}

	var nextCursor string
	if hasMore && len(page) > 0 {
		nextCursor = page[len(page)-1]
	}
	return files, nextCursor, hasMore, nil
}
