package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	request "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
)

var AutoCodePackage = new(autoCodePackage)

type autoCodePackage struct{}

// Create 创建包信息
// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePackage) Create(ctx context.Context, info *request.SysAutoCodePackageCreate) error {
	if info.PackageName == "autocode" || info.PackageName == "system" || info.PackageName == "example" || info.PackageName == "" {
		return errors.New("不能使用已保留的package name")
	}
	if !errors.Is(global.GVA_DB.Where("package_name = ?", info.PackageName).First(&model.SysAutoCodePackage{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同PackageName")
	}
	// TODO 创建enter
	return global.GVA_DB.WithContext(ctx).Create(&info).Error
}

// Delete 获取所有包
// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePackage) Delete(ctx context.Context, info common.GetById) error {
	err := global.GVA_DB.WithContext(ctx).Delete(&model.SysAutoCodePackage{}, info.Uint()).Error
	if err != nil {
		return errors.Wrap(err, "删除失败!")
	}
	return nil
}

// All 获取所有包
// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePackage) All(ctx context.Context) (entities []model.SysAutoCodePackage, err error) {
	err = global.GVA_DB.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, errors.Wrap(err, "获取所有包失败!")
	}
	return entities, nil
}

// Templates 获取所有模版文件夹
// @author: [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePackage) Templates(ctx context.Context) ([]string, error) {
	templates := make([]string, 0)
	entries, err := os.ReadDir("resource")
	if err != nil {
		return nil, errors.Wrap(err, "读取模版文件夹失败!")
	}
	for i := 0; i < len(entries); i++ {
		if entries[i].IsDir() {
			templates = append(templates, entries[i].Name())
		}
	}
	return templates, nil
}
