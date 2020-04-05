package model

import (
	"errors"
	"gin-vue-admin/global"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"strings"
)

type CasbinModel struct {
	ID          uint   `json:"id" gorm:"column:_id"`
	Ptype       string `json:"ptype" gorm:"column:ptype"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}

// 供入参使用
type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// 供入参使用
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

// 更新权限
func (c *CasbinModel) CasbinPUpdate(AuthorityId string, casbinInfos []CasbinInfo) error {
	c.clearCasbin(0, AuthorityId)
	for _, v := range casbinInfos {
		cm := CasbinModel{
			ID:          0,
			Ptype:       "p",
			AuthorityId: AuthorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		addflag := c.AddCasbin(cm)
		if addflag == false {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

// API更新随动
func (c *CasbinModel) CasbinApiUpdate(oldPath string, newPath string) error {
	var cs []CasbinModel
	err := global.GVA_DB.Table("casbin_rule").Where("v1 = ?", oldPath).Find(&cs).Update("v1", newPath).Error
	return err
}

//添加权限
func (c *CasbinModel) AddCasbin(cm CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy(cm.AuthorityId, cm.Path, cm.Method)
}

//获取权限列表
func (c *CasbinModel) GetPolicyPathByAuthorityId(AuthorityId string) []string {
	e := Casbin()
	var pathList []string
	list := e.GetFilteredPolicy(0, AuthorityId)
	for _, v := range list {
		pathList = append(pathList, v[1])
	}
	return pathList
}

//清除匹配的权限
func (c *CasbinModel) clearCasbin(v int, p string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p)

}

// 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	//剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}

//持久化到数据库  引入自定义规则
func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(global.GVA_DB)
	e := casbin.NewEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}
