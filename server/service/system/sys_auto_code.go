package system

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	ast2 "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"

	"github.com/flipped-aurora/gin-vue-admin/server/resource/autocode_template/subcontract"
	cp "github.com/otiai10/copy"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"gorm.io/gorm"
)

const (
	autoPath           = "autocode_template/"
	autocodePath       = "resource/autocode_template"
	plugPath           = "resource/plug_template"
	packageService     = "service/%s/enter.go"
	packageServiceName = "service"
	packageRouter      = "router/%s/enter.go"
	packageRouterName  = "router"
	packageAPI         = "api/v1/%s/enter.go"
	packageAPIName     = "api/v1"
)

type autoPackage struct {
	path string
	temp string
	name string
}

var (
	packageInjectionMap map[string]astInjectionMeta
	injectionPaths      []injectionMeta
)

func Init(Package string) {
	injectionPaths = []injectionMeta{
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SApi, Package), "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRouter, Package), "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SService, Package), "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}

	packageInjectionMap = map[string]astInjectionMeta{
		packageServiceName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "service", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "ServiceGroup",
			structNameF:  "%sServiceGroup",
		},
		packageRouterName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "router", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "RouterGroup",
			structNameF:  "%s",
		},
		packageAPIName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "api/v1", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "ApiGroup",
			structNameF:  "%sApiGroup",
		},
	}
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // 带格式化的
}

type astInjectionMeta struct {
	path         string
	importCodeF  string
	structNameF  string
	packageNameF string
	groupName    string
}

type tplData struct {
	template         *template.Template
	autoPackage      string
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
}

type AutoCodeService struct{}

var AutoCodeServiceApp = new(AutoCodeService)

// @author: [songzhibin97](https://github.com/songzhibin97)
// @function: PreviewTemp
// @description: 预览创建代码
// @param: model.AutoCodeStruct
// @return: map[string]string, error

func (autoCodeService *AutoCodeService) PreviewTemp(autoCode system.AutoCodeStruct) (map[string]string, error) {
	makeDictTypes(&autoCode)
	for i := range autoCode.Fields {
		if autoCode.Fields[i].FieldType == "time.Time" {
			autoCode.HasTimer = true
		}
		if autoCode.Fields[i].Require {
			autoCode.NeedValid = true
		}
		if autoCode.Fields[i].Sort {
			autoCode.NeedSort = true
		}
		if autoCode.Fields[i].FieldType == "picture" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "video" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "richtext" {
			autoCode.HasRichText = true
		}
		if autoCode.Fields[i].FieldType == "pictures" {
			autoCode.HasPic = true
			autoCode.NeedJSON = true
		}
		if autoCode.Fields[i].FieldType == "file" {
			autoCode.HasFile = true
			autoCode.NeedJSON = true
		}
	}
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
		data, err := io.ReadAll(f)
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

// @author: [piexlmax](https://github.com/piexlmax)
// @function: CreateTemp
// @description: 创建代码
// @param: model.AutoCodeStruct
// @return: err error

func (autoCodeService *AutoCodeService) CreateTemp(autoCode system.AutoCodeStruct, ids ...uint) (err error) {
	makeDictTypes(&autoCode)
	for i := range autoCode.Fields {
		if autoCode.Fields[i].FieldType == "time.Time" {
			autoCode.HasTimer = true
		}
		if autoCode.Fields[i].Require {
			autoCode.NeedValid = true
		}
		if autoCode.Fields[i].Sort {
			autoCode.NeedSort = true
		}
		if autoCode.Fields[i].FieldType == "picture" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "video" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "richtext" {
			autoCode.HasRichText = true
		}
		if autoCode.Fields[i].FieldType == "pictures" {
			autoCode.NeedJSON = true
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "file" {
			autoCode.NeedJSON = true
			autoCode.HasFile = true
		}
	}
	// 增加判断: 重复创建struct
	if autoCode.AutoMoveFile && AutoCodeHistoryServiceApp.Repeat(autoCode.BusinessDB, autoCode.StructName, autoCode.Package) {
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
		Init(autoCode.Package)
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

		{
			// 在gorm.go 注入 自动迁移
			path := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SInitialize, "gorm.go")
			varDB := utils.MaheHump(autoCode.BusinessDB)
			ast2.AddRegisterTablesAst(path, "RegisterTables", autoCode.Package, varDB, autoCode.BusinessDB, autoCode.StructName)
		}

		{
			// router.go 注入 自动迁移
			path := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SInitialize, "router.go")
			ast2.AddRouterCode(path, "Routers", autoCode.Package, autoCode.StructName)
		}
		// 给各个enter进行注入
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
				autoCode.Package,
				autoCode.BusinessDB,
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
				autoCode.Package,
				autoCode.BusinessDB,
			)
		}
	}
	if err != nil {
		return err
	}
	if autoCode.AutoMoveFile {
		return system.ErrAutoMove
	}
	return nil
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetAllTplFile
// @description: 获取 pathName 文件夹下所有 tpl 文件
// @param: pathName string, fileList []string
// @return: []string, error

