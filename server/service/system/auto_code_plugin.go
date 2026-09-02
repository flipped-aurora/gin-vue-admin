package system

import (
	"bytes"
	"context"
	"fmt"
	goast "go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	pluginUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	ast "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/mholt/archives"
	cp "github.com/otiai10/copy"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var AutoCodePlugin = new(autoCodePlugin)

type autoCodePlugin struct{}

// Install 插件安装
func (s *autoCodePlugin) Install(file *multipart.FileHeader) (web, server int, err error) {
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

	// 在临时目录创建目标文件
	// 使用完整路径拼接的好处：明确文件位置，避免路径混乱
	out, err := os.Create(GVAPLUGPINATH + file.Filename)
	if err != nil {
		return -1, -1, err
	}

	// 将上传的文件内容复制到临时文件
	// 使用io.Copy的好处：高效处理大文件，自动管理缓冲区，避免内存溢出
	_, err = io.Copy(out, src)
	if err != nil {
		out.Close()
		return -1, -1, err
	}

	// 立即关闭文件，确保数据写入磁盘并释放文件句柄
	// 必须在解压前关闭，否则在Windows系统上会导致文件被占用无法解压
	err = out.Close()
	if err != nil {
		return -1, -1, err
	}

	paths, err := utils.Unzip(GVAPLUGPINATH+file.Filename, GVAPLUGPINATH)
	paths = filterFile(paths)
	var webIndex = -1
	var serverIndex = -1
	webPlugin := ""
	serverPlugin := ""
	serverPackage := ""
	serverRootName := ""

	for i := range paths {
		paths[i] = filepath.ToSlash(paths[i])
		pathArr := strings.Split(paths[i], "/")
		ln := len(pathArr)

		if ln < 4 {
			continue
		}
		if pathArr[2]+"/"+pathArr[3] == `server/plugin` {
			if len(serverPlugin) == 0 {
				serverPlugin = filepath.Join(pathArr[0], pathArr[1], pathArr[2], pathArr[3])
			}
			if serverRootName == "" && ln > 1 && pathArr[1] != "" {
				serverRootName = pathArr[1]
			}
			if ln > 4 && serverPackage == "" && pathArr[4] != "" {
				serverPackage = pathArr[4]
			}
		}
		if pathArr[2]+"/"+pathArr[3] == `web/plugin` && len(webPlugin) == 0 {
			webPlugin = filepath.Join(pathArr[0], pathArr[1], pathArr[2], pathArr[3])
		}
	}
	if len(serverPlugin) == 0 && len(webPlugin) == 0 {
		zap.L().Error("非标准插件，请按照文档自动迁移使用")
		return webIndex, serverIndex, errors.New("非标准插件，请按照文档自动迁移使用")
	}

	if len(serverPlugin) != 0 {
		if serverPackage == "" {
			serverPackage = serverRootName
		}
		err = installation(serverPlugin, global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Server)
		if err != nil {
			return webIndex, serverIndex, err
		}
		err = ensurePluginRegisterImport(serverPackage)
		if err != nil {
			return webIndex, serverIndex, err
		}
	}

	if len(webPlugin) != 0 {
		err = installation(webPlugin, global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Web)
		if err != nil {
			return webIndex, serverIndex, err
		}
	}

	return 1, 1, err
}

func installation(path string, formPath string, toPath string) error {
	arr := strings.Split(filepath.ToSlash(path), "/")
	ln := len(arr)
	if ln < 3 {
		return errors.New("arr")
	}
	name := arr[ln-3]

	var form = filepath.Join(global.GVA_CONFIG.AutoCode.Root, formPath, path)
	var to = filepath.Join(global.GVA_CONFIG.AutoCode.Root, toPath, "plugin")
	_, err := os.Stat(to + name)
	if err == nil {
		zap.L().Error("autoPath 已存在同名插件，请自行手动安装", zap.String("to", to))
		return errors.New(toPath + "已存在同名插件，请自行手动安装")
	}
	return cp.Copy(form, to, cp.Options{Skip: skipMacSpecialDocument})
}

func ensurePluginRegisterImport(packageName string) error {
	module := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module)
	if module == "" {
		return errors.New("autocode module is empty")
	}
	if packageName == "" {
		return errors.New("plugin package is empty")
	}

	registerPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "register.go")
	src, err := os.ReadFile(registerPath)
	if err != nil {
		return err
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, registerPath, src, parser.ParseComments)
	if err != nil {
		return err
	}

	importPath := fmt.Sprintf("%s/plugin/%s", module, packageName)
	if ast.CheckImport(astFile, importPath) {
		return nil
	}

	importSpec := &goast.ImportSpec{
		Name: goast.NewIdent("_"),
		Path: &goast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", importPath)},
	}
	var importDecl *goast.GenDecl
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*goast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok == token.IMPORT {
			importDecl = genDecl
			break
		}
	}
	if importDecl == nil {
		astFile.Decls = append([]goast.Decl{
			&goast.GenDecl{
				Tok:   token.IMPORT,
				Specs: []goast.Spec{importSpec},
			},
		}, astFile.Decls...)
	} else {
		importDecl.Specs = append(importDecl.Specs, importSpec)
	}

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)

	return os.WriteFile(registerPath, bf.Bytes(), 0666)
}

