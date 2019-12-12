package sysModel

import (
	"errors"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"main/init/qmsql"
	"strings"
)

type CasbinModel struct {
	ID          uint           `json:"id" gorm:"column:_id"`
	Ptype       string        `json:"ptype" gorm:"column:ptype"`
	AuthorityId string        `json:"rolename" gorm:"column:v0"`
	Path        string        `json:"path" gorm:"column:v1"`
	Method      string        `json:"method" gorm:"column:v2"`
}



// 更新权限
func (c *CasbinModel) CasbinPUpdata(AuthorityId string,Paths []string)error{
	c.clearCasbin(0,AuthorityId)
	for _,v:=range Paths{
		cm:= CasbinModel{
			ID:          0,
			Ptype:       "p",
			AuthorityId: AuthorityId,
			Path:        v,
			Method:      "POST",
		}
		addflag := c.AddCasbin(cm)
		if(addflag == false){
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

// API更新随动
func (c *CasbinModel) CasbinApiUpdata(oldPath string,newPath string)error{
	var cs []CasbinModel
	err := qmsql.DEFAULTDB.Table("casbin_rule").Where("v1 = ?",oldPath).Find(&cs).Update("v1", newPath).Error
	return err
}

//添加权限
func (c *CasbinModel) AddCasbin(cm CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy( cm.AuthorityId, cm.Path, cm.Method)
}

//获取权限列表
func (c *CasbinModel) GetPolicyPathByAuthorityId(AuthorityId string)[]string {
	e := Casbin()
	var pathList []string
	list := e.GetFilteredPolicy(0, AuthorityId)
	for _,v:=range list{
		pathList = append(pathList, v[1])
	}
	return pathList
}

//清除匹配的权限
func (c *CasbinModel) clearCasbin(v int,p string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v,p)

}

// 自定义规则函数
func ParamsMatch(key1 string, key2 string) bool {
	k1arr := strings.Split(key1,"?")
	return k1arr[0] == key2
}

// 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}

//持久化到数据库  引入自定义规则
func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(qmsql.DEFAULTDB)
	e := casbin.NewEnforcer("./static/rbacmodel/rbac_model.conf", a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	e.LoadPolicy()
	return e
}
