package system

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct {
}

// @Tags AutoCode
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AutoHistoryByID true "删除回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /autoCode/delSysHistory [post]
func (autoApi *AutoCodeApi) DelSysHistory(c *gin.Context) {
	var id systemReq.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	err := autoCodeHistoryService.DeletePage(id.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithMessage("删除成功", c)

}

// @Tags AutoCode
// @Summary 查询回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SysAutoHistory true "查询回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getSysHistory [post]
func (autoApi *AutoCodeApi) GetSysHistory(c *gin.Context) {
	var search systemReq.SysAutoHistory
	_ = c.ShouldBindJSON(&search)
	err, list, total := autoCodeHistoryService.GetSysHistoryPage(search.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     search.Page,
			PageSize: search.PageSize,
		}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AutoHistoryByID true "回滚自动生成代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"回滚成功"}"
// @Router /autoCode/rollback [post]
func (autoApi *AutoCodeApi) RollBack(c *gin.Context) {
	var id systemReq.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if err := autoCodeHistoryService.RollBack(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("回滚成功", c)
}

// @Tags AutoCode
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AutoHistoryByID true "获取meta信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getMeta [post]
func (autoApi *AutoCodeApi) GetMeta(c *gin.Context) {
	var id systemReq.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if v, err := autoCodeHistoryService.GetMeta(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(gin.H{"meta": v}, "获取成功", c)
	}

}

// @Tags AutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "预览创建代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.GVA_LOG.Error("预览失败!", zap.Any("err", err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var apiIds []uint
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.GVA_LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Any("err", err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return
		} else {
			apiIds = ids
		}
	}
	err := autoCodeService.CreateTemp(a, apiIds...)
	if err != nil {
		if errors.Is(err, system.AutoMoveErr) {
			c.Writer.Header().Add("success", "true")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	err, tables := autoCodeService.GetTables(dbName)
	if err != nil {
		global.GVA_LOG.Error("查询table失败!", zap.Any("err", err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(c *gin.Context) {
	if err, dbs := autoCodeService.GetDB(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	if err, columns := autoCodeService.GetColumn(tableName, dbName); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}
