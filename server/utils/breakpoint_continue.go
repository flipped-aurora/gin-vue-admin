package utils

import (
	"io/ioutil"
	"os"
	"strconv"
)

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

const breakpointDir = "./breakpointDir/"
const finishDir = "./fileDir/"

func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (error, string) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err, path
	}
	err, pathc := makeFileContent(content, fileName, path, contentNumber)
	return err, pathc

}

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // "可以继续上传"
	} else {
		return false // "切片不完整，废弃"
	}
}

func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (error, string) {
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err, path
	} else {
		_, err = f.Write(content)
		if err != nil {
			return err, path
		}
	}
	return nil, path
}

func MakeFile(fileName string, FileMd5 string) (error, string) {
	rd, err := ioutil.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return err, finishDir + fileName
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, _ := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	for k := range rd {
		content, _ := ioutil.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return err, finishDir + fileName
		}
	}
	defer fd.Close()
	return nil, finishDir + fileName
}

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
