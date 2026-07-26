package utils

import (
	"errors"
	"go/token"
	"path/filepath"
	"regexp"
	"strings"
)

var pluginNamePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

// ValidatePluginName restricts plugin names to lowercase ASCII Go identifiers.
func ValidatePluginName(name string) error {
	if !pluginNamePattern.MatchString(name) || token.IsKeyword(name) {
		return errors.New("invalid plugin name")
	}
	return nil
}

// JoinWithinRoot joins path elements while ensuring the result stays below root.
func JoinWithinRoot(root string, elems ...string) (string, error) {
	if strings.TrimSpace(root) == "" {
		return "", errors.New("path root is empty")
	}
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return "", errors.New("failed to resolve path root")
	}
	for _, elem := range elems {
		if filepath.IsAbs(elem) || filepath.VolumeName(elem) != "" || strings.HasPrefix(elem, "/") || strings.HasPrefix(elem, `\`) {
			return "", errors.New("path escapes root")
		}
	}
	parts := append([]string{rootAbs}, elems...)
	targetAbs, err := filepath.Abs(filepath.Join(parts...))
	if err != nil {
		return "", errors.New("failed to resolve target path")
	}
	rel, err := filepath.Rel(rootAbs, targetAbs)
	if err != nil || filepath.IsAbs(rel) || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", errors.New("path escapes root")
	}
	return targetAbs, nil
}
