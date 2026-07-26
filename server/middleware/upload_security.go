package middleware

import (
	"mime"
	"path"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
)

// UploadResponseHeaders prevents uploaded files from being MIME-sniffed or
// executing inline unless they are known-safe raster or media formats.
func UploadResponseHeaders(publicPath string) gin.HandlerFunc {
	prefix := "/" + strings.Trim(strings.ReplaceAll(publicPath, `\`, "/"), "/")
	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		if requestPath == prefix || strings.HasPrefix(requestPath, prefix+"/") {
			c.Header("X-Content-Type-Options", "nosniff")
			filename := path.Base(requestPath)
			if !upload.CanServeInline(filename) {
				c.Header("Content-Disposition", mime.FormatMediaType("attachment", map[string]string{"filename": filename}))
			}
		}
		c.Next()
	}
}
