package system

import (
	"fmt"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
)

var ViewAuthorityMenuPostgres = new(viewAuthorityMenuPostgres)

type viewAuthorityMenuPostgres struct{}

func (a *viewAuthorityMenuPostgres) TableName() string {
	var entity system.SysMenu
	return entity.TableName()
}

func (a *viewAuthorityMenuPostgres) Initialize() error {
	var entity AuthorityMenus
	sql := `
	CREATE VIEW @table_name as
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
	sql = strings.ReplaceAll(sql, "@table_name", a.TableName())
	sql = strings.ReplaceAll(sql, "@menus", "sys_base_menus")
	sql = strings.ReplaceAll(sql, "@authorities_menus", entity.TableName())
	if err := global.GVA_DB.Exec(sql).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"视图创建失败!")
	}
	return nil
}

func (a *viewAuthorityMenuPostgres) CheckDataExist() bool {
	err1 := global.GVA_DB.Find(&[]system.SysMenu{}).Error
	err2 := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist", global.GVA_CONFIG.Pgsql.Dbname, a.TableName()))
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
