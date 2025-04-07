package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/goccy/go-json"
	"io"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

// GetDB
// @Tags      AutoCode
// @Summary   获取当前所有数据库
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前所有数据库"
// @Router    /autoCode/getDB [get]
func (autoApi *AutoCodeApi) GetDB(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbs, err := autoCodeService.Database(businessDB).GetDB(businessDB)
	var dbList []map[string]interface{}
	for _, db := range global.GVA_CONFIG.DBList {
		var item = make(map[string]interface{})
		item["aliasName"] = db.AliasName
		item["dbName"] = db.Dbname
		item["disable"] = db.Disable
		item["dbtype"] = db.Type
		dbList = append(dbList, item)
	}
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs, "dbList": dbList}, "获取成功", c)
	}
}

// GetTables
// @Tags      AutoCode
// @Summary   获取当前数据库所有表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前数据库所有表"
// @Router    /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.Query("dbName")
	businessDB := c.Query("businessDB")
	if dbName == "" {
		dbName = *global.GVA_ACTIVE_DBNAME
		if businessDB != "" {
			for _, db := range global.GVA_CONFIG.DBList {
				if db.AliasName == businessDB {
					dbName = db.Dbname
				}
			}
		}
	}

	tables, err := autoCodeService.Database(businessDB).GetTables(businessDB, dbName)
	if err != nil {
		global.GVA_LOG.Error("查询table失败!", zap.Error(err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// GetColumn
// @Tags      AutoCode
// @Summary   获取当前表所有字段
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前表所有字段"
// @Router    /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbName := c.Query("dbName")
	if dbName == "" {
		dbName = *global.GVA_ACTIVE_DBNAME
		if businessDB != "" {
			for _, db := range global.GVA_CONFIG.DBList {
				if db.AliasName == businessDB {
					dbName = db.Dbname
				}
			}
		}
	}
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database(businessDB).GetColumn(businessDB, tableName, dbName)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}

func (autoApi *AutoCodeApi) LLMAuto(c *gin.Context) {
	var llm common.JSONMap
	err := c.ShouldBindJSON(&llm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if global.GVA_CONFIG.AutoCode.AiPath == "" {
		response.FailWithMessage("请先前往插件市场个人中心获取AiPath并填入config.yaml中", c)
		return
	}

	path := strings.ReplaceAll(global.GVA_CONFIG.AutoCode.AiPath, "{FUNC}", fmt.Sprintf("api/chat/%s", llm["mode"]))
	res, err := request.HttpRequest(
		path,
		"POST",
		nil,
		nil,
		llm,
	)
	if err != nil {
		global.GVA_LOG.Error("大模型生成失败!", zap.Error(err))
		response.FailWithMessage("大模型生成失败"+err.Error(), c)
		return
	}
	var resStruct response.Response
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		global.GVA_LOG.Error("大模型生成失败!", zap.Error(err))
		response.FailWithMessage("大模型生成失败"+err.Error(), c)
		return
	}
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		global.GVA_LOG.Error("大模型生成失败!", zap.Error(err))
		response.FailWithMessage("大模型生成失败"+err.Error(), c)
		return
	}

	if resStruct.Code == 7 {
		global.GVA_LOG.Error("大模型生成失败!"+resStruct.Msg, zap.Error(err))
		response.FailWithMessage("大模型生成失败"+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
}
