package stringsx

import (
	"strings"
)

// GetSuffix 获取后缀名 GetSuffix("1.txt",".") -> txt
func GetSuffix(s string, stp string) string {
	// 找到最后一个 '.' 的位置
	lastDotIndex := strings.LastIndex(s, stp)
	if lastDotIndex == -1 {
		return ""
	}

	// 获取 '.' 后面的字符
	extension := s[lastDotIndex+1:]
	return extension
}
