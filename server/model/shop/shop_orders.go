// 自动生成模板ShopOrders
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// shopOrders表 结构体  ShopOrders
type ShopOrders struct {
	global.GVA_MODEL            //ID字段
	AppId            string     `json:"AppId" form:"AppId" gorm:"column:AppId;comment:直连商户申请的公众号或移动应用appid;size:20;"`                                                                                                                                //直连商户申请的公众号或移动应用appid
	MchId            string     `json:"MchId" form:"MchId" gorm:"column:MchId;comment:直连商户的商户号;size:20;"`                                                                                                                                            //直连商户的商户号
	OutTradeNo       string     `json:"OutTradeNo" form:"OutTradeNo" gorm:"column:OutTradeNo;comment:商户系统内部订单号;size:24;"`                                                                                                                            //商户系统内部订单号
	TransactionId    string     `json:"TransactionId" form:"TransactionId" gorm:"column:TransactionId;comment:微信支付系统生成的订单号;size:60;"`                                                                                                                //微信支付系统生成的订单号
	TradeType        string     `json:"TradeType" form:"TradeType" gorm:"column:TradeType;comment:交易类型 枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付;size:16;"`                                                   //交易类型 枚举值：JSAPI：公众号支付, NATIVE：扫码支付, APP：APP支付, MICROPAY：付款码支付, MWEB：H5支付, FACEPAY：刷脸支付
	TradeState       string     `json:"TradeState" form:"TradeState" gorm:"column:TradeState;comment:交易状态 枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败);size:16;"` //交易状态 枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, CLOSED：已关闭, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)
	BankType         string     `json:"BankType" form:"BankType" gorm:"column:BankType;comment:银行类型;size:32;"`                                                                                                                                       //银行类型
	Attach           string     `json:"Attach" form:"Attach" gorm:"column:Attach;comment:附加数据;size:255;"`                                                                                                                                            //附加数据
	SuccessTime      *time.Time `json:"SuccessTime" form:"SuccessTime" gorm:"column:SuccessTime;comment:支付完成时间;"`                                                                                                                                    //支付完成时间
	OpenId           string     `json:"OpenId" form:"OpenId" gorm:"column:OpenId;comment:用户在直连商户下的唯一标识;size:50;"`                                                                                                                                    //用户在直连商户appid下的唯一标识
	SubOpenId        string     `json:"SubOpenId" form:"SubOpenId" gorm:"column:SubOpenId;comment:用户在sub_appid下的标识;size:50;"`                                                                                                                        //用户在sub_appid下的标识
	Total            *int64     `json:"Total" form:"Total" gorm:"column:Total;type:bigint(20);comment:订单总金额，单位为分;size:20;"`                                                                                                                          //订单总金额，单位为分
	PayerTotal       *int64     `json:"PayerTotal" form:"PayerTotal" gorm:"column:PayerTotal;type:bigint(20);comment:用户支付金额，单位为分;size:20;"`                                                                                                          //用户支付金额，单位为分
	UserId           *int       `json:"UserId" form:"UserId" gorm:"column:UserId;type:int(10);comment:用户id 关联;size:10;"`                                                                                                                             //用户id 关联
	AddrId           *int       `json:"AddrId" form:"AddrId" gorm:"column:AddrId;type:int(10);comment:地址id 关联;size:10;"`                                                                                                                             //地址id 关联
	GoodsId          int        `json:"GoodsId" form:"GoodsId" gorm:"column:GoodsId;type:int(10);comment:商品ID 关联;size:10;"`                                                                                                                          //商品ID 关联
	ShopID           string     `json:"ShopID" form:"ShopID" gorm:"column:ShopID;comment:店铺ID 关联;size:32;"`                                                                                                                                          //店铺ID 关联
	TransportId      *int       `json:"TransportId" form:"TransportId" gorm:"column:TransportId;type:int(10);comment:配送ID 关联;size:10;"`                                                                                                              //配送ID 关联
	Endtime          *time.Time `json:"Endtime" form:"Endtime" gorm:"column:Endtime;comment:订单结束时间;"`                                                                                                                                                //订单结束时间
	GoodsTitle       string     `json:"GoodsTitle" form:"GoodsTitle" gorm:"column:GoodsTitle;comment:商品标题;size:255;"`                                                                                                                                //商品标题
	GoodsPrice       *int64     `json:"GoodsPrice" form:"GoodsPrice" gorm:"column:GoodsPrice;type:bigint(20);comment:商品价格;size:20;"`                                                                                                                 //商品价格
	PostFee          *int64     `json:"PostFee" form:"PostFee" gorm:"column:PostFee;type:bigint(10);comment:商品邮费，单位为分;size:10;"`                                                                                                                     //商品邮费，单位为分
	CallUrl          string     `json:"CallUrl" form:"CallUrl" gorm:"column:CallUrl;comment:第三方回调地址;size:255;"`
	PayMent          string     `json:"PayMent" form:"PayMent" gorm:"column:PayMent;comment:支付方式，1=微信，2=支付宝;size:2;"` //支付方式 1 =微信 2 =支付宝 0=未知
	DeviceId         string     `json:"DeviceId" form:"DeviceId" gorm:"column:DeviceId;comment:设备ID;size:32;"`        //设备ID                                                                                                                                                                                    //设备ID
}

// TableName shopOrders表 ShopOrders自定义表名 shop_orders
func (ShopOrders) TableName() string {
	return "shop_orders"
}
