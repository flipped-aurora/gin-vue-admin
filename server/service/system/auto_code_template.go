package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
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
	var entity model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(ctx).Where("package_name = ?", info.Package).First(&entity).Error
	if err != nil {
		return errors.Wrap(err, "查询包失败!")
	}
	templates, _, err := AutoCodePackage.TemplateMap(ctx, entity, info)
	if err != nil {
		return err
	}
	for key, create := range templates {
		var files *template.Template
		files, err = template.ParseFiles(key)
		if err != nil {
			return errors.Wrapf(err, "[filpath:%s]读取模版文件失败!", key)
		}
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
	err = global.GVA_DB.WithContext(ctx).Create(&history).Error
	if err != nil {
		return errors.Wrap(err, "创建历史记录失败!")
	}
	return nil
}

// Preview 预览自动化代码
func (s *autoCodeTemplate) Preview(ctx context.Context, info request.AutoCode) (map[string]string, error) {
	return nil, nil
}
