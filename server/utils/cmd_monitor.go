package utils

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"path/filepath"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@struct_name: Watch
//@description: 监控对象

type Watch struct {
	*fsnotify.Watcher
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: NewWatch
//@description: Watch的实例化方法
//@return: *Watch

func NewWatch() *Watch {
	obj, _ := fsnotify.NewWatcher()
	return &Watch{obj}
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@object: w *Watch
//@function: Watch
//@description: 监控对象
//@param: path string, t *T
//@return: error

func (w *Watch) Watch(path string, t *T) error {
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
				// 判断事件
				switch {
				case even.Op&fsnotify.Create == fsnotify.Create:
					//这里获取新创建文件的信息，如果是目录，则加入监控中
					fmt.Println("创建文件 : ", even.Name)
					//t.AddTask()
					_ = w.Add(even.Name)
				case even.Op&fsnotify.Write == fsnotify.Write:
					fmt.Println("修改文件 : ", even.Name)
					w.addTask(t, even.Name)
				case even.Op&fsnotify.Remove == fsnotify.Remove || even.Op&fsnotify.Rename == fsnotify.Rename:
					fmt.Println("删除或重命名文件 : ", even.Name)
					_ = w.Remove(even.Name)
					w.addTask(t, even.Name)
				}
			case err = <-w.Errors:
				fmt.Println("even Error:", err)
				c <- err
				return
			}
		}
	}()
	return <-c

}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@object: w *Watch
//@function: watchDir
//@description: 处理监控目录
//@param: path string
//@return: error

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
			// 判断是否可监控的文件
			if chickPower(fPath) {
				err = w.watchFile(fPath)
				if err != nil {
					return err
				}
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

//@author: [songzhibin97](https://github.com/songzhibin97)
//@object: w *Watch
//@function: watchDir
//@description: 处理监控单文件
//@param: path string
//@return: error

func (w *Watch) watchFile(path string) error {
	var err error
	if chickPower(path) {
		err = w.Add(path)
	}
	return err
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: chickPower
//@description: 判断是否在可控范围内
//@param: path string
//@return: error

func chickPower(name string) bool {
	name = filepath.Ext(name)
	return name == ".go" || name == ".yaml"
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@object: w *Watch
//@function: addTask
//@description: 偏函数 简化发送任务
//@param: path string
//@return: error

func (w *Watch) addTask(t *T, name string) {
	if chickPower(name) {
		fmt.Println("Add Task->>>>>>")
		t.AddTask()
	}
}
