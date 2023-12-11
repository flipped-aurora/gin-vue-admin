// 自动生成模板ShopGoods
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// shopGoods表 结构体  ShopGoods
type ShopGoods struct {
	global.GVA_MODEL
	GoodsShopid     string     `json:"goodsShopid" form:"goodsShopid" gorm:"column:goods_shopid;comment:店铺id;size:255;"`                                         //店铺id
	GoodsMacid      string     `json:"goodsMacid" form:"goodsMacid" gorm:"column:goods_Macid;comment:设备id;size:255;"`                                            //设备ID
	GoodsName       string     `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名称;size:1024;"`                                              //商品名称
	GoodsDesc       string     `json:"goodsDesc" form:"goodsDesc" gorm:"column:goods_desc;comment:商品描述;size:2048;"`                                              //商品描述
	GoodsPrice      *int64     `json:"goodsPrice" form:"goodsPrice" gorm:"column:goods_price;type:bigint(20);comment:商品价格，单位为：分;size:20;"`                       //商品价格
	GoodsMaketPrice *int64     `json:"goodsMaketPrice,bigint" form:"goodsMaketPrice" gorm:"column:goods_maket_price;type:bigint(20);comment:市场价，单位为：分;size:20;"` //市场价
	GoodsStock      *int       `json:"goodsStock,int" form:"goodsStock" gorm:"column:goods_stock;type:int(10);comment:商品库存;size:10;"`                            //商品库存
	GoodsImg        string     `json:"goodsImg" form:"goodsImg" gorm:"column:goods_img;comment:商品图片链接;size:255;"`                                                //图片链接
	GoodsSellTime   *time.Time `json:"goodsSellTime" form:"goodsSellTime" gorm:"column:goods_sell_time;comment:上架时间;"`                                           //上架时间
	GoodsStatus     string     `json:"goodsStatus" form:"goodsStatus" gorm:"column:goods_status;type:smallint(2);comment:商品状态，0待审核，1正常售卖，2已下架，3已售罄;size:10;"`    //商品状态
	GoodsPostfee    *int64     `json:"goodsPostfee" form:"goodsPostfee" gorm:"column:goods_postfee;type:bigint(10);comment:商品邮费，单位为分;size:19;"`                  //商品邮费
}

// TableName shopGoods表 ShopGoods自定义表名 shop_goods
func (ShopGoods) TableName() string {
	return "shop_goods"
}
