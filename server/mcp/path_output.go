package mcpTool

import (
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func outputPath(path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return ""
	}

	root := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Root)
	if root != "" {
		normalizedRoot := normalizeFilesystemPath(root)
		normalizedPath := normalizeFilesystemPath(path)

		if rel, err := filepath.Rel(normalizedRoot, normalizedPath); err == nil && !isParentTraversal(rel) {
			return filepath.ToSlash(filepath.Clean(rel))
		}
	}

	return filepath.ToSlash(filepath.Clean(path))
}

func normalizeFilesystemPath(path string) string {
	cleaned := filepath.Clean(path)
	if absPath, err := filepath.Abs(cleaned); err == nil {
		return absPath
	}
	return cleaned
}

func isParentTraversal(path string) bool {
	path = filepath.Clean(path)
	if path == ".." {
		return true
	}
	return strings.HasPrefix(path, ".."+string(filepath.Separator))
}
