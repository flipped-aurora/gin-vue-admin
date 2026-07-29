package system

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/autocode"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	utilsAst "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/pkg/errors"
)

var AutoCodeTemplate = new(autoCodeTemplate)
var autoCodeCreateMu sync.Mutex

type autoCodeTemplate struct{}

func autoCodeCreateContext(ctx context.Context) context.Context {
	return context.WithoutCancel(ctx)
}

func (s *autoCodeTemplate) checkPackage(Pkg string, template string) (err error) {
	switch template {
	case "package":
		apiEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", Pkg, "enter.go")
		_, err = os.Stat(apiEnter)
		if err != nil {
			return fmt.Errorf("package结构异常,缺少api/v1/%s/enter.go", Pkg)
		}
		serviceEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", Pkg, "enter.go")
		_, err = os.Stat(serviceEnter)
		if err != nil {
			return fmt.Errorf("package结构异常,缺少service/%s/enter.go", Pkg)
		}
		routerEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", Pkg, "enter.go")
		_, err = os.Stat(routerEnter)
		if err != nil {
			return fmt.Errorf("package结构异常,缺少router/%s/enter.go", Pkg)
		}
	case "plugin":
		pluginEnter := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", Pkg, "plugin.go")
		_, err = os.Stat(pluginEnter)
		if err != nil {
			return fmt.Errorf("plugin结构异常,缺少plugin/%s/plugin.go", Pkg)
		}
	}
	return nil
}

// Create 创建生成自动化代码
func (s *autoCodeTemplate) Create(ctx context.Context, info request.AutoCode) error {
	autoCodeCreateMu.Lock()
	defer autoCodeCreateMu.Unlock()

	// Frontend files are published last, but their Vite refresh may still cancel
	// the HTTP request before the handler returns. The accepted create task must finish.
	createCtx := autoCodeCreateContext(ctx)
	var autoPkg model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(createCtx).Where("package_name = ?", info.Package).First(&autoPkg).Error
	if err != nil {
		return errors.Wrap(err, "查询包失败!")
	}
	err = s.checkPackage(info.Package, autoPkg.Template)
	if err != nil {
		return err
	}
	exists, err := autoCodeIdentityExists(createCtx, global.GVA_DB, info)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("已经创建过此数据结构,请勿重复创建!")
	}

	layout, err := newAutoCodeTaskLayout(
		global.GVA_CONFIG.AutoCode.Root,
		global.GVA_CONFIG.AutoCode.Server,
		global.GVA_CONFIG.AutoCode.WebRoot(),
	)
	if err != nil {
		return err
	}
	generated, templates, injections, err := s.generate(createCtx, info, autoPkg)
	if err != nil {
		return err
	}
	history := info.History()
	history.Templates = templates
	history.Injections = make(map[string]string, len(injections))
	for key, value := range injections {
		bytes, marshalErr := json.Marshal(value)
		if marshalErr != nil {
			return errors.Wrapf(marshalErr, "序列化自动代码注入信息 %s 失败", key)
		}
		history.Injections[key] = string(bytes)
	}
	files := make(map[string][]byte, len(generated))
	for target, builder := range generated {
		files[target] = []byte(builder.String())
	}
	fileTask, err := prepareAutoCodeFileTask(layout, files)
	if err != nil {
		return err
	}
	return commitAutoCodeFileTask(fileTask, publishPreparedAutoCodeFile, func() error {
		return persistAutoCodeMetadata(createCtx, global.GVA_DB, info, autoPkg.Template, history)
	})
}

// Preview 预览自动化代码
func (s *autoCodeTemplate) Preview(ctx context.Context, info request.AutoCode) (map[string]string, error) {
	var entity model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(ctx).Where("package_name = ?", info.Package).First(&entity).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询包失败!")
	}
	// 增加判断: 重复创建struct 或者重复的简称
	if AutocodeHistory.Repeat(ctx, info.BusinessDB, info.StructName, info.Abbreviation, info.Package) && !info.IsAdd {
		return nil, errors.New("已经创建过此数据结构或重复简称,请勿重复创建!")
	}

	preview := make(map[string]string)
	codes, _, _, err := s.generate(ctx, info, entity)
	if err != nil {
		return nil, err
	}
	for key, writer := range codes {
		if len(key) > len(global.GVA_CONFIG.AutoCode.Root) {
			key, _ = filepath.Rel(global.GVA_CONFIG.AutoCode.Root, key)
		}
		// 获取key的后缀 取消.
		suffix := filepath.Ext(key)[1:]
		var builder strings.Builder
		builder.WriteString("```" + suffix + "\n\n")
		builder.WriteString(writer.String())
		builder.WriteString("\n\n```")
		preview[key] = builder.String()
	}
	return preview, nil
}

func (s *autoCodeTemplate) generate(ctx context.Context, info request.AutoCode, entity model.SysAutoCodePackage) (map[string]strings.Builder, map[string]string, map[string]utilsAst.Ast, error) {
	templates, asts, _, err := AutoCodePackage.templates(ctx, entity, info, false)
	if err != nil {
		return nil, nil, nil, err
	}
	code := make(map[string]strings.Builder)
	for key, create := range templates {
		var files *template.Template
		files, err = template.New(filepath.Base(key)).Funcs(autocode.GetTemplateFuncMap()).ParseFiles(key)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "[filpath:%s]读取模版文件失败!", key)
		}
		var builder strings.Builder
		err = files.Execute(&builder, info)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "[filpath:%s]生成文件失败!", create)
		}
		code[create] = builder
	} // 生成文件
	injectedCode, injections, err := renderAutoCodeInjections(info, asts)
	if err != nil {
		return nil, nil, nil, err
	}
	for key, builder := range injectedCode {
		code[key] = builder
	}
	// 注入代码
	return code, templates, injections, nil
}

func renderAutoCodeInjections(info request.AutoCode, asts map[string]utilsAst.Ast) (map[string]strings.Builder, map[string]utilsAst.Ast, error) {
	code := make(map[string]strings.Builder, len(asts))
	injections := make(map[string]utilsAst.Ast, len(asts))
	for key, value := range asts {
		keys := strings.Split(key, "=>")
		if len(keys) == 2 {
			if keys[1] == utilsAst.TypePluginInitializeV2 {
				continue
			}
			if info.OnlyTemplate {
				if keys[1] == utilsAst.TypePackageInitializeGorm || keys[1] == utilsAst.TypePluginInitializeGorm {
					continue
				}
			}
			if !info.AutoMigrate {
				if keys[1] == utilsAst.TypePackageInitializeGorm || keys[1] == utilsAst.TypePluginInitializeGorm {
					continue
				}
			}
			var builder strings.Builder
			parsed, err := value.Parse("", &builder)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "[filepath:%s]解析注入目标失败", keys[0])
			}
			if parsed == nil {
				return nil, nil, errors.Errorf("[filepath:%s]解析注入目标为空", keys[0])
			}
			if err = value.Injection(parsed); err != nil {
				return nil, nil, errors.Wrapf(err, "[filepath:%s]注入代码失败", keys[0])
			}
			if err = value.Format("", &builder, parsed); err != nil {
				return nil, nil, errors.Wrapf(err, "[filepath:%s]格式化注入代码失败", keys[0])
			}
			code[keys[0]] = builder
			injections[keys[1]] = value
			fmt.Println(keys[0], "注入成功!")
		}
	}
	return code, injections, nil
}

func (s *autoCodeTemplate) AddFunc(ctx context.Context, info request.AutoFunc) error {
	autoPkg := model.SysAutoCodePackage{}
	err := global.GVA_DB.WithContext(ctx).First(&autoPkg, "package_name = ?", info.Package).Error
	if err != nil {
		return err
	}
	if autoPkg.Template != "package" {
		info.IsPlugin = true
	}
	err = s.addTemplateToFile("api.go", info)
	if err != nil {
		return err
	}
	err = s.addTemplateToFile("server.go", info)
	if err != nil {
		return err
	}
	err = s.addTemplateToFile("api.js", info)
	if err != nil {
		return err
	}
	return s.addTemplateToAst("router", info)
}

func (s *autoCodeTemplate) GetApiAndServer(ctx context.Context, info request.AutoFunc) (map[string]string, error) {
	autoPkg := model.SysAutoCodePackage{}
	err := global.GVA_DB.WithContext(ctx).First(&autoPkg, "package_name = ?", info.Package).Error
	if err != nil {
		return nil, err
	}
	if autoPkg.Template != "package" {
		info.IsPlugin = true
	}

	apiStr, err := s.getTemplateStr("api.go", info)
	if err != nil {
		return nil, err
	}
	serverStr, err := s.getTemplateStr("server.go", info)
	if err != nil {
		return nil, err
	}
	jsStr, err := s.getTemplateStr("api.js", info)
	if err != nil {
		return nil, err
	}
	return map[string]string{"api": apiStr, "server": serverStr, "js": jsStr}, nil

}

func (s *autoCodeTemplate) getTemplateStr(t string, info request.AutoFunc) (string, error) {
	tempPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", "function", t+".tpl")
	files, err := template.New(filepath.Base(tempPath)).Funcs(autocode.GetTemplateFuncMap()).ParseFiles(tempPath)
	if err != nil {
		return "", errors.Wrapf(err, "[filepath:%s]读取模版文件失败!", tempPath)
	}
	var builder strings.Builder
	err = files.Execute(&builder, info)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.Wrapf(err, "[filpath:%s]生成文件失败!", tempPath)
	}
	return builder.String(), nil
}

func (s *autoCodeTemplate) addTemplateToAst(t string, info request.AutoFunc) error {
	if !isSafeFileName(info.HumpPackageName) {
		return fmt.Errorf("文件名包含非法字符，拒绝写入")
	}
	tPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", info.Package, info.HumpPackageName+".go")
	funcName := fmt.Sprintf("Init%sRouter", info.StructName)

	routerStr := "RouterWithoutAuth"
	if info.IsAuth {
		routerStr = "Router"
	}

	stmtStr := fmt.Sprintf("%s%s.%s(\"%s\", %sApi.%s)", info.Abbreviation, routerStr, info.Method, info.Router, info.Abbreviation, info.FuncName)
	if info.IsPlugin {
		tPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, "router", info.HumpPackageName+".go")
		stmtStr = fmt.Sprintf("group.%s(\"%s\", api%s.%s)", info.Method, info.Router, info.StructName, info.FuncName)
		funcName = "Init"
	}

	src, err := os.ReadFile(tPath)
	if err != nil {
		return err
	}

	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, 0)
	if err != nil {
		return err
	}
	funcDecl := utilsAst.FindFunction(astFile, funcName)
	stmtNode := utilsAst.CreateStmt(stmtStr)

	if info.IsAuth {
		for i := 0; i < len(funcDecl.Body.List); i++ {
			st := funcDecl.Body.List[i]
			// 使用类型断言来检查stmt是否是一个块语句
			if blockStmt, ok := st.(*ast.BlockStmt); ok {
				// 如果是，插入代码 跳出
				blockStmt.List = append(blockStmt.List, stmtNode)
				break
			}
		}
	} else {
		for i := len(funcDecl.Body.List) - 1; i >= 0; i-- {
			st := funcDecl.Body.List[i]
			// 使用类型断言来检查stmt是否是一个块语句
			if blockStmt, ok := st.(*ast.BlockStmt); ok {
				// 如果是，插入代码 跳出
				blockStmt.List = append(blockStmt.List, stmtNode)
				break
			}
		}
	}

	// 创建一个新的文件
	f, err := os.Create(tPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := format.Node(f, fileSet, astFile); err != nil {
		return err
	}
	return err
}

func (s *autoCodeTemplate) addTemplateToFile(t string, info request.AutoFunc) error {
	if !isSafeFileName(info.HumpPackageName) || !isSafeFileName(info.PackageName) {
		return fmt.Errorf("文件名包含非法字符，拒绝写入")
	}
	getTemplateStr, err := s.getTemplateStr(t, info)
	if err != nil {
		return err
	}
	var target string

	switch t {
	case "api.go":
		if info.IsAi && info.ApiFunc != "" {
			getTemplateStr = info.ApiFunc
		}
		target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", info.Package, info.HumpPackageName+".go")
	case "server.go":
		if info.IsAi && info.ServerFunc != "" {
			getTemplateStr = info.ServerFunc
		}
		target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", info.Package, info.HumpPackageName+".go")
	case "api.js":
		if info.IsAi && info.JsFunc != "" {
			getTemplateStr = info.JsFunc
		}
		target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "api", info.Package, info.PackageName+".js")
	}
	if info.IsPlugin {
		switch t {
		case "api.go":
			target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, "api", info.HumpPackageName+".go")
		case "server.go":
			target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, "service", info.HumpPackageName+".go")
		case "api.js":
			target = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", info.Package, "api", info.PackageName+".js")
		}
	}

	// 打开文件，如果不存在则返回错误
	file, err := os.OpenFile(target, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入内容
	_, err = fmt.Fprintln(file, getTemplateStr)
	if err != nil {
		fmt.Printf("写入文件失败: %s\n", err.Error())
		return err
	}

	return nil
}
