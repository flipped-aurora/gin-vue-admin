package system

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/pkg/errors"
	"go/token"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var AutoCodePackage = new(autoCodePackage)

type autoCodePackage struct{}

// Create 创建包信息
// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodePackage) Create(ctx context.Context, info *request.SysAutoCodePackageCreate) error {
	switch {
	case info.Template == "":
		return errors.New("模板不能为空!")
	case info.Template == "page":
		return errors.New("page为表单生成器!")
	case info.PackageName == "":
		return errors.New("PackageName不能为空!")
	case info.PackageName == "preview":
		return errors.New("preview为预览代码生成器保留关键字!")
	case token.IsKeyword(info.PackageName):
		return errors.Errorf("%s为go的关键字!", info.PackageName)
	case info.Template == "package":
		if info.PackageName == "autocode" || info.PackageName == "system" || info.PackageName == "example" {
			return errors.New("不能使用已保留的package name")
		}
	default:
		break
	}
	if !errors.Is(global.GVA_DB.Where("package_name = ?", info.PackageName).First(&model.SysAutoCodePackage{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同PackageName")
	}
	create := info.Create()
	err := global.GVA_DB.WithContext(ctx).Create(&create).Error
	if err != nil {
		return errors.Wrap(err, "创建失败!")
	}
	code := info.AutoCode()
	_, asts, err := s.templates(ctx, create, code)
	if err != nil {
		return err
	}
	for key, value := range asts { // key 为 模版绝对路径
		var (
			enterPath    string
			templatePath string
		)
		v1, o2 := value.(*ast.PackageModuleEnter)
		if o2 {
			switch v1.Type {
			case ast.TypePackageApiModuleEnter, ast.TypePackageRouterModuleEnter, ast.TypePackageServiceModuleEnter:
				enterPath = v1.Path
				templatePath = v1.TemplatePath
			default:
				continue
			}
		}
		v2, o2 := value.(*ast.PluginEnter)
		if o2 {
			switch v1.Type {
			case ast.TypePluginApiEnter, ast.TypePluginRouterEnter, ast.TypePluginServiceEnter:
				enterPath = v2.Path
				templatePath = v2.TemplatePath
			default:
				continue
			}
		}
		if enterPath == "" && templatePath == "" {
			continue
		}
		var files *template.Template
		files, err = template.ParseFiles(templatePath)
		if err != nil {
			return errors.Wrapf(err, "[type:%s][filepath:%s]读取模版文件失败!", key, templatePath)
		}
		err = os.MkdirAll(filepath.Dir(enterPath), os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "[type:%s][filepath:%s]创建文件夹失败!", key, enterPath)
		}
		var file *os.File
		file, err = os.Create(enterPath)
		if err != nil {
			return errors.Wrapf(err, "[type:%s][filepath:%s]创建文件夹失败!", key, enterPath)
		}
		err = files.Execute(file, code)
		if err != nil {
			return errors.Wrapf(err, "[type:%s][filepath:%s]生成失败!", key, enterPath)
		}
		fmt.Printf("[type:%s][filepath:%s]生成成功!\n", key, enterPath)
	}
	return nil
}

// Delete 删除包记录
// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodePackage) Delete(ctx context.Context, info common.GetById) error {
	err := global.GVA_DB.WithContext(ctx).Delete(&model.SysAutoCodePackage{}, info.Uint()).Error
	if err != nil {
		return errors.Wrap(err, "删除失败!")
	}
	return nil
}

// All 获取所有包
// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodePackage) All(ctx context.Context) (entities []model.SysAutoCodePackage, err error) {
	err = global.GVA_DB.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, errors.Wrap(err, "获取所有包失败!")
	}
	return entities, nil
}

// Templates 获取所有模版文件夹
// @author: [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodePackage) Templates(ctx context.Context) ([]string, error) {
	templates := make([]string, 0)
	entries, err := os.ReadDir("resource")
	if err != nil {
		return nil, errors.Wrap(err, "读取模版文件夹失败!")
	}
	for i := 0; i < len(entries); i++ {
		if entries[i].IsDir() {
			if entries[i].Name() == "page" {
				continue
			} // page 为表单生成器
			if entries[i].Name() == "preview" {
				continue
			} // preview 为预览代码生成器的代码
			templates = append(templates, entries[i].Name())
		}
	}
	return templates, nil
}

