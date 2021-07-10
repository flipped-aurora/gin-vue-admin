package model

import "gin-vue-admin/global"

// 自动迁移代码记录,用于回滚,重放使用

type SysAutoCodeHistory struct {
	global.GVA_MODEL
	TableName     string
	AutoCodeMeta  string `gorm:"type:text"` // 其他meta信息 path;path
	InjectionMeta string `gorm:"type:text"` // 注入的内容 RouterPath@functionName@RouterString;
	ApiIDs        string // api表注册内容
	Flag          int    // 表示对应状态 0 代表创建, 1 代表回滚 ...
}
