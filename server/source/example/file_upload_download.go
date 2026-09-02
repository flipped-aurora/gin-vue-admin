package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderExaFile = system.InitOrderInternal + 1

type initExaFileMysql struct{}

// auto run
func init() {
	system.RegisterInit(initOrderExaFile, &initExaFileMysql{})
}

func (i *initExaFileMysql) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&example.ExaFileUploadAndDownload{})
}

func (i *initExaFileMysql) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&example.ExaFileUploadAndDownload{})
}

func (i *initExaFileMysql) InitializerName() string {
	return example.ExaFileUploadAndDownload{}.TableName()
}

func (i *initExaFileMysql) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []example.ExaFileUploadAndDownload{
		{Name: "10.png", Url: "https://qmplusimg.henrongyi.top/gvalogo.png", Tag: "png", Key: "158787308910.png"},
		{Name: "logo.png", Url: "https://qmplusimg.henrongyi.top/1576554439myAvatar.png", Tag: "png", Key: "1587973709logo.png"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, example.ExaFileUploadAndDownload{}.TableName()+"表数据初始化失败!")
	}
	return ctx, nil
}

func (i *initExaFileMysql) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	lookup := example.ExaFileUploadAndDownload{Name: "logo.png", Key: "1587973709logo.png"}
	if errors.Is(db.First(&lookup, &lookup).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
