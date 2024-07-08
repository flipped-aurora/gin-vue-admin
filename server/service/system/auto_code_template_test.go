package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		SkipInitializeWithVersion: true,
		DefaultStringSize:         191,
	}))
	if err != nil {
		panic(err)
	}
	global.GVA_DB = db
}

func TestSysAutoCode_PackageTemplate(t *testing.T) {
	info := request.AutoCode{
		Package:         "todo",
		PackageName:     "user",
		Abbreviation:    "user",
		HumpPackageName: "user",
	}
	code, err := AutoCodeTemplate.PackageTemplate(context.Background(), info)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(code)
}
