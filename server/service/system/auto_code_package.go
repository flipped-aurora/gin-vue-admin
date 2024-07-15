package system

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"github.com/pkg/errors"
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
	var paths []string
	switch {
	case info.Template == "":
		return errors.New("模板不能为空!")
	case info.Template == "page":
		return errors.New("page为表单生成器!")
	case info.PackageName == "":
		return errors.New("PackageName不能为空!")
	case info.Template == "package":
		if info.PackageName == "autocode" || info.PackageName == "system" || info.PackageName == "example" {
			return errors.New("不能使用已保留的package name")
		}
		paths = append(paths, filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", info.PackageName))
		paths = append(paths, filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", info.PackageName))
		paths = append(paths, filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", info.PackageName))
	default:
		paths = append(paths, filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.PackageName))
	}
	for i := 0; i < len(paths); i++ {
		_, err := os.Stat(paths[i])
		if !os.IsNotExist(err) {
			return errors.Errorf("[PackageName:%s]已存在!", info.PackageName)
		}
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
	_, enter, err := s.templates(ctx, create, code)
	if err != nil {
		return err
	}
	for key, values := range enter { // key 为 模版绝对路径
		for k2, value := range values { // k2 为需要生成enter的路径
			if k2 == "enter" {
				var files *template.Template
				files, err = template.ParseFiles(key)
				if err != nil {
					return errors.Wrapf(err, "[filepath:%s]读取模版文件失败!", key)
				}
				err = os.MkdirAll(filepath.Dir(value), os.ModePerm)
				if err != nil {
					return errors.Wrap(err, "创建文件夹失败!")
				}
				var file *os.File
				file, err = os.Create(value)
				if err != nil {
					return errors.Wrapf(err, "[filepath:%s]创建文件失败!", value)
				}
				err = files.Execute(file, code)
				if err != nil {
					return errors.Wrapf(err, "[filepath:%s]生成失败!", value)
				}
				fmt.Printf("[filepath:%s]enter生成成功!\n", value)
			}
		}
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

func (s *autoCodePackage) templates(ctx context.Context, entity model.SysAutoCodePackage, info request.AutoCode) (code map[string]string, enter map[string]map[string]string, err error) {
	code = make(map[string]string)
	enter = make(map[string]map[string]string)
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
								child := make(map[string]string, 2)
								isApi := strings.Index(secondDirs[j].Name(), "api")
								isRouter := strings.Index(secondDirs[j].Name(), "router")
								isService := strings.Index(secondDirs[j].Name(), "service")
								if isApi != -1 {
									child[ast.TypePackageApiEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", strings.TrimSuffix(threeDirs[k].Name(), ext))
									child[ast.TypePackageApiModuleEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", entity.PackageName, info.HumpPackageName+".go")
								}
								if isRouter != -1 {
									child[ast.TypePackageRouterEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									child[ast.TypePackageRouterModuleEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, info.HumpPackageName+".go")
								}
								if isService != -1 {
									child[ast.TypePackageServiceEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
									child[ast.TypePackageServiceModuleEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), entity.PackageName, info.HumpPackageName+".go")
								}
								enter[four] = child
								continue
							}
							if entity.PackageName == "preview" {
								create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "resource", entity.PackageName, templateDirs[i].Name(), secondDirs[j].Name(), info.HumpPackageName+".go")
							}
							code[four] = create
							continue
						}
						if hasEnter != -1 {
							child := make(map[string]string, 2)
							isApi := strings.Index(secondDirs[j].Name(), "api")
							isRouter := strings.Index(secondDirs[j].Name(), "router")
							isService := strings.Index(secondDirs[j].Name(), "service")
							if isApi != -1 {
								child[ast.TypePluginApiEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							if isRouter != -1 {
								child[ast.TypePluginRouterEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							if isService != -1 {
								child[ast.TypePluginServiceEnter] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", entity.PackageName, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							}
							enter[four] = child
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
	return code, enter, nil
}
