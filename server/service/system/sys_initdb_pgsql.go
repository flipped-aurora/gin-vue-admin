package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/source/example"
	"github.com/flipped-aurora/gin-vue-admin/server/source/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"path/filepath"
)

// writePgsqlConfig pgsql 回写配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) writePgsqlConfig(pgsql config.Pgsql) error {
	global.GVA_CONFIG.Pgsql = pgsql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	global.GVA_VP.Set("jwt.signing-key", uuid.NewV4())
	return global.GVA_VP.WriteConfig()
}

// createPgsqlDatabase 根据页面传递的数据库名 创建数据库
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) createPgsqlDatabase(dsn string, dbName string) error {
	ctx := context.Background()
	_sql := "CREATE DATABASE " + dbName
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close(ctx)
	}()
	if err = db.Ping(ctx); err != nil {
		return err
	}
	_, err = db.Exec(ctx, _sql)
	return err
}
func (initDBService *InitDBService) initPgsqlDB(conf request.InitDB) error {
	dsn := conf.PgsqlEmptyDsn()
	if err := initDBService.createPgsqlDatabase(dsn, conf.DBName); err != nil {
		return err
	} // 创建数据库

	pgsqlConfig := conf.ToPgsqlConfig()
	if pgsqlConfig.Dbname == "" {
		return nil
	} // 如果没有数据库名, 则跳出初始化数据

	if db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  pgsqlConfig.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		global.GVA_DB = db
	}

	if err := initDBService.initTables(); err != nil {
		global.GVA_DB = nil
		return err
	}

	if err := initDBService.initPgsqlData(); err != nil {
		global.GVA_DB = nil
		return err
	}

	if err := initDBService.writePgsqlConfig(pgsqlConfig); err != nil {
		return err
	}

	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}

// initPgsqlData pgsql 初始化数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) initPgsqlData() error {
	return model.MysqlDataInitialize(
		system.Api,
		system.User,
		system.Casbin,
		system.BaseMenu,
		system.Authority,
		system.Dictionary,
		system.UserAuthority,
		system.DataAuthorities,
		system.AuthoritiesMenus,
		system.DictionaryDetail,
		system.ViewAuthorityMenuPostgres,

		example.File,
	)
}
