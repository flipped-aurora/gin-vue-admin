package system

import (
	"bytes"
	"context"
	"fmt"
	goast "go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
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

func (s *autoCodePlugin) Install(file *multipart.FileHeader, parentPlugin string) (web, server int, err error) {
	if parentPlugin == "" {
		return s.installTopLevel(file)
	}
	if err = utils.ValidatePluginName(parentPlugin); err != nil {
		return -1, -1, err
	}
	return s.installSubPlugin(file, parentPlugin)
}

// Install 插件安装
func (s *autoCodePlugin) installTopLevel(file *multipart.FileHeader) (web, server int, err error) {
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

func (s *autoCodePlugin) installSubPlugin(file *multipart.FileHeader, parentPlugin string) (web, server int, err error) {
	serverRoot, err := utils.JoinWithinRoot(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server)
	if err != nil {
		return -1, -1, err
	}
	if err = os.MkdirAll(serverRoot, 0755); err != nil {
		return -1, -1, err
	}
	tempDir, err := os.MkdirTemp(serverRoot, "gva-plugin-")
	if err != nil {
		return -1, -1, err
	}
	defer os.RemoveAll(tempDir)

	src, err := file.Open()
	if err != nil {
		return -1, -1, err
	}
	defer src.Close()

	archiveName := filepath.Base(strings.TrimSpace(file.Filename))
	if archiveName == "." || archiveName == string(filepath.Separator) {
		return -1, -1, errors.New("invalid plugin archive name")
	}
	archivePath, err := utils.JoinWithinRoot(tempDir, archiveName)
	if err != nil {
		return -1, -1, err
	}
	out, err := os.Create(archivePath)
	if err != nil {
		return -1, -1, err
	}
	if _, err = io.Copy(out, src); err != nil {
		out.Close()
		return -1, -1, err
	}
	if err = out.Close(); err != nil {
		return -1, -1, err
	}

	paths, err := utils.Unzip(archivePath, tempDir)
	if err != nil {
		return -1, -1, err
	}
	archive, err := findPluginArchive(tempDir, filterFile(paths))
	if err != nil {
		return -1, -1, err
	}
	if archive.server != nil && archive.web != nil && archive.server.name != archive.web.name {
		return -1, -1, errors.New("server and web plugin names must match")
	}

	var serverTarget, webTarget string
	if archive.server != nil {
		serverTarget, err = pluginInstallRoot(global.GVA_CONFIG.AutoCode.Server, parentPlugin)
		if err != nil {
			return -1, -1, err
		}
		if err = ensurePluginTargetAvailable(serverTarget, archive.server.name); err != nil {
			return -1, -1, err
		}
	}
	if archive.web != nil {
		webTarget, err = pluginInstallRoot(global.GVA_CONFIG.AutoCode.Web, parentPlugin)
		if err != nil {
			return -1, -1, err
		}
		if err = ensurePluginTargetAvailable(webTarget, archive.web.name); err != nil {
			return -1, -1, err
		}
	}

	if archive.server != nil {
		if err = prepareSubPluginSource(archive.server.path(), parentPlugin, archive.server.name); err != nil {
			return -1, -1, err
		}
		if err = copyPluginArchive(archive.server.root, serverTarget); err != nil {
			return -1, -1, err
		}
		if err = ensureParentSubPluginRegistration(parentPlugin, archive.server.name); err != nil {
			return -1, -1, err
		}
		server = 1
	}
	if archive.web != nil {
		if err = prepareSubPluginWebSource(archive.web.path(), parentPlugin, archive.web.name); err != nil {
			return -1, server, err
		}
		if err = copyPluginArchive(archive.web.root, webTarget); err != nil {
			return -1, server, err
		}
		web = 1
	}
	return web, server, nil
}

type pluginArchive struct {
	server *pluginArchivePart
	web    *pluginArchivePart
}

type pluginArchivePart struct {
	root string
	name string
}

func (p *pluginArchivePart) path() string {
	return filepath.Join(p.root, p.name)
}

func findPluginArchive(tempDir string, paths []string) (pluginArchive, error) {
	archive := pluginArchive{}
	for _, path := range paths {
		relativePath, err := filepath.Rel(tempDir, path)
		if err != nil {
			return pluginArchive{}, err
		}
		parts := strings.Split(filepath.ToSlash(relativePath), "/")
		for index := 0; index+2 < len(parts); index++ {
			if parts[index+1] != "plugin" || (parts[index] != "server" && parts[index] != "web") {
				continue
			}
			name := parts[index+2]
			if err := utils.ValidatePluginName(name); err != nil {
				return pluginArchive{}, err
			}
			part := &pluginArchivePart{
				root: filepath.Join(tempDir, filepath.FromSlash(strings.Join(parts[:index+2], "/"))),
				name: name,
			}
			if parts[index] == "server" {
				if archive.server != nil && (archive.server.name != part.name || archive.server.root != part.root) {
					return pluginArchive{}, errors.New("plugin archive contains multiple server plugins")
				}
				archive.server = part
			} else {
				if archive.web != nil && (archive.web.name != part.name || archive.web.root != part.root) {
					return pluginArchive{}, errors.New("plugin archive contains multiple web plugins")
				}
				archive.web = part
			}
			break
		}
	}
	if archive.server == nil && archive.web == nil {
		return pluginArchive{}, errors.New("invalid plugin archive")
	}
	return archive, nil
}

func pluginInstallRoot(component, parentPlugin string) (string, error) {
	root, err := utils.JoinWithinRoot(global.GVA_CONFIG.AutoCode.Root, component, "plugin")
	if err != nil || parentPlugin == "" {
		return root, err
	}
	parentPath, err := utils.JoinWithinRoot(root, parentPlugin)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(parentPath)
	if err != nil {
		return "", errors.Wrap(err, "parent plugin does not exist")
	}
	if !info.IsDir() {
		return "", errors.New("parent plugin is not a directory")
	}
	return utils.JoinWithinRoot(parentPath, "subPlugin")
}

func ensurePluginTargetAvailable(root, pluginName string) error {
	target, err := utils.JoinWithinRoot(root, pluginName)
	if err != nil {
		return err
	}
	if _, err = os.Stat(target); err == nil {
		return errors.New("plugin already exists")
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}

func copyPluginArchive(sourceRoot, destinationRoot string) error {
	if err := os.MkdirAll(destinationRoot, 0755); err != nil {
		return err
	}
	return cp.Copy(sourceRoot, destinationRoot, cp.Options{Skip: skipMacSpecialDocument})
}

func prepareSubPluginSource(childPath, parentPlugin, childPlugin string) error {
	module := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module)
	if module == "" {
		return errors.New("autocode module is empty")
	}
	oldImportPrefix := fmt.Sprintf("%s/plugin/%s", module, childPlugin)
	newImportPrefix := fmt.Sprintf("%s/plugin/%s/subPlugin/%s", module, parentPlugin, childPlugin)
	err := filepath.Walk(childPath, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		return rewriteGoImportPaths(path, oldImportPrefix, newImportPrefix)
	})
	if err != nil {
		return err
	}
	entryPath := filepath.Join(childPath, "plugin.go")
	if err = suppressSubPluginSelfRegistration(entryPath); err != nil {
		return err
	}
	menuPath := filepath.Join(childPath, "initialize", "menu.go")
	if _, statErr := os.Stat(menuPath); statErr == nil {
		return rewriteSubPluginMenuComponentPaths(menuPath, parentPlugin, childPlugin)
	} else if !os.IsNotExist(statErr) {
		return statErr
	}
	return nil
}

func prepareSubPluginWebSource(childPath, parentPlugin, childPlugin string) error {
	return filepath.Walk(childPath, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() || !isSubPluginWebSourceFile(path) {
			return nil
		}
		return rewriteSubPluginWebImportPaths(path, parentPlugin, childPlugin)
	})
}

