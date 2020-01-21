package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/dbModel"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]

func BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
	} else {
		f, err := FileHeader.Open()
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
		} else {
			cen, _ := ioutil.ReadAll(f)
			defer f.Close()
			if flag := servers.CheckMd5(cen, chunkMd5); flag {
				err, file := new(dbModel.ExaFile).FindOrCreateFile(fileMd5, fileName, chunkTotal)
				if err != nil {
					servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
				} else {
					err, pathc := servers.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
					if err != nil {
						servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
					} else {
						err = file.CreateFileChunk(pathc, chunkNumber)
						if err != nil {
							servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
						} else {
							servers.ReportFormat(c, true, "切片创建成功", gin.H{})
						}
					}
				}
			} else {
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file params file true "查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	err, file := new(dbModel.ExaFile).FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查找失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查找成功", gin.H{"file": file})
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file params file true "查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := servers.MakeFile(fileName, fileMd5)
	if err != nil {
		servers.ReportFormat(c, true, fmt.Sprintf("文件创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "文件创建成功", gin.H{"filePath": filePath})
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file params file true "查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func RemoveChunk(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := servers.RemoveChunk(fileMd5)
	err = new(dbModel.ExaFile).DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		servers.ReportFormat(c, true, fmt.Sprintf("缓存切片删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "缓存切片删除成功", gin.H{})
	}
}
