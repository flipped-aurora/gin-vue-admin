package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// FileMove: 文件移动供外部调用
// src: 源位置 绝对路径相对路径都可以
// dst: 目标位置 绝对路径相对路径都可以 dst 必须为文件夹
func FileMove(src string, dst string) (err error) {
	if dst == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(dst) && !filepath.IsAbs(src) {
		return errors.New(dst + " or " + src + " path is not abs")
	}
	// TODO 判断文件夹是否存在,不存在mkdir
	return os.Rename(src, dst)
}