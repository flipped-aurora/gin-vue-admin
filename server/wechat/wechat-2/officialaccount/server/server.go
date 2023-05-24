package server

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/message"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// Server struct
type Server struct {
	*context.Context
	Writer  http.ResponseWriter
	Request *http.Request

	skipValidate bool

	openID string

	messageHandler func(*message.MixMessage) *message.Reply

	RequestRawXMLMsg  []byte
	RequestMsg        *message.MixMessage
	ResponseRawXMLMsg []byte
	ResponseMsg       interface{}

	isSafeMode    bool
	isJSONContent bool
	random        []byte
	nonce         string
	timestamp     int64
}

// NewServer init
func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

// SkipValidate set skip validate
func (srv *Server) SkipValidate(skip bool) {
	srv.skipValidate = skip
}

// Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	if !srv.Validate() {
		log.Error("Validate Signature Failed.")
		return fmt.Errorf("请求校验失败")
	}

	echostr, exists := srv.GetQuery("echostr")
	if exists {
		srv.String(echostr)
		return nil
	}

	response, err := srv.handleRequest()
	if err != nil {
		return err
	}
	// 非安全模式下，请求处理方法返回为nil则直接回复success给微信服务器
	if response == nil && !srv.isSafeMode {
		srv.String("success")
		return nil
	}

	// debug print request msg
	log.Debugf("request msg =%s", string(srv.RequestRawXMLMsg))

	return srv.buildResponse(response)
}

// Validate 校验请求是否合法
func (srv *Server) Validate() bool {
	if srv.skipValidate {
		return true
	}
	timestamp := srv.Query("timestamp")
	nonce := srv.Query("nonce")
	signature := srv.Query("signature")
	log.Debugf("validate signature, timestamp=%s, nonce=%s", timestamp, nonce)
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

// HandleRequest 处理微信的请求
func (srv *Server) handleRequest() (reply *message.Reply, err error) {
	// set isSafeMode
	srv.isSafeMode = false
	encryptType := srv.Query("encrypt_type")
	if encryptType == "aes" {
		srv.isSafeMode = true
	}

	// set request contentType
	contentType := srv.Request.Header.Get("Content-Type")
	srv.isJSONContent = strings.Contains(contentType, "application/json")

	// set openID
	srv.openID = srv.Query("openid")

	var msg interface{}
	msg, err = srv.getMessage()
	if err != nil {
		return
	}
	mixMessage, success := msg.(*message.MixMessage)
	if !success {
		err = errors.New("消息类型转换失败")
	}
	srv.RequestMsg = mixMessage
	reply = srv.messageHandler(mixMessage)
	return
}

// GetOpenID return openID
func (srv *Server) GetOpenID() string {
	return srv.openID
}

// getMessage 解析微信返回的消息
func (srv *Server) getMessage() (interface{}, error) {
	var rawXMLMsgBytes []byte
	var err error
	if srv.isSafeMode {
		encryptedXMLMsg, dataErr := srv.getEncryptBody()
		if dataErr != nil {
			return nil, dataErr
		}

		// 验证消息签名
		timestamp := srv.Query("timestamp")
		srv.timestamp, err = strconv.ParseInt(timestamp, 10, 32)
		if err != nil {
			return nil, err
		}
		nonce := srv.Query("nonce")
		srv.nonce = nonce
		msgSignature := srv.Query("msg_signature")
		msgSignatureGen := util.Signature(srv.Token, timestamp, nonce, encryptedXMLMsg.EncryptedMsg)
		if msgSignature != msgSignatureGen {
			return nil, fmt.Errorf("消息不合法，验证签名失败")
		}

		// 解密
		srv.random, rawXMLMsgBytes, err = util.DecryptMsg(srv.AppID, encryptedXMLMsg.EncryptedMsg, srv.EncodingAESKey)
		if err != nil {
			return nil, fmt.Errorf("消息解密失败, err=%v", err)
		}
	} else {
		rawXMLMsgBytes, err = io.ReadAll(srv.Request.Body)
		if err != nil {
			return nil, fmt.Errorf("从body中解析xml失败, err=%v", err)
		}
	}

	srv.RequestRawXMLMsg = rawXMLMsgBytes

	return srv.parseRequestMessage(rawXMLMsgBytes)
}

func (srv *Server) getEncryptBody() (*message.EncryptedXMLMsg, error) {
	var encryptedXMLMsg = &message.EncryptedXMLMsg{}
	if srv.isJSONContent {
		if err := json.NewDecoder(srv.Request.Body).Decode(encryptedXMLMsg); err != nil {
			return nil, fmt.Errorf("从body中解析json失败,err=%v", err)
		}
	} else {
		if err := xml.NewDecoder(srv.Request.Body).Decode(encryptedXMLMsg); err != nil {
			return nil, fmt.Errorf("从body中解析xml失败,err=%v", err)
		}
	}
	return encryptedXMLMsg, nil
}

func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg *message.MixMessage, err error) {
	msg = &message.MixMessage{}
	if !srv.isJSONContent {
		err = xml.Unmarshal(rawXMLMsgBytes, msg)
		return
	}
	// parse json
	err = json.Unmarshal(rawXMLMsgBytes, msg)
	if err != nil {
		return
	}
	// nonstandard json, 目前小程序订阅消息返回数据格式不标准，订阅消息模板单个List返回是对象，多个List返回是数组。
	if msg.MsgType == message.MsgTypeEvent {
		listData := gjson.Get(string(rawXMLMsgBytes), "List")
		if listData.IsObject() {
			listItem := message.SubscribeMsgPopupEvent{}
			if parseErr := json.Unmarshal([]byte(listData.Raw), &listItem); parseErr != nil {
				return msg, parseErr
			}
			msg.SetSubscribeMsgPopupEvents([]message.SubscribeMsgPopupEvent{listItem})
		} else if listData.IsArray() {
			listItems := make([]message.SubscribeMsgPopupEvent, 0)
			if parseErr := json.Unmarshal([]byte(listData.Raw), &listItems); parseErr != nil {
				return msg, parseErr
			}
			msg.SetSubscribeMsgPopupEvents(listItems)
		}
	}
	return
}

// SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(*message.MixMessage) *message.Reply) {
	srv.messageHandler = handler
}

func (srv *Server) buildResponse(reply *message.Reply) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		// do nothing
		return nil
	}
	msgType := reply.MsgType
	switch msgType {
	case message.MsgTypeText:
	case message.MsgTypeImage:
	case message.MsgTypeVoice:
	case message.MsgTypeVideo:
	case message.MsgTypeMusic:
	case message.MsgTypeNews:
	case message.MsgTypeTransfer:
	default:
		err = message.ErrUnsupportReply
		return
	}

	msgData := reply.MsgData
	value := reflect.ValueOf(msgData)
	// msgData must be a ptr
	kind := value.Kind().String()
	if kind != "ptr" {
		return message.ErrUnsupportReply
	}

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(srv.RequestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(srv.RequestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

	params[0] = reflect.ValueOf(msgType)
	value.MethodByName("SetMsgType").Call(params)

	params[0] = reflect.ValueOf(util.GetCurrTS())
	value.MethodByName("SetCreateTime").Call(params)

	srv.ResponseMsg = msgData
	srv.ResponseRawXMLMsg, err = xml.Marshal(msgData)
	return
}

// Send 将自定义的消息发送
func (srv *Server) Send() (err error) {
	replyMsg := srv.ResponseMsg
	log.Debugf("response msg =%+v", replyMsg)
	if srv.isSafeMode {
		// 安全模式下对消息进行加密
		var encryptedMsg []byte
		encryptedMsg, err = util.EncryptMsg(srv.random, srv.ResponseRawXMLMsg, srv.AppID, srv.EncodingAESKey)
		if err != nil {
			return
		}
		// TODO 如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		timestampStr := strconv.FormatInt(timestamp, 10)
		msgSignature := util.Signature(srv.Token, timestampStr, srv.nonce, string(encryptedMsg))
		replyMsg = message.ResponseEncryptedXMLMsg{
			EncryptedMsg: string(encryptedMsg),
			MsgSignature: msgSignature,
			Timestamp:    timestamp,
			Nonce:        srv.nonce,
		}
	}
	if replyMsg != nil {
		srv.XML(replyMsg)
	}
	return
}
