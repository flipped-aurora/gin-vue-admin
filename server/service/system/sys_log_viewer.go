package system

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	pathpkg "path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

const (
	DefaultLogChunkLines = 500
	MaxLogChunkBytes     = 2 * 1024 * 1024
	logReadBlockSize     = 64 * 1024
)

var (
	ErrInvalidLogMonth    = errors.New("日志月份格式不正确")
	ErrInvalidLogDate     = errors.New("日志日期格式不正确")
	ErrInvalidLogPath     = errors.New("日志文件路径不合法")
	ErrLogFileNotFound    = errors.New("日志文件不存在")
	ErrLogFileUnreadable  = errors.New("日志文件不可读取")
	ErrLogRootUnavailable = errors.New("日志目录不可读取")
)

type LogViewerService struct{}

func (s *LogViewerService) ListDates(ctx context.Context, month string) (result systemRes.LogDateList, err error) {
	result = systemRes.LogDateList{Month: month, Dates: make([]systemRes.LogDateItem, 0)}
	if err = validateLogMonth(month); err != nil {
		return result, err
	}

	logRoot, exists, err := openConfiguredLogRoot()
	if err != nil {
		return result, err
	}
	if !exists {
		return result, nil
	}
	defer logRoot.Close()

	entries, err := fs.ReadDir(logRoot.FS(), ".")
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	for _, entry := range entries {
		if err = ctx.Err(); err != nil {
			return result, err
		}
		if entry.Type()&os.ModeSymlink != 0 || !entry.IsDir() || !strings.HasPrefix(entry.Name(), month+"-") {
			continue
		}
		if validateLogDate(entry.Name()) != nil {
			continue
		}
		count, countErr := countLogFiles(ctx, logRoot, entry.Name())
		if countErr != nil {
			return result, countErr
		}
		if count == 0 {
			continue
		}
		result.Dates = append(result.Dates, systemRes.LogDateItem{Date: entry.Name(), FileCount: count})
	}

	sort.Slice(result.Dates, func(i, j int) bool {
		return result.Dates[i].Date < result.Dates[j].Date
	})
	return result, nil
}

func (s *LogViewerService) ListFiles(ctx context.Context, date string) (result systemRes.LogFileList, err error) {
	result = systemRes.LogFileList{Date: date, Files: make([]systemRes.LogFileItem, 0)}
	if err = validateLogDate(date); err != nil {
		return result, err
	}

	logRoot, exists, err := openConfiguredLogRoot()
	if err != nil {
		return result, err
	}
	if !exists {
		return result, nil
	}
	defer logRoot.Close()

	info, err := logRoot.Lstat(date)
	if errors.Is(err, fs.ErrNotExist) {
		return result, nil
	}
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return result, nil
	}
	dateRoot, err := logRoot.OpenRoot(date)
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	defer dateRoot.Close()

	err = fs.WalkDir(dateRoot.FS(), ".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		if !strings.EqualFold(filepath.Ext(entry.Name()), ".log") {
			return nil
		}
		fileInfo, infoErr := entry.Info()
		if infoErr != nil {
			return infoErr
		}
		if !fileInfo.Mode().IsRegular() {
			return nil
		}
		result.Files = append(result.Files, systemRes.LogFileItem{
			Path:       path,
			Name:       entry.Name(),
			Size:       fileInfo.Size(),
			ModifiedAt: fileInfo.ModTime(),
		})
		return nil
	})
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}

	sort.Slice(result.Files, func(i, j int) bool {
		return result.Files[i].Path < result.Files[j].Path
	})
	return result, nil
}

