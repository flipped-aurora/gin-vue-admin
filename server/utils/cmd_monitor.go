package utils

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Watch: 监控对象
type Watch struct {
	*fsnotify.Watcher
}

func NewWatch() *Watch {
	obj, _ := fsnotify.NewWatcher()
	return &Watch{obj}
}

// Watch: 监控对象
func (w *Watch) Watch(path string) error {
	// 先转化为绝对路径
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	// 判断是单个文件还是目录文件
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	// 判断是否是目录 添加监控
	if fileInfo.IsDir() {
		// dir
		err = w.watchDir(path)

	} else {
		err = w.watchFile(path)

	}
	if err != nil {
		return err
	}
	c := make(chan error)
	// 启动监控
	go func() {
		for {
			select {
			case even, ok := <-w.Events:
				if !ok {
					// close
					fmt.Println("Errors close")
					c <- errors.New("errors close")
					return
				}
				// 判断时间
				fmt.Println("even", even)
				switch {
				// todo 待处理
				case even.Op&fsnotify.Create == fsnotify.Create:
					//这里获取新创建文件的信息，如果是目录，则加入监控中
					fmt.Println("创建文件 : ", even.Name)
					_ = w.Add(even.Name)
				case even.Op&fsnotify.Write == fsnotify.Write:
					fmt.Println("修改 : ", even.Name)
				case even.Op&fsnotify.Remove == fsnotify.Remove:
					fmt.Println("删除 : ", even.Name)
					_ = w.Remove(even.Name)
				case even.Op&fsnotify.Rename == fsnotify.Rename:
					fmt.Println("重命名 : ", even.Name)
					_ = w.Remove(even.Name)
				}
			case err = <-w.Errors:
				c <- err
				return
			}
		}
	}()

	return <-c
}

// watchDir: 处理监控目录
func (w *Watch) watchDir(path string) error {
	// 先将自己添加到监控
	err := w.Add(path)

	if err != nil {
		return err
	}
	fileSlice, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range fileSlice {
		fPath := filepath.Join(path, f.Name())
		if !f.IsDir() {
			// todo 这里加一些条件筛选添加监控的内容
			err = w.watchFile(fPath)
			if err != nil {
				return err
			}
		} else {
			err := w.watchDir(fPath)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// watchDir: 处理监控单文件
func (w *Watch) watchFile(path string) error {
	return w.Add(path)
}
