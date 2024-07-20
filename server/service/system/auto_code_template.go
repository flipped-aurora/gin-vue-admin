package system

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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

var AutoCodeTemplate = new(autoCodeTemplate)

type autoCodeTemplate struct{}

// Create 创建生成自动化代码
func (s *autoCodeTemplate) Create(ctx context.Context, info request.AutoCode) error {
	history := info.History()
	var autoPkg model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(ctx).Where("package_name = ?", info.Package).First(&autoPkg).Error
	if err != nil {
		return errors.Wrap(err, "查询包失败!")
	}

	// 增加判断: 重复创建struct
	if AutocodeHistory.Repeat(info.BusinessDB, info.StructName, info.Package) {
		return errors.New("已经创建过此数据结构,请勿重复创建!")
	}

	// 自动创建api
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
	}

	// 自动创建menu
	if info.AutoCreateMenuToSql {
		var entity model.SysBaseMenu
		err := global.GVA_DB.WithContext(ctx).First(&entity, "name = ?", info.Abbreviation).Error
		if err == nil {
			return errors.New("存在相同的菜单路由，请关闭自动创建菜单功能")
		}
		entity = info.Menu(autoPkg.Template)
		err = global.GVA_DB.WithContext(ctx).Create(&entity).Error
		if err != nil {
			return errors.Wrap(err, "创建菜单失败!")
		}
		history.MenuID = entity.ID
	}

	generate, templates, injections, err := s.generate(ctx, info, autoPkg)
	if err != nil {
		return err
	}
	for key, builder := range generate {
		err = os.MkdirAll(filepath.Dir(key), os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]创建文件夹失败!", key)
		}
		err = os.WriteFile(key, []byte(builder.String()), 0666)
		if err != nil {
			return errors.Wrapf(err, "[filepath:%s]写入文件失败!", key)
		}
	}

	// 创建历史记录
	history.Templates = templates
	history.Injections = make(map[string]string, len(injections))
	for key, value := range injections {
		bytes, _ := json.Marshal(value)
		history.Injections[key] = string(bytes)
	}
	err = AutocodeHistory.Create(ctx, history)
	if err != nil {
		return err
	}
	return nil
}

// Preview 预览自动化代码
func (s *autoCodeTemplate) Preview(ctx context.Context, info request.AutoCode) (map[string]string, error) {
	var entity model.SysAutoCodePackage
	err := global.GVA_DB.WithContext(ctx).Where("package_name = ?", info.Package).First(&entity).Error
	if err != nil {
		return nil, errors.Wrap(err, "查询包失败!")
	}
	codes := make(map[string]strings.Builder)
	preview := make(map[string]string)
	codes, _, _, err = s.generate(ctx, info, entity)
	if err != nil {
		return nil, err
	}
	for key, writer := range codes {
		if len(key) > len(global.GVA_CONFIG.AutoCode.Root) {
			key, _ = filepath.Rel(global.GVA_CONFIG.AutoCode.Root, key)
		}
		var builder strings.Builder
		builder.WriteString("```\n\n")
		builder.WriteString(writer.String())
		builder.WriteString("\n\n```")
		preview[key] = builder.String()
	}
	return preview, nil
}

func (s *autoCodeTemplate) generate(ctx context.Context, info request.AutoCode, entity model.SysAutoCodePackage) (map[string]strings.Builder, map[string]string, map[string]ast.Ast, error) {
	templates, asts, _, err := AutoCodePackage.templates(ctx, entity, info)
	if err != nil {
		return nil, nil, nil, err
	}
	code := make(map[string]strings.Builder)
	for key, create := range templates {
		var files *template.Template
		files, err = template.ParseFiles(key)
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
	injections := make(map[string]ast.Ast, len(asts))
	if info.AutoMigrate {
		for key, value := range asts {
			keys := strings.Split(key, "=>")
			if len(keys) == 2 {
				var builder strings.Builder
				parse, _ := value.Parse("", &builder)
				if parse != nil {
					_ = value.Injection(parse)
					err = value.Format("", &builder, parse)
					if err != nil {
						return nil, nil, nil, err
					}
					code[keys[0]] = builder
					injections[keys[1]] = value
					fmt.Println(keys[0], "注入成功!")
				}
			}
		}
	} // 注入代码
	return code, templates, injections, nil
}
