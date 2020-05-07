package service

import (
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"os"
	"text/template"
)

// @title    CreateTemp
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22）
// @param     autoCode        model.AutoCodeStruct
// @return    err             error

func CreateTemp(autoCode model.AutoCodeStruct) (err error) {
	basePath := "./resource/template"
	modelTmpl, err := template.ParseFiles(basePath + "/te/model.go.tpl")
	if err != nil {
		return err
	}
	apiTmpl, err := template.ParseFiles(basePath + "/te/api.go.tpl")
	if err != nil {
		return err
	}
	routerTmpl, err := template.ParseFiles(basePath + "/te/router.go.tpl")
	if err != nil {
		return err
	}
	serviceTmpl, err := template.ParseFiles(basePath + "/te/service.go.tpl")
	if err != nil {
		return err
	}
	feapiTmpl, err := template.ParseFiles(basePath + "/fe/api.js.tpl")
	if err != nil {
		return err
	}
	feTableTmpl, err := template.ParseFiles(basePath + "/fe/table.vue.tpl")
	if err != nil {
		return err
	}
	readmeTmpl, err := template.ParseFiles(basePath + "/readme.txt.tpl")
	if err != nil {
		return err
	}
	//自动化总目录
	_autoCode := "./autoCode/"
	//自动化后台代码目录
	_te := "./autoCode/te/"
	_dir := _te + autoCode.PackageName
	_modeldir := _te + autoCode.PackageName + "/model"
	_apidir := _te + autoCode.PackageName + "/api"
	_routerdir := _te + autoCode.PackageName + "/router"
	_servicedir := _te + autoCode.PackageName + "/service"
	//自动化前台代码目录
	_fe := "./autoCode/fe/"
	_fe_dir := _fe + autoCode.PackageName
	_fe_apidir := _fe + autoCode.PackageName + "/api"
	_fe_tabledir := _fe + autoCode.PackageName + "/table"
	err = utils.CreateDir(_autoCode, _te, _dir, _modeldir, _apidir, _routerdir, _servicedir, _fe, _fe_dir, _fe_apidir, _fe_tabledir)
	if err != nil {
		return err
	}
	model, err := os.OpenFile(_te+autoCode.PackageName+"/model/"+autoCode.PackageName+".go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	api, err := os.OpenFile(_te+autoCode.PackageName+"/api/"+autoCode.PackageName+".go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	router, err := os.OpenFile(_te+autoCode.PackageName+"/router/"+autoCode.PackageName+".go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	service, err := os.OpenFile(_te+autoCode.PackageName+"/service/"+autoCode.PackageName+".go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	feapi, err := os.OpenFile(_fe+autoCode.PackageName+"/api/"+autoCode.PackageName+".js", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	fetable, err := os.OpenFile(_fe+autoCode.PackageName+"/table/"+autoCode.PackageName+".vue", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	readme, err := os.OpenFile(_autoCode+"readme.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	// 生成代码
	{
		err = modelTmpl.Execute(model, autoCode)
		if err != nil {
			return err
		}
		err = apiTmpl.Execute(api, autoCode)
		if err != nil {
			return err
		}
		err = routerTmpl.Execute(router, autoCode)
		if err != nil {
			return err
		}
		err = serviceTmpl.Execute(service, autoCode)
		if err != nil {
			return err
		}
		err = feapiTmpl.Execute(feapi, autoCode)
		if err != nil {
			return err
		}
		err = feTableTmpl.Execute(fetable, autoCode)
		if err != nil {
			return err
		}
		err = readmeTmpl.Execute(readme, autoCode)
		if err != nil {
			return err
		}
	}
	_ = model.Close()
	_ = api.Close()
	_ = router.Close()
	_ = service.Close()
	_ = feapi.Close()
	_ = fetable.Close()
	_ = readme.Close()
	fileList := []string{
		_te + autoCode.PackageName + "/model/" + autoCode.PackageName + ".go",
		_te + autoCode.PackageName + "/api/" + autoCode.PackageName + ".go",
		_te + autoCode.PackageName + "/router/" + autoCode.PackageName + ".go",
		_te + autoCode.PackageName + "/service/" + autoCode.PackageName + ".go",
		_fe + autoCode.PackageName + "/api/" + autoCode.PackageName + ".js",
		_fe + autoCode.PackageName + "/table/" + autoCode.PackageName + ".vue",

		_autoCode + "readme.txt",
	}
	err = utils.ZipFiles("./ginvueadmin.zip", fileList, ".", ".")
	if err != nil {
		return err
	}
	err = os.RemoveAll(_autoCode)
	if err != nil {
		return err
	}
	return nil
}
