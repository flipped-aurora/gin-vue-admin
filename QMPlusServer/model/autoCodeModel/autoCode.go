package autoCodeModel

import (
	"fmt"
	"gin-vue-admin/tools"
	"html/template"
	"os"
)

//开发中功能，若您发现这块代码可以研究，可以无视
type AutoCodeStruct struct {
	StructName   string      `json:"structName"`
	PackageName  string      `json:"packageName"`
	Abbreviation string      `json:"abbreviation"`
	Components   []Component `json:"components"`
}

type Component struct {
	ComponentName       string       `json:"componentName"`
	ComponentType       string       `json:"componentType"`
	ComponentJson       string       `json:"componentJson"`
	Ismultiple          bool         `json:"isMultiple"`
	ComponentShowType   string       `json:"componentShowType"`
	NideDictionary      bool         `json:"nideDictionary"`
	DictionaryName      string       `json:"dictionaryName"`
	ComponentDictionary []Dictionary `json:"dictionary"`
}

type Dictionary struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func Temp() {
	modelTmpl, modelTplErr := template.ParseFiles("../../tpl/te/model.go.tpl")
	apiTmpl, apiTplErr := template.ParseFiles("../../tpl/te/api.go.tpl")
	routerTmpl, routerTplErr := template.ParseFiles("../../tpl/te/router.go.tpl")
	feapiTmpl, feapiTplErr := template.ParseFiles("../../tpl/fe/api.js.tpl")

	fmt.Println(modelTplErr,apiTplErr,routerTplErr,feapiTplErr)

	a1 := Component{
		ComponentName:       "TestComponent",
		ComponentType:       "string",
		ComponentJson:       "testComponent",
		Ismultiple:          false,
		ComponentShowType:   "",
		NideDictionary:      false,
		DictionaryName:      "",
		ComponentDictionary: nil,
	}
	a2 := Component{
		ComponentName:       "TestBigComponent",
		ComponentType:       "int",
		ComponentJson:       "testBigComponent",
		Ismultiple:          false,
		ComponentShowType:   "",
		NideDictionary:      false,
		DictionaryName:      "",
		ComponentDictionary: nil,
	}
	a := AutoCodeStruct{
		StructName:   "Test",
		PackageName:  "autocode",
		Abbreviation: "t",
		Components:   []Component{a1, a2},
	}
	_autoCode := "../../autoCode/"
	_te := "../../autoCode/te/"
	_dir := "../../autoCode/te/" + a.PackageName
	_modeldir := "../../autoCode/te/" + a.PackageName+"/model"
	_apidir := "../../autoCode/te/" + a.PackageName+"/api"
	_routerdir := "../../autoCode/te/" + a.PackageName+"/router"
	_fe := "../../autoCode/fe/"
	_fe_dir:="../../autoCode/fe/" + a.PackageName
	_fe_apidir:="../../autoCode/fe/" + a.PackageName+"/api"
	mkerr := createDir(_autoCode,_te,_dir,_modeldir,_apidir,_routerdir,_fe,_fe_dir,_fe_apidir)
	fmt.Print(mkerr)
	model, _ := os.OpenFile("../../autoCode/te/"+a.PackageName+"/model/model.go", os.O_CREATE|os.O_WRONLY, 0755)
	api, _ := os.OpenFile("../../autoCode/te/"+a.PackageName+"/api/api.go", os.O_CREATE|os.O_WRONLY, 0755)
	router, _ := os.OpenFile("../../autoCode/te/"+a.PackageName+"/router/router.go", os.O_CREATE|os.O_WRONLY, 0755)
	feapi ,_ := os.OpenFile("../../autoCode/fe/"+a.PackageName+"/api/api.js", os.O_CREATE|os.O_WRONLY, 0755)
	modelErr := modelTmpl.Execute(model, a)
	apiErr := apiTmpl.Execute(api, a)
	routerErr := routerTmpl.Execute(router, a)
	feapiErr := feapiTmpl.Execute(feapi, a)

	fmt.Println(modelErr,apiErr,routerErr,feapiErr)
}



//批量创建文件夹
func createDir(dirs ...string)(err error){
	for _,v:=range dirs{
		exist, err := tools.PathExists(v)
		if err != nil {
			//log.L.Info(fmt.Sprintf("get dir error![%v]\n", err))
			return err
		}
		if exist {
			//log.L.Info(fmt.Sprintf("has dir![%v]\n"+_dir))
		} else {
			//log.L.Info(fmt.Sprintf("no dir![%v]\n"+_dir))
			// 创建文件夹
			err = os.Mkdir(v, os.ModePerm)
			if err != nil {
				//log.L.Error(fmt.Sprintf("mkdir error![%v]\n",err))
			} else {
				//log.L.Info("mkdir success!\n")
			}
		}
	}
	return err
}