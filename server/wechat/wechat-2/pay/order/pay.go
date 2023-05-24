package order

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/pay/config"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
var payGateway = "https://api.mch.weixin.qq.com/pay/unifiedorder"

// SUCCESS 表示支付成功
const SUCCESS = "SUCCESS"

// Order struct extends context
type Order struct {
	*config.Config
}

// NewOrder return an instance of order package
func NewOrder(cfg *config.Config) *Order {
	order := Order{cfg}
	return &order
}

// Params was NEEDED when request Unified order
// 传入的参数，用于生成 prepay_id 的必需参数
type Params struct {
	TotalFee   string
	CreateIP   string
	Body       string
	OutTradeNo string
	TimeExpire string // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。
	OpenID     string
	TradeType  string
	SignType   string
	Detail     string
	Attach     string
	GoodsTag   string
	NotifyURL  string
}

// Config 是传出用于 js sdk 用的参数
type Config struct {
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	PrePayID  string `json:"prePayId"`
	SignType  string `json:"signType"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
}

// ConfigForApp 是传出用于 app sdk 用的参数
type ConfigForApp struct {
	AppID     string `json:"appid"`
	MchID     string `json:"partnerid"` // 微信支付分配的商户号
	PrePayID  string `json:"prepayid"`
	Package   string `json:"package"`
	NonceStr  string `json:"nonceStr"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}

// PreOrder 是 Unified order 接口的返回
type PreOrder struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid,omitempty"`
	MchID      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrePayID   string `xml:"prepay_id,omitempty"`
	CodeURL    string `xml:"code_url,omitempty"`
	MWebURL    string `xml:"mweb_url,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
}

// payRequest 接口请求参数
type payRequest struct {
	AppID          string `xml:"appid"`                 // 公众账号ID
	MchID          string `xml:"mch_id"`                // 商户号
	DeviceInfo     string `xml:"device_info,omitempty"` // 设备号
	NonceStr       string `xml:"nonce_str"`             // 随机字符串
	Sign           string `xml:"sign"`                  // 签名
	SignType       string `xml:"sign_type,omitempty"`   // 签名类型
	Body           string `xml:"body"`                  // 商品描述
	Detail         string `xml:"detail,omitempty"`      // 商品详情
	Attach         string `xml:"attach,omitempty"`      // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`          // 商户订单号
	FeeType        string `xml:"fee_type,omitempty"`    // 标价币种
	TotalFee       string `xml:"total_fee"`             // 标价金额
	SpbillCreateIP string `xml:"spbill_create_ip"`      // 终端IP
	TimeStart      string `xml:"time_start,omitempty"`  // 交易起始时间
	TimeExpire     string `xml:"time_expire,omitempty"` // 交易结束时间
	GoodsTag       string `xml:"goods_tag,omitempty"`   // 订单优惠标记
	NotifyURL      string `xml:"notify_url"`            // 通知地址
	TradeType      string `xml:"trade_type"`            // 交易类型
	ProductID      string `xml:"product_id,omitempty"`  // 商品ID
	LimitPay       string `xml:"limit_pay,omitempty"`   // 指定支付方式
	OpenID         string `xml:"openid,omitempty"`      // 用户标识
	SceneInfo      string `xml:"scene_info,omitempty"`  // 场景信息

	XMLName struct{} `xml:"xml"`
}

func (req *payRequest) BridgePayRequest(p *Params, AppID, MchID, nonceStr, sign string) *payRequest {
	request := payRequest{
		AppID:          AppID,
		MchID:          MchID,
		NonceStr:       nonceStr,
		Sign:           sign,
		Body:           p.Body,
		OutTradeNo:     p.OutTradeNo,
		TotalFee:       p.TotalFee,
		SpbillCreateIP: p.CreateIP,
		NotifyURL:      p.NotifyURL,
		TradeType:      p.TradeType,
		OpenID:         p.OpenID,
		SignType:       p.SignType,
		Detail:         p.Detail,
		Attach:         p.Attach,
		GoodsTag:       p.GoodsTag,
	}
	return &request
}