func isSubPluginWebSourceFile(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".js", ".mjs", ".cjs", ".ts", ".mts", ".cts", ".jsx", ".tsx", ".vue":
		return true
	default:
		return false
	}
}

func rewriteSubPluginWebImportPaths(path, parentPlugin, childPlugin string) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	legacyPrefix := fmt.Sprintf("@/plugin/%s", childPlugin)
	nestedPrefix := fmt.Sprintf("@/plugin/%s/subPlugin/%s", parentPlugin, childPlugin)
	rewritten := rewriteSubPluginWebAliasPrefix(string(source), legacyPrefix, nestedPrefix)
	if rewritten == string(source) {
		return nil
	}
	return os.WriteFile(path, []byte(rewritten), 0666)
}

func rewriteSubPluginWebAliasPrefix(source, legacyPrefix, nestedPrefix string) string {
	var rewritten strings.Builder
	rewritten.Grow(len(source) + len(nestedPrefix))

	position := 0
	for {
		index := strings.Index(source[position:], legacyPrefix)
		if index < 0 {
			rewritten.WriteString(source[position:])
			return rewritten.String()
		}
		index += position
		prefixEnd := index + len(legacyPrefix)
		if strings.HasPrefix(source[index:], nestedPrefix) || (prefixEnd < len(source) && isPluginNameCharacter(source[prefixEnd])) {
			rewritten.WriteString(source[position:prefixEnd])
			position = prefixEnd
			continue
		}

		rewritten.WriteString(source[position:index])
		rewritten.WriteString(nestedPrefix)
		position = prefixEnd
	}
}

