package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var AutoCodeTemplate = new(autoCodeTemplate)

type autoCodeTemplate struct{}

// Create 创建生成自动化代码
func (s *autoCodeTemplate) Create(ctx context.Context, info request.AutoCode) error {
	history := info.History()
	if AutocodeHistory.Repeat(info.BusinessDB, info.StructName, info.Package) {
		return errors.New("已经创建过此数据结构,请勿重复创建!")
	} // 增加判断: 重复创建struct
	if info.AutoCreateApiToSql {
		apis := info.Apis()
		err := global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			for _, v := range apis {
				var api model.SysApi
				err := tx.Where("path = ? AND method = ?", v.Path, v.Method).First(&api).Error
				if err == nil {
					return errors.New("存在相同的API，请关闭自动创建API功能")
				}
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err = tx.Create(&v).Error; err != nil { // 遇到错误时回滚事务
						return err
					}
					history.ApiIDs = append(history.ApiIDs, v.ID)
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	} // 自动创建api
	if info.AutoCreateMenuToSql {
		var entity model.SysBaseMenu
		err := global.GVA_DB.WithContext(ctx).First(&entity, "name = ?", info.Abbreviation).Error
		if err == nil {
			return errors.New("存在相同的菜单路由，请关闭自动创建菜单功能")
		}
		entity = info.Menu()
		err = global.GVA_DB.WithContext(ctx).Create(&entity).Error
		if err != nil {
			return errors.Wrap(err, "创建菜单失败!")
		}
		history.MenuID = entity.ID
	} // 自动创建menu
	templates, err := s.PackageTemplate(ctx, info)
	if err != nil {
		return err
	}
	for key, create := range templates {
		var files *template.Template
		files, err = template.ParseFiles(key)
		if err != nil {
			return errors.Wrapf(err, "[filpath:%s]读取模版文件失败!", key)
		}
		create = filepath.Join(global.GVA_CONFIG.AutoCode.WebRoot(), create)
		_, err = os.Stat(create)
		if !os.IsNotExist(err) {
			continue
		}
		var file *os.File
		file, err = os.Create(create)
		if err != nil {
			return errors.Wrapf(err, "[filpath:%s]创建文件失败!", create)
		}
		err = files.Execute(file, info)
		if err != nil {
			return errors.Wrapf(err, "[filpath:%s]生成文件失败!", create)
		}
	} // 生成文件
	{
		// router_biz.go 路由注册
		if info.AutoMigrate {

		} // gorm_biz.go gorm自动迁移
		// plugin.go 插件注册
		// (api||router||service)/enter.go 注入
	} // 注入代码
	{

	} // 记录历史
	return nil
}

// Preview 预览自动化代码
func (s *autoCodeTemplate) Preview(ctx context.Context, info request.AutoCode) (map[string]string, error) {
	return nil, nil
}

func (s *autoCodeTemplate) PackageTemplate(ctx context.Context, info request.AutoCode) (code map[string]string, err error) {
	var entity model.SysAutoCodePackage
	err = global.GVA_DB.First("package_name = ?", info.PackageName).First(&entity).Error
	if err != nil {
		return nil, errors.Wrap(err, "获取模版信息失败!")
	}
	code = make(map[string]string)
	templateDir := filepath.Join("resource", entity.Template)
	templateDirs, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", templateDir)
	}
	length := len(templateDirs)
	if length >= 2 && length > 3 {
		return nil, errors.New("不符合模版格式!")
	}
	for i := 0; i < length; i++ {
		second := filepath.Join(templateDir, templateDirs[i].Name())
		if !templateDirs[i].IsDir() && entity.Template != "package" {
			return nil, errors.Errorf("[filpath:%s]非法模版文件!", second)
		}
		switch templateDirs[i].Name() {
		case "server":
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}
			for j := 0; j < len(secondDirs); j++ {
				three := filepath.Join(second, secondDirs[j].Name())
				if !secondDirs[j].IsDir() {
					return nil, errors.Errorf("[filpath:%s]非法模版文件!", three)
				}
				switch secondDirs[j].Name() {
				case "api", "router", "service":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						api := strings.Index(threeDirs[k].Name(), "api")
						enter := strings.Index(threeDirs[k].Name(), "enter")
						router := strings.Index(threeDirs[k].Name(), "router")
						service := strings.Index(threeDirs[k].Name(), "service")
						if router == -1 && api == -1 && service == -1 && enter == -1 {
							return nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), info.HumpPackageName+".go")
						if enter != -1 {
							create = filepath.Join(global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
						}
						code[four] = create
					}
				case "gen", "config", "initialize":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						gen := strings.Index(threeDirs[k].Name(), "gen")
						config := strings.Index(threeDirs[k].Name(), "config")
						initialize := strings.Index(threeDirs[k].Name(), "initialize")
						if gen == -1 && config == -1 && initialize == -1 {
							return nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
						code[four] = create
					}
				case "model":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							var fourDirs []os.DirEntry
							fourDirs, err = os.ReadDir(four)
							if err != nil {
								return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", four)
							}
							for l := 0; l < len(fourDirs); l++ {
								five := filepath.Join(four, fourDirs[l].Name())
								if fourDirs[l].IsDir() {
									return nil, errors.Errorf("[filpath:%s]非法模版文件夹!", five)
								}
								ext := filepath.Ext(five)
								if ext != ".template" && ext != ".tpl" {
									return nil, errors.Errorf("[filpath:%s]非法模版后缀!", five)
								}
								hasRequest := strings.Index(fourDirs[l].Name(), "request")
								if hasRequest == -1 {
									return nil, errors.Errorf("[filpath:%s]非法模版文件!", five)
								}
								create := filepath.Join(global.GVA_CONFIG.AutoCode.Server, threeDirs[k].Name(), info.HumpPackageName+".go")
								code[four] = create
							}
							continue
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						hasModel := strings.Index(threeDirs[k].Name(), "model")
						if hasModel == -1 {
							return nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(global.GVA_CONFIG.AutoCode.Server, secondDirs[j].Name(), strings.TrimSuffix(threeDirs[k].Name(), ext))
						code[four] = create
					}
				default:
					return nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		case "web":
			var secondDirs []os.DirEntry
			secondDirs, err = os.ReadDir(second)
			if err != nil {
				return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", second)
			}
			for j := 0; j < len(secondDirs); j++ {
				three := filepath.Join(second, secondDirs[j].Name())
				if !secondDirs[j].IsDir() {
					return nil, errors.Errorf("[filpath:%s]非法模版文件!", three)
				}
				switch secondDirs[j].Name() {
				case "api", "form", "view", "table":
					var threeDirs []os.DirEntry
					threeDirs, err = os.ReadDir(three)
					if err != nil {
						return nil, errors.Wrapf(err, "读取模版文件夹[%s]失败!", three)
					}
					for k := 0; k < len(threeDirs); k++ {
						four := filepath.Join(three, threeDirs[k].Name())
						if threeDirs[k].IsDir() {
							return nil, errors.Errorf("[filpath:%s]非法模版文件夹!", four)
						}
						ext := filepath.Ext(four)
						if ext != ".template" && ext != ".tpl" {
							return nil, errors.Errorf("[filpath:%s]非法模版后缀!", four)
						}
						api := strings.Index(threeDirs[k].Name(), "api")
						form := strings.Index(threeDirs[k].Name(), "form")
						view := strings.Index(threeDirs[k].Name(), "view")
						table := strings.Index(threeDirs[k].Name(), "table")
						if api == -1 && form == -1 && view == -1 && table == -1 {
							return nil, errors.Errorf("[filpath:%s]非法模版文件!", four)
						}
						create := filepath.Join(secondDirs[j].Name(), info.Package, info.PackageName+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
						if (view != -1 || table != -1) && entity.Template == "package" {
							formPath := filepath.Join(three, "form.vue"+ext)
							value, ok := code[formPath]
							if ok {
								value = filepath.Join(global.GVA_CONFIG.AutoCode.WebRoot(), secondDirs[j].Name(), info.Package, info.PackageName+"Form"+filepath.Ext(strings.TrimSuffix(threeDirs[k].Name(), ext)))
								code[formPath] = value
							}
						}
						code[four] = create
					}
				default:
					return nil, errors.Errorf("[filpath:%s]非法模版文件夹!", three)
				}
			}
		default:
			if entity.Template != "package" {
				return nil, errors.Errorf("[filpath:%s]不符合模版格式!", second)
			}
		}
	}
	return code, nil
}
