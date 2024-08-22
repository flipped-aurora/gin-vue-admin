package compress

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// DeCompress 解压文件 返回解压的目录
func DeCompress(zipFile, dest string) (absPath string, err error) {
	// 打开zip文件
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}

	defer func() {
		err := reader.Close()
		if err != nil {
			log.Fatalf("[unzip]: close reader %s", err.Error())
		}
	}()

	var (
		//first string // 记录第一次的解压的名字
		order int = 0
	)

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return "", err
		}
		filename := filepath.Join(dest, file.Name)
		//记录第一次的名字
		if order == 0 {
			//first = filename
		}
		order += 1
		//fmt.Println(getDir(filename))
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(filename, 0755)
			if err != nil {
				return "", err
			}
		} else {
			w, err := os.Create(filename)
			if err != nil {
				return "", err
			}
			//defer w.Close()
			_, err = io.Copy(w, rc)
			if err != nil {
				return "", err
			}
			iErr := w.Close()
			if iErr != nil {
				log.Fatalf("[unzip]: close io %s", iErr.Error())
			}
			fErr := rc.Close()
			if fErr != nil {
				log.Fatalf("[unzip]: close io %s", fErr.Error())
			}
		}
	}
	return strings.TrimSuffix(zipFile, ".zip"), nil
}
