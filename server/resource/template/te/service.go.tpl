package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    Create{{.StructName}}
// @description   create a {{.StructName}}
// @param     {{.Abbreviation}}               model.{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @return    err             error

func Create{{.StructName}}({{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = global.GVA_DB.Create(&{{.Abbreviation}}).Error
	return err
}

// @title    Delete{{.StructName}}
// @description   delete a {{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     {{.Abbreviation}}               model.{{.StructName}}
// @return                    error

func Delete{{.StructName}}({{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = global.GVA_DB.Delete({{.Abbreviation}}).Error
	return err
}

// @title    Update{{.StructName}}
// @description   update a {{.StructName}}
// @param     {{.Abbreviation}}          *model.{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @return                    error

func Update{{.StructName}}({{.Abbreviation}} *model.{{.StructName}}) (err error) {
	err = global.GVA_DB.Save({{.Abbreviation}}).Error
	return err
}

// @title    Get{{.StructName}}
// @description   get the info of a {{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    {{.StructName}}        {{.StructName}}

func Get{{.StructName}}(id uint) (err error, {{.Abbreviation}} model.{{.StructName}}) {
	err = global.GVA_DB.Where("id = ?", id).First(&{{.Abbreviation}}).Error
	return
}

// @title    Get{{.StructName}}InfoList
// @description   get {{.StructName}} list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func Get{{.StructName}}InfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var {{.Abbreviation}}s []model.{{.StructName}}
	err = db.Find(&{{.Abbreviation}}s).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
	return err, {{.Abbreviation}}s, total
}