func filterFile(paths []string) []string {
	np := make([]string, 0, len(paths))
	for _, path := range paths {
		if ok, _ := skipMacSpecialDocument(nil, path, ""); ok {
			continue
		}
		np = append(np, path)
	}
	return np
}

func skipMacSpecialDocument(_ os.FileInfo, src, _ string) (bool, error) {
	if strings.Contains(src, ".DS_Store") || strings.Contains(src, "__MACOSX") {
		return true, nil
	}
	return false, nil
}

func (s *autoCodePlugin) PubPlug(plugName string) (zipPath string, err error) {
	if plugName == "" {
		return "", errors.New("插件名称不能为空")
	}

	// 防止路径穿越
	plugName = filepath.Clean(plugName)

	webPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", plugName)
	serverPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", plugName)
	// 创建一个新的zip文件

	// 判断目录是否存在
	_, err = os.Stat(webPath)
	if err != nil {
		return "", errors.New("web路径不存在")
	}
	_, err = os.Stat(serverPath)
	if err != nil {
		return "", errors.New("server路径不存在")
	}

	fileName := plugName + ".zip"
	// 创建一个新的zip文件
	files, err := archives.FilesFromDisk(context.Background(), nil, map[string]string{
		webPath:    plugName + "/web/plugin/" + plugName,
		serverPath: plugName + "/server/plugin/" + plugName,
	})

	// create the output file we'll write to
	out, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer out.Close()

	// we can use the CompressedArchive type to gzip a tarball
	// (compression is not required; you could use Tar directly)
	format := archives.CompressedArchive{
		//Compression: archives.Gz{},
		Archival: archives.Zip{},
	}

	// create the archive
	err = format.Archive(context.Background(), out, files)
	if err != nil {
		return
	}

	return filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fileName), nil
}

func (s *autoCodePlugin) InitMenu(menuInfo request.InitMenu) (err error) {
	menuPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", menuInfo.PlugName, "initialize", "menu.go")
	src, err := os.ReadFile(menuPath)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	arrayAst := ast.FindArray(astFile, "model", "SysBaseMenu")
	var menus []system.SysBaseMenu

	parentMenu := []system.SysBaseMenu{
		{
			ParentId:  0,
			Path:      menuInfo.PlugName + "Menu",
			Name:      menuInfo.PlugName + "Menu",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      0,
			Meta: system.Meta{
				Title: menuInfo.ParentMenu,
				Icon:  "school",
			},
		},
	}

	// 查询菜单及其关联的参数和按钮
	err = global.GVA_DB.Preload("Parameters").Preload("MenuBtn").Find(&menus, "id in (?)", menuInfo.Menus).Error
	if err != nil {
		return err
	}
	menus = append(parentMenu, menus...)
	menuExpr := ast.CreateMenuStructAst(menus)
	arrayAst.Elts = *menuExpr

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)

	os.WriteFile(menuPath, bf.Bytes(), 0666)
	return nil
}

func (s *autoCodePlugin) InitAPI(apiInfo request.InitApi) (err error) {
	apiPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", apiInfo.PlugName, "initialize", "api.go")
	src, err := os.ReadFile(apiPath)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	arrayAst := ast.FindArray(astFile, "model", "SysApi")
	var apis []system.SysApi
	err = global.GVA_DB.Find(&apis, "id in (?)", apiInfo.APIs).Error
	if err != nil {
		return err
	}
	apisExpr := ast.CreateApiStructAst(apis)
	arrayAst.Elts = *apisExpr

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)

	os.WriteFile(apiPath, bf.Bytes(), 0666)
	return nil
}

func (s *autoCodePlugin) InitDictionary(dictInfo request.InitDictionary) (err error) {
	dictPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", dictInfo.PlugName, "initialize", "dictionary.go")
	src, err := os.ReadFile(dictPath)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	arrayAst := ast.FindArray(astFile, "model", "SysDictionary")
	var dictionaries []system.SysDictionary
	err = global.GVA_DB.Preload("SysDictionaryDetails").Find(&dictionaries, "id in (?)", dictInfo.Dictionaries).Error
	if err != nil {
		return err
	}
	dictExpr := ast.CreateDictionaryStructAst(dictionaries)
	arrayAst.Elts = *dictExpr

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)

	os.WriteFile(dictPath, bf.Bytes(), 0666)
	return nil
}

