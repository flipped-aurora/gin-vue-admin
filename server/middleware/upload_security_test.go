package middleware

import (
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadResponseHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(UploadResponseHeaders("/uploads/file"))
	engine.GET("/*path", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	tests := []struct {
		path       string
		nosniff    bool
		attachment bool
	}{
		{path: "/uploads/file/photo.PNG", nosniff: true},
		{path: "/uploads/file/movie.mp4", nosniff: true},
		{path: "/uploads/file/manual.pdf", nosniff: true, attachment: true},
		{path: "/uploads/file/legacy.html", nosniff: true, attachment: true},
		{path: "/uploads/file/vector.svg", nosniff: true, attachment: true},
		{path: "/uploads/files/legacy.html"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			resp := httptest.NewRecorder()
			engine.ServeHTTP(resp, req)

			gotNosniff := resp.Header().Get("X-Content-Type-Options") == "nosniff"
			if gotNosniff != tt.nosniff {
				t.Errorf("nosniff = %v, want %v", gotNosniff, tt.nosniff)
			}
			disposition := resp.Header().Get("Content-Disposition")
			gotAttachment := false
			if disposition != "" {
				mediaType, _, err := mime.ParseMediaType(disposition)
				if err != nil {
					t.Fatalf("invalid Content-Disposition %q: %v", disposition, err)
				}
				gotAttachment = mediaType == "attachment"
			}
			if gotAttachment != tt.attachment {
				t.Errorf("attachment = %v, want %v (header %q)", gotAttachment, tt.attachment, disposition)
			}
		})
	}
}
