package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initSecurityConfig struct{}

const initOrderSecurityConfig = system.InitOrderSystem + 100

// auto run
func init() {
	system.RegisterInit(initOrderSecurityConfig, &initSecurityConfig{})
}

func (i *initSecurityConfig) InitializerName() string {
	return sysModel.SysSecurityConfig{}.TableName()
}

func (i *initSecurityConfig) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysSecurityConfig{})
}

func (i *initSecurityConfig) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysSecurityConfig{})
}

func (i *initSecurityConfig) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	cfg := sysModel.DefaultSecurityConfig(global.GVA_CONFIG.Captcha)
	cfg.ID = 1
	if err := db.Create(&cfg).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysSecurityConfig{}.TableName()+"默认配置初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), cfg)
	return next, nil
}

func (i *initSecurityConfig) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("id = ?", 1).First(&sysModel.SysSecurityConfig{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
