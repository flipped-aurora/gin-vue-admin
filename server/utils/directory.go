package utils

import (
	"os"
	"path/filepath"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    createDir
// @description   批量创建文件夹
// @auth                     （2020/04/05  20:22）
// @param     dirs            string
// @return    err             error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			//log.L.Info(fmt.Sprintf("get dir error![%v]\n", err))
			return err
		}
		if exist {
			//log.L.Info(fmt.Sprintf("has dir![%v]\n"+_dir))
		} else {
			//log.L.Info(fmt.Sprintf("no dir![%v]\n"+_dir))
			// 创建文件夹
			err = os.Mkdir(v, os.ModePerm)
			if err != nil {
				//log.L.Error(fmt.Sprintf("mkdir error![%v]\n",err))
			} else {
				//log.L.Info("mkdir success!\n")
			}
		}
	}
	return err
}

// @title cwd
// @description 获取当前工作目录
// @return string

func CWD() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}
