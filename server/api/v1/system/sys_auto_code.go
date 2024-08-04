package system

import (
	"io"
	"strings"

	"github.com/goccy/go-json"

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
// @Router    /autoCode/getDatabase [get]
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
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs, "dbList": dbList}, global.Translate("general.getDataSuccess"), c)
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
		dbName = global.GVA_CONFIG.Mysql.Dbname
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
		global.GVA_LOG.Error(global.Translate("sys_auto_code.queryTablesFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.queryTablesFail"), c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, global.Translate("general.getDataSuccess"), c)
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
		dbName = global.GVA_CONFIG.Mysql.Dbname
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
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, global.Translate("general.getDataSuccess"), c)
	}
}

func (autoApi *AutoCodeApi) LLMAuto(c *gin.Context) {
	prompt := c.Query("prompt")
	mode := c.Query("mode")
	params := make(map[string]string)
	params["prompt"] = prompt
	params["mode"] = mode
	path := strings.ReplaceAll(global.GVA_CONFIG.AutoCode.AiPath, "{FUNC}", "api/chat/ai")
	res, err := request.HttpRequest(
		path,
		"POST",
		nil,
		params,
		nil,
	)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_auto_code.largeModelCreationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.largeModelCreationFailErr")+" "+err.Error(), c)
		return
	}
	var resStruct response.Response
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_auto_code.largeModelCreationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.largeModelCreationFailErr")+" "+err.Error(), c)
		return
	}
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_auto_code.largeModelCreationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.largeModelCreationFailErr")+" "+err.Error(), c)
		return
	}

	if resStruct.Code == 7 {
		global.GVA_LOG.Error(global.Translate("sys_auto_code.largeModelCreationFail")+resStruct.Msg, zap.Error(err))
		response.FailWithMessage(global.Translate("sys_auto_code.largeModelCreationFailErr")+" "+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
}
