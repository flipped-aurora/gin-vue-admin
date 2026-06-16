package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
)

const initOrderCli = initOrderCasbin + 1

type initCli struct{}

func init() {
	system.RegisterInit(initOrderCli, &initCli{})
}

func (i *initCli) InitializerName() string {
	return "sys_cli"
}

func (i *initCli) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysCli{}, &sysModel.SysCliApi{})
}

func (i *initCli) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&sysModel.SysCli{}) && m.HasTable(&sysModel.SysCliApi{}) && m.HasColumn(&sysModel.SysCli{}, "command")
}

func (i *initCli) InitializeData(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initCli) DataInserted(ctx context.Context) bool {
	return true
}