func (s *autoCodePlugin) Remove(pluginName string, pluginType string) (err error) {
	// 1. 删除前端代码
	if pluginType == "web" || pluginType == "full" {
		webDir := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", pluginName)
		err = os.RemoveAll(webDir)
		if err != nil {
			return errors.Wrap(err, "删除前端插件目录失败")
		}
	}

	// 2. 删除后端代码
	if pluginType == "server" || pluginType == "full" {
		serverDir := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", pluginName)
		err = os.RemoveAll(serverDir)
		if err != nil {
			return errors.Wrap(err, "删除后端插件目录失败")
		}

		// 移除注册
		removePluginRegisterImport(pluginName)
	}

	// 通过utils 获取 api 菜单 字典
	apis, menus, dicts := pluginUtils.GetPluginData(pluginName)

	// 3. 删除菜单 (递归删除)
	if len(menus) > 0 {
		for _, menu := range menus {
			var dbMenu system.SysBaseMenu
			if err := global.GVA_DB.Where("name = ?", menu.Name).First(&dbMenu).Error; err == nil {
				// 获取该菜单及其所有子菜单的ID
				var menuIds []int
				GetMenuIds(dbMenu, &menuIds)
				// 逆序删除，先删除子菜单
				for i := len(menuIds) - 1; i >= 0; i-- {
					err := BaseMenuServiceApp.DeleteBaseMenu(menuIds[i])
					if err != nil {
						zap.L().Error("删除菜单失败", zap.Int("id", menuIds[i]), zap.Error(err))
					}
				}
			}
		}
	}

	// 4. 删除API
	if len(apis) > 0 {
		for _, api := range apis {
			var dbApi system.SysApi
			if err := global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&dbApi).Error; err == nil {
				err := ApiServiceApp.DeleteApi(dbApi)
				if err != nil {
					zap.L().Error("删除API失败", zap.String("path", api.Path), zap.Error(err))
				}
			}
		}
	}

	// 5. 删除字典
	if len(dicts) > 0 {
		for _, dict := range dicts {
			var dbDict system.SysDictionary
			if err := global.GVA_DB.Where("type = ?", dict.Type).First(&dbDict).Error; err == nil {
				err := DictionaryServiceApp.DeleteSysDictionary(dbDict)
				if err != nil {
					zap.L().Error("删除字典失败", zap.String("type", dict.Type), zap.Error(err))
				}
			}
		}
	}

	return nil
}

func GetMenuIds(menu system.SysBaseMenu, ids *[]int) {
	*ids = append(*ids, int(menu.ID))
	var children []system.SysBaseMenu
	global.GVA_DB.Where("parent_id = ?", menu.ID).Find(&children)
	for _, child := range children {
        // 先递归收集子菜单
		GetMenuIds(child, ids)
	}
}

func removePluginRegisterImport(packageName string) error {
	module := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module)
	if module == "" {
		return errors.New("autocode module is empty")
	}
	if packageName == "" {
		return errors.New("plugin package is empty")
	}

	registerPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "register.go")
	src, err := os.ReadFile(registerPath)
	if err != nil {
		return err
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, registerPath, src, parser.ParseComments)
	if err != nil {
		return err
	}

	importPath := fmt.Sprintf("%s/plugin/%s", module, packageName)
    importLit := fmt.Sprintf("%q", importPath)

    // 移除 import
	var newDecls []goast.Decl
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*goast.GenDecl)
		if !ok {
			newDecls = append(newDecls, decl)
			continue
		}
		if genDecl.Tok == token.IMPORT {
			var newSpecs []goast.Spec
			for _, spec := range genDecl.Specs {
				importSpec, ok := spec.(*goast.ImportSpec)
				if !ok {
					newSpecs = append(newSpecs, spec)
					continue
				}
				if importSpec.Path.Value != importLit {
					newSpecs = append(newSpecs, spec)
				}
			}
            // 如果还有其他import，保留该 decl
			if len(newSpecs) > 0 {
				genDecl.Specs = newSpecs
				newDecls = append(newDecls, genDecl)
			}
		} else {
			newDecls = append(newDecls, decl)
		}
	}
	astFile.Decls = newDecls

	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)

	return os.WriteFile(registerPath, bf.Bytes(), 0666)
}
