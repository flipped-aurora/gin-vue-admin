// 自动生成模板{{.StructName}}
package {{.PackageName}}

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type {{.StructName}} struct {
      gorm.Model {{range .Components}}
      {{.ComponentName}}  {{.ComponentType}} `json:"{{.ComponentJson}}"`    {{ end }}
}

// 创建{{.StructName}}
func ({{.Abbreviation}} *{{StructName}})Create{{.StructName}}()(err error){
        err = qmsql.DEFAULTDB.Create({{.Abbreviation}}).Error
        return err
}

// 删除{{.StructName}}
func ({{.Abbreviation}} *{{StructName}})Delete{{.StructName}}()(err error){
        err = qmsql.DEFAULTDB.Delete({{.Abbreviation}}).Error
        return err
}

// 更新{{.StructName}}
func ({{.Abbreviation}} *{{StructName}})Update{{.StructName}}()(err error, re{{.Abbreviation}} {{.StructName}}){

}

func ({{.Abbreviation}} *{{StructName}})Create(){

}