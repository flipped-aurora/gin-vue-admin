package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct{}

// InitDB
// @Tags     InitDB
// @Summary  初始化用户数据库
// @Produce  application/json
// @Param    data  body      request.InitDB                  true  "初始化数据库参数"
// @Success  200   {object}  response.Response{data=string}  "初始化用户数据库"
// @Router   /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error(global.Translate("init.dbAlreadyExist"))
		response.FailWithMessage(global.Translate("init.dbAlreadyExist"), c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_initdb.paramVerifyFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_initdb.paramVerifyFailErr"), c)
		return
	}

	// added by mohamed hassan to allow multilanguage support
	if dbInfo.Language != "" {
		if dbInfo.Language == "en" || dbInfo.Language == "zh" || dbInfo.Language == "ar" {
			global.GVA_CONFIG.Language.Language = dbInfo.Language // set intial database and system langauge here
		} else {
			global.GVA_CONFIG.Language.Language = "en" // set defualt language to initialize database and system to english
		}
	}
	// end of adding

	if err := initDBService.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error(global.Translate("sys_initdb.autoCreateDBFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_initdb.autoCreateDBFailErr"), c)
		return
	}
	response.OkWithData(global.Translate("sys_initdb.autoCreateDBSuccess"), c)
}

// CheckDB
// @Tags     CheckDB
// @Summary  初始化用户数据库
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]interface{},msg=string}  "初始化用户数据库"
// @Router   /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "init.initDB"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "init.dbAlreadyInit"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, global.Translate(message), c)
}
