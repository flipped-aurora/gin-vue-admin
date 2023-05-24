package kf

import (
	"encoding/xml"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// SignatureOptions 微信服务器验证参数
type SignatureOptions struct {
	Signature string `form:"msg_signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
}

// VerifyURL 验证请求参数是否合法并返回解密后的消息内容
//
//	 //Gin框架的使用示例
//		r.GET("/v1/event/callback", func(c *gin.Context) {
//			options := kf.SignatureOptions{}
//			//获取回调的的校验参数
//			if = c.ShouldBindQuery(&options); err != nil {
//				c.String(http.StatusUnauthorized, "参数解析失败")
//			}
//			// 调用VerifyURL方法校验当前请求，如果合法则把解密后的内容作为响应返回给微信服务器
//			echo, err := kfClient.VerifyURL(options)
//			if err == nil {
//				c.String(http.StatusOK, echo)
//			} else {
//				c.String(http.StatusUnauthorized, "非法请求来源")
//			}
//		})
func (r *Client) VerifyURL(options SignatureOptions) (string, error) {
	if options.Signature != util.Signature(r.ctx.Token, options.TimeStamp, options.Nonce, options.EchoStr) {
		return "", NewSDKErr(40015)
	}
	_, bData, err := util.DecryptMsg(r.corpID, options.EchoStr, r.encodingAESKey)
	if err != nil {
		return "", NewSDKErr(40016)
	}

	return string(bData), nil
}

// 原始回调消息内容
type callbackOriginMessage struct {
	ToUserName string // 企业微信的CorpID，当为第三方套件回调事件时，CorpID的内容为suiteid
	AgentID    string // 接收的应用id，可在应用的设置页面获取
	Encrypt    string // 消息结构体加密后的字符串
}

// CallbackMessage 微信客服回调消息
type CallbackMessage struct {
	ToUserName string `json:"to_user_name"` // 微信客服组件ID
	CreateTime int    `json:"create_time"`  // 消息创建时间，unix时间戳
	MsgType    string `json:"msgtype"`      // 消息的类型，此时固定为 event
	Event      string `json:"event"`        // 事件的类型，此时固定为 kf_msg_or_event
	Token      string `json:"token"`        // 调用拉取消息接口时，需要传此token，用于校验请求的合法性
}

// GetCallbackMessage 获取回调事件中的消息内容
//
//	 //Gin框架的使用示例
//		r.POST("/v1/event/callback", func(c *gin.Context) {
//			var (
//				message kf.CallbackMessage
//				body []byte
//			)
//			// 读取原始消息内容
//			body, err = c.GetRawData()
//			if err != nil {
//				c.String(http.StatusInternalServerError, err.Error())
//				return
//			}
//			// 解析原始数据
//			message, err = kfClient.GetCallbackMessage(body)
//			if err != nil {
//				c.String(http.StatusInternalServerError, "消息获取失败")
//				return
//			}
//			fmt.Println(message)
//			c.String(200, "ok")
//		})
func (r *Client) GetCallbackMessage(encryptedMsg []byte) (msg CallbackMessage, err error) {
	var origin callbackOriginMessage
	if err = xml.Unmarshal(encryptedMsg, &origin); err != nil {
		return msg, err
	}
	_, bData, err := util.DecryptMsg(r.corpID, origin.Encrypt, r.encodingAESKey)
	if err != nil {
		return msg, NewSDKErr(40016)
	}
	if err = xml.Unmarshal(bData, &msg); err != nil {
		return msg, err
	}
	return msg, err
}
