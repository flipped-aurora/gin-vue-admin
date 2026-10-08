package example

import (
	"errors"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	exampleService "github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

// CreateScanUploadToken 创建短时、一次性的扫码上传凭证
// @Tags      ExaFileUploadAndDownload
// @Summary   创建扫码上传凭证
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object{classId=int}  true  "文件分类 ID"
// @Success   200   {object}  response.Response{data=object{token=string,expiresAt=int64},msg=string}
// @Router    /fileUploadAndDownload/createScanUploadToken [post]
func (b *FileUploadAndDownloadApi) CreateScanUploadToken(c *gin.Context) {
	var req struct {
		ClassID int `json:"classId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.ClassID < 0 {
		response.FailWithMessage("classId 参数错误", c)
		return
	}
	token, expiresAt, err := fileUploadAndDownloadService.IssueScanUploadToken(req.ClassID)
	if err != nil {
		global.GVA_LOG.Error("创建扫码上传凭证失败!", zap.Error(err))
		response.FailWithMessage("创建扫码上传凭证失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"token":     token,
		"expiresAt": expiresAt.Unix(),
	}, "创建成功", c)
}

// UploadFileByScanToken 使用一次性扫码凭证上传文件
// @Tags      ExaFileUploadAndDownload
// @Summary   使用扫码上传凭证上传文件
// @accept    multipart/form-data
// @Produce   application/json
// @Param     x-upload-token  header    string  true  "扫码上传凭证"
// @Param     file            formData  file    true  "上传文件"
// @Success   200             {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}
// @Router    /fileUploadAndDownload/uploadByScanToken [post]
func (b *FileUploadAndDownloadApi) UploadFileByScanToken(c *gin.Context) {
	classID, err := fileUploadAndDownloadService.ConsumeScanUploadToken(c.GetHeader("x-upload-token"))
	if err != nil {
		if !errors.Is(err, exampleService.ErrInvalidScanUploadToken) {
			global.GVA_LOG.Error("读取扫码上传凭证失败!", zap.Error(err))
		}
		response.NoAuth("扫码上传凭证无效或已过期，请重新扫码", c)
		return
	}
	b.uploadFile(c, classID)
}

// UploadFile
// @Tags      ExaFileUploadAndDownload
// @Summary   上传文件示例
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	classId, _ := strconv.Atoi(c.DefaultPostForm("classId", "0"))
	b.uploadFile(c, classId)
}

func (b *FileUploadAndDownloadApi) uploadFile(c *gin.Context, classId int) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave, classId) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("上传文件失败!", zap.Error(err))
		response.FailWithMessage("上传文件失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.EditFileName(file)
	if err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// DeleteFile
// @Tags      ExaFileUploadAndDownload
// @Summary   删除文件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      example.ExaFileUploadAndDownload  true  "传入文件里面id即可"
// @Success   200   {object}  response.Response{msg=string}     "删除文件"
// @Router    /fileUploadAndDownload/deleteFile [post]
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetFileList
// @Tags      ExaFileUploadAndDownload
// @Summary   分页文件列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.ExaAttachmentCategorySearch                                        true  "页码, 每页大小, 分类id"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router    /fileUploadAndDownload/getFileList [post]
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.ExaAttachmentCategorySearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
// @Tags      ExaFileUploadAndDownload
// @Summary   导入URL
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      example.ExaFileUploadAndDownload  true  "对象"
// @Success   200   {object}  response.Response{msg=string}     "导入URL"
// @Router    /fileUploadAndDownload/importURL [post]
func (b *FileUploadAndDownloadApi) ImportURL(c *gin.Context) {
	var file []example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.ImportURL(&file); err != nil {
		global.GVA_LOG.Error("导入URL失败!", zap.Error(err))
		response.FailWithMessage("导入URL失败", c)
		return
	}
	response.OkWithMessage("导入URL成功", c)
}
