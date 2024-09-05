package osx

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	err = destination.Sync()
	if err != nil {
		return err
	}

	return nil
}

// CopyDirectory 复制目录，1 源目录存在的目录（文件），目标目录不存在，会copy文件到目标目录，2 目标目录存在，源目录不存在，会保留目标目录的文件
// srcDir 源目录
// dstDir 目标目录
func CopyDirectory(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			err = os.MkdirAll(dstPath, info.Mode())
			if err != nil && !os.IsExist(err) {
				return err
			}
		} else {
			err = copyFile(path, dstPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
