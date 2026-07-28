package system

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestLogViewerApiRejectsMissingQuery(t *testing.T) {
	root := useAPILogRoot(t)
	_ = root
	router := newLogViewerTestRouter()

	for _, path := range []string{"/dates", "/files", "/content"} {
		t.Run(path, func(t *testing.T) {
			response := performLogViewerRequest(t, router, path)
			require.Equal(t, float64(7), response["code"])
			require.Equal(t, "参数错误", response["msg"])
		})
	}
}

func TestLogViewerApiReturnsDates(t *testing.T) {
	root := useAPILogRoot(t)
	writeAPILog(t, root, "2026-07-27/info.log", "ready\n")
	router := newLogViewerTestRouter()

	response := performLogViewerRequest(t, router, "/dates?month=2026-07")
	require.Equal(t, float64(0), response["code"])
	data := response["data"].(map[string]any)
	require.Equal(t, "2026-07", data["month"])
	dates := data["dates"].([]any)
	require.Len(t, dates, 1)
	require.Equal(t, "2026-07-27", dates[0].(map[string]any)["date"])
	require.Equal(t, float64(1), dates[0].(map[string]any)["fileCount"])
}

func TestLogViewerApiReturnsFilesAndContent(t *testing.T) {
	root := useAPILogRoot(t)
	writeAPILog(t, root, "2026-07-27/mongo/info.log", "first\nsecond\n")
	router := newLogViewerTestRouter()

	filesResponse := performLogViewerRequest(t, router, "/files?date=2026-07-27")
	require.Equal(t, float64(0), filesResponse["code"])
	filesData := filesResponse["data"].(map[string]any)
	files := filesData["files"].([]any)
	require.Len(t, files, 1)
	require.Equal(t, "mongo/info.log", files[0].(map[string]any)["path"])

	query := url.Values{"date": {"2026-07-27"}, "path": {"mongo/info.log"}}
	contentResponse := performLogViewerRequest(t, router, "/content?"+query.Encode())
	require.Equal(t, float64(0), contentResponse["code"])
	contentData := contentResponse["data"].(map[string]any)
	require.Equal(t, "first\nsecond\n", contentData["content"])
	require.Equal(t, float64(2), contentData["lineCount"])
	require.Equal(t, false, contentData["hasMore"])
}

func TestLogViewerApiDoesNotLeakAbsolutePath(t *testing.T) {
	root := useAPILogRoot(t)
	writeAPILog(t, root, "2026-07-27/info.log", "ready\n")
	router := newLogViewerTestRouter()

	query := url.Values{
		"date": {"2026-07-27"},
		"path": {"../secret.log"},
	}
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/content?"+query.Encode(), nil)
	router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.NotContains(t, recorder.Body.String(), root)
	var response map[string]any
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &response))
	require.Equal(t, float64(7), response["code"])
	require.Equal(t, "日志文件路径不合法", response["msg"])
}

func newLogViewerTestRouter() *gin.Engine {
	testutil.InitNopLogger()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	api := &LogViewerApi{}
	router.GET("/dates", api.GetLogDates)
	router.GET("/files", api.GetLogFiles)
	router.GET("/content", api.GetLogContent)
	return router
}

func performLogViewerRequest(t *testing.T, router *gin.Engine, path string) map[string]any {
	t.Helper()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, path, nil)
	router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	var response map[string]any
	require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &response))
	return response
}

func useAPILogRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	previous := global.GVA_CONFIG.Zap.Director
	global.GVA_CONFIG.Zap.Director = root
	t.Cleanup(func() {
		global.GVA_CONFIG.Zap.Director = previous
	})
	return root
}

func writeAPILog(t *testing.T, root, relativePath, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
}
