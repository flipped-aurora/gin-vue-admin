package system

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	request "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
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
	length := len(templateDirs)
	if length >= 2 && length > 3 {
		return nil, nil, errors.New("不符合模版格式!")
	}
	for i := 0; i < length; i++ {
		second := filepath.Join(templateDir, templateDirs[i].Name())
		if !templateDirs[i].IsDir() && entity.Template != "package" {
			return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", second)
		}
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
					if secondDirs[j].Name() == "plugin.go.template" || secondDirs[j].Name() == "main.go.template" {
						ext := filepath.Ext(three)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", three)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.ServerRoot(entity.Template, info.Package, ""), templateDirs[i].Name(), strings.TrimSuffix(secondDirs[j].Name(), ext))
						code[three] = create
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
						create := filepath.Join(global.GVA_CONFIG.AutoCode.ServerRoot(entity.Template, info.Package, secondDirs[j].Name()), info.HumpPackageName+".go")
						if hasEnter != -1 {
							child := make(map[string]string, 2)
							isApi := strings.Index(secondDirs[j].Name(), "api")
							isRouter := strings.Index(secondDirs[j].Name(), "router")
							isService := strings.Index(secondDirs[j].Name(), "service")
							if entity.Template == "package" {
								if isApi != -1 {
									child["api"] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", "enter.go")
									child["enter"] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "v1", info.Package, strings.TrimSuffix(threeDirs[k].Name(), ext))
								}
								if isRouter != -1 {
									child["router"] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "enter.go")
									child["enter"] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), info.Package, strings.TrimSuffix(threeDirs[k].Name(), ext))
								}
								if isService != -1 {
									child["service"] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), "enter.go")
									child["enter"] = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), info.Package, strings.TrimSuffix(threeDirs[k].Name(), ext))
								}
								enter[four] = child
								continue
							}
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", info.Package, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
							child["enter"] = create
							enter[four] = child
							continue
						} // enter.go
						code[four] = create
					}
				case "gen", "config", "initialize":
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
						config := strings.Index(threeDirs[k].Name(), "config")
						initialize := strings.Index(threeDirs[k].Name(), "initialize")
						if gen == -1 && config == -1 && initialize == -1 {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.ServerRoot(entity.Template, info.Package, secondDirs[j].Name()), strings.TrimSuffix(threeDirs[k].Name(), ext))
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
								create := filepath.Join(global.GVA_CONFIG.AutoCode.ServerRoot(entity.Template, info.Package, secondDirs[j].Name()), threeDirs[k].Name(), info.HumpPackageName+".go")
								code[four] = create
							}
							continue
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						hasModel := strings.Index(threeDirs[k].Name(), "model")
						if hasModel == -1 {
							return nil, nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.ServerRoot(entity.Template, info.Package, secondDirs[j].Name()), info.HumpPackageName+".go")
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
						create := filepath.Join(global.GVA_CONFIG.AutoCode.WebRoot(entity.Template, info.Package, secondDirs[j].Name()), info.PackageName+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
						if (view != -1 || table != -1) && entity.Template == "package" {
							formPath := filepath.Join(three, "form.vue"+ext)
							value, ok := code[formPath]
							if ok {
								value = filepath.Join(global.GVA_CONFIG.AutoCode.WebRoot(entity.Template, info.Package, secondDirs[j].Name()), info.PackageName+"Form"+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
								code[formPath] = value
							}
						}
						code[four] = create
					}
				default:
					return nil, nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		default:
			if entity.Template != "package" {
				return nil, nil, errors.Errorf("[filpath:%s]不符合模版格式!", second)
			}
		}
	}
	return code, enter, nil
}
