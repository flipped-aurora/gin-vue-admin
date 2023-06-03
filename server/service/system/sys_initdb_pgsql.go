package system

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/gookit/color"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgsqlInitHandler struct{}

func NewPgsqlInitHandler() *PgsqlInitHandler {
	return &PgsqlInitHandler{}
}

// WriteConfig pgsql 回写配置
func (h PgsqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Pgsql)
	if !ok {
		return errors.New("postgresql config invalid")
	}
	global.GVA_CONFIG.System.DbType = "pgsql"
	global.GVA_CONFIG.Pgsql = c
	global.GVA_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	return global.GVA_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 pg
func (h PgsqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "pgsql" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToPgsqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.PgsqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE %s;", c.Dbname)
	if err = createDatabase(dsn, "pgx", createSql); err != nil {
		return nil, err
	} // 创建数据库

	var db *gorm.DB
	if db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h PgsqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h PgsqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for i := 0; i < len(inits); i++ {
		if inits[i].DataInserted(next) {
			color.Info.Printf(InitDataExist, Pgsql, inits[i].InitializerName())
			continue
		}
		if n, err := inits[i].InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Pgsql, inits[i].InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Pgsql, inits[i].InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Pgsql)
	return nil
}
