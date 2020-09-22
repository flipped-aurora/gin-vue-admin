package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strconv"
)

// 保存文件切片路径
func SaveChunk(uploader model.ExaSimpleUploader) (err error) {
	return global.GVA_DB.Create(uploader).Error
}

// 检查文件是否已经上传过
func CheckFileMd5(md5 string) (err error, uploads []model.ExaSimpleUploader, isDone bool) {
	err = global.GVA_DB.Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
	isDone = errors.Is(global.GVA_DB.First(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound)
	return err, uploads, !isDone
}

// 合并文件
func MergeFileMd5(md5 string, fileName string) (err error) {
	finishDir := "./finish/"
	dir := "./chunk/" + md5
	//如果文件上传成功 不做后续操作 通知成功即可
	if !errors.Is(global.GVA_DB.First(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound) {
		return nil
	}

	//打开切片文件夹
	rd, err := ioutil.ReadDir(dir)
	_ = os.MkdirAll(finishDir, os.ModePerm)
	//创建目标文件
	fd, _ := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(dir + "/" + fileName + strconv.Itoa(k+1))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
		}
	}
	//关闭文件
	defer fd.Close()

	if err != nil {
		return err
	}
	//创建事务
	tx := global.GVA_DB.Begin()
	//删除切片信息
	err = tx.Delete(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, false).Error
	// 添加文件信息
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	err = tx.Create(&model.ExaSimpleUploader{
		Identifier: md5,
		IsDone:     true,
		FilePath:   finishDir + fileName,
		Filename:   fileName,
	}).Error
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	tx.Commit()
	//清除切片
	err = os.RemoveAll(dir)
	return
}
