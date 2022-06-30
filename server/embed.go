package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	//go:embed resource
	//go:embed config.yaml
	//go:embed config.docker.yaml
	resource embed.FS
)

var Embed = new(_embed)

type _embed struct{}

// RestoreFolder
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [WangLeonard](https://github.com/WangLeonard)
func (e *_embed) RestoreFolder(dir string) {
	entries, err := resource.ReadDir(dir)
	if err != nil {
		fmt.Println("[embed restore resource file]: err:", err)
		return
	}
	for i := 0; i < len(entries); i++ {
		if entries[i].IsDir() {
			e.RestoreFile(filepath.Join(dir, entries[i].Name()), entries[i])
			continue
		}
		e.RestoreFile(entries[i].Name(), entries[i])
	}
}

// RestoreFile
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [WangLeonard](https://github.com/WangLeonard)
func (e *_embed) RestoreFile(path string, entry fs.DirEntry) {
	_, err := os.Stat(path)
	if entry.IsDir() { // 文件夹
		if os.IsNotExist(err) { // 判断 path 变量的文件夹存在, 不存在则创建文件夹
			fmt.Printf("[embed restore mkdir] dir:%s\n", path)
			err = os.Mkdir(path, os.ModePerm) // 创建文件夹, 权限为 os.ModePerm 可自行修改
			if err != nil {
				fmt.Printf("[embed restore mkdir] err:%v\n", err)
				return
			}
		}
		var entries []fs.DirEntry
		entries, err = resource.ReadDir(path) // 读取文件夹的文件和文件夹数据
		if err != nil {
			return
		}
		for i := 0; i < len(entries); i++ {
			_, err = os.Stat(entries[i].Name()) // 获取子文件夹的信息
			dirPath := filepath.Join(path, entries[i].Name())
			if os.IsNotExist(err) && entries[i].IsDir() { // 判断子文件夹是否存在, 这里有可能是文件,所以要加上是否为文件夹
				fmt.Println("[embed restore mkdir] dir:", dirPath)
				err = os.Mkdir(dirPath, os.ModePerm) // 创建文件夹, 权限为 os.ModePerm 可自行修改
				if err != nil {
					fmt.Println("[embed restore mkdir] err:", err)
					return
				}
			}
			e.RestoreFile(dirPath, entries[i])
		}
	}

	if os.IsNotExist(err) && !entry.IsDir() { // 文件
		var src fs.File
		src, err = resource.Open(path) // 根据path从embed的到文件数据
		if err != nil {
			fmt.Println("[embed restore resource open file] open embed file failed, err:", err)
			return
		}
		var dst *os.File
		dst, err = os.Create(path) // 创建本地文件的 writer
		if err != nil {
			fmt.Println("[embed restore os create file] write err:", err)
			return
		}
		_, err = io.Copy(dst, src) // 把embed的数据复制到本地
		if err != nil {
			fmt.Println("[embed restore io copy file] writer file failed, err:", err)
			return
		}
		defer func() { // 关闭文件流
			_ = src.Close()
			_ = dst.Close()
		}()
		return
	}
	fmt.Println("[embed restore resource file] file exist, path:", path)
}
