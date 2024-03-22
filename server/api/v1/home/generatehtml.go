package home

import (
	"fmt"
	"html/template"
	"os"
)

var (
	//生成的Html保存目录
	htmlOutPath = "temp/"
	//静态文件模版目录
	templatePath = "templates/"
)

// 生成静态文件的方法 htmlname 模板文件  tempname 缓存文件
func GetGenerateHtml(data any, htmlname, tempname string) string {
	//1.获取html生成路径
	fileName := tempname
	//2.判断静态文件是否存在
	if exist(templatePath + htmlOutPath + tempname) {
		return fileName
	} else {
		temp := []string{templatePath + htmlname, templatePath + "header.html", templatePath + "left1.html", templatePath + "left2.html", templatePath + "footer.html"}
		//1.获取模版
		contenstTmp, err := template.New(htmlname).ParseFiles(temp...)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		//4.生成静态文件
		err = generateStaticHtml(contenstTmp, templatePath+htmlOutPath+tempname, data)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fileName = ""
	}
	return fileName
}

// 生成静态文件
func generateStaticHtml(template *template.Template, fileName string, data any) error {

	//2.生成静态文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	err = template.Execute(file, &data)
	return err
}

// 判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
