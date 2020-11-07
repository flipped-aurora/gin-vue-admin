package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// fileMove: 文件移动
// src: 源位置 需要传入绝对路径
// dst: 目标位置 需要传入绝对路径
func fileMove(src string, dst string) error {
	if !filepath.IsAbs(dst) && !filepath.IsAbs(src) {
		return errors.New(fmt.Sprintf("%s or %s path is not abs", dst, src))
	}
	return os.Rename(src, dst)
}

// FileMove: 文件移动供外部调用
// src: 源位置 绝对路径相对路径都可以
// dst: 目标位置 绝对路径相对路径都可以 dst 必须为文件夹
func FileMove(src string, dst string) error {
	var err error
	// 转化为绝对路径
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	// 判断传入的是否是目录
	oSrc, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !oSrc.IsDir() {
		return errors.New(fmt.Sprintf("%s is not Dir", src))
	}
	oDst, err := os.Stat(dst)
	if err != nil {
		return err
	}
	if !oDst.IsDir() {
		return errors.New(fmt.Sprintf("%s is not Dir", dst))
	}
	// 遍历指定目录下所有文件
	f, err := ioutil.ReadDir(src)
	for _, file := range f {
		nDst := filepath.Join(dst, file.Name())
		nSrc := filepath.Join(src, file.Name())
		err = fileMove(nSrc, nDst)
		if err != nil {
			return err
		}
	}
	return err
}

// GetAutoPath 根据生成路径生成移动路径
func GetAutoPath(path string) string {
	return filepath.Base(filepath.Dir(path))
}