func (s *autoCodePackage) templates(ctx context.Context, entity model.SysAutoCodePackage, info request.AutoCode) (code map[string]string, asts map[string]ast.Ast, err error) {
	code = make(map[string]string)
	asts = make(map[string]ast.Ast)
	templateDir := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.Template)
	templateDirs, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", templateDir)
	}
	for i := 0; i < len(templateDirs); i++ {
		second := filepath.Join(templateDir, templateDirs[i].Name())
		switch templateDirs[i].Name() {
		case "server":
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}
			for j := 0; j < len(secondDirs); j++ {
				three := filepath.Join(second, secondDirs[j].Name())
				if !secondDirs[j].IsDir() {
					ext := filepath.Ext(secondDirs[j].Name())
					if ext != ".template" && ext != ".tpl" {
						return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", three)
					}
					name := strings.TrimSuffix(secondDirs[j].Name(), ext)
					if name == "main.go" || name == "plugin.go" {
						pluginInitialize := &ast.PluginInitialize{
							Type:       ast.TypePluginInitialize,
							Path:       filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize"),
							PluginPath: filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, name),
							ImportPath: fmt.Sprintf(`"%s/plugin/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
						}
						asts[ast.TypePluginInitialize] = pluginInitialize
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, name)
						if entity.PackageName == "preview" {
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), name)
						}
						code[second] = create
						continue
					}
					return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", three)
				}
				switch secondDirs[j].Name() {
				case "api", "router", "service":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						api := strings.Index(threeDirs[k].Name(), "api")
						hasEnter := strings.Index(threeDirs[k].Name(), "enter")
						router := strings.Index(threeDirs[k].Name(), "router")
						service := strings.Index(threeDirs[k].Name(), "service")
						if router == -1 && api == -1 && service == -1 && hasEnter == -1 {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if entity.Template == "package" {
							create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, info.HumpPackageName+".go")
							if api != -1 {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", entity.PackageName, info.HumpPackageName+".go")
							}
							if hasEnter != -1 {
								isApi := strings.Index(secondDirs[j].Name(), "api")
								isRouter := strings.Index(secondDirs[j].Name(), "router")
								isService := strings.Index(secondDirs[j].Name(), "service")
								if isApi != -1 {
									path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", "enter.go")
									importPath := fmt.Sprintf(`"%s/%s/%s/%s"`, global.GVA_CONFIG.AutoCode.Module, "api", "v1", entity.PackageName)
									packageEnter := &ast.PackageEnter{
										Type:              ast.TypePackageApiEnter,
										Path:              path,
										ImportPath:        importPath,
										StructName:        utils.FirstUpper(entity.PackageName) + "ApiGroup",
										PackageName:       entity.PackageName,
										PackageStructName: "ApiGroup",
									}
									if entity.PackageName == "preview" {
										packageEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), "v1", strings.TrimSuffix(threeDirs[k].Name(), ext))
									}
									asts[ast.TypePackageApiEnter] = packageEnter
									path = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", entity.PackageName, "enter.go")
									importPath = fmt.Sprintf(`"%s/service"`, global.GVA_CONFIG.AutoCode.Module)
									packageApiModuleEnter := &ast.PackageModuleEnter{
										Type:         ast.TypePackageApiModuleEnter,
										Path:         path,
										ImportPath:   importPath,
										StructName:   info.StructName + "Api",
										AppName:      "ServiceGroup",
										GroupName:    utils.FirstUpper(entity.PackageName) + "ServiceGroup",
										ModuleName:   info.Abbreviation,
										PackageName:  "service",
										ServiceName:  utils.FirstUpper(entity.PackageName) + "Service",
										TemplatePath: four,
									}
									if entity.PackageName == "preview" {
										packageApiModuleEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									}
									asts[ast.TypePackageApiModuleEnter] = packageApiModuleEnter
								}
								if isRouter != -1 {
									path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "enter.go")
									importPath := fmt.Sprintf(`"%s/%s/%s"`, global.GVA_CONFIG.AutoCode.Module, secondDirs[j].Name(), entity.PackageName)
									packageEnter := &ast.PackageEnter{
										Type:              ast.TypePackageRouterEnter,
										Path:              path,
										ImportPath:        importPath,
										StructName:        utils.FirstUpper(entity.PackageName),
										PackageName:       entity.PackageName,
										PackageStructName: "RouterGroup",
									}
									if entity.PackageName == "preview" {
										packageEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									}
									asts[ast.TypePackageRouterEnter] = packageEnter
									path = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, "enter.go")
									importPath = fmt.Sprintf(`api "%s/api/v1"`, global.GVA_CONFIG.AutoCode.Module)
									packageRouterModuleEnter := &ast.PackageModuleEnter{
										Type:         ast.TypePackageRouterModuleEnter,
										Path:         path,
										ImportPath:   importPath,
										StructName:   info.StructName + "Router",
										AppName:      "ApiGroupApp",
										GroupName:    utils.FirstUpper(entity.PackageName) + "ApiGroup",
										ModuleName:   info.Abbreviation + "Api",
										PackageName:  "api",
										PreviewPath:  "api",
										ServiceName:  info.StructName + "Api",
										TemplatePath: four,
									}
									if entity.PackageName == "preview" {
										packageRouterModuleEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									}
									asts[ast.TypePackageRouterModuleEnter] = packageRouterModuleEnter
									packageInitializeRouter := &ast.PackageInitializeRouter{
										Type:         ast.TypePackageInitializeRouter,
										Path:         filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router_biz.go"),
										ImportPath:   fmt.Sprintf(`"%s/router/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
										AppName:      "RouterGroupApp",
										GroupName:    utils.FirstUpper(entity.PackageName),
										PackageName:  entity.PackageName,
										FunctionName: "Init" + info.StructName + "Router",
									}
									if entity.PackageName == "preview" {
										packageInitializeRouter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, "initialize", "router_biz.go")
									}
									asts[ast.TypePackageInitializeRouter] = packageInitializeRouter
								}
								if isService != -1 {
									path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									importPath := fmt.Sprintf(`"%s/service/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName)
									packageEnter := &ast.PackageEnter{
										Type:              ast.TypePackageServiceEnter,
										Path:              path,
										ImportPath:        importPath,
										StructName:        info.StructName + "ServiceGroup",
										PackageName:       entity.PackageName,
										PackageStructName: "ServiceGroup",
									}
									asts[ast.TypePackageServiceEnter] = packageEnter
									path = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, "enter.go")
									packageServiceModuleEnter := &ast.PackageModuleEnter{
										Type:         ast.TypePackageServiceModuleEnter,
										Path:         path,
										StructName:   info.StructName + "Service",
										TemplatePath: four,
									}
									if entity.PackageName == "preview" {
										packageServiceModuleEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), "enter.go")
									}
									asts[ast.TypePackageServiceModuleEnter] = packageServiceModuleEnter
								}
								continue
							}
							if entity.PackageName == "preview" {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.HumpPackageName+".go")
							}
							code[four] = create
							continue
						}
						if hasEnter != -1 {
							isApi := strings.Index(secondDirs[j].Name(), "api")
							isRouter := strings.Index(secondDirs[j].Name(), "router")
							isService := strings.Index(secondDirs[j].Name(), "service")
							if isRouter != -1 {
								pluginRouterEnter := &ast.PluginEnter{
									Type:            ast.TypePluginApiEnter,
									Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									ImportPath:      fmt.Sprintf(`"%s/plugin/%s/api"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
									ModuleName:      "api" + info.StructName,
									GroupName:       "Api",
									PackageName:     "api",
									ServiceName:     info.StructName,
									TemplatePath:    four,
								}
								if entity.PackageName == "preview" {
									pluginRouterEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
								}
								asts[ast.TypePluginRouterEnter] = pluginRouterEnter
							}
							if isApi != -1 {
								pluginApiEnter := &ast.PluginEnter{
									Type:            ast.TypePluginApiEnter,
									Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									ImportPath:      fmt.Sprintf(`"%s/plugin/%s/service"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
									ModuleName:      "service" + info.StructName,
									GroupName:       "Service",
									PackageName:     "service",
									ServiceName:     info.StructName,
									TemplatePath:    four,
								}
								if entity.PackageName == "preview" {
									pluginApiEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
								}
								asts[ast.TypePluginApiEnter] = pluginApiEnter
							}
							if isService != -1 {
								pluginServiceEnter := &ast.PluginEnter{
									Type:            ast.TypePluginServiceEnter,
									Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
									TemplatePath:    four,
								}
								if entity.PackageName == "preview" {
									pluginServiceEnter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
								}
								asts[ast.TypePluginServiceEnter] = pluginServiceEnter
							}
							create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), info.HumpPackageName+".go")
							if entity.PackageName == "preview" {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.HumpPackageName+".go")
							}
							code[four] = create
							continue
						} // enter.go
					}
				case "gen", "config", "initialize", "plugin":
					if entity.Template == "package" {
						continue
					} // package模板不需要生成gen, config, initialize
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						gen := strings.Index(threeDirs[k].Name(), "gen")
						api := strings.Index(threeDirs[k].Name(), "api")
						menu := strings.Index(threeDirs[k].Name(), "menu")
						viper := strings.Index(threeDirs[k].Name(), "viper")
						plugin := strings.Index(threeDirs[k].Name(), "plugin")
						config := strings.Index(threeDirs[k].Name(), "config")
						router := strings.Index(threeDirs[k].Name(), "router")
						hasGorm := strings.Index(threeDirs[k].Name(), "gorm")
						if gen != -1 && api != -1 && menu != -1 && viper != -1 && plugin != -1 && config != -1 && router != -1 && hasGorm != -1 {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if gen != -1 {
							pluginGen := &ast.PluginGen{
								Type:        ast.TypePluginGen,
								Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:  fmt.Sprintf(`"%s/plugin/%s/model"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
								StructName:  info.StructName,
								PackageName: "model",
								IsNew:       true,
							}
							if entity.PackageName == "preview" {
								pluginGen.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							asts[ast.TypePluginGen] = pluginGen
						}
						if hasGorm != -1 {
							pluginInitializeGorm := &ast.PluginInitializeGorm{
								Type:        ast.TypePluginInitializeGorm,
								Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:  fmt.Sprintf(`"%s/plugin/%s/model"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
								StructName:  info.StructName,
								PackageName: "model",
								IsNew:       true,
							}
							if entity.PackageName == "preview" {
								pluginInitializeGorm.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							asts[ast.TypePluginInitializeGorm] = pluginInitializeGorm
						}
						if router != -1 {
							pluginInitializeRouter := &ast.PluginInitializeRouter{
								Type:                 ast.TypePluginInitializeGorm,
								Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:           fmt.Sprintf(`"%s/plugin/%s/router"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
								AppName:              "Router",
								GroupName:            info.StructName,
								PackageName:          "router",
								FunctionName:         "Init",
								LeftRouterGroupName:  "public",
								RightRouterGroupName: "private",
							}
							if entity.PackageName == "preview" {
								pluginInitializeRouter.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							asts[ast.TypePluginInitializeGorm] = pluginInitializeRouter
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
						if entity.PackageName == "preview" {
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
						}
						code[four] = create
					}
				case "model":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							var fourDirs []os.DirEntry
							fourDirs, err = os.ReadDir(four)
							if err != nil {
								return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", four)
							}
							for l := 0; l < len(fourDirs); l++ {
								five := filepath.Join(four, fourDirs[l].Name())
								if fourDirs[l].IsDir() {
									return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", five)
								}
								ext := filepath.Ext(five)
								if ext != ".template" && ext != ".tpl" {
									return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", five)
								}
								hasRequest := strings.Index(fourDirs[l].Name(), "request")
								if hasRequest == -1 {
									return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", five)
								}
								if entity.Template == "package" {
									create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, threeDirs[k].Name(), info.HumpPackageName+".go")
									if entity.PackageName == "preview" {
										create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), threeDirs[k].Name(), info.HumpPackageName+".go")
									}
									code[five] = create
									continue
								}
								create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), info.HumpPackageName+".go")
								if entity.PackageName == "preview" {
									create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.HumpPackageName+".go")
								}
								code[four] = create
							}
							continue
						}
						ext := filepath.Ext(threeDirs[k].Name())
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						hasModel := strings.Index(threeDirs[k].Name(), "model")
						if hasModel == -1 {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if entity.Template == "package" {
							packageInitializeGorm := &ast.PackageInitializeGorm{
								Type:        ast.TypePackageInitializeGorm,
								Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "gorm_biz.go"),
								ImportPath:  fmt.Sprintf(`"%s/model/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
								StructName:  info.StructName,
								PackageName: entity.PackageName,
								IsNew:       true,
							}
							if entity.PackageName == "preview" {
								packageInitializeGorm.PreviewPath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, info.HumpPackageName+".go")
							if entity.PackageName == "preview" {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.HumpPackageName+".go")
							}
							code[four] = create
							continue
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), info.HumpPackageName+".go")
						if entity.PackageName == "preview" {
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.HumpPackageName+".go")
						}
						code[four] = create
					}
				default:
					return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		case "web":
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}
			for j := 0; j < len(secondDirs); j++ {
				three := filepath.Join(second, secondDirs[j].Name())
				if !secondDirs[j].IsDir() {
					return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", three)
				}
				switch secondDirs[j].Name() {
				case "api", "form", "view", "table":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						api := strings.Index(threeDirs[k].Name(), "api")
						form := strings.Index(threeDirs[k].Name(), "form")
						view := strings.Index(threeDirs[k].Name(), "view")
						table := strings.Index(threeDirs[k].Name(), "table")
						if api == -1 && form == -1 && view == -1 && table == -1 {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if entity.Template == "package" {
							if view != -1 || table != -1 {
								formPath := filepath.Join(three, "form.vue"+ext)
								value, ok := code[formPath]
								if ok {
									value = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), secondDirs[j].Name(), info.PackageName, info.Abbreviation+"Form"+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
									if entity.PackageName == "preview" {
										value = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.Abbreviation+"Form"+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
									}
									code[formPath] = value
								}
							}
							create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), secondDirs[j].Name(), info.PackageName, info.Abbreviation+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
							if entity.PackageName == "preview" {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.Abbreviation+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
							}
							code[four] = create
							continue
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), "plugin", secondDirs[j].Name(), info.Abbreviation+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
						if entity.PackageName == "preview" {
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.Abbreviation+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
						}
						code[four] = create
					}
				default:
					return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		case "readme.txt.tpl":
			continue
		default:
			return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", second)
		}
	}
	return code, asts, nil
}
