package utils

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// DetectMIME 推断上传文件的 MIME 类型。
// 优先用 header 自带的 Content-Type；缺失时按扩展名推断；再不行读前 512 字节嗅探。
func DetectMIME(header *multipart.FileHeader) string {
	// 1. header 自带
	if ct := header.Header.Get("Content-Type"); ct != "" {
		return ct
	}
	// 2. 扩展名推断
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ct := mimeExt(ext); ct != "" {
		return ct
	}
	// 3. 嗅探前 512 字节
	f, err := header.Open()
	if err != nil {
		return "application/octet-stream"
	}
	defer f.Close()
	buf := make([]byte, 512)
	n, _ := f.Read(buf)
	return http.DetectContentType(buf[:n])
}

// mimeExt 按扩展名映射常见 MIME（仅补充 net/http 不覆盖的几个常见类型）。
func mimeExt(ext string) string {
	switch ext {
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".mp4":
		return "video/mp4"
	case ".mp3":
		return "audio/mpeg"
	case ".csv":
		return "text/csv"
	default:
		return ""
	}
}
