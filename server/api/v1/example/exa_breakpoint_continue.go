package example

import (
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/example"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {object} response.Response{msg=string} "断点续传到服务器"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func (b *FileUploadAndDownloadApi) BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Error(err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	cen, _ := io.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.GVA_LOG.Error("检查md5失败!", zap.Error(err))
		response.FailWithMessage("检查md5失败", c)
		return
	}
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找或创建记录失败!", zap.Error(err))
		response.FailWithMessage("查找或创建记录失败", c)
		return
	}
	pathC, err := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("断点续传失败!", zap.Error(err))
		response.FailWithMessage("断点续传失败", c)
		return
	}

	if err = fileUploadAndDownloadService.CreateFileChunk(file.ID, pathC, chunkNumber); err != nil {
		global.GVA_LOG.Error("创建文件记录失败!", zap.Error(err))
		response.FailWithMessage("创建文件记录失败", c)
		return
	}
	response.OkWithMessage("切片创建成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {object} response.Response{data=exampleRes.FileResponse,msg=string} "查找文件,返回包括文件详情"
// @Router /fileUploadAndDownload/findFile [post]
func (b *FileUploadAndDownloadApi) FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败", c)
	} else {
		response.OkWithDetailed(exampleRes.FileResponse{File: file}, "查找成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {object} response.Response{data=exampleRes.FilePathResponse,msg=string} "创建文件,返回包括文件路径"
// @Router /fileUploadAndDownload/findFile [post]
func (b *FileUploadAndDownloadApi) BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath, err := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("文件创建失败!", zap.Error(err))
		response.FailWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	} else {
		response.OkWithDetailed(exampleRes.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {object} response.Response{msg=string} "删除切片"
// @Router /fileUploadAndDownload/removeChunk [post]
func (b *FileUploadAndDownloadApi) RemoveChunk(c *gin.Context) {
	var file example.ExaFile
	_ = c.ShouldBindJSON(&file)
	err := utils.RemoveChunk(file.FileMd5)
	if err != nil {
		global.GVA_LOG.Error("缓存切片删除失败!", zap.Error(err))
		return
	}
	err = fileUploadAndDownloadService.DeleteFileChunk(file.FileMd5, file.FilePath)
	if err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("缓存切片删除成功", c)
	}
}