func (autoCodeService *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(pathName)
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

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetDB
// @description: 获取指定数据库和指定数据表的所有字段名,类型值等
// @param: tableName string, dbName string
// @return: err error, Columns []request.ColumnReq

func (autoCodeService *AutoCodeService) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
	}
}

// @author: [SliverHorn](https://github.com/SliverHorn)
// @author: [songzhibin97](https://github.com/songzhibin97)
// @function: addAutoMoveFile
// @description: 生成对应的迁移文件路径
// @param: *tplData
// @return: null

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
				fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRouter, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SApi, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SService, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SModel, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRequest, data.autoPackage), base)
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

// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
// @function: CreateApi
// @description: 自动创建api数据,
// @param: a *model.AutoCodeStruct
// @return: err error

func (autoCodeService *AutoCodeService) AutoCreateApi(a *system.AutoCodeStruct) (ids []uint, err error) {
	apiList := []system.SysApi{
		{
			Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
			Description: "新增" + a.Description,
			ApiGroup:    a.Description,
			Method:      "POST",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
			Description: "删除" + a.Description,
			ApiGroup:    a.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
			Description: "批量删除" + a.Description,
			ApiGroup:    a.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
			Description: "更新" + a.Description,
			ApiGroup:    a.Description,
			Method:      "PUT",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
			Description: "根据ID获取" + a.Description,
			ApiGroup:    a.Description,
			Method:      "GET",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
			Description: "获取" + a.Description + "列表",
			ApiGroup:    a.Description,
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
	tplFileList, err := autoCodeService.GetAllTplFile(autocodePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList)) // 当文件夹下存在多个tpl文件时，改为map更合理
	// 根据文件路径生成 tplData 结构体，待填充数据
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value, autoPackage: autoCode.Package})
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
		trimBase := strings.TrimPrefix(value.locationPath, autocodePath+"/")
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
		ast2.ImportForAutoEnter(meta.path, meta.funcName, code)
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}

