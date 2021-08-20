//+build !windows

package plugin

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
)

// LoadPlugin 加载插件 传入path
func LoadPlugin(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		fileSlice, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, ff := range fileSlice {
			if !ff.IsDir() && filepath.Ext(ff.Name()) == ".so" {
				if err = loadPlugin(path, ff); err != nil {
					return err
				}
			} else if ff.IsDir() {
				_ = LoadPlugin(filepath.Join(path, ff.Name()))
			}
		}
		return nil
	} else {
		return loadPlugin(path, fileInfo)
	}
}

func loadPlugin(path string, f fs.FileInfo) error {
	if filepath.Ext(f.Name()) == ".so" {
		fPath := filepath.Join(path, f.Name())
		// 加载插件
		p, err := plugin.Open(fPath)
		if err != nil {
			fmt.Println("loadPlugin err ", err)
			return err
		}
		// 判断是否满足协议
		// 要满足 OnlyFuncName && 实现 Plugin 接口
		if v, err := p.Lookup(OnlyFuncName); err != nil {
			fmt.Println("loadPlugin err ", err)
			return err
		} else if _, ok := v.(Plugin); !ok {
			fmt.Println("loadPlugin err ", fmt.Sprintf("path:%s 没有实现 %s 接口", filepath.Base(fPath), OnlyFuncName))
			return errors.New("没有实现指定接口")
		} else {

		}
		fmt.Println("loadPlugin add ", filepath.Base(fPath))
		ManagementPlugin.SetPlugin(filepath.Base(fPath), p)
	}
	return nil
}