// BridgeConfig get js bridge config
func (o *Order) BridgeConfig(p *Params) (cfg Config, err error) {
	var (
		buffer    strings.Builder
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	)
	order, err := o.PrePayOrder(p)
	if err != nil {
		return
	}
	buffer.WriteString("appId=")
	buffer.WriteString(order.AppID)
	buffer.WriteString("&nonceStr=")
	buffer.WriteString(order.NonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString("prepay_id=" + order.PrePayID)
	buffer.WriteString("&signType=")
	buffer.WriteString(p.SignType)
	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&key=")
	buffer.WriteString(o.Key)

	sign, err := util.CalculateSign(buffer.String(), p.SignType, o.Key)
	if err != nil {
		return
	}
	// 签名
	cfg.PaySign = sign
	cfg.NonceStr = order.NonceStr
	cfg.Timestamp = timestamp
	cfg.PrePayID = order.PrePayID
	cfg.SignType = p.SignType
	cfg.Package = "prepay_id=" + order.PrePayID
	return
}

// BridgeAppConfig get app bridge config
func (o *Order) BridgeAppConfig(p *Params) (cfg ConfigForApp, err error) {
	var (
		timestamp = strconv.FormatInt(util.GetCurrTS(), 10)
		noncestr  = util.RandomStr(32)
		_package  = "Sign=WXPay"
	)
	order, err := o.PrePayOrder(p)
	if err != nil {
		return
	}

	result := map[string]string{
		"appid":     order.AppID,
		"partnerid": order.MchID,
		"prepayid":  order.PrePayID,
		"package":   _package,
		"noncestr":  noncestr,
		"timestamp": timestamp,
	}
	// 签名
	sign, err := util.ParamSign(result, o.Key)
	if err != nil {
		return
	}
	result["sign"] = sign
	cfg = ConfigForApp{
		AppID:     result["appid"],
		MchID:     result["partnerid"],
		PrePayID:  result["prepayid"],
		Package:   result["package"],
		NonceStr:  result["noncestr"],
		Timestamp: result["timestamp"],
		Sign:      result["sign"],
	}
	return
}

// PrePayOrder return data for invoke wechat payment
func (o *Order) PrePayOrder(p *Params) (payOrder PreOrder, err error) {
	nonceStr := util.RandomStr(32)

	// 通知地址
	if len(p.NotifyURL) == 0 {
		p.NotifyURL = o.NotifyURL // 默认使用order.NotifyURL
	}

	// 签名类型
	if p.SignType == "" {
		p.SignType = util.SignTypeMD5
	}

	param := map[string]string{
		"appid":            o.AppID,
		"body":             p.Body,
		"mch_id":           o.MchID,
		"nonce_str":        nonceStr,
		"out_trade_no":     p.OutTradeNo,
		"spbill_create_ip": p.CreateIP,
		"total_fee":        p.TotalFee,
		"trade_type":       p.TradeType,
		"openid":           p.OpenID,
		"sign_type":        p.SignType,
		"detail":           p.Detail,
		"attach":           p.Attach,
		"goods_tag":        p.GoodsTag,
		"notify_url":       p.NotifyURL,
	}

	if p.TimeExpire != "" {
		// 如果有传入交易结束时间
		param["time_expire"] = p.TimeExpire
	}

	sign, err := util.ParamSign(param, o.Key)
	if err != nil {
		return
	}
	request := new(payRequest).BridgePayRequest(p, o.AppID, o.MchID, nonceStr, sign)
	if len(p.TimeExpire) > 0 {
		// 如果有传入交易结束时间
		request.TimeExpire = p.TimeExpire
	}
	rawRet, err := util.PostXML(payGateway, request)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &payOrder)
	if err != nil {
		return
	}
	if payOrder.ReturnCode == SUCCESS {
		// pay success
		if payOrder.ResultCode == SUCCESS {
			err = nil
			return
		}
		err = errors.New(payOrder.ErrCode + payOrder.ErrCodeDes)
		return
	}
	err = errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [sign : " + sign + "]")
	return
}

// PrePayID will request wechat merchant api and request for a pre-payment order id
func (o *Order) PrePayID(p *Params) (prePayID string, err error) {
	order, err := o.PrePayOrder(p)
	if err != nil {
		return
	}
	if order.PrePayID == "" {
		err = errors.New("empty prepayid")
	}
	prePayID = order.PrePayID
	return
}
