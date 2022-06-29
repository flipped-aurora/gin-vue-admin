package utils

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type EmbedFSBatch []embed.FS

func (b EmbedFSBatch) RestoreAll() {
	for _, efs := range b {
		restoreDirectory(efs, ".")
	}
}

func NewEmbedFSBatch(es ...embed.FS) EmbedFSBatch {
	return es
}

// restoreDirectory
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [WangLeonard](https://github.com/WangLeonard)
func restoreDirectory(efs embed.FS, dir string) {
	entries, err := efs.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "[embed restore folder]: err:", err)
		return
	}
	for i := 0; i < len(entries); i++ {
		if !entries[i].IsDir() {
			restoreFile(efs, filepath.Join(dir, entries[i].Name()), entries[i])
			continue
		}

		nextDir := filepath.Join(dir, entries[i].Name())
		_, _ = fmt.Fprintf(os.Stderr, "[embed restore mkdir] dir:%s\n", nextDir)
		err = os.MkdirAll(nextDir, os.ModePerm) // 创建文件夹, 权限为 os.ModePerm 可自行修改
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[embed restore mkdir] err:%v\n", err)
			return
		}
		restoreDirectory(efs, nextDir)
	}
}

// restoreFile
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [WangLeonard](https://github.com/WangLeonard)
func restoreFile(efs embed.FS, path string, entry fs.DirEntry) {
	if entry.IsDir() { // 文件夹
		restoreDirectory(efs, path)
		return
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		var src fs.File
		src, err = efs.Open(path) // 根据path从embed的到文件数据
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[embed restore resource open file] open embed file %s failed, err: %v", path, err)
			return
		}
		_, _ = fmt.Fprintf(os.Stderr, "[embed restore create file] file:%s\n", path)
		var dst *os.File
		dst, err = os.Create(path) // 创建本地文件的 writer
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "[embed restore os create file] write err:", err)
			return
		}
		_, err = io.Copy(dst, src) // 把embed的数据复制到本地
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "[embed restore io copy file] writer file failed, err:", err)
			return
		}
		defer func() { // 关闭文件流
			_ = src.Close()
			_ = dst.Close()
		}()
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, "[embed restore resource file] file exist, path:", path)
}
