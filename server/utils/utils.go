package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

var (
	lastTimestamp int64  // 上次生成订单号的时间戳
	counter       uint32 // 当前时间戳内已经生成的订单数
	//nodeID        uint32     = 1 // 全局唯一节点 ID（示例值，根据实际情况进行赋值）
	mutex sync.Mutex // 互斥锁
)

// GetOrderNumber 生成订单号15位
func GetOrderNumber() string {
	timestamp := time.Now().Unix() // 获取当前时间戳

	mutex.Lock()
	defer mutex.Unlock()

	if timestamp < lastTimestamp {
		panic("Clock moved backwards")
	}

	if timestamp > lastTimestamp {
		lastTimestamp = timestamp
		counter = 0
	} else {
		counter++
		if counter >= 9999 {
			time.Sleep(time.Millisecond) // 如果当前时间戳内的订单数达到上限，则等待一段时间
		}
	}
	orderNumber := fmt.Sprintf("5%d%04d", lastTimestamp, counter)
	return orderNumber
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

// SetNotifyData 微信设置订单数据
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

// ReplaceAttach 处理棉花糖机器的附加参数
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