func (autoCodeService *AutoCodeService) CreateAutoCode(s *system.SysAutoCode) error {
	if s.PackageName == "autocode" || s.PackageName == "system" || s.PackageName == "example" || s.PackageName == "" {
		return errors.New("不能使用已保留的package name")
	}
	if !errors.Is(global.GVA_DB.Where("package_name = ?", s.PackageName).First(&system.SysAutoCode{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同PackageName")
	}
	if e := autoCodeService.CreatePackageTemp(s.PackageName); e != nil {
		return e
	}
	return global.GVA_DB.Create(&s).Error
}

func (autoCodeService *AutoCodeService) GetPackage() (pkgList []system.SysAutoCode, err error) {
	err = global.GVA_DB.Find(&pkgList).Error
	return pkgList, err
}

func (autoCodeService *AutoCodeService) DelPackage(a system.SysAutoCode) error {
	return global.GVA_DB.Delete(&a).Error
}

func (autoCodeService *AutoCodeService) CreatePackageTemp(packageName string) error {
	Init(packageName)
	pendingTemp := []autoPackage{{
		path: packageService,
		name: packageServiceName,
		temp: string(subcontract.Server),
	}, {
		path: packageRouter,
		name: packageRouterName,
		temp: string(subcontract.Router),
	}, {
		path: packageAPI,
		name: packageAPIName,
		temp: string(subcontract.API),
	}}
	for i, s := range pendingTemp {
		pendingTemp[i].path = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, filepath.Clean(fmt.Sprintf(s.path, packageName)))
	}
	// 选择模板
	for _, s := range pendingTemp {
		err := os.MkdirAll(filepath.Dir(s.path), 0755)
		if err != nil {
			return err
		}

		f, err := os.Create(s.path)
		if err != nil {
			return err
		}

		defer f.Close()

		temp, err := template.New("").Parse(s.temp)
		if err != nil {
			return err
		}
		err = temp.Execute(f, struct {
			PackageName string `json:"package_name"`
		}{packageName})
		if err != nil {
			return err
		}
	}
	// 创建完成后在对应的位置插入结构代码
	for _, v := range pendingTemp {
		meta := packageInjectionMap[v.name]
		if err := ast2.ImportReference(meta.path, fmt.Sprintf(meta.importCodeF, v.name, packageName), fmt.Sprintf(meta.structNameF, utils.FirstUpper(packageName)), fmt.Sprintf(meta.packageNameF, packageName), meta.groupName); err != nil {
			return err
		}
	}
	return nil
}

// CreatePlug 自动创建插件模板
func (autoCodeService *AutoCodeService) CreatePlug(plug system.AutoPlugReq) error {
	// 检查列表参数是否有效
	plug.CheckList()
	tplFileList, _ := autoCodeService.GetAllTplFile(plugPath, nil)
	for _, tpl := range tplFileList {
		temp, err := template.ParseFiles(tpl)
		if err != nil {
			zap.L().Error("parse err", zap.String("tpl", tpl), zap.Error(err))
			return err
		}
		pathArr := strings.SplitAfter(tpl, "/")
		if strings.Index(pathArr[2], "tpl") < 0 {
			dirPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SPlug, plug.Snake+"/"+pathArr[2]))
			os.MkdirAll(dirPath, 0755)
		}
		file := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SPlug, plug.Snake+"/"+tpl[len(plugPath):len(tpl)-4]))
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			zap.L().Error("open file", zap.String("tpl", tpl), zap.Error(err), zap.Any("plug", plug))
			return err
		}
		defer f.Close()

		err = temp.Execute(f, plug)
		if err != nil {
			zap.L().Error("exec err", zap.String("tpl", tpl), zap.Error(err), zap.Any("plug", plug))
			return err
		}
	}
	return nil
}

