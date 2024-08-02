package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initExcelTemplate struct{}

const initOrderExcelTemplate = initOrderDictDetail + 1

// auto run
func init() {
	system.RegisterInit(initOrderExcelTemplate, &initExcelTemplate{})
}

func (i initExcelTemplate) InitializerName() string {
	return "sys_export_templates"
}

func (i *initExcelTemplate) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysExportTemplate{})
}

func (i *initExcelTemplate) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysExportTemplate{})
}

func (i *initExcelTemplate) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []sysModel.SysExportTemplate{
		{
			Name:       "api",
			TableName:  "sys_apis",
			TemplateID: "api",
			TemplateInfo: `{
"path":"路径",
"method":"方法（大写）",
"description":"方法介绍",
"api_group":"方法分组"
}`,
		},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "sys_export_templates"+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initExcelTemplate) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&sysModel.SysExportTemplate{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
