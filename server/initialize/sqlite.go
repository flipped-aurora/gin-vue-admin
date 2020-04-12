package initialize

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//初始化数据库并产生数据库全局变量
func Sqlite() {
	admin := global.GVA_CONFIG.Sqlite
	if db, err := gorm.Open("sqlite3", fmt.Sprintf("%s?%s", admin.Path,admin.Config)); err != nil {
		global.GVA_LOG.Error("DEFAULTDB数据库启动异常", err)
	} else {
		global.GVA_DB = db
		global.GVA_DB.LogMode(admin.LogMode)
	}
}
