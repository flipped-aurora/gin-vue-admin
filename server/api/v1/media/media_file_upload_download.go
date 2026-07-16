package media

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/media/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type FileUploadAndDownloadApi struct{}

// UploadFile
// @Tags      FileUploadAndDownload
// @Summary   上传文件示例
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=exampleRes.FileUploadAndDownloadResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file media.FileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	classId, _ := strconv.Atoi(c.DefaultPostForm("classId", "0"))
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("接收文件失败!")
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(c.Request.Context(), header, noSave, classId, utils.GetUserID(c)) // 文件上传后拿到文件路径
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("上传文件失败!")
		response.FailWithMessage("上传文件失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.FileUploadAndDownloadResponse{File: file}, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
// @Tags      FileUploadAndDownload
// @Summary   编辑文件名或者备注
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      media.FileUploadAndDownload  true  "传入文件id和文件名或备注"
// @Success   200   {object}  response.Response{msg=string}     "编辑文件名或者备注"
// @Router    /fileUploadAndDownload/editFileName [post]
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file media.FileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.EditFileName(c.Request.Context(), file)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("编辑失败!")
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// DeleteFile
// @Tags      FileUploadAndDownload
// @Summary   删除文件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      media.FileUploadAndDownload  true  "传入文件里面id即可"
// @Success   200   {object}  response.Response{msg=string}     "删除文件"
// @Router    /fileUploadAndDownload/deleteFile [post]
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file media.FileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.DeleteFile(c.Request.Context(), file); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("删除失败!")
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetFileList
// @Tags      FileUploadAndDownload
// @Summary   分页文件列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AttachmentCategorySearch                                        true  "页码, 每页大小, 分类id"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router    /fileUploadAndDownload/getFileList [post]
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.AttachmentCategorySearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(c.Request.Context(), pageInfo)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取失败!")
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// ImportURL
// @Tags      FileUploadAndDownload
// @Summary   导入URL
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      []media.FileUploadAndDownload  true  "对象数组"
// @Success   200   {object}  response.Response{msg=string}       "导入URL"
// @Router    /fileUploadAndDownload/importURL [post]
func (b *FileUploadAndDownloadApi) ImportURL(c *gin.Context) {
	var file []media.FileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.ImportURL(c.Request.Context(), &file); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("导入URL失败!")
		response.FailWithMessage("导入URL失败", c)
		return
	}
	response.OkWithMessage("导入URL成功", c)
}

// DeleteFiles 批量删除文件
// @Tags      FileUploadAndDownload
// @Summary   批量删除文件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      request.FileBatchDeleteReq  true  "文件ID列表"
// @Success   200   {object}  response.Response{data=object,msg=string}  "批量删除结果"
// @Router    /fileUploadAndDownload/deleteFiles [post]
func (b *FileUploadAndDownloadApi) DeleteFiles(c *gin.Context) {
	var req request.FileBatchDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	failedIDs, err := fileUploadAndDownloadService.DeleteFiles(c.Request.Context(), req.IDs)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("批量删除失败!")
		response.FailWithMessage("批量删除失败", c)
		return
	}
	if len(failedIDs) > 0 {
		response.OkWithDetailed(gin.H{"failedIds": failedIDs}, "部分文件删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// FindFile 查找单个文件详情
// @Tags      FileUploadAndDownload
// @Summary   查找单个文件详情
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     id  query  int  true  "文件ID"
// @Success   200   {object}  response.Response{data=media.FileUploadAndDownload,msg=string}  "文件详情"
// @Router    /fileUploadAndDownload/findFile [get]
func (b *FileUploadAndDownloadApi) FindFile(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("文件ID非法", c)
		return
	}
	file, err := fileUploadAndDownloadService.GetFileDetail(c.Request.Context(), uint(id))
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("查询失败!")
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(file, c)
}

// ListOssFiles 列举 OSS 存储桶内真实文件
// @Tags      FileUploadAndDownload
// @Summary   列举存储桶文件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      request.ListOssFilesReq  true  "前缀/游标/每页数量"
// @Success   200   {object}  response.Response{data=object,msg=string}  "存储桶文件列表"
// @Router    /fileUploadAndDownload/listOssFiles [post]
func (b *FileUploadAndDownloadApi) ListOssFiles(c *gin.Context) {
	var req request.ListOssFilesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	files, nextCursor, hasMore, err := fileUploadAndDownloadService.ListOssFiles(c.Request.Context(), req.Prefix, req.Cursor, req.Limit)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("列举存储桶文件失败!")
		response.FailWithMessage("列举存储桶文件失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"list": files, "nextCursor": nextCursor, "hasMore": hasMore}, "获取成功", c)
}
