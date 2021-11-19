package system

import "github.com/gookit/color"

type InitDBFunc interface {
	Init() (err error)
}

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	AuthorityMenu   = "\n[%v] --> %v 视图已存在!\n"
	InitDataExist   = "\n[%v] --> %v 表的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 表初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 表初始数据成功!\n"
)

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

// MysqlDataInitialize Mysql 初始化接口使用封装
// Author [SliverHorn](https://github.com/SliverHorn)
func MysqlDataInitialize(inits ...InitData) error {
	var entity SysMenu
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == entity.TableName() {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, Mysql, inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, Mysql, inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, err)
			return err
		} else {
			color.Info.Printf(InitDataSuccess, Mysql, inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, Mysql)
	return nil
}

// PgsqlDataInitialize Pgsql 初始化接口使用封装
// Author [SliverHorn](https://github.com/SliverHorn)
func PgsqlDataInitialize(inits ...InitData) error {
	var entity SysMenu
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == entity.TableName() {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, Pgsql, inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, Pgsql, inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, Pgsql, err)
			continue
		} else {
			color.Info.Printf(InitDataSuccess, Pgsql, inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, Pgsql)
	return nil
}
