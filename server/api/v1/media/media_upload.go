package media

import (
	"io"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	mediaRes "github.com/flipped-aurora/gin-vue-admin/server/model/media/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type MediaUploadApi struct{}

// Init 初始化上传(秒传/续传探测)
// @Tags     MediaUpload
// @Summary  初始化大文件上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body request.UploadInitReq true "初始化参数"
// @Success  200 {object} response.Response{data=mediaRes.UploadInitResp}
// @Router   /mediaUpload/init [post]
func (a *MediaUploadApi) Init(c *gin.Context) {
	var req request.UploadInitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if max := global.GVA_CONFIG.Media.MaxFileSize; max > 0 && req.FileSize > max {
		response.FailWithMessage("文件超过大小上限", c)
		return
	}
	res, err := mediaUploadService.Init(c.Request.Context(), utils.GetUserID(c), req)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("init失败")
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// Chunk 上传一个分片
// @Tags     MediaUpload
// @Summary  上传分片
// @Security ApiKeyAuth
// @accept   multipart/form-data
// @Produce  application/json
// @Router   /mediaUpload/chunk [post]
func (a *MediaUploadApi) Chunk(c *gin.Context) {
	uploadID, _ := strconv.Atoi(c.PostForm("uploadId"))
	index, _ := strconv.Atoi(c.PostForm("chunkIndex"))
	chunkHash := c.PostForm("chunkHash")
	fh, err := c.FormFile("chunk")
	if err != nil {
		response.FailWithMessage("接收分片失败", c)
		return
	}
	f, err := fh.Open()
	if err != nil {
		response.FailWithMessage("分片读取失败", c)
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		response.FailWithMessage("分片读取失败", c)
		return
	}
	if err := mediaUploadService.SaveChunk(c.Request.Context(), utils.GetUserID(c), uint(uploadID), index, chunkHash, data); err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("收片失败")
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("ok", c)
}

// Complete 完成合并
// @Tags     MediaUpload
// @Summary  完成上传并合并入库
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data body request.UploadCompleteReq true "完成参数"
// @Success  200 {object} response.Response{data=mediaRes.UploadCompleteResp}
// @Router   /mediaUpload/complete [post]
func (a *MediaUploadApi) Complete(c *gin.Context) {
	var req request.UploadCompleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	m, err := mediaUploadService.Complete(c.Request.Context(), utils.GetUserID(c), req.UploadID)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("合并失败")
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(mediaRes.UploadCompleteResp{Media: m}, c)
}

// Cancel 取消上传
// @Tags     MediaUpload
// @Summary  取消上传
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    uploadId path int true "uploadId"
// @Router   /mediaUpload/{uploadId} [delete]
func (a *MediaUploadApi) Cancel(c *gin.Context) {
	uploadID, _ := strconv.Atoi(c.Param("uploadId"))
	if err := mediaUploadService.Cancel(c.Request.Context(), utils.GetUserID(c), uint(uploadID)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("已取消", c)
}
