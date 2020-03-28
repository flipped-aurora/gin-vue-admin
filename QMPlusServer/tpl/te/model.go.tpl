// 自动生成模板{{.StructName}}
package {{.PackageName}}

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type {{.StructName}} struct {
      gorm.Model {{range .Components}}
      {{.ComponentName}}  {{.ComponentType}} `json:"{{.ComponentJson}}"`    {{ end }}
}

// 创建{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})Create{{.StructName}}()(err error){
        err = qmsql.DEFAULTDB.Create({{.Abbreviation}}).Error
        return err
}

// 删除{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})Delete{{.StructName}}()(err error){
        err = qmsql.DEFAULTDB.Delete({{.Abbreviation}}).Error
        return err
}

// 更新{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})Update{{.StructName}}()(err error, re{{.Abbreviation}} {{.StructName}}){
        err = qmsql.DEFAULTDB.Save({{.Abbreviation}}).Error
        return err, *{{.Abbreviation}}
}

// 根据ID查看单条{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})FindById()(err error,re{{.Abbreviation}} {{.StructName}}){
    err = qmsql.DEFAULTDB.Where("id = ?",{{.Abbreviation}}.ID).First(&re{{.Abbreviation}}).Error
    return err,re{{.Abbreviation}}
}

// 分页获取{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})GetInfoList(info modelInterface.PageInfo)(err error, list interface{}, total int){
    	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
    	err, db, total := servers.PagingServer({{.Abbreviation}}, info)
    	if err != nil {
    		return
    	} else {
    		var re{{.StructName}}List []{{.StructName}}
    		err = db.Find(&re{{.StructName}}List).Error
    		return err, re{{.StructName}}List, total
    	}
}