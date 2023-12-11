package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopGoodsService struct {
}

// CreateShopGoods 创建shopGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) CreateShopGoods(shopGoods *shop.ShopGoods) (err error) {
	err = global.GVA_DB.Create(shopGoods).Error
	return err
}

// DeleteShopGoods 删除shopGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) DeleteShopGoods(shopGoods shop.ShopGoods) (err error) {
	err = global.GVA_DB.Delete(&shopGoods).Error
	return err
}

// DeleteShopGoodsByIds 批量删除shopGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) DeleteShopGoodsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shop.ShopGoods{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShopGoods 更新shopGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) UpdateShopGoods(shopGoods shop.ShopGoods) (err error) {
	err = global.GVA_DB.Save(&shopGoods).Error
	return err
}

// GetShopGoods 根据id获取shopGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) GetShopGoods(id uint) (shopGoods shop.ShopGoods, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shopGoods).Error
	return
}

// GetShopGoodsInfoList 分页获取shopGoods表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopGoodsService *ShopGoodsService) GetShopGoodsInfoList(info shopReq.ShopGoodsSearch) (list []shop.ShopGoods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shop.ShopGoods{})
	var shopGoodss []shop.ShopGoods
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GoodsShopid != "" {
		db = db.Where("goods_shopid = ?", info.GoodsShopid)
	}
	if info.GoodsMacid != "" {
		db = db.Where("goods_Macid = ?", info.GoodsMacid)
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name = ?", info.GoodsName)
	}
	if info.GoodsStatus != "" {
		db = db.Where("goods_status = ?", info.GoodsStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&shopGoodss).Error
	return shopGoodss, total, err
}
