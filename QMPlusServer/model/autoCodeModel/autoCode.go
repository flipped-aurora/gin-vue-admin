package autoCodeModel

import (
	"fmt"
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
	tmpl, err := template.ParseFiles("../../tpl/te/struct.go.tpl")
	fmt.Println(tmpl, err)
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

	_dir := "../" + a.PackageName
	exist, err := pathExists(_dir)
	if err != nil {
		//log.L.Info(fmt.Sprintf("get dir error![%v]\n", err))
		return
	}
	if exist {
		//log.L.Info(fmt.Sprintf("has dir![%v]\n"+_dir))
	} else {
		//log.L.Info(fmt.Sprintf("no dir![%v]\n"+_dir))
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			//log.L.Error(fmt.Sprintf("mkdir error![%v]\n",err))
		} else {
			//log.L.Info("mkdir success!\n")
		}
	}
	file, err := os.OpenFile("../"+a.PackageName+"/struct.go", os.O_CREATE|os.O_WRONLY, 0755)
	err = tmpl.Execute(file, a)
}

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
