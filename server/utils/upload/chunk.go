package upload

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func chunkRoot() string {
	d := global.GVA_CONFIG.Media.ChunkDir
	if d == "" {
		d = "uploads/chunks"
	}
	return d
}

// ChunkDir 某次上传的暂存目录
func ChunkDir(uploadID uint) string {
	return filepath.Join(chunkRoot(), fmt.Sprintf("%d", uploadID))
}

// SaveChunkFile 写入一个分片,返回写入字节数(覆盖写,幂等)
func SaveChunkFile(uploadID uint, index int, data []byte) (int64, error) {
	dir := ChunkDir(uploadID)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return 0, err
	}
	p := filepath.Join(dir, fmt.Sprintf("%d.part", index))
	if err := os.WriteFile(p, data, 0o644); err != nil {
		return 0, err
	}
	return int64(len(data)), nil
}

// MergeChunks 按 0..total-1 顺序流式合并到 dstPath,返回成品 md5
func MergeChunks(uploadID uint, total int, dstPath string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
		return "", err
	}
	out, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	h := md5.New()
	w := io.MultiWriter(out, h)
	for i := 0; i < total; i++ {
		p := filepath.Join(ChunkDir(uploadID), fmt.Sprintf("%d.part", i))
		in, err := os.Open(p)
		if err != nil {
			return "", fmt.Errorf("缺少分片 %d: %w", i, err)
		}
		if _, err := io.Copy(w, in); err != nil {
			in.Close()
			return "", err
		}
		in.Close()
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// RemoveUploadDir 删除某次上传的暂存目录
func RemoveUploadDir(uploadID uint) error {
	return os.RemoveAll(ChunkDir(uploadID))
}

// ReceivedIndexes 扫描暂存目录已存在的分片索引(机会式恢复用,可选)
func ReceivedIndexes(uploadID uint) []int {
	entries, err := os.ReadDir(ChunkDir(uploadID))
	if err != nil {
		return nil
	}
	var idx []int
	for _, e := range entries {
		var i int
		if _, err := fmt.Sscanf(e.Name(), "%d.part", &i); err == nil {
			idx = append(idx, i)
		}
	}
	sort.Ints(idx)
	return idx
}

// BuildFileHeader 从本地文件构造 *multipart.FileHeader(内存安全:大文件落临时盘)。
// 返回 cleanup 用于释放临时表单。
func BuildFileHeader(localPath, fieldName, fileName string) (*multipart.FileHeader, func(), error) {
	src, err := os.Open(localPath)
	if err != nil {
		return nil, nil, err
	}
	defer src.Close()

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	part, err := mw.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, nil, err
	}
	if _, err = io.Copy(part, src); err != nil {
		return nil, nil, err
	}
	if err = mw.Close(); err != nil {
		return nil, nil, err
	}
	reader := multipart.NewReader(body, mw.Boundary())
	form, err := reader.ReadForm(1 << 20) // 1MB 以上落临时盘,避免大文件 OOM
	if err != nil {
		return nil, nil, err
	}
	files := form.File[fieldName]
	if len(files) == 0 {
		_ = form.RemoveAll()
		return nil, nil, fmt.Errorf("构造 FileHeader 失败")
	}
	return files[0], func() { _ = form.RemoveAll() }, nil
}
