package system

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/stretchr/testify/require"
)

func TestLogViewerListDatesAndFiles(t *testing.T) {
	root := useLogRoot(t)
	writeTestLog(t, root, "2026-07-25/info.log", "info\n")
	writeTestLog(t, root, "2026-07-26/error.log", "error\n")
	writeTestLog(t, root, "2026-07-27/info.log", "info\n")
	writeTestLog(t, root, "2026-07-27/mongo/info.log", "mongo\n")
	writeTestLog(t, root, "2026-07-27/readme.txt", "ignored\n")
	writeTestLog(t, root, "invalid-date/info.log", "ignored\n")
	require.NoError(t, os.MkdirAll(filepath.Join(root, "2026-07-27", "directory.log"), 0o755))

	service := &LogViewerService{}
	dates, err := service.ListDates(context.Background(), "2026-07")
	require.NoError(t, err)
	require.Equal(t, "2026-07", dates.Month)
	require.Equal(t, []systemRes.LogDateItem{
		{Date: "2026-07-25", FileCount: 1},
		{Date: "2026-07-26", FileCount: 1},
		{Date: "2026-07-27", FileCount: 2},
	}, dates.Dates)

	files, err := service.ListFiles(context.Background(), "2026-07-27")
	require.NoError(t, err)
	require.Equal(t, "2026-07-27", files.Date)
	require.Equal(t, []string{"info.log", "mongo/info.log"}, logFilePaths(files.Files))
	require.Equal(t, int64(len("info\n")), files.Files[0].Size)
	require.False(t, files.Files[0].ModifiedAt.IsZero())
}

func TestLogViewerListDatesSkipsDirectoriesWithoutReadableLogs(t *testing.T) {
	root := useLogRoot(t)
	require.NoError(t, os.MkdirAll(filepath.Join(root, "2026-07-27"), 0o755))
	writeTestLog(t, root, "2026-07-27/readme.txt", "ignored\n")
	service := &LogViewerService{}

	dates, err := service.ListDates(context.Background(), "2026-07")
	require.NoError(t, err)
	require.Empty(t, dates.Dates)
}

func TestLogViewerListRejectsInvalidDatesAndHandlesMissingRoot(t *testing.T) {
	root := useLogRoot(t)
	service := &LogViewerService{}

	_, err := service.ListDates(context.Background(), "2026-13")
	require.ErrorIs(t, err, ErrInvalidLogMonth)
	_, err = service.ListFiles(context.Background(), "2026-02-30")
	require.ErrorIs(t, err, ErrInvalidLogDate)

	global.GVA_CONFIG.Zap.Director = filepath.Join(root, "missing")
	dates, err := service.ListDates(context.Background(), "2026-07")
	require.NoError(t, err)
	require.Empty(t, dates.Dates)
	files, err := service.ListFiles(context.Background(), "2026-07-27")
	require.NoError(t, err)
	require.Empty(t, files.Files)
}

func TestLogViewerReadLatestAndOlderChunks(t *testing.T) {
	root := useLogRoot(t)
	var builder strings.Builder
	for line := 1; line <= 1205; line++ {
		fmt.Fprintf(&builder, "line-%04d\n", line)
	}
	path := writeTestLog(t, root, "2026-07-27/info.log", builder.String())
	service := &LogViewerService{}

	latest, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "info.log",
	})
	require.NoError(t, err)
	require.Equal(t, 500, latest.LineCount)
	require.True(t, strings.HasPrefix(latest.Content, "line-0706\n"))
	require.True(t, strings.HasSuffix(latest.Content, "line-1205\n"))
	require.True(t, latest.HasMore)
	require.False(t, latest.LimitedByBytes)

	older, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date:   "2026-07-27",
		Path:   "info.log",
		Cursor: &latest.NextCursor,
	})
	require.NoError(t, err)
	require.Equal(t, 500, older.LineCount)
	require.True(t, strings.HasPrefix(older.Content, "line-0206\n"))
	require.True(t, strings.HasSuffix(older.Content, "line-0705\n"))
	require.NotContains(t, older.Content, "line-0706")
	require.True(t, older.HasMore)

	oldest, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date:   "2026-07-27",
		Path:   "info.log",
		Cursor: &older.NextCursor,
	})
	require.NoError(t, err)
	require.Equal(t, 205, oldest.LineCount)
	require.True(t, strings.HasPrefix(oldest.Content, "line-0001\n"))
	require.True(t, strings.HasSuffix(oldest.Content, "line-0205\n"))
	require.False(t, oldest.HasMore)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	require.NoError(t, err)
	_, err = file.WriteString("line-1206\n")
	require.NoError(t, err)
	require.NoError(t, file.Close())

	refreshed, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "info.log",
	})
	require.NoError(t, err)
	require.Contains(t, refreshed.Content, "line-1206\n")

	staleCursor := refreshed.Size + 1024
	fromStaleCursor, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date:   "2026-07-27",
		Path:   "info.log",
		Cursor: &staleCursor,
	})
	require.NoError(t, err)
	require.Equal(t, refreshed.Content, fromStaleCursor.Content)
}

func TestLogViewerReadContentLimitsBytesAndHandlesEmptyFile(t *testing.T) {
	root := useLogRoot(t)
	writeTestLog(t, root, "2026-07-27/huge.log", strings.Repeat("x", MaxLogChunkBytes+1024))
	writeTestLog(t, root, "2026-07-27/empty.log", "")
	service := &LogViewerService{}

	huge, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "huge.log",
	})
	require.NoError(t, err)
	require.Len(t, huge.Content, MaxLogChunkBytes)
	require.Equal(t, 1, huge.LineCount)
	require.True(t, huge.LimitedByBytes)
	require.True(t, huge.HasMore)

	empty, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "empty.log",
	})
	require.NoError(t, err)
	require.Empty(t, empty.Content)
	require.Zero(t, empty.LineCount)
	require.False(t, empty.HasMore)
}

func TestReadLogRangeUsesActualBytesAfterTruncate(t *testing.T) {
	path := filepath.Join(t.TempDir(), "truncated.log")
	require.NoError(t, os.WriteFile(path, []byte("short"), 0o644))
	file, err := os.Open(path)
	require.NoError(t, err)
	t.Cleanup(func() { _ = file.Close() })

	data, err := readLogRange(file, 0, 32)
	require.NoError(t, err)
	require.Equal(t, "short", string(data))
	require.NotContains(t, data, byte(0))
}

func TestLogViewerRejectsUnsafePaths(t *testing.T) {
	root := useLogRoot(t)
	writeTestLog(t, root, "2026-07-27/mongo/info.log", "safe\n")
	require.NoError(t, os.MkdirAll(filepath.Join(root, "2026-07-27", "directory.log"), 0o755))
	service := &LogViewerService{}

	unsafePaths := []string{
		"", "../secret.log", "mongo\\info.log", "/tmp/secret.log", "C:/secret.log",
		"mongo/./info.log", "mongo/../info.log", "mongo//info.log", "readme.txt", "directory.log",
	}
	for _, unsafePath := range unsafePaths {
		t.Run(strings.ReplaceAll(unsafePath, "/", "_"), func(t *testing.T) {
			_, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
				Date: "2026-07-27",
				Path: unsafePath,
			})
			require.Error(t, err)
		})
	}

	missing, err := service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "missing.log",
	})
	require.Empty(t, missing.Content)
	require.ErrorIs(t, err, ErrLogFileNotFound)

	negativeCursor := int64(-1)
	_, err = service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date:   "2026-07-27",
		Path:   "mongo/info.log",
		Cursor: &negativeCursor,
	})
	require.ErrorIs(t, err, ErrInvalidLogPath)
}

func TestLogViewerRejectsSymlinks(t *testing.T) {
	root := useLogRoot(t)
	outside := writeTestLog(t, root, "outside/secret.log", "secret\n")
	dateDir := filepath.Join(root, "2026-07-27")
	require.NoError(t, os.MkdirAll(dateDir, 0o755))
	link := filepath.Join(dateDir, "linked.log")
	if err := os.Symlink(outside, link); err != nil {
		t.Skipf("symlink creation is unavailable: %v", err)
	}

	service := &LogViewerService{}
	files, err := service.ListFiles(context.Background(), "2026-07-27")
	require.NoError(t, err)
	require.Empty(t, files.Files)

	_, err = service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "linked.log",
	})
	require.ErrorIs(t, err, ErrInvalidLogPath)
}

func TestLogViewerRejectsSymlinkedDateDirectory(t *testing.T) {
	root := useLogRoot(t)
	outsideRoot := t.TempDir()
	writeTestLog(t, outsideRoot, "secret.log", "secret\n")
	if err := os.Symlink(outsideRoot, filepath.Join(root, "2026-07-27")); err != nil {
		t.Skipf("directory symlink creation is unavailable: %v", err)
	}

	service := &LogViewerService{}
	dates, err := service.ListDates(context.Background(), "2026-07")
	require.NoError(t, err)
	require.Empty(t, dates.Dates)

	files, err := service.ListFiles(context.Background(), "2026-07-27")
	require.NoError(t, err)
	require.Empty(t, files.Files)

	_, err = service.ReadContent(context.Background(), systemReq.LogContentQuery{
		Date: "2026-07-27",
		Path: "secret.log",
	})
	require.ErrorIs(t, err, ErrInvalidLogPath)
}

func useLogRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	previous := global.GVA_CONFIG.Zap.Director
	global.GVA_CONFIG.Zap.Director = root
	t.Cleanup(func() {
		global.GVA_CONFIG.Zap.Director = previous
	})
	return root
}

func writeTestLog(t *testing.T, root, relativePath, content string) string {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
	return path
}

func logFilePaths(files []systemRes.LogFileItem) []string {
	paths := make([]string, 0, len(files))
	for _, file := range files {
		paths = append(paths, file.Path)
	}
	return paths
}
