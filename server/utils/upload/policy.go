package upload

import (
	"errors"
	"path/filepath"
	"strings"
)

var allowedUploadExtensions = map[string]struct{}{
	".jpg": {}, ".jpeg": {}, ".png": {}, ".gif": {}, ".webp": {}, ".bmp": {}, ".ico": {}, ".avif": {},
	".mp3": {}, ".wav": {}, ".ogg": {}, ".m4a": {}, ".flac": {}, ".aac": {},
	".mp4": {}, ".webm": {}, ".mov": {}, ".avi": {}, ".mkv": {},
	".txt": {}, ".md": {}, ".csv": {}, ".json": {}, ".log": {}, ".pdf": {},
	".doc": {}, ".docx": {}, ".xls": {}, ".xlsx": {}, ".ppt": {}, ".pptx": {},
	".zip": {}, ".rar": {}, ".7z": {}, ".tar": {}, ".gz": {}, ".tgz": {}, ".bin": {},
}

var inlineUploadExtensions = map[string]struct{}{
	".jpg": {}, ".jpeg": {}, ".png": {}, ".gif": {}, ".webp": {}, ".bmp": {}, ".ico": {}, ".avif": {},
	".mp3": {}, ".wav": {}, ".ogg": {}, ".m4a": {}, ".flac": {}, ".aac": {},
	".mp4": {}, ".webm": {}, ".mov": {}, ".avi": {}, ".mkv": {},
}

func validatedExtension(filename string) (string, error) {
	if filename == "" || strings.TrimSpace(filename) != filename || strings.ContainsAny(filename, `/\`) {
		return "", errors.New("file extension is not allowed")
	}
	ext := strings.ToLower(filepath.Ext(filename))
	if _, ok := allowedUploadExtensions[ext]; !ok {
		return "", errors.New("file extension is not allowed")
	}
	return ext, nil
}

// ValidateFileExtension rejects active content and unknown upload types.
func ValidateFileExtension(filename string) error {
	_, err := validatedExtension(filename)
	return err
}

// CanServeInline limits inline responses to raster images and audio/video files.
func CanServeInline(filename string) bool {
	ext, err := validatedExtension(filename)
	if err != nil {
		return false
	}
	_, ok := inlineUploadExtensions[ext]
	return ok
}
