package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
)

type LogViewerApi struct{}

// GetLogDates
// @Tags      LogViewer
// @Summary   获取指定月份存在日志的日期
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     month  query     string  true  "月份(YYYY-MM)"
// @Success   200    {object}  response.Response{data=systemRes.LogDateList,msg=string}  "返回日志日期和文件数量"
// @Router    /logViewer/dates [get]
func (l *LogViewerApi) GetLogDates(c *gin.Context) {
	var query systemReq.LogMonthQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	var data systemRes.LogDateList
	data, err := logViewerService.ListDates(c.Request.Context(), query.Month)
	if err != nil {
		failLogViewerRequest(c, err, "获取日志日期失败")
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetLogFiles
// @Tags      LogViewer
// @Summary   获取指定日期的日志文件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     date  query     string  true  "日期(YYYY-MM-DD)"
// @Success   200   {object}  response.Response{data=systemRes.LogFileList,msg=string}  "返回递归日志文件列表"
// @Router    /logViewer/files [get]
func (l *LogViewerApi) GetLogFiles(c *gin.Context) {
	var query systemReq.LogDateQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	var data systemRes.LogFileList
	data, err := logViewerService.ListFiles(c.Request.Context(), query.Date)
	if err != nil {
		failLogViewerRequest(c, err, "获取日志文件失败")
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetLogContent
// @Tags      LogViewer
// @Summary   分块读取日志文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     date    query     string  true   "日期(YYYY-MM-DD)"
// @Param     path    query     string  true   "日期目录内的日志相对路径"
// @Param     cursor  query     int     false  "向前读取的字节游标"
// @Success   200     {object}  response.Response{data=systemRes.LogContent,msg=string}  "返回最多500行且不超过2MiB的日志内容"
// @Router    /logViewer/content [get]
func (l *LogViewerApi) GetLogContent(c *gin.Context) {
	var query systemReq.LogContentQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	var data systemRes.LogContent
	data, err := logViewerService.ReadContent(c.Request.Context(), query)
	if err != nil {
		failLogViewerRequest(c, err, "读取日志内容失败")
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func failLogViewerRequest(c *gin.Context, err error, logMessage string) {
	logger.WithCtx(c.Request.Context()).Mod("log-viewer").Err(err).Error(logMessage)
	message := "读取日志失败"
	switch {
	case errors.Is(err, systemService.ErrInvalidLogMonth):
		message = "日志月份格式不正确"
	case errors.Is(err, systemService.ErrInvalidLogDate):
		message = "日志日期格式不正确"
	case errors.Is(err, systemService.ErrInvalidLogPath):
		message = "日志文件路径不合法"
	case errors.Is(err, systemService.ErrLogFileNotFound):
		message = "日志文件不存在"
	case errors.Is(err, systemService.ErrLogFileUnreadable):
		message = "日志文件不可读取"
	case errors.Is(err, systemService.ErrLogRootUnavailable):
		message = "日志目录不可读取"
	}
	response.FailWithMessage(message, c)
}
