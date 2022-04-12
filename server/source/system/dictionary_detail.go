package system

import (
	"context"
	"fmt"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderDictDetail = initOrderDict + 1

type initDictDetail struct{}

// auto run
func init() {
	system.RegisterInit(initOrderDictDetail, &initDictDetail{})
}

func (i *initDictDetail) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysDictionaryDetail{})
}

func (i *initDictDetail) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysDictionaryDetail{})
}

func (i initDictDetail) InitializerName() string {
	return sysModel.SysDictionaryDetail{}.TableName()
}

func (i *initDictDetail) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	dicts, ok := ctx.Value(initDict{}.InitializerName()).([]sysModel.SysDictionary)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext,
			fmt.Sprintf("未找到 %s 表初始化数据", sysModel.SysDictionary{}.TableName()))
	}
	True := true
	dicts[0].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "男", Value: 1, Status: &True, Sort: 1},
		{Label: "女", Value: 2, Status: &True, Sort: 2},
	}

	dicts[1].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "smallint", Value: 1, Status: &True, Sort: 1},
		{Label: "mediumint", Value: 2, Status: &True, Sort: 2},
		{Label: "int", Value: 3, Status: &True, Sort: 3},
		{Label: "bigint", Value: 4, Status: &True, Sort: 4},
	}

	dicts[2].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "date", Status: &True},
		{Label: "time", Value: 1, Status: &True, Sort: 1},
		{Label: "year", Value: 2, Status: &True, Sort: 2},
		{Label: "datetime", Value: 3, Status: &True, Sort: 3},
		{Label: "timestamp", Value: 5, Status: &True, Sort: 5},
	}
	dicts[3].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "float", Status: &True},
		{Label: "double", Value: 1, Status: &True, Sort: 1},
		{Label: "decimal", Value: 2, Status: &True, Sort: 2},
	}

	dicts[4].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "char", Status: &True},
		{Label: "varchar", Value: 1, Status: &True, Sort: 1},
		{Label: "tinyblob", Value: 2, Status: &True, Sort: 2},
		{Label: "tinytext", Value: 3, Status: &True, Sort: 3},
		{Label: "text", Value: 4, Status: &True, Sort: 4},
		{Label: "blob", Value: 5, Status: &True, Sort: 5},
		{Label: "mediumblob", Value: 6, Status: &True, Sort: 6},
		{Label: "mediumtext", Value: 7, Status: &True, Sort: 7},
		{Label: "longblob", Value: 8, Status: &True, Sort: 8},
		{Label: "longtext", Value: 9, Status: &True, Sort: 9},
	}

	dicts[5].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "tinyint", Status: &True},
	}
	for _, dict := range dicts {
		if err := db.Model(&dict).Association("SysDictionaryDetails").
			Replace(dict.SysDictionaryDetails); err != nil {
			return ctx, errors.Wrap(err, sysModel.SysDictionaryDetail{}.TableName()+"表数据初始化失败!")
		}
	}
	return ctx, nil
}

func (i *initDictDetail) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var dict sysModel.SysDictionary
	if err := db.Preload("SysDictionaryDetails").
		First(&dict, &sysModel.SysDictionary{Name: "数据库bool类型"}).Error; err != nil {
		return false
	}
	return len(dict.SysDictionaryDetails) > 0 && dict.SysDictionaryDetails[0].Label == "tinyint"
}
