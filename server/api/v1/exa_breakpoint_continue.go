package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
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
		response.Result(response.SUCCESS, nil, fmt.Sprintf("%v", err), c)
	} else {
		f, err := FileHeader.Open()
		if err != nil {
			response.Result(response.ERROR, nil, fmt.Sprintf("%v", err), c)
		} else {
			cen, _ := ioutil.ReadAll(f)
			defer f.Close()
			if flag := utils.CheckMd5(cen, chunkMd5); flag {
				err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
				if err != nil {
					response.Result(response.ERROR, nil, fmt.Sprintf("%v", err), c)
				} else {
					err, pathc := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
					if err != nil {
						response.Result(response.ERROR, nil, fmt.Sprintf("%v", err), c)
					} else {
						err = service.CreateFileChunk(file.ID, pathc, chunkNumber)
						if err != nil {
							response.Result(response.ERROR, nil, fmt.Sprintf("%v", err), c)
						} else {
							response.Result(response.SUCCESS, nil, "切片创建成功", c)
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
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		response.Result(response.ERROR, nil, fmt.Sprintf("查找失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"file": file}, "查找成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		response.Result(response.ERROR, gin.H{"filePath": filePath}, fmt.Sprintf("文件创建失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"filePath": filePath}, "文件创建成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func RemoveChunk(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		response.Result(response.ERROR, gin.H{"filePath": filePath}, fmt.Sprintf("缓存切片删除失败：%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{"filePath": filePath}, "缓存切片删除成功", c)
	}
}
