package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	_ = ctx
	entities := []model.SysApi{
		// 代码生成器
		{Path: "/autoCode/getDB", Description: "获取数据库列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/getTables", Description: "获取数据表列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/getColumn", Description: "获取字段列表", ApiGroup: "代码生成器", Method: "GET"},
		{Path: "/autoCode/preview", Description: "预览自动代码", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/createTemp", Description: "生成自动代码", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/pubPlug", Description: "打包插件", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/installPlugin", Description: "安装插件", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/removePlugin", Description: "移除插件", ApiGroup: "代码生成器", Method: "POST"},
		{Path: "/autoCode/getPluginList", Description: "获取插件列表", ApiGroup: "代码生成器", Method: "GET"},

		// 模板配置
		{Path: "/autoCode/getPackage", Description: "获取自动化包列表", ApiGroup: "模板配置", Method: "POST"},
		{Path: "/autoCode/delPackage", Description: "删除自动化包", ApiGroup: "模板配置", Method: "POST"},
		{Path: "/autoCode/createPackage", Description: "创建自动化包", ApiGroup: "模板配置", Method: "POST"},
		{Path: "/autoCode/getTemplates", Description: "获取模板列表", ApiGroup: "模板配置", Method: "GET"},

		// 代码生成器历史
		{Path: "/autoCode/getMeta", Description: "获取自动代码历史元数据", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/rollback", Description: "回滚自动代码历史", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/delSysHistory", Description: "删除自动代码历史", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/getSysHistory", Description: "获取自动代码历史列表", ApiGroup: "代码生成器历史", Method: "POST"},
		{Path: "/autoCode/addFunc", Description: "追加自动代码方法", ApiGroup: "代码生成器历史", Method: "POST"},
	}
	utils.RegisterApis(entities...)
}
