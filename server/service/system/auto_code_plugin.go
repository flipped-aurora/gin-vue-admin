package system

import (
	"bytes"
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/mholt/archives"
	cp "github.com/otiai10/copy"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
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
	webPlugin := ""
	serverPlugin := ""

	for i := range paths {
		paths[i] = filepath.ToSlash(paths[i])
		pathArr := strings.Split(paths[i], "/")
		ln := len(pathArr)

		if ln < 4 {
			continue
		}
		if pathArr[2]+"/"+pathArr[3] == `server/plugin` && len(serverPlugin) == 0 {
			serverPlugin = filepath.Join(pathArr[0], pathArr[1], pathArr[2], pathArr[3])
		}
		if pathArr[2]+"/"+pathArr[3] == `web/plugin` && len(webPlugin) == 0 {
			webPlugin = filepath.Join(pathArr[0], pathArr[1], pathArr[2], pathArr[3])
		}
	}
	if len(serverPlugin) == 0 && len(webPlugin) == 0 {
		zap.L().Error(global.Translate("sys_auto_code.nonStandardPlugin"))
		return webIndex, serverIndex, errors.New(global.Translate("sys_auto_code.nonStandardPlugin"))
	}

	if len(serverPlugin) != 0 {
		err = installation(serverPlugin, global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Server)
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
		zap.L().Error(global.Translate("sys_auto_code.autoPathExists"), zap.String("to", to))
		return errors.New(toPath + global.Translate("sys_auto_code.duplicatePlugin"))
	}
	return cp.Copy(form, to, cp.Options{Skip: skipMacSpecialDocument})
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
		return "", errors.New(global.Translate("sys_auto_code.pluginNameRequired"))
	}

	// 防止路径穿越
	plugName = filepath.Clean(plugName)

	webPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", plugName)
	serverPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", plugName)
	// 创建一个新的zip文件

	// 判断目录是否存在
	_, err = os.Stat(webPath)
	if err != nil {
		return "", errors.New(global.Translate("sys_auto_code.webPathNotExist"))
	}
	_, err = os.Stat(serverPath)
	if err != nil {
		return "", errors.New(global.Translate("sys_auto_code.serverPathNotExist"))
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
