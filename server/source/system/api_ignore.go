package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApiIgnore struct{}

const initOrderApiIgnore = initOrderApi + 1

// auto run
func init() {
	system.RegisterInit(initOrderApiIgnore, &initApiIgnore{})
}

func (i *initApiIgnore) InitializerName() string {
	return sysModel.SysIgnoreApi{}.TableName()
}

func (i *initApiIgnore) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysIgnoreApi{})
}

func (i *initApiIgnore) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysIgnoreApi{})
}

func (i *initApiIgnore) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysIgnoreApi{
		{Method: "GET", Path: "/swagger/*any"},
		{Method: "GET", Path: "/api/freshCasbin"},
		{Method: "GET", Path: "/uploads/file/*filepath"},
		{Method: "GET", Path: "/health"},
		{Method: "HEAD", Path: "/uploads/file/*filepath"},
		{Method: "POST", Path: "/autoCode/llmAuto"},
		{Method: "POST", Path: "/system/reloadSystem"},
		{Method: "POST", Path: "/base/login"},
		{Method: "POST", Path: "/base/captcha"},
		{Method: "POST", Path: "/init/initdb"},
		{Method: "POST", Path: "/init/checkdb"},
		{Method: "GET", Path: "/info/getInfoDataSource"},
		{Method: "GET", Path: "/info/getInfoPublic"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysIgnoreApi{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApiIgnore) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/swagger/*any", "GET").
		First(&sysModel.SysIgnoreApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