func (autoCodeService *AutoCodeService) InstallPlugin(file *multipart.FileHeader) (web, server int, err error) {
	const GVAPLUGPINATH = "./gva-plug-temp/"
	defer os.RemoveAll(GVAPLUGPINATH)
	_, err = os.Stat(GVAPLUGPINATH)
	if os.IsNotExist(err) {
		os.Mkdir(GVAPLUGPINATH, os.ModePerm)
	}

	src, err := file.Open()
	if err != nil {
		return -1, -1, err
	}
	defer src.Close()

	out, err := os.Create(GVAPLUGPINATH + file.Filename)
	if err != nil {
		return -1, -1, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	paths, err := utils.Unzip(GVAPLUGPINATH+file.Filename, GVAPLUGPINATH)
	paths = filterFile(paths)
	var webIndex = -1
	var serverIndex = -1
	for i := range paths {
		paths[i] = filepath.ToSlash(paths[i])
		pathArr := strings.Split(paths[i], "/")
		ln := len(pathArr)
		if ln < 2 {
			continue
		}
		if pathArr[ln-2] == "server" && pathArr[ln-1] == "plugin" {
			serverIndex = i
		}
		if pathArr[ln-2] == "web" && pathArr[ln-1] == "plugin" {
			webIndex = i
		}
	}
	if webIndex == -1 && serverIndex == -1 {
		zap.L().Error("非标准插件，请按照文档自动迁移使用")
		return webIndex, serverIndex, errors.New("非标准插件，请按照文档自动迁移使用")
	}

	if webIndex != -1 {
		err = installation(paths[webIndex], global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Web)
		if err != nil {
			return webIndex, serverIndex, err
		}
	}

	if serverIndex != -1 {
		err = installation(paths[serverIndex], global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Server)
	}
	return webIndex, serverIndex, err
}

func installation(path string, formPath string, toPath string) error {
	arr := strings.Split(filepath.ToSlash(path), "/")
	ln := len(arr)
	if ln < 3 {
		return errors.New("arr")
	}
	name := arr[ln-3]

	var form = filepath.ToSlash(global.GVA_CONFIG.AutoCode.Root + formPath + "/" + path)
	var to = filepath.ToSlash(global.GVA_CONFIG.AutoCode.Root + toPath + "/plugin/")
	_, err := os.Stat(to + name)
	if err == nil {
		zap.L().Error("autoPath 已存在同名插件，请自行手动安装", zap.String("to", to))
		return errors.New(toPath + "已存在同名插件，请自行手动安装")
	}
	return cp.Copy(form, to, cp.Options{Skip: skipMacSpecialDocument})
}

func filterFile(paths []string) []string {
	np := make([]string, 0, len(paths))
	for _, path := range paths {
		if ok, _ := skipMacSpecialDocument(path); ok {
			continue
		}
		np = append(np, path)
	}
	return np
}

func skipMacSpecialDocument(src string) (bool, error) {
	if strings.Contains(src, ".DS_Store") || strings.Contains(src, "__MACOSX") {
		return true, nil
	}
	return false, nil
}

func (autoCodeService *AutoCodeService) PubPlug(plugName string) (zipPath string, err error) {
	if plugName == "" {
		return "", errors.New("插件名称不能为空")
	}

	// 防止路径穿越
	plugName = filepath.Clean(plugName)

	webPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", plugName)
	serverPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", plugName)
	// 创建一个新的zip文件

	// 判断目录是否存在
	webInfo, err := os.Stat(webPath)
	if err != nil {
		return "", errors.New("web路径不存在")
	}
	serverInfo, err := os.Stat(serverPath)
	if err != nil {
		return "", errors.New("server路径不存在")
	}

	fileName := plugName + ".zip"
	// 创建一个新的zip文件
	zipFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()

	// 创建一个zip写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 创建一个新的文件头
	webHeader, err := zip.FileInfoHeader(webInfo)
	if err != nil {
		return
	}

	// 创建一个新的文件头
	serverHeader, err := zip.FileInfoHeader(serverInfo)
	if err != nil {
		return
	}

	webHeader.Name = filepath.Join(plugName, "web", "plugin")
	serverHeader.Name = filepath.Join(plugName, "server", "plugin")

	// 将文件添加到zip归档中
	_, err = zipWriter.CreateHeader(serverHeader)
	_, err = zipWriter.CreateHeader(webHeader)

	// 遍历webPath目录并将所有非隐藏文件添加到zip归档中
	err = filepath.Walk(webPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过隐藏文件
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// 创建一个新的文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 将文件头的名称设置为文件的相对路径
		rel, _ := filepath.Rel(webPath, path)
		header.Name = filepath.Join(plugName, "web", "plugin", plugName, rel)

		// 将文件添加到zip归档中
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// 打开文件并将其内容复制到zip归档中
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	// 遍历serverPath目录并将所有非隐藏文件添加到zip归档中
	err = filepath.Walk(serverPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过隐藏文件和目录
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// 创建一个新的文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 将文件头的名称设置为文件的相对路径
		rel, _ := filepath.Rel(serverPath, path)
		header.Name = filepath.Join(plugName, "server", "plugin", plugName, rel)
		// 将文件添加到zip归档中
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// 打开文件并将其内容复制到zip归档中
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return
	}
	return filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fileName), nil
}
