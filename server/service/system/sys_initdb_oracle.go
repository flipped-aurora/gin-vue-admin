package system

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/gookit/color"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/dzwvip/oracle"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OracleInitHandler struct{}

func NewOracleInitHandler() *OracleInitHandler {
	return &OracleInitHandler{}
}

// WriteConfig oracle回写配置
func (h OracleInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Oracle)
	if !ok {
		return errors.New("oracle config invalid")
	}
	global.GVA_CONFIG.System.DbType = "oracle"
	global.GVA_CONFIG.Oracle = c
	global.GVA_CONFIG.JWT.SigningKey = uuid.New().String()
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	global.GVA_ACTIVE_DBNAME = &c.Dbname
	return global.GVA_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 oracle
func (h OracleInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "oracle" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToOracleConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.OracleEmptyDsn()
	createSql := fmt.Sprintf("CREATE USER %s IDENTIFIED BY %s DEFAULT TABLESPACE USERS;", c.Dbname, c.Password)
	if err = createDatabase(dsn, "oracle", createSql); err != nil {
		return nil, err
	} // 创建用户

	var db *gorm.DB
	if db, err = gorm.Open(oracle.New(oracle.Config{
		DSN: c.Dsn(), // DSN data source name
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h OracleInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h OracleInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, "Oracle", init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, "Oracle", init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, "Oracle", init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, "Oracle")
	return nil
}
