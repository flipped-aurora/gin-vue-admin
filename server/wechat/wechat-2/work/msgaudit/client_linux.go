//go:build linux && cgo && msgaudit
// +build linux,cgo,msgaudit

// Package msgaudit only for linux
package msgaudit

// #cgo LDFLAGS: -L${SRCDIR}/lib -lWeWorkFinanceSdk_C
// #cgo CFLAGS: -I ./lib/
// #include <stdlib.h>
// #include "WeWorkFinanceSdk_C.h"
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/config"
)

// Client 会话存档
type Client struct {
	ptr        *C.WeWorkFinanceSdk_t
	privateKey string
}

// NewClient 初始会话会话存档实例
/**
* 初始化函数
* Return值=0表示该API调用成功
*
* @param [in]  sdk			NewSdk返回的sdk指针
* @param [in]  corpid       调用企业的企业id，例如：wwd08c8exxxx5ab44d，可以在企业微信管理端--我的企业--企业信息查看
* @param [in]  secret		 聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
* @param [in]  privateKey    消息加密私钥，可以在企业微信管理端--管理工具--消息加密公钥查看对用公钥，私钥一般由自己保存
*
*
* @return 返回是否初始化成功
*      0   - 成功
*      !=0 - 失败
 */
func NewClient(cfg *config.Config) (*Client, error) {
	ptr := C.NewSdk()
	corpIDC := C.CString(cfg.CorpID)
	corpSecretC := C.CString(cfg.CorpSecret)
	defer func() {
		C.free(unsafe.Pointer(corpIDC))
		C.free(unsafe.Pointer(corpSecretC))
	}()
	retC := C.Init(ptr, corpIDC, corpSecretC)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	return &Client{
		ptr:        ptr,
		privateKey: cfg.RasPrivateKey,
	}, nil
}

// Free 释放SDK实例是可调用该方法释放内存
func (s *Client) Free() {
	if s.ptr == nil {
		return
	}
	C.DestroySdk(s.ptr)
	s.ptr = nil
}

// GetChatData 拉取聊天记录函数
/**
* 拉取聊天记录函数
*
*
* @param [in]  seq             从指定的seq开始拉取消息，注意的是返回的消息从seq+1开始返回，seq为之前接口返回的最大seq值。首次使用请使用seq:0
* @param [in]  limit           一次拉取的消息条数，最大值1000条，超过1000条会返回错误
* @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
* @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
* @param [in]  timeout         超时时间，单位秒
* @return chatDatas       返回本次拉取消息的数据，slice结构体.内容包括errcode/errmsg，以及每条消息内容。示例如下：

{"errcode":0,"errmsg":"ok","chatdata":[{"seq":196,"msgid":"CAQQ2fbb4QUY0On2rYSAgAMgip/yzgs=","publickey_ver":3,"encrypt_random_key":"ftJ+uz3n/z1DsxlkwxNgE+mL38H42/KCvN8T60gbbtPD+Rta1hKTuQPzUzO6Hzne97MgKs7FfdDxDck/v8cDT6gUVjA2tZ/M7euSD0L66opJ/IUeBtpAtvgVSD5qhlaQjvfKJc/zPMGNK2xCLFYqwmQBZXbNT7uA69Fflm512nZKW/piK2RKdYJhRyvQnA1ISxK097sp9WlEgDg250fM5tgwMjujdzr7ehK6gtVBUFldNSJS7ndtIf6aSBfaLktZgwHZ57ONewWq8GJe7WwQf1hwcDbCh7YMG8nsweEwhDfUz+u8rz9an+0lgrYMZFRHnmzjgmLwrR7B/32Qxqd79A==","encrypt_chat_msg":"898WSfGMnIeytTsea7Rc0WsOocs0bIAerF6de0v2cFwqo9uOxrW9wYe5rCjCHHH5bDrNvLxBE/xOoFfcwOTYX0HQxTJaH0ES9OHDZ61p8gcbfGdJKnq2UU4tAEgGb8H+Q9n8syRXIjaI3KuVCqGIi4QGHFmxWenPFfjF/vRuPd0EpzUNwmqfUxLBWLpGhv+dLnqiEOBW41Zdc0OO0St6E+JeIeHlRZAR+E13Isv9eS09xNbF0qQXWIyNUi+ucLr5VuZnPGXBrSfvwX8f0QebTwpy1tT2zvQiMM2MBugKH6NuMzzuvEsXeD+6+3VRqL"}]}
*/
func (s *Client) GetChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) ([]ChatData, error) {
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	chatSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeSlice(chatSlice)
	}()

	if s.ptr == nil {
		return nil, NewSDKErr(10002)
	}

	retC := C.GetChatData(s.ptr, C.ulonglong(seq), C.uint(limit), proxyC, passwdC, C.int(timeout), chatSlice)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	buf := s.GetContentFromSlice(chatSlice)

	var data ChatDataResponse
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	if data.ErrCode != 0 {
		return nil, data.Error
	}
	return data.ChatDataList, nil
}

// GetRawChatData 拉取聊天记录函数
/**
* 拉取聊天记录函数
*
*
* @param [in]  seq             从指定的seq开始拉取消息，注意的是返回的消息从seq+1开始返回，seq为之前接口返回的最大seq值。首次使用请使用seq:0
* @param [in]  limit           一次拉取的消息条数，最大值1000条，超过1000条会返回错误
* @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
* @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
* @param [in]  timeout         超时时间，单位秒
* @return chatDatas       返回本次拉取消息的数据，slice结构体.内容包括errcode/errmsg，以及每条消息内容。示例如下：

{"errcode":0,"errmsg":"ok","chatdata":[{"seq":196,"msgid":"CAQQ2fbb4QUY0On2rYSAgAMgip/yzgs=","publickey_ver":3,"encrypt_random_key":"ftJ+uz3n/z1DsxlkwxNgE+mL38H42/KCvN8T60gbbtPD+Rta1hKTuQPzUzO6Hzne97MgKs7FfdDxDck/v8cDT6gUVjA2tZ/M7euSD0L66opJ/IUeBtpAtvgVSD5qhlaQjvfKJc/zPMGNK2xCLFYqwmQBZXbNT7uA69Fflm512nZKW/piK2RKdYJhRyvQnA1ISxK097sp9WlEgDg250fM5tgwMjujdzr7ehK6gtVBUFldNSJS7ndtIf6aSBfaLktZgwHZ57ONewWq8GJe7WwQf1hwcDbCh7YMG8nsweEwhDfUz+u8rz9an+0lgrYMZFRHnmzjgmLwrR7B/32Qxqd79A==","encrypt_chat_msg":"898WSfGMnIeytTsea7Rc0WsOocs0bIAerF6de0v2cFwqo9uOxrW9wYe5rCjCHHH5bDrNvLxBE/xOoFfcwOTYX0HQxTJaH0ES9OHDZ61p8gcbfGdJKnq2UU4tAEgGb8H+Q9n8syRXIjaI3KuVCqGIi4QGHFmxWenPFfjF/vRuPd0EpzUNwmqfUxLBWLpGhv+dLnqiEOBW41Zdc0OO0St6E+JeIeHlRZAR+E13Isv9eS09xNbF0qQXWIyNUi+ucLr5VuZnPGXBrSfvwX8f0QebTwpy1tT2zvQiMM2MBugKH6NuMzzuvEsXeD+6+3VRqL"}]}
*/
func (s *Client) GetRawChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) (ChatDataResponse, error) {
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	chatSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeSlice(chatSlice)
	}()

	if s.ptr == nil {
		return ChatDataResponse{}, NewSDKErr(10002)
	}

	retC := C.GetChatData(s.ptr, C.ulonglong(seq), C.uint(limit), proxyC, passwdC, C.int(timeout), chatSlice)
	ret := int(retC)
	if ret != 0 {
		return ChatDataResponse{}, NewSDKErr(ret)
	}
	buf := s.GetContentFromSlice(chatSlice)

	var data ChatDataResponse
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return ChatDataResponse{}, err
	}
	return data, nil
}

// DecryptData 解析密文.企业微信自有解密内容
/**
* @brief 解析密文.企业微信自有解密内容
* @param [in]  encrypt_key, getchatdata返回的encrypt_random_key,使用企业自持对应版本秘钥RSA解密后的内容
* @param [in]  encrypt_msg, getchatdata返回的encrypt_chat_msg
* @param [out] msg, 解密的消息明文
* @return 返回是否调用成功
*      0   - 成功
*      !=0 - 失败
 */
func (s *Client) DecryptData(encryptRandomKey string, encryptMsg string) (msg ChatMessage, err error) {
	encryptKey, err := util.RSADecryptBase64(s.privateKey, encryptRandomKey)
	if err != nil {
		return msg, err
	}
	encryptKeyC := C.CString(string(encryptKey))
	encryptMsgC := C.CString(encryptMsg)
	msgSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(encryptKeyC))
		C.free(unsafe.Pointer(encryptMsgC))
		C.FreeSlice(msgSlice)
	}()

	retC := C.DecryptData(encryptKeyC, encryptMsgC, msgSlice)
	ret := int(retC)
	if ret != 0 {
		return msg, NewSDKErr(ret)
	}
	buf := s.GetContentFromSlice(msgSlice)

	// handle illegal escape character in text
	for i := 0; i < len(buf); {
		if buf[i] < 0x20 {
			buf = append(buf[:i], buf[i+1:]...)
			continue
		}
		i++
	}

	var baseMessage BaseMessage
	err = json.Unmarshal(buf, &baseMessage)
	if err != nil {
		return msg, err
	}

	msg.ID = baseMessage.MsgID
	msg.From = baseMessage.From
	msg.ToList = baseMessage.ToList
	msg.Action = baseMessage.Action
	msg.Type = baseMessage.MsgType
	msg.originData = buf
	return msg, err
}

// GetMediaData 拉取媒体消息函数
/**
 * 拉取媒体消息函数
 * Return值=0表示该API调用成功
 *
 *
 * @param [in]  sdk				NewSdk返回的sdk指针
 * @param [in]  sdkFileid		从GetChatData返回的聊天消息中，媒体消息包括的sdkfileid
 * @param [in]  proxy			使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
 * @param [in]  passwd			代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
 * @param [in]  indexbuf		媒体消息分片拉取，需要填入每次拉取的索引信息。首次不需要填写，默认拉取512k，后续每次调用只需要将上次调用返回的outindexbuf填入即可。
 * @param [in]  timeout			超时时间，单位秒
 * @param [out] media_data		返回本次拉取的媒体数据.MediaData结构体.内容包括data(数据内容)/outindexbuf(下次索引)/is_finish(拉取完成标记)

 *
 * @return 返回是否调用成功
 *      0   - 成功
 *      !=0 - 失败
 */
func (s *Client) GetMediaData(indexBuf string, sdkFileID string, proxy string, passwd string, timeout int) (*MediaData, error) {
	indexBufC := C.CString(indexBuf)
	sdkFileIDC := C.CString(sdkFileID)
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	mediaDataC := C.NewMediaData()
	defer func() {
		C.free(unsafe.Pointer(indexBufC))
		C.free(unsafe.Pointer(sdkFileIDC))
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeMediaData(mediaDataC)
	}()

	if s.ptr == nil {
		return nil, NewSDKErr(10002)
	}

	retC := C.GetMediaData(s.ptr, indexBufC, sdkFileIDC, proxyC, passwdC, C.int(timeout), mediaDataC)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	return &MediaData{
		OutIndexBuf: C.GoString(C.GetOutIndexBuf(mediaDataC)),
		Data:        C.GoBytes(unsafe.Pointer(C.GetData(mediaDataC)), C.GetDataLen(mediaDataC)),
		IsFinish:    int(C.IsMediaDataFinish(mediaDataC)) == 1,
	}, nil
}

// GetContentFromSlice 从切片内获取内容
func (s *Client) GetContentFromSlice(slice *C.struct_Slice_t) []byte {
	return C.GoBytes(unsafe.Pointer(C.GetContentFromSlice(slice)), C.GetSliceLen(slice))
}
