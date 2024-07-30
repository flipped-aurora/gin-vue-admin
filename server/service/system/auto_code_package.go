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
	case token.IsKeyword(info.PackageName):
		return errors.Errorf("%s为go的关键字!", info.PackageName)
	case info.Template == "package":
		if info.PackageName == "system" || info.PackageName == "example" {
			return errors.New("不能使用已保留的package name")
		}
	default:
		break
	}
	if !errors.Is(global.GVA_DB.Where("package_name = ? and template = ?", info.PackageName, info.Template).First(&model.SysAutoCodePackage{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同PackageName")
	}
	create := info.Create()
	return global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&create).Error
		if err != nil {
			return errors.Wrap(err, "创建失败!")
		}
		code := info.AutoCode()
		_, asts, creates, err := s.templates(ctx, create, code)
		if err != nil {
			return err
		}
		for key, value := range creates { // key 为 模版绝对路径
			var files *template.Template
			files, err = template.ParseFiles(key)
			if err != nil {
				return errors.Wrapf(err, "[filepath:%s]读取模版文件失败!", key)
			}
			err = os.MkdirAll(filepath.Dir(value), os.ModePerm)
			if err != nil {
				return errors.Wrapf(err, "[filepath:%s]创建文件夹失败!", value)
			}
			var file *os.File
			file, err = os.Create(value)
			if err != nil {
				return errors.Wrapf(err, "[filepath:%s]创建文件夹失败!", value)
			}
			err = files.Execute(file, code)
			_ = file.Close()
			if err != nil {
				return errors.Wrapf(err, "[filepath:%s]生成失败!", value)
			}
			fmt.Printf("[template:%s][filepath:%s]生成成功!\n", key, value)
		}
		for key, value := range asts {
			keys := strings.Split(key, "=>")
			if len(keys) == 2 {
				switch keys[1] {
				case ast.TypePluginInitializeV2, ast.TypePackageApiEnter, ast.TypePackageRouterEnter, ast.TypePackageServiceEnter:
					file, _ := value.Parse("", nil)
					if file != nil {
						err = value.Injection(file)
						if err != nil {
							return err
						}
						err = value.Format("", nil, file)
						if err != nil {
							return err
						}
					}
					fmt.Printf("[type:%s]注入成功!\n", key)
				}
			}
		}
		return nil
	})
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
			if entries[i].Name() == "function" {
				continue
			} // function 为函数生成器
			if entries[i].Name() == "preview" {
				continue
			} // preview 为预览代码生成器的代码
			templates = append(templates, entries[i].Name())
		}
	}
	return templates, nil
}

