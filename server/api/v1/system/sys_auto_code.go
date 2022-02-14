package system

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

// PreviewTemp
// @Tags AutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "预览创建代码"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "预览创建后的代码"
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
		global.GVA_LOG.Error(global.Translate("sys_auto_code.previewFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.previewFailErr"), c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, global.Translate("sys_auto_code.previewSuccess"), c)
	}
}

// CreateTemp
// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"general.createSuccss"}"
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
			global.GVA_LOG.Error(global.Translate("sys_auto_code.autoCreateFail"), zap.Error(err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape(global.Translate("sys_auto_code.autoCreateFail")))
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

// GetDB
// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取当前所有数据库"
// @Router /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(c *gin.Context) {
	dbs, err := autoCodeService.Database().GetDB()
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	}
	response.OkWithDetailed(gin.H{"dbs": dbs}, global.Translate("general.getDataSuccess"), c)
}

// GetTables
// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取当前数据库所有表"
// @Router /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tables, err := autoCodeService.Database().GetTables(dbName)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_auto_code.queryTablesFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.queryTablesFail"), c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, global.Translate("general.getDataSuccess"), c)
	}
}

// GetColumn
// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取当前表所有字段"
// @Router /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database().GetColumn(tableName, dbName)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	}
	response.OkWithDetailed(gin.H{"columns": columns}, global.Translate("general.getDataSuccess"), c)
}