func isPluginNameCharacter(value byte) bool {
	return value == '_' || (value >= 'a' && value <= 'z') || (value >= '0' && value <= '9')
}

func rewriteGoImportPaths(path, oldPrefix, newPrefix string) error {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	changed := false
	for _, spec := range astFile.Imports {
		importPath, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return err
		}
		if importPath == oldPrefix || strings.HasPrefix(importPath, oldPrefix+"/") {
			spec.Path.Value = strconv.Quote(newPrefix + strings.TrimPrefix(importPath, oldPrefix))
			changed = true
		}
	}
	if !changed {
		return nil
	}
	return writeGoFile(path, fileSet, astFile)
}

func suppressSubPluginSelfRegistration(path string) error {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	registryAliases := map[string]bool{}
	registryPath := fmt.Sprintf("%s/utils/plugin/v2", strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module))
	for _, spec := range astFile.Imports {
		importPath, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return err
		}
		if importPath != registryPath {
			continue
		}
		if spec.Name == nil {
			registryAliases["plugin"] = true
		} else {
			registryAliases[spec.Name.Name] = true
		}
	}

	changed := false
	decls := make([]goast.Decl, 0, len(astFile.Decls))
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*goast.FuncDecl)
		if !ok || funcDecl.Name.Name != "init" || funcDecl.Body == nil {
			decls = append(decls, decl)
			continue
		}
		statements := make([]goast.Stmt, 0, len(funcDecl.Body.List))
		for _, statement := range funcDecl.Body.List {
			if isPluginRegistryRegisterStatement(statement, registryAliases) {
				changed = true
				continue
			}
			statements = append(statements, statement)
		}
		if len(statements) > 0 {
			funcDecl.Body.List = statements
			decls = append(decls, funcDecl)
		} else if len(funcDecl.Body.List) > 0 {
			changed = true
		}
	}
	if !changed {
		return nil
	}
	astFile.Decls = decls
	return writeGoFile(path, fileSet, astFile)
}

func isPluginRegistryRegisterStatement(statement goast.Stmt, registryAliases map[string]bool) bool {
	expressionStatement, ok := statement.(*goast.ExprStmt)
	if !ok {
		return false
	}
	call, ok := expressionStatement.X.(*goast.CallExpr)
	if !ok {
		return false
	}
	selector, ok := call.Fun.(*goast.SelectorExpr)
	if !ok || selector.Sel.Name != "Register" {
		return false
	}
	ident, ok := selector.X.(*goast.Ident)
	return ok && registryAliases[ident.Name]
}

func rewriteSubPluginMenuComponentPaths(path, parentPlugin, childPlugin string) error {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	legacyPrefix := fmt.Sprintf("plugin/%s/", childPlugin)
	nestedPrefix := fmt.Sprintf("plugin/%s/subPlugin/%s/", parentPlugin, childPlugin)
	changed := false
	goast.Inspect(astFile, func(node goast.Node) bool {
		literal, ok := node.(*goast.BasicLit)
		if !ok || literal.Kind != token.STRING {
			return true
		}
		value, err := strconv.Unquote(literal.Value)
		if err != nil || strings.HasPrefix(value, nestedPrefix) || !strings.HasPrefix(value, legacyPrefix) {
			return true
		}
		literal.Value = strconv.Quote(nestedPrefix + strings.TrimPrefix(value, legacyPrefix))
		changed = true
		return true
	})
	if !changed {
		return nil
	}
	return writeGoFile(path, fileSet, astFile)
}

