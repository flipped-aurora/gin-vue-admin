package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type AutoCodeExampleService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAutoCodeExample
//@description: 创建自动化示例数据
//@param: autoCodeExampleService *AutoCodeExampleService
//@return: err error

func (autoCodeExampleService *AutoCodeExampleService) CreateAutoCodeExample(autoCodeExample autocode.AutoCodeExample) (err error) {
	err = global.GVA_DB.Create(&autoCodeExample).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAutoCodeExample
//@description: 删除自动化示例数据
//@param: autoCodeExample autocode.AutoCodeExample
//@return: err error

func (autoCodeExampleService *AutoCodeExampleService) DeleteAutoCodeExample(autoCodeExample autocode.AutoCodeExample) (err error) {
	err = global.GVA_DB.Delete(&autoCodeExample).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAutoCodeExample
//@description: 更新自动化示例数据
//@param: autoCodeExample *autocode.AutoCodeExample
//@return: err error

func (autoCodeExampleService *AutoCodeExampleService) UpdateAutoCodeExample(autoCodeExample *autocode.AutoCodeExample) (err error) {
	err = global.GVA_DB.Save(autoCodeExample).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAutoCodeExampleDetail
//@description: 根据id获取自动化示例数据
//@param: id uint
//@return: err error, autoCodeExample autocode.AutoCodeExample

func (autoCodeExampleService *AutoCodeExampleService) GetAutoCodeExample(id uint) (err error, autoCodeExample autocode.AutoCodeExample) {
	err = global.GVA_DB.Where("id = ?", id).First(&autoCodeExample).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAutoCodeExampleInfoList
//@description: 分页获取自动化示例数据
//@param: info request.AutoCodeExampleSearch
//@return: err error, list interface{}, total int64

func (autoCodeExampleService *AutoCodeExampleService) GetAutoCodeExampleInfoList(info request.AutoCodeExampleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.AutoCodeExample{})
	var autoCodeExamples []autocode.AutoCodeExample
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.AutoCodeExampleField != "" {
		db = db.Where("label LIKE ?", "%"+info.AutoCodeExampleField+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&autoCodeExamples).Error
	return err, autoCodeExamples, total
}