func (s *LogViewerService) ReadContent(ctx context.Context, query systemReq.LogContentQuery) (result systemRes.LogContent, err error) {
	result.Date = query.Date
	result.Path = query.Path
	if err = ctx.Err(); err != nil {
		return result, err
	}
	if query.Cursor != nil && *query.Cursor < 0 {
		return result, ErrInvalidLogPath
	}

	file, info, err := openValidatedLogFile(query.Date, query.Path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	end := info.Size()
	if query.Cursor != nil && *query.Cursor <= end {
		end = *query.Cursor
	}
	start, limitedByBytes, err := findLogChunkStart(file, end)
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrLogFileUnreadable, err)
	}

	data, err := readLogRange(file, start, end)
	if err != nil {
		return result, fmt.Errorf("%w: %v", ErrLogFileUnreadable, err)
	}
	if int64(len(data)) < end-start {
		currentInfo, statErr := file.Stat()
		if statErr != nil || currentInfo.Size() < start {
			if statErr != nil {
				return result, fmt.Errorf("%w: %v", ErrLogFileUnreadable, statErr)
			}
			return result, ErrLogFileUnreadable
		}
		info = currentInfo
	}

	result.Content = string(data)
	result.LineCount = countLogicalLines(data)
	result.NextCursor = start
	result.HasMore = start > 0
	result.LimitedByBytes = limitedByBytes
	result.Size = info.Size()
	result.ModifiedAt = info.ModTime()
	return result, nil
}

func readLogRange(file *os.File, start, end int64) ([]byte, error) {
	if end <= start {
		return []byte{}, nil
	}
	data := make([]byte, int(end-start))
	n, err := file.ReadAt(data, start)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	return data[:n], nil
}

func configuredLogRoot() (root string, exists bool, err error) {
	director := global.GVA_CONFIG.Zap.Director
	if strings.TrimSpace(director) == "" {
		return "", false, ErrLogRootUnavailable
	}
	root, err = filepath.Abs(director)
	if err != nil {
		return "", false, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	info, err := os.Stat(root)
	if errors.Is(err, fs.ErrNotExist) {
		return root, false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	if !info.IsDir() {
		return "", false, ErrLogRootUnavailable
	}
	return root, true, nil
}

func openConfiguredLogRoot() (root *os.Root, exists bool, err error) {
	rootPath, exists, err := configuredLogRoot()
	if err != nil || !exists {
		return nil, exists, err
	}
	root, err = os.OpenRoot(rootPath)
	if err != nil {
		return nil, false, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	return root, true, nil
}

func validateLogMonth(month string) error {
	parsed, err := time.Parse("2006-01", month)
	if err != nil || parsed.Format("2006-01") != month {
		return ErrInvalidLogMonth
	}
	return nil
}

func validateLogDate(date string) error {
	parsed, err := time.Parse("2006-01-02", date)
	if err != nil || parsed.Format("2006-01-02") != date {
		return ErrInvalidLogDate
	}
	return nil
}

func countLogFiles(ctx context.Context, logRoot *os.Root, date string) (count int, err error) {
	dateInfo, err := logRoot.Lstat(date)
	if err != nil || dateInfo.Mode()&os.ModeSymlink != 0 || !dateInfo.IsDir() {
		return 0, err
	}
	dateRoot, err := logRoot.OpenRoot(date)
	if err != nil {
		return 0, err
	}
	defer dateRoot.Close()

	err = fs.WalkDir(dateRoot.FS(), ".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), ".log") {
			return nil
		}
		info, infoErr := entry.Info()
		if infoErr != nil {
			return infoErr
		}
		if info.Mode().IsRegular() {
			count++
		}
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrLogRootUnavailable, err)
	}
	return count, nil
}

func openValidatedLogFile(date, apiPath string) (file *os.File, info os.FileInfo, err error) {
	if err = validateLogDate(date); err != nil {
		return nil, nil, err
	}
	segments, err := validateLogAPIPath(apiPath)
	if err != nil {
		return nil, nil, err
	}
	logRoot, exists, err := openConfiguredLogRoot()
	if err != nil {
		return nil, nil, err
	}
	if !exists {
		return nil, nil, ErrLogFileNotFound
	}
	defer logRoot.Close()

	dateInfo, err := logRoot.Lstat(date)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil, ErrLogFileNotFound
	}
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrLogFileUnreadable, err)
	}
	if dateInfo.Mode()&os.ModeSymlink != 0 || !dateInfo.IsDir() {
		return nil, nil, ErrInvalidLogPath
	}

	dateRoot, err := logRoot.OpenRoot(date)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrLogFileUnreadable, err)
	}
	defer dateRoot.Close()

	relativePath := ""
	var validatedInfo os.FileInfo
	for index, segment := range segments {
		relativePath = filepath.Join(relativePath, segment)
		validatedInfo, err = dateRoot.Lstat(relativePath)
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil, ErrLogFileNotFound
		}
		if err != nil {
			return nil, nil, fmt.Errorf("%w: %v", ErrLogFileUnreadable, err)
		}
		if validatedInfo.Mode()&os.ModeSymlink != 0 {
			return nil, nil, ErrInvalidLogPath
		}
		if index < len(segments)-1 && !validatedInfo.IsDir() {
			return nil, nil, ErrInvalidLogPath
		}
	}
	if validatedInfo == nil || !validatedInfo.Mode().IsRegular() {
		return nil, nil, ErrInvalidLogPath
	}

	file, err = dateRoot.Open(relativePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil, ErrLogFileNotFound
		}
		return nil, nil, fmt.Errorf("%w: %v", ErrLogFileUnreadable, err)
	}
	openedInfo, statErr := file.Stat()
	if statErr != nil || !openedInfo.Mode().IsRegular() || !os.SameFile(validatedInfo, openedInfo) {
		file.Close()
		if statErr != nil {
			return nil, nil, fmt.Errorf("%w: %v", ErrLogFileUnreadable, statErr)
		}
		return nil, nil, ErrInvalidLogPath
	}
	return file, openedInfo, nil
}

func validateLogAPIPath(apiPath string) ([]string, error) {
	if apiPath == "" || strings.Contains(apiPath, "\\") || strings.Contains(apiPath, ":") || pathpkg.IsAbs(apiPath) {
		return nil, ErrInvalidLogPath
	}
	segments := strings.Split(apiPath, "/")
	for _, segment := range segments {
		if segment == "" || segment == "." || segment == ".." {
			return nil, ErrInvalidLogPath
		}
	}
	if !strings.EqualFold(pathpkg.Ext(apiPath), ".log") {
		return nil, ErrInvalidLogPath
	}
	return segments, nil
}

func findLogChunkStart(file *os.File, end int64) (start int64, limitedByBytes bool, err error) {
	if end <= 0 {
		return 0, false, nil
	}

	lastByte := []byte{0}
	if _, err = file.ReadAt(lastByte, end-1); err != nil {
		return 0, false, err
	}
	targetNewlines := DefaultLogChunkLines
	if lastByte[0] == '\n' {
		targetNewlines++
	}

	position := end
	scanned := int64(0)
	newlines := 0
	for position > 0 && scanned < MaxLogChunkBytes {
		readSize := int64(logReadBlockSize)
		if readSize > position {
			readSize = position
		}
		if remaining := int64(MaxLogChunkBytes) - scanned; readSize > remaining {
			readSize = remaining
		}
		blockStart := position - readSize
		block := make([]byte, int(readSize))
		n, readErr := file.ReadAt(block, blockStart)
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			return 0, false, readErr
		}
		for index := n - 1; index >= 0; index-- {
			if block[index] != '\n' {
				continue
			}
			newlines++
			if newlines == targetNewlines {
				return blockStart + int64(index) + 1, false, nil
			}
		}
		scanned += int64(n)
		position = blockStart
	}
	if position > 0 {
		return end - int64(MaxLogChunkBytes), true, nil
	}
	return 0, false, nil
}

func countLogicalLines(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	count := bytes.Count(data, []byte{'\n'})
	if data[len(data)-1] != '\n' {
		count++
	}
	return count
}
