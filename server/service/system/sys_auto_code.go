package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"gorm.io/gorm"
)

const (
	autoPath = "autocode_template/"
	basePath = "resource/template"
)

var injectionPaths []injectionMeta

func Init() {
	if len(injectionPaths) != 0 {
		return
	}
	injectionPaths = []injectionMeta{
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SInitialize, "gorm.go"),
			funcName:    "MysqlTables",
			structNameF: "autocode.%s{},",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SInitialize, "router.go"),
			funcName:    "Routers",
			structNameF: "autocodeRouter.Init%sRouter(PrivateGroup)",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SApi, "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SRouter, "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SService, "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // 带格式化的
}

type tplData struct {
	template         *template.Template
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
}

type AutoCodeService struct{}

var AutoCodeServiceApp = new(AutoCodeService)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: PreviewTemp
//@description: 预览创建代码
//@param: model.AutoCodeStruct
//@return: map[string]string, error

func (autoCodeService *AutoCodeService) PreviewTemp(autoCode system.AutoCodeStruct) (map[string]string, error) {
	makeDictTypes(&autoCode)
	dataList, _, needMkdir, err := autoCodeService.getNeedList(&autoCode)
	if err != nil {
		return nil, err
	}

	// 写入文件前，先创建文件夹
	if err = utils.CreateDir(needMkdir...); err != nil {
		return nil, err
	}

	// 创建map
	ret := make(map[string]string)

	// 生成map
	for _, value := range dataList {
		ext := ""
		if ext = filepath.Ext(value.autoCodePath); ext == ".txt" {
			continue
		}
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return nil, err
		}
		if err = value.template.Execute(f, autoCode); err != nil {
			return nil, err
		}
		_ = f.Close()
		f, err = os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_RDONLY, 0o755)
		if err != nil {
			return nil, err
		}
		builder := strings.Builder{}
		builder.WriteString("```")

		if ext != "" && strings.Contains(ext, ".") {
			builder.WriteString(strings.Replace(ext, ".", "", -1))
		}
		builder.WriteString("\n\n")
		data, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		builder.Write(data)
		builder.WriteString("\n\n```")

		pathArr := strings.Split(value.autoCodePath, string(os.PathSeparator))
		ret[pathArr[1]+"-"+pathArr[3]] = builder.String()
		_ = f.Close()

	}
	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	return ret, nil
}

