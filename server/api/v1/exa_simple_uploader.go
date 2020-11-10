package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传插件版示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /simpleUploader/upload [post]
func SimpleUploaderUpload(c *gin.Context) {
	var chunk model.ExaSimpleUploader
	_, header, err := c.Request.FormFile("file")
	chunk.Filename = c.PostForm("filename")
	chunk.ChunkNumber = c.PostForm("chunkNumber")
	chunk.CurrentChunkSize = c.PostForm("currentChunkSize")
	chunk.Identifier = c.PostForm("identifier")
	chunk.TotalSize = c.PostForm("totalSize")
	chunk.TotalChunks = c.PostForm("totalChunks")
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		if err := utils.CreateDir(chunkDir); err != nil {
			global.GVA_LOG.Error("创建目录失败!", zap.Any("err", err))
		}
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err = c.SaveUploadedFile(header, chunkPath)
	if err != nil {
		global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		response.FailWithMessage("切片创建失败", c)
		return
	}
	chunk.CurrentChunkPath = chunkPath
	err = service.SaveChunk(chunk)
	if err != nil {
		global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		response.FailWithMessage("切片创建失败", c)
		return
	} else {
		response.OkWithMessage("切片创建成功", c)
	}
}

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /simpleUploader/checkFileMd5 [get]
func CheckFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	err, chunks, isDone := service.CheckFileMd5(md5)
	if err != nil {
		global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
		response.FailWithMessage("md5读取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"chunks": chunks,
			"isDone": isDone,
		},"查询成功", c)
	}
}

// @Tags SimpleUploader
// @Summary 合并文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"合并成功"}"
// @Router /simpleUploader/mergeFileMd5 [get]
func MergeFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	fileName := c.Query("fileName")
	err := service.MergeFileMd5(md5, fileName)
	if err != nil {
		global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
		response.FailWithMessage("md5读取失败", c)
	} else {
		response.OkWithMessage("合并成功", c)
	}
}