func (s *autoCodePackage) templates(ctx context.Context, entity model.SysAutoCodePackage, info request.AutoCode) (code map[string]string, asts map[string]ast.Ast, creates map[string]string, err error) {
	code = make(map[string]string)
	asts = make(map[string]ast.Ast)
	creates = make(map[string]string)
	templateDir := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.Template)
	templateDirs, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", templateDir)
	}
	for i := 0; i < len(templateDirs); i++ {
		second := filepath.Join(templateDir, templateDirs[i].Name())
		switch templateDirs[i].Name() {
		case "server":
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}
			for j := 0; j < len(secondDirs); j++ {
				if secondDirs[j].Name() == ".DS_Store" {
					continue
				}
				three := filepath.Join(second, secondDirs[j].Name())
				if !secondDirs[j].IsDir() {
					ext := filepath.Ext(secondDirs[j].Name())
					if ext != ".template" && ext != ".tpl" {
						return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", three)
					}
					name := strings.TrimSuffix(secondDirs[j].Name(), ext)
					if name == "main.go" || name == "plugin.go" {
						pluginInitialize := &ast.PluginInitializeV2{
							Type:        ast.TypePluginInitializeV2,
							Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, name),
							PluginPath:  filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "plugin_biz_v2.go"),
							ImportPath:  fmt.Sprintf(`"%s/plugin/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
							PackageName: entity.PackageName,
						}
						asts[pluginInitialize.PluginPath+"=>"+pluginInitialize.Type.String()] = pluginInitialize
						creates[three] = pluginInitialize.Path
						continue
					}
					return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", three)
				}
				switch secondDirs[j].Name() {
				case "api", "router", "service":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						if threeDirs[k].Name() == ".DS_Store" {
							continue
						}
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						api := strings.Index(threeDirs[k].Name(), "api")
						hasEnter := strings.Index(threeDirs[k].Name(), "enter")
						router := strings.Index(threeDirs[k].Name(), "router")
						service := strings.Index(threeDirs[k].Name(), "service")
						if router == -1 && api == -1 && service == -1 && hasEnter == -1 {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
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
									packageApiEnter := &ast.PackageEnter{
										Type:              ast.TypePackageApiEnter,
										Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", "enter.go"),
										ImportPath:        fmt.Sprintf(`"%s/%s/%s/%s"`, global.GVA_CONFIG.AutoCode.Module, "api", "v1", entity.PackageName),
										StructName:        utils.FirstUpper(entity.PackageName) + "ApiGroup",
										PackageName:       entity.PackageName,
										PackageStructName: "ApiGroup",
									}
									asts[packageApiEnter.Path+"=>"+packageApiEnter.Type.String()] = packageApiEnter
									packageApiModuleEnter := &ast.PackageModuleEnter{
										Type:        ast.TypePackageApiModuleEnter,
										Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", entity.PackageName, "enter.go"),
										ImportPath:  fmt.Sprintf(`"%s/service"`, global.GVA_CONFIG.AutoCode.Module),
										StructName:  info.StructName + "Api",
										AppName:     "ServiceGroupApp",
										GroupName:   utils.FirstUpper(entity.PackageName) + "ServiceGroup",
										ModuleName:  info.Abbreviation + "Service",
										PackageName: "service",
										ServiceName: info.StructName + "Service",
									}
									asts[packageApiModuleEnter.Path+"=>"+packageApiModuleEnter.Type.String()] = packageApiModuleEnter
									creates[four] = packageApiModuleEnter.Path
								}
								if isRouter != -1 {
									packageRouterEnter := &ast.PackageEnter{
										Type:              ast.TypePackageRouterEnter,
										Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "enter.go"),
										ImportPath:        fmt.Sprintf(`"%s/%s/%s"`, global.GVA_CONFIG.AutoCode.Module, secondDirs[j].Name(), entity.PackageName),
										StructName:        utils.FirstUpper(entity.PackageName),
										PackageName:       entity.PackageName,
										PackageStructName: "RouterGroup",
									}
									asts[packageRouterEnter.Path+"=>"+packageRouterEnter.Type.String()] = packageRouterEnter
									packageRouterModuleEnter := &ast.PackageModuleEnter{
										Type:        ast.TypePackageRouterModuleEnter,
										Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, "enter.go"),
										ImportPath:  fmt.Sprintf(`api "%s/api/v1"`, global.GVA_CONFIG.AutoCode.Module),
										StructName:  info.StructName + "Router",
										AppName:     "ApiGroupApp",
										GroupName:   utils.FirstUpper(entity.PackageName) + "ApiGroup",
										ModuleName:  info.Abbreviation + "Api",
										PackageName: "api",
										ServiceName: info.StructName + "Api",
									}
									creates[four] = packageRouterModuleEnter.Path
									asts[packageRouterModuleEnter.Path+"=>"+packageRouterModuleEnter.Type.String()] = packageRouterModuleEnter
									packageInitializeRouter := &ast.PackageInitializeRouter{
										Type:                 ast.TypePackageInitializeRouter,
										Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router_biz.go"),
										ImportPath:           fmt.Sprintf(`"%s/router"`, global.GVA_CONFIG.AutoCode.Module),
										AppName:              "RouterGroupApp",
										GroupName:            utils.FirstUpper(entity.PackageName),
										ModuleName:           entity.PackageName + "Router",
										PackageName:          "router",
										FunctionName:         "Init" + info.StructName + "Router",
										LeftRouterGroupName:  "privateGroup",
										RightRouterGroupName: "publicGroup",
									}
									asts[packageInitializeRouter.Path+"=>"+packageInitializeRouter.Type.String()] = packageInitializeRouter
								}
								if isService != -1 {
									path := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									importPath := fmt.Sprintf(`"%s/service/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName)
									packageServiceEnter := &ast.PackageEnter{
										Type:              ast.TypePackageServiceEnter,
										Path:              path,
										ImportPath:        importPath,
										StructName:        utils.FirstUpper(entity.PackageName) + "ServiceGroup",
										PackageName:       entity.PackageName,
										PackageStructName: "ServiceGroup",
									}
									asts[packageServiceEnter.Path+"=>"+packageServiceEnter.Type.String()] = packageServiceEnter
									packageServiceModuleEnter := &ast.PackageModuleEnter{
										Type:       ast.TypePackageServiceModuleEnter,
										Path:       filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, "enter.go"),
										StructName: info.StructName + "Service",
									}
									asts[packageServiceModuleEnter.Path+"=>"+packageServiceModuleEnter.Type.String()] = packageServiceModuleEnter
									creates[four] = packageServiceModuleEnter.Path
								}
								continue
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
									Type:            ast.TypePluginRouterEnter,
									Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									ImportPath:      fmt.Sprintf(`"%s/plugin/%s/api"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
									ModuleName:      "api" + info.StructName,
									GroupName:       "Api",
									PackageName:     "api",
									ServiceName:     info.StructName,
								}
								asts[pluginRouterEnter.Path+"=>"+pluginRouterEnter.Type.String()] = pluginRouterEnter
								creates[four] = pluginRouterEnter.Path
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
								}
								asts[pluginApiEnter.Path+"=>"+pluginApiEnter.Type.String()] = pluginApiEnter
								creates[four] = pluginApiEnter.Path
							}
							if isService != -1 {
								pluginServiceEnter := &ast.PluginEnter{
									Type:            ast.TypePluginServiceEnter,
									Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
									StructName:      info.StructName,
									StructCamelName: info.Abbreviation,
								}
								asts[pluginServiceEnter.Path+"=>"+pluginServiceEnter.Type.String()] = pluginServiceEnter
								creates[four] = pluginServiceEnter.Path
							}
							continue
						} // enter.go
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), info.HumpPackageName+".go")
						code[four] = create
					}
				case "gen", "config", "initialize", "plugin", "response":
					if entity.Template == "package" {
						continue
					} // package模板不需要生成gen, config, initialize
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						if threeDirs[k].Name() == ".DS_Store" {
							continue
						}
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						gen := strings.Index(threeDirs[k].Name(), "gen")
						api := strings.Index(threeDirs[k].Name(), "api")
						menu := strings.Index(threeDirs[k].Name(), "menu")
						viper := strings.Index(threeDirs[k].Name(), "viper")
						plugin := strings.Index(threeDirs[k].Name(), "plugin")
						config := strings.Index(threeDirs[k].Name(), "config")
						router := strings.Index(threeDirs[k].Name(), "router")
						hasGorm := strings.Index(threeDirs[k].Name(), "gorm")
						response := strings.Index(threeDirs[k].Name(), "response")
						if gen != -1 && api != -1 && menu != -1 && viper != -1 && plugin != -1 && config != -1 && router != -1 && hasGorm != -1 && response != -1 {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if api != -1 || menu != -1 || viper != -1 || response != -1 || plugin != -1 || config != -1 {
							creates[four] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
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
							asts[pluginGen.Path+"=>"+pluginGen.Type.String()] = pluginGen
							creates[four] = pluginGen.Path
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
							asts[pluginInitializeGorm.Path+"=>"+pluginInitializeGorm.Type.String()] = pluginInitializeGorm
							creates[four] = pluginInitializeGorm.Path
						}
						if router != -1 {
							pluginInitializeRouter := &ast.PluginInitializeRouter{
								Type:                 ast.TypePluginInitializeRouter,
								Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext)),
								ImportPath:           fmt.Sprintf(`"%s/plugin/%s/router"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
								AppName:              "Router",
								GroupName:            info.StructName,
								PackageName:          "router",
								FunctionName:         "Init",
								LeftRouterGroupName:  "public",
								RightRouterGroupName: "private",
							}
							asts[pluginInitializeRouter.Path+"=>"+pluginInitializeRouter.Type.String()] = pluginInitializeRouter
							creates[four] = pluginInitializeRouter.Path
						}
					}
				case "model":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						if threeDirs[k].Name() == ".DS_Store" {
							continue
						}
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							var fourDirs []os.DirEntry
							fourDirs, err = os.ReadDir(four)
							if err != nil {
								return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", four)
							}
							for l := 0; l < len(fourDirs); l++ {
								if fourDirs[l].Name() == ".DS_Store" {
									continue
								}
								five := filepath.Join(four, fourDirs[l].Name())
								if fourDirs[l].IsDir() {
									return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", five)
								}
								ext := filepath.Ext(five)
								if ext != ".template" && ext != ".tpl" {
									return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", five)
								}
								hasRequest := strings.Index(fourDirs[l].Name(), "request")
								if hasRequest == -1 {
									return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", five)
								}
								create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), threeDirs[k].Name(), info.HumpPackageName+".go")
								if entity.Template == "package" {
									create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, threeDirs[k].Name(), info.HumpPackageName+".go")
								}
								code[five] = create
							}
							continue
						}
						ext := filepath.Ext(threeDirs[k].Name())
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						hasModel := strings.Index(threeDirs[k].Name(), "model")
						if hasModel == -1 {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), info.HumpPackageName+".go")
						if entity.Template == "package" {
							packageInitializeGorm := &ast.PackageInitializeGorm{
								Type:        ast.TypePackageInitializeGorm,
								Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "gorm_biz.go"),
								ImportPath:  fmt.Sprintf(`"%s/model/%s"`, global.GVA_CONFIG.AutoCode.Module, entity.PackageName),
								Business:    info.BusinessDB,
								StructName:  info.StructName,
								PackageName: entity.PackageName,
								IsNew:       true,
							}
							code[four] = packageInitializeGorm.Path
							asts[packageInitializeGorm.Path+"=>"+packageInitializeGorm.Type.String()] = packageInitializeGorm
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, info.HumpPackageName+".go")
						}
						code[four] = create
					}
				default:
					return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		case "web":
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}
			for j := 0; j < len(secondDirs); j++ {
				if secondDirs[j].Name() == ".DS_Store" {
					continue
				}
				three := filepath.Join(second, secondDirs[j].Name())
				if !secondDirs[j].IsDir() {
					return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", three)
				}
				switch secondDirs[j].Name() {
				case "api", "form", "view", "table":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, nil, nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						if threeDirs[k].Name() == ".DS_Store" {
							continue
						}
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						api := strings.Index(threeDirs[k].Name(), "api")
						form := strings.Index(threeDirs[k].Name(), "form")
						view := strings.Index(threeDirs[k].Name(), "view")
						table := strings.Index(threeDirs[k].Name(), "table")
						if api == -1 && form == -1 && view == -1 && table == -1 {
							return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						if entity.Template == "package" {
							if view != -1 || table != -1 {
								formPath := filepath.Join(three, "form.vue"+ext)
								value, ok := code[formPath]
								if ok {
									value = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), secondDirs[j].Name(), entity.PackageName, info.PackageName, info.PackageName+"Form"+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
									code[formPath] = value
								}
							}
							create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), secondDirs[j].Name(), entity.PackageName, info.PackageName, info.PackageName+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
							if api != -1 {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), secondDirs[j].Name(), entity.PackageName, info.PackageName+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
							}
							code[four] = create
							continue
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.WebRoot(), "plugin", entity.PackageName, secondDirs[j].Name(), info.Abbreviation+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
						code[four] = create
					}
				default:
					return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		case "readme.txt.tpl", "readme.txt.template":
			continue
		default:
			if templateDirs[i].Name() == ".DS_Store" {
				continue
			}
			return nil, nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", second)
		}
	}
	return code, asts, creates, nil
}
