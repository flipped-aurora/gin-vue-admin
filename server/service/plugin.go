package service

import (
	"github.com/eyotang/gin-vue-admin/server/global"
	"github.com/eyotang/gin-vue-admin/server/model"
	"github.com/eyotang/gin-vue-admin/server/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateProductPlugin
//@description: 创建ProductPlugin记录
//@param: productPlugin model.ProductPlugin
//@return: err error

func CreateProductPlugin(productPlugin model.ProductPlugin) (err error) {
	err = global.GVA_DB.Create(&productPlugin).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProductPlugin
//@description: 删除ProductPlugin记录
//@param: productPlugin model.ProductPlugin
//@return: err error

func DeleteProductPlugin(productPlugin model.ProductPlugin) (err error) {
	err = global.GVA_DB.Delete(&productPlugin).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteProductPluginByIds
//@description: 批量删除ProductPlugin记录
//@param: ids request.IdsReq
//@return: err error

func DeleteProductPluginByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ProductPlugin{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateProductPlugin
//@description: 更新ProductPlugin记录
//@param: productPlugin *model.ProductPlugin
//@return: err error

func UpdateProductPlugin(productPlugin model.ProductPlugin) (err error) {
	err = global.GVA_DB.Save(&productPlugin).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProductPlugin
//@description: 根据id获取ProductPlugin记录
//@param: id uint
//@return: err error, productPlugin model.ProductPlugin

func GetProductPlugin(id uint) (err error, productPlugin model.ProductPlugin) {
	err = global.GVA_DB.Where("id = ?", id).First(&productPlugin).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetProductPluginInfoList
//@description: 分页获取ProductPlugin记录
//@param: info request.ProductPluginSearch
//@return: err error, list interface{}, total int64

func GetProductPluginInfoList(info request.ProductPluginSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ProductPlugin{})
    var productPlugins []model.ProductPlugin
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&productPlugins).Error
	return err, productPlugins, total
}