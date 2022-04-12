package system

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
)

const initOrderMenuViewPg = initOrderMenuAuthority + 1

type initMenuViewPg struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuViewPg, &initMenuViewPg{})
}

func (i initMenuViewPg) InitializerName() string {
	return fmt.Sprintf("postgresql 视图<%s>", sysModel.SysMenu{}.TableName())
}

func (i *initMenuViewPg) InitializeData(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initMenuViewPg) DataInserted(ctx context.Context) bool {
	return true // ignore
}

func (a *initMenuViewPg) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "pgsql" {
		return ctx, nil // ignore
	}
	joinTableName := db.Model(&sysModel.SysAuthority{}).Association("SysBaseMenus").Relationship.JoinTable.Name

	sql := `
	CREATE OR REPLACE VIEW @table_name as
	select @menus.id                       as id,
		   @menus.path                     as path,
		   @menus.name                     as name,
		   @menus.icon                     as icon,
		   @menus.sort                     as sort,
		   @menus.title                    as title,
		   @menus.hidden                   as hidden,
		   @menus.parent_id                as parent_id,
		   @menus.component                as component,
		   @menus.keep_alive               as keep_alive,
		   @menus.created_at               as created_at,
		   @menus.updated_at               as updated_at,
		   @menus.deleted_at               as deleted_at,
		   @menus.menu_level               as menu_level,
		   @menus.default_menu             as default_menu,
		   @menus.close_tab                as close_tab,
		   @authorities_menus.sys_base_menu_id      as menu_id,
		   @authorities_menus.sys_authority_authority_id as authority_id
	from (@authorities_menus join @menus on ((@authorities_menus.sys_base_menu_id = @menus.id)));`
	sql = strings.ReplaceAll(sql, "@table_name", sysModel.SysMenu{}.TableName())
	sql = strings.ReplaceAll(sql, "@menus", sysModel.SysBaseMenu{}.TableName())
	sql = strings.ReplaceAll(sql, "@authorities_menus", joinTableName)
	if err := global.GVA_DB.Exec(sql).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysMenu{}.TableName()+"视图创建失败!")
	}
	return ctx, nil
}

func (a *initMenuViewPg) TableCreated(ctx context.Context) bool {
	err1 := global.GVA_DB.Find(&[]sysModel.SysMenu{}).Error
	err2 := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist",
		global.GVA_CONFIG.Pgsql.Dbname, sysModel.SysMenu{}.TableName()))
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