func makeDictTypes(autoCode *system.AutoCodeStruct) {
	DictTypeM := make(map[string]string)
	for _, v := range autoCode.Fields {
		if v.DictType != "" {
			DictTypeM[v.DictType] = ""
		}
	}

	for k := range DictTypeM {
		autoCode.DictTypes = append(autoCode.DictTypes, k)
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTemp
//@description: 创建代码
//@param: model.AutoCodeStruct
//@return: err error

func (autoCodeService *AutoCodeService) CreateTemp(autoCode system.AutoCodeStruct, ids ...uint) (err error) {
	makeDictTypes(&autoCode)
	// 增加判断: 重复创建struct
	if autoCode.AutoMoveFile && AutoCodeHistoryServiceApp.Repeat(autoCode.StructName) {
		return RepeatErr
	}
	dataList, fileList, needMkdir, err := autoCodeService.getNeedList(&autoCode)
	if err != nil {
		return err
	}
	meta, _ := json.Marshal(autoCode)
	// 写入文件前，先创建文件夹
	if err = utils.CreateDir(needMkdir...); err != nil {
		return err
	}

	// 生成文件
	for _, value := range dataList {
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, autoCode); err != nil {
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	bf := strings.Builder{}
	idBf := strings.Builder{}
	injectionCodeMeta := strings.Builder{}
	for _, id := range ids {
		idBf.WriteString(strconv.Itoa(int(id)))
		idBf.WriteString(";")
	}
	if autoCode.AutoMoveFile { // 判断是否需要自动转移
		Init()
		for index := range dataList {
			autoCodeService.addAutoMoveFile(&dataList[index])
		}
		// 判断目标文件是否都可以移动
		for _, value := range dataList {
			if utils.FileExist(value.autoMoveFilePath) {
				return errors.New(fmt.Sprintf("目标文件已存在:%s\n", value.autoMoveFilePath))
			}
		}
		for _, value := range dataList { // 移动文件
			if err := utils.FileMove(value.autoCodePath, value.autoMoveFilePath); err != nil {
				return err
			}
		}
		err = injectionCode(autoCode.StructName, &injectionCodeMeta)
		if err != nil {
			return
		}
		// 保存生成信息
		for _, data := range dataList {
			if len(data.autoMoveFilePath) != 0 {
				bf.WriteString(data.autoMoveFilePath)
				bf.WriteString(";")
			}
		}

		if global.GVA_CONFIG.AutoCode.TransferRestart {
			go func() {
				_ = utils.Reload()
			}()
		}
	} else { // 打包
		if err = utils.ZipFiles("./ginvueadmin.zip", fileList, ".", "."); err != nil {
			return err
		}
	}
	if autoCode.AutoMoveFile || autoCode.AutoCreateApiToSql {
		if autoCode.TableName != "" {
			err = AutoCodeHistoryServiceApp.CreateAutoCodeHistory(
				string(meta),
				autoCode.StructName,
				autoCode.Description,
				bf.String(),
				injectionCodeMeta.String(),
				autoCode.TableName,
				idBf.String(),
			)
		} else {
			err = AutoCodeHistoryServiceApp.CreateAutoCodeHistory(
				string(meta),
				autoCode.StructName,
				autoCode.Description,
				bf.String(),
				injectionCodeMeta.String(),
				autoCode.StructName,
				idBf.String(),
			)
		}
	}
	if err != nil {
		return err
	}
	if autoCode.AutoMoveFile {
		return system.AutoMoveErr
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllTplFile
//@description: 获取 pathName 文件夹下所有 tpl 文件
//@param: pathName string, fileList []string
//@return: []string, error

func (autoCodeService *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = autoCodeService.GetAllTplFile(pathName+"/"+fi.Name(), fileList)
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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDB
//@description: 获取指定数据库和指定数据表的所有字段名,类型值等
//@param: tableName string, dbName string
//@return: err error, Columns []request.ColumnReq

func (autoCodeService *AutoCodeService) DropTable(tableName string) error {
	return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: addAutoMoveFile
//@description: 生成对应的迁移文件路径
//@param: *tplData
//@return: null

func (autoCodeService *AutoCodeService) addAutoMoveFile(data *tplData) {
	base := filepath.Base(data.autoCodePath)
	fileSlice := strings.Split(data.autoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server,
				global.GVA_CONFIG.AutoCode.SRouter, base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SApi, base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SService, base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SModel, base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SRequest, base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Web, global.GVA_CONFIG.AutoCode.WApi, base)
		} else if strings.Contains(fileSlice[n-2], "form") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Web, global.GVA_CONFIG.AutoCode.WForm, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Web, global.GVA_CONFIG.AutoCode.WTable, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), base)
		}
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: CreateApi
//@description: 自动创建api数据,
//@param: a *model.AutoCodeStruct
//@return: err error

func (autoCodeService *AutoCodeService) AutoCreateApi(a *system.AutoCodeStruct) (ids []uint, err error) {
	apiList := []system.SysApi{
		{
			Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
			Description: "新增" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "POST",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
			Description: "删除" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
			Description: "批量删除" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
			Description: "更新" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "PUT",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
			Description: "根据ID获取" + a.Description,
			ApiGroup:    a.Abbreviation,
			Method:      "GET",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
			Description: "获取" + a.Description + "列表",
			ApiGroup:    a.Abbreviation,
			Method:      "GET",
		},
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range apiList {
			var api system.SysApi
			if errors.Is(tx.Where("path = ? AND method = ?", v.Path, v.Method).First(&api).Error, gorm.ErrRecordNotFound) {
				if err = tx.Create(&v).Error; err != nil { // 遇到错误时回滚事务
					return err
				} else {
					ids = append(ids, v.ID)
				}
			}
		}
		return nil
	})
	return ids, err
}

func (autoCodeService *AutoCodeService) getNeedList(autoCode *system.AutoCodeStruct) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	// 去除所有空格
	utils.TrimSpace(autoCode)
	for _, field := range autoCode.Fields {
		utils.TrimSpace(field)
	}
	// 获取 basePath 文件夹下所有tpl文件
	tplFileList, err := autoCodeService.GetAllTplFile(basePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList)) // 当文件夹下存在多个tpl文件时，改为map更合理
	// 根据文件路径生成 tplData 结构体，待填充数据
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value})
	}
	// 生成 *Template, 填充 template 字段
	for index, value := range dataList {
		dataList[index].template, err = template.ParseFiles(value.locationPath)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
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
				var fileName string
				if origFileName[firstDot:] != ".go" {
					fileName = autoCode.PackageName + origFileName[firstDot:]
				} else {
					fileName = autoCode.HumpPackageName + origFileName[firstDot:]
				}

				dataList[index].autoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator], autoCode.PackageName,
					origFileName[:firstDot], fileName)
			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].autoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].autoCodePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

// injectionCode 封装代码注入
func injectionCode(structName string, bf *strings.Builder) error {
	for _, meta := range injectionPaths {
		code := fmt.Sprintf(meta.structNameF, structName)
		if err := utils.AutoInjectionCode(meta.path, meta.funcName, code); err != nil {
			return err
		}
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}
