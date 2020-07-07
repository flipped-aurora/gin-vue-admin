package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type tplData struct {
	template     *template.Template
	locationPath string
	autoCodePath string
}

// @title    CreateTemp
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22）
// @param     autoCode        model.AutoCodeStruct
// @return    err             error

func CreateTemp(autoCode model.AutoCodeStruct) (err error) {
	basePath := "resource/template"
	// 获取 basePath 文件夹下所有tpl文件
	tplFileList, err := GetAllTplFile(basePath, nil)
	if err != nil {
		return err
	}
	dataList := make([]tplData, 0, len(tplFileList))
	fileList := make([]string, 0, len(tplFileList))
	needMkdir := make([]string, 0, len(tplFileList)) // 当文件夹下存在多个tpl文件时，改为map更合理
	// 根据文件路径生成 tplData 结构体，待填充数据
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value})
	}
	// 生成 *Template, 填充 template 字段
	for index, value := range dataList {
		dataList[index].template, err = template.ParseFiles(value.locationPath)
		if err != nil {
			return err
		}
	}

	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/fe/api.js.tpl -> autoCode/fe/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
	autoPath := "autoCode/"
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.locationPath, basePath+"/")
		if trimBase == "readme.txt.tpl" {
			dataList[index].autoCodePath = autoPath + "readme.txt"
			continue
		}

		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				dataList[index].autoCodePath = autoPath + trimBase[:lastSeparator] + "/" + autoCode.PackageName + "/" +
					origFileName[:firstDot] + "/" + autoCode.PackageName + origFileName[firstDot:]
			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].autoCodePath, "/"); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].autoCodePath[:lastSeparator])
		}
	}

	// 写入文件前，先创建文件夹
	if err = utils.CreateDir(needMkdir...); err != nil {
		return err
	}

	// 生成文件
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, autoCode); err != nil {
			return err
		}
		_ = f.Close()
	}

	// 生成压缩包
	if err := utils.ZipFiles("./ginvueadmin.zip", fileList, ".", "."); err != nil {
		return err
	}

	// 移除中间文件
	if err := os.RemoveAll(autoPath); err != nil {
		return err
	}
	return nil
}

// GetAllTplFile 用来获取 pathName 文件夹下所有 tpl 文件
func GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = GetAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}

func GetTables(dbName string) (err error, TableNames []request.TableReq) {
	err = global.GVA_DB.Raw("select table_name as table_name from information_schema.tables where table_schema = ?", dbName).Scan(&TableNames).Error
	return err, TableNames
}

func GetDB() (err error, DBNames []request.DBReq) {
	err = global.GVA_DB.Raw("SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;").Scan(&DBNames).Error
	return err, DBNames
}

func GetColume(tableName string, dbName string) (err error, Columes []request.ColumeReq) {
	err = global.GVA_DB.Raw("SELECT COLUMN_NAME colume_name,DATA_TYPE data_type,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT colume_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", tableName, dbName).Scan(&Columes).Error
	return err, Columes
}
