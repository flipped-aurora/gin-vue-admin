package service

import (
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"html/template"
	"os"
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
	feapiTmpl, err := template.ParseFiles(basePath + "/fe/api.js.tpl")
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
	//自动化前台代码目录
	_fe := "./autoCode/fe/"
	_fe_dir := _fe + autoCode.PackageName
	_fe_apidir := _fe + autoCode.PackageName + "/api"
	err = utils.CreateDir(_autoCode, _te, _dir, _modeldir, _apidir, _routerdir, _fe, _fe_dir, _fe_apidir)
	if err != nil {
		return err
	}
	model, err := os.OpenFile(_te+autoCode.PackageName+"/model/model.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	api, err := os.OpenFile(_te+autoCode.PackageName+"/api/api.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	router, err := os.OpenFile(_te+autoCode.PackageName+"/router/router.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	feapi, err := os.OpenFile(_fe+autoCode.PackageName+"/api/api.js", os.O_CREATE|os.O_WRONLY, 0755)
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
		err = feapiTmpl.Execute(feapi, autoCode)
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
	_ = feapi.Close()
	_ = readme.Close()
	fileList := []string{
		_te + autoCode.PackageName + "/model/model.go",
		_te + autoCode.PackageName + "/api/api.go",
		_te + autoCode.PackageName + "/router/router.go",
		_fe + autoCode.PackageName + "/api/api.js",
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