func ensureParentSubPluginRegistration(parentPlugin, childPlugin string) error {
	module := strings.TrimSpace(global.GVA_CONFIG.AutoCode.Module)
	if module == "" {
		return errors.New("autocode module is empty")
	}
	parentPath, err := pluginPath(global.GVA_CONFIG.AutoCode.Server, parentPlugin, "plugin.go")
	if err != nil {
		return err
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, parentPath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	childImportPath := fmt.Sprintf("%s/plugin/%s/subPlugin/%s", module, parentPlugin, childPlugin)
	alias := "subPlugin_" + childPlugin
	imported := false
	for _, spec := range astFile.Imports {
		importPath, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return err
		}
		if importPath != childImportPath {
			continue
		}
		imported = true
		if spec.Name == nil {
			alias = childPlugin
		} else {
			alias = spec.Name.Name
		}
	}
	if alias == "_" || alias == "." {
		return errors.New("child plugin import cannot be blank or dot imported")
	}
	if !imported {
		importSpec := &goast.ImportSpec{
			Name: goast.NewIdent(alias),
			Path: &goast.BasicLit{Kind: token.STRING, Value: strconv.Quote(childImportPath)},
		}
		if importDecl := firstImportDeclaration(astFile); importDecl == nil {
			astFile.Decls = append([]goast.Decl{
				&goast.GenDecl{Tok: token.IMPORT, Specs: []goast.Spec{importSpec}},
			}, astFile.Decls...)
		} else {
			importDecl.Specs = append(importDecl.Specs, importSpec)
		}
	}

	registerFunc := findPluginRegisterFunction(astFile)
	if registerFunc == nil {
		return errors.New("parent plugin does not implement Register")
	}
	if hasChildPluginRegisterCall(registerFunc.Body, alias) {
		return nil
	}
	groupName := "group"
	if registerFunc.Type.Params != nil && len(registerFunc.Type.Params.List) > 0 && len(registerFunc.Type.Params.List[0].Names) > 0 {
		groupName = registerFunc.Type.Params.List[0].Names[0].Name
	}
	registerFunc.Body.List = append(registerFunc.Body.List, &goast.ExprStmt{
		X: &goast.CallExpr{
			Fun: &goast.SelectorExpr{
				X:   &goast.SelectorExpr{X: goast.NewIdent(alias), Sel: goast.NewIdent("Plugin")},
				Sel: goast.NewIdent("Register"),
			},
			Args: []goast.Expr{goast.NewIdent(groupName)},
		},
	})
	return writeGoFile(parentPath, fileSet, astFile)
}

func firstImportDeclaration(astFile *goast.File) *goast.GenDecl {
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*goast.GenDecl)
		if ok && genDecl.Tok == token.IMPORT {
			return genDecl
		}
	}
	return nil
}

func findPluginRegisterFunction(astFile *goast.File) *goast.FuncDecl {
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*goast.FuncDecl)
		if ok && funcDecl.Name.Name == "Register" && funcDecl.Body != nil {
			return funcDecl
		}
	}
	return nil
}

func hasChildPluginRegisterCall(body *goast.BlockStmt, alias string) bool {
	for _, statement := range body.List {
		expressionStatement, ok := statement.(*goast.ExprStmt)
		if !ok {
			continue
		}
		call, ok := expressionStatement.X.(*goast.CallExpr)
		if !ok {
			continue
		}
		registerSelector, ok := call.Fun.(*goast.SelectorExpr)
		if !ok || registerSelector.Sel.Name != "Register" {
			continue
		}
		pluginSelector, ok := registerSelector.X.(*goast.SelectorExpr)
		if !ok || pluginSelector.Sel.Name != "Plugin" {
			continue
		}
		ident, ok := pluginSelector.X.(*goast.Ident)
		if ok && ident.Name == alias {
			return true
		}
	}
	return false
}

func writeGoFile(path string, fileSet *token.FileSet, astFile *goast.File) error {
	var out bytes.Buffer
	if err := format.Node(&out, fileSet, astFile); err != nil {
		return err
	}
	return os.WriteFile(path, out.Bytes(), 0666)
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
