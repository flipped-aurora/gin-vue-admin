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
	if err := utils.ValidatePluginName(packageName); err != nil {
		return err
	}
	module := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module)
	if module == "" {
		return errors.New("autocode module is empty")
	}
	serverPluginRoot, err := utils.JoinWithinRoot(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin")
	if err != nil {
		return err
	}
	registerPath, err := utils.JoinWithinRoot(serverPluginRoot, "register.go")
	if err != nil {
		return err
	}
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
	if err := printer.Fprint(bf, fileSet, astFile); err != nil {
		return err
	}

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
	if err = utils.ValidatePluginName(plugName); err != nil {
		return "", err
	}
	webPath, err := pluginPath(global.GVA_CONFIG.AutoCode.Web, plugName)
	if err != nil {
		return "", err
	}
	serverPath, err := pluginPath(global.GVA_CONFIG.AutoCode.Server, plugName)
	if err != nil {
		return "", err
	}
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

func (s *autoCodePlugin) InitMenu(ctx context.Context, menuInfo request.InitMenu) (err error) {
	menuPath, fileSet, astFile, arrayAst, err := loadPluginInitializeArray(menuInfo.PlugName, "menu.go", "SysBaseMenu")
	if err != nil {
		return err
	}
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
	err = global.GVA_DB.WithContext(ctx).Preload("Parameters").Preload("MenuBtn").Find(&menus, "id in (?)", menuInfo.Menus).Error
	if err != nil {
		return err
	}
	menus = append(parentMenu, menus...)
	menuExpr := ast.CreateMenuStructAst(menus)
	arrayAst.Elts = *menuExpr

	return writePluginInitializeFile(menuPath, fileSet, astFile)
}

func (s *autoCodePlugin) InitAPI(ctx context.Context, apiInfo request.InitApi) (err error) {
	apiPath, fileSet, astFile, arrayAst, err := loadPluginInitializeArray(apiInfo.PlugName, "api.go", "SysApi")
	if err != nil {
		return err
	}
	var apis []system.SysApi
	err = global.GVA_DB.WithContext(ctx).Find(&apis, "id in (?)", apiInfo.APIs).Error
	if err != nil {
		return err
	}
	apisExpr := ast.CreateApiStructAst(apis)
	arrayAst.Elts = *apisExpr

	return writePluginInitializeFile(apiPath, fileSet, astFile)
}

func (s *autoCodePlugin) InitDictionary(ctx context.Context, dictInfo request.InitDictionary) (err error) {
	dictPath, fileSet, astFile, arrayAst, err := loadPluginInitializeArray(dictInfo.PlugName, "dictionary.go", "SysDictionary")
	if err != nil {
		return err
	}
	var dictionaries []system.SysDictionary
	err = global.GVA_DB.WithContext(ctx).Preload("SysDictionaryDetails").Find(&dictionaries, "id in (?)", dictInfo.Dictionaries).Error
	if err != nil {
		return err
	}
	dictExpr := ast.CreateDictionaryStructAst(dictionaries)
	arrayAst.Elts = *dictExpr

	return writePluginInitializeFile(dictPath, fileSet, astFile)
}

func (s *autoCodePlugin) Remove(ctx context.Context, pluginName string, pluginType string) (err error) {
	if err = utils.ValidatePluginName(pluginName); err != nil {
		return err
	}
	if pluginType != "web" && pluginType != "server" && pluginType != "full" {
		return errors.New("invalid plugin type")
	}

	var webDir, serverDir string
	if pluginType == "web" || pluginType == "full" {
		webDir, err = pluginPath(global.GVA_CONFIG.AutoCode.Web, pluginName)
		if err != nil {
			return err
		}
	}
	if pluginType == "server" || pluginType == "full" {
		serverDir, err = pluginPath(global.GVA_CONFIG.AutoCode.Server, pluginName)
		if err != nil {
			return err
		}
	}

	// 1. 删除前端代码
	if pluginType == "web" || pluginType == "full" {
		err = os.RemoveAll(webDir)
		if err != nil {
			return errors.Wrap(err, "删除前端插件目录失败")
		}
	}

	// 2. 删除后端代码
	if pluginType == "server" || pluginType == "full" {
		err = os.RemoveAll(serverDir)
		if err != nil {
			return errors.Wrap(err, "删除后端插件目录失败")
		}

		// 移除注册
		if err = removePluginRegisterImport(pluginName); err != nil {
			return errors.Wrap(err, "移除插件注册失败")
		}
	}

	// 通过utils 获取 api 菜单 字典
	apis, menus, dicts := pluginUtils.GetPluginData(pluginName)

	// DB 清理阶段脱离请求取消:上面 1/2 的文件删除不可逆且已完成,若客户端此刻
	// 断连,请求 ctx 取消会让下面的清理全部快速失败(循环内错误只记日志不中断),
	// 菜单/API/字典静默残留;且进程重启后插件不再注册,GetPluginData 拿不到
	// 数据,孤儿再也无法通过 Remove 清掉。WithoutCancel 保留链路字段、剥离取消。
	cleanupCtx := context.WithoutCancel(ctx)

	// 3. 删除菜单 (递归删除)
	if len(menus) > 0 {
		for _, menu := range menus {
			var dbMenu system.SysBaseMenu
			if err := global.GVA_DB.WithContext(cleanupCtx).Where("name = ?", menu.Name).First(&dbMenu).Error; err == nil {
				// 获取该菜单及其所有子菜单的ID
				var menuIds []int
				GetMenuIds(cleanupCtx, dbMenu, &menuIds)
				// 逆序删除，先删除子菜单
				for i := len(menuIds) - 1; i >= 0; i-- {
					err := BaseMenuServiceApp.DeleteBaseMenu(cleanupCtx, menuIds[i])
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
			if err := global.GVA_DB.WithContext(cleanupCtx).Where("path = ? AND method = ?", api.Path, api.Method).First(&dbApi).Error; err == nil {
				err := ApiServiceApp.DeleteApi(cleanupCtx, dbApi)
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
			if err := global.GVA_DB.WithContext(cleanupCtx).Where("type = ?", dict.Type).First(&dbDict).Error; err == nil {
				err := DictionaryServiceApp.DeleteSysDictionary(cleanupCtx, dbDict)
				if err != nil {
					zap.L().Error("删除字典失败", zap.String("type", dict.Type), zap.Error(err))
				}
			}
		}
	}

	return nil
}

func GetMenuIds(ctx context.Context, menu system.SysBaseMenu, ids *[]int) {
	*ids = append(*ids, int(menu.ID))
	var children []system.SysBaseMenu
	global.GVA_DB.WithContext(ctx).Where("parent_id = ?", menu.ID).Find(&children)
	for _, child := range children {
		// 先递归收集子菜单
		GetMenuIds(ctx, child, ids)
	}
}

func removePluginRegisterImport(packageName string) error {
	if err := utils.ValidatePluginName(packageName); err != nil {
		return err
	}
	module := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module)
	if module == "" {
		return errors.New("autocode module is empty")
	}
	serverPluginRoot, err := utils.JoinWithinRoot(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin")
	if err != nil {
		return err
	}
	registerPath, err := utils.JoinWithinRoot(serverPluginRoot, "register.go")
	if err != nil {
		return err
	}
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
	if err := printer.Fprint(bf, fileSet, astFile); err != nil {
		return err
	}

	return os.WriteFile(registerPath, bf.Bytes(), 0666)
}

func pluginPath(component, pluginName string, elems ...string) (string, error) {
	if err := utils.ValidatePluginName(pluginName); err != nil {
		return "", err
	}
	root, err := utils.JoinWithinRoot(global.GVA_CONFIG.AutoCode.Root, component, "plugin")
	if err != nil {
		return "", err
	}
	parts := append([]string{pluginName}, elems...)
	return utils.JoinWithinRoot(root, parts...)
}

func loadPluginInitializeArray(pluginName, fileName, selectorName string) (string, *token.FileSet, *goast.File, *goast.CompositeLit, error) {
	path, err := pluginPath(global.GVA_CONFIG.AutoCode.Server, pluginName, "initialize", fileName)
	if err != nil {
		return "", nil, nil, nil, err
	}
	src, err := os.ReadFile(path)
	if err != nil {
		return "", nil, nil, nil, errors.Wrap(err, "读取插件初始化文件失败")
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, path, src, parser.ParseComments)
	if err != nil {
		return "", nil, nil, nil, errors.Wrap(err, "解析插件初始化文件失败")
	}
	arrayAst := ast.FindArray(astFile, "model", selectorName)
	if arrayAst == nil {
		return "", nil, nil, nil, errors.Errorf("插件初始化文件缺少 []model.%s 数组", selectorName)
	}
	return path, fileSet, astFile, arrayAst, nil
}

func writePluginInitializeFile(path string, fileSet *token.FileSet, astFile *goast.File) error {
	var out bytes.Buffer
	if err := printer.Fprint(&out, fileSet, astFile); err != nil {
		return errors.Wrap(err, "打印插件初始化文件失败")
	}
	if err := os.WriteFile(path, out.Bytes(), 0666); err != nil {
		return errors.Wrap(err, "写入插件初始化文件失败")
	}
	return nil
}
