package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	SUCCESS        = "SUCCESS"        //支付成功
	REFUND         = "REFUND"         //转入退款
	NOTPAY         = "NOTPAY"         //未支付
	CLOSED         = "CLOSED"         //已关闭
	REVOKED        = "REVOKED"        //已撤销
	USERPAYING     = "USERPAYING"     //支付中
	PAYERROR       = "PAYERROR"       //支付失败
	WAIT_BUYER_PAY = "WAIT_BUYER_PAY" //交易创建
	TRADE_CLOSED   = "TRADE_CLOSED"   //交易超时
	TRADE_SUCCESS  = "TRADE_SUCCESS"  //支付成功
	TRADE_FINISHED = "TRADE_FINISHED" //交易结束
)

const (
	Wxpay  = "1"
	Alipay = "2"
)

var num int64

// 生成24位订单号
// 前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
func GetOrderNumber() string {
	t := time.Now()
	s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

// SetOrderData 组装数据 注意：js下单需要openid  native不需要
func SetOrderData(appID, mchID, attach, callUrl, openid string) (res *shop.ShopOrders) {
	var order shop.ShopOrders
	order.AppId = appID
	order.MchId = mchID
	order.Attach = attach
	order.CallUrl = callUrl
	order.OpenId = openid
	return &order
}

func SetNotifyData(order *shop.ShopOrders, decode *wechat.V3DecryptResult) {
	order.TransactionId = decode.TransactionId
	order.TradeType = decode.TradeType
	order.TradeState = decode.TradeState
	order.BankType = decode.BankType
	order.OpenId = decode.Payer.Openid
	order.SubOpenId = decode.Payer.SubOpenid
	order.PayMent = Wxpay
	payerTotal := int64(decode.Amount.PayerTotal)
	order.PayerTotal = &payerTotal
	order.AppId = decode.Appid
	order.MchId = decode.Mchid
	order.Attach = decode.Attach
	t, err := time.Parse(time.RFC3339, decode.SuccessTime)
	if err != nil {
		return
	}
	order.SuccessTime = &t
	return

}

func HttpPost(url, outTradeNo string, code int) {
	data := make(map[string]string)
	data["tradeno"] = outTradeNo
	data["code"] = strconv.Itoa(code)
	bytesData, _ := json.Marshal(data)
	client := &http.Client{}
	rep, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	rep.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resq, _ := client.Do(rep)
	if resq != nil {
		defer resq.Body.Close()
	}
	return
}

func HttpGet(url string) string {
	// 创建一个HTTP客户端
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return ""
	}
	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时发生错误:", err)
		return ""
	}
	defer resp.Body.Close()
	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return ""
	}
	return string(body)
}

func ReplaceAttach(s string) (str *gopay.BodyMap, arr error) {
	attach := strings.Replace(s, "?attach=null", "", 1) //有多余数据 字符串替换处理一下
	arrstr := strings.Split(attach, ",")                //把字符串用逗号分割
	//?attach=ZZ9KDT845Z,13,571434121592954,1314,1226862,1369596012470",
	if len(arrstr) < 6 {
		return nil, errors.New("数据解析失败")
	}
	m := make(gopay.BodyMap)
	m.Set("macid", arrstr[0])
	m.Set("goodsid", arrstr[1])
	m.Set("outTreadNo", arrstr[2])
	m.Set("money", arrstr[3])
	m.Set("agencyNo", arrstr[4])
	m.Set("hlMerchantId", arrstr[5])
	return &m, nil
}

// {ts=1699067732664, hlMerchantId=1369596012470, key=625031257, outTreadNo=569621434041902, macid=ZZ9KDT845Z, agencyNo=1226862}
func ParseRequest(str string) (s *request.RequestMhtData) {
	// 去除大括号
	str = strings.TrimLeft(str, "{")
	str = strings.TrimRight(str, "}")
	// 按逗号分割字符串
	pairs := strings.Split(str, ",")
	// 初始化结构体实例
	data := &request.RequestMhtData{}
	// 提取键值对并填充到结构体中
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair != "" {
			parts := strings.Split(pair, "=")
			if len(parts) != 2 {
				continue
			}
			key := parts[0]
			value := parts[1]
			switch key {
			case "ts":
				data.Ts = value
			case "hlMerchantId":
				data.HlMerchantId = value
			case "key":
				data.Key = value
			case "outTreadNo":
				data.OutTreadNo = value
			case "macid":
				data.Macid = value
			case "agencyNo":
				data.AgencyNo = value
			case "patternId":
				data.PatternId = value
			}
		}
	}
	// 打印解析后的数据
	return data
}
