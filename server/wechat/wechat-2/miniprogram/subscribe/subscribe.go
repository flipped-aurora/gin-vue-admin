package subscribe

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 发送订阅消息
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	subscribeSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"

	// 获取当前帐号下的个人模板列表
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
	getTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"

	// 添加订阅模板
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
	addTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate"

	// 删除私有模板
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html
	delTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate"

	// 统一服务消息
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
	uniformMessageSend = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
)

// Subscribe 订阅消息
type Subscribe struct {
	*context.Context
}

// NewSubscribe 实例化
func NewSubscribe(ctx *context.Context) *Subscribe {
	return &Subscribe{Context: ctx}
}

// Message 订阅消息请求参数
type Message struct {
	ToUser           string               `json:"touser"`            // 必选，接收者（用户）的 openid
	TemplateID       string               `json:"template_id"`       // 必选，所需下发的订阅模板id
	Page             string               `json:"page"`              // 可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]*DataItem `json:"data"`              // 必选, 模板内容
	MiniprogramState string               `json:"miniprogram_state"` // 可选，跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string               `json:"lang"`              // 入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value interface{} `json:"value"`
	Color string      `json:"color"`
}

// TemplateItem template item
type TemplateItem struct {
	PriTmplID string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int64  `json:"type"`
}

// TemplateList template list
type TemplateList struct {
	util.CommonError
	Data []TemplateItem `json:"data"`
}

// Send 发送订阅消息
func (s *Subscribe) Send(msg *Message) (err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "Send")
}

// ListTemplates 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func (s *Subscribe) ListTemplates() (*TemplateList, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", getTemplateURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}
	templateList := TemplateList{}
	err = util.DecodeWithError(response, &templateList, "ListTemplates")
	if err != nil {
		return nil, err
	}
	return &templateList, nil
}

// UniformMessage 统一服务消息
type UniformMessage struct {
	ToUser           string `json:"touser"`
	WeappTemplateMsg struct {
		TemplateID      string               `json:"template_id"`
		Page            string               `json:"page"`
		FormID          string               `json:"form_id"`
		Data            map[string]*DataItem `json:"data"`
		EmphasisKeyword string               `json:"emphasis_keyword"`
	} `json:"weapp_template_msg"`
	MpTemplateMsg struct {
		Appid       string `json:"appid"`
		TemplateID  string `json:"template_id"`
		URL         string `json:"url"`
		Miniprogram struct {
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"miniprogram"`
		Data map[string]*DataItem `json:"data"`
	} `json:"mp_template_msg"`
}

// UniformSend 发送统一服务消息
func (s *Subscribe) UniformSend(msg *UniformMessage) (err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uniformMessageSend, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "UniformSend")
}

type resSubscribeAdd struct {
	util.CommonError

	TemplateID string `json:"priTmplId"`
}

// Add 添加订阅消息模板
func (s *Subscribe) Add(ShortID string, kidList []int, sceneDesc string) (templateID string, err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	var msg = struct {
		TemplateIDShort string `json:"tid"`
		SceneDesc       string `json:"sceneDesc"`
		KidList         []int  `json:"kidList"`
	}{TemplateIDShort: ShortID, SceneDesc: sceneDesc, KidList: kidList}
	uri := fmt.Sprintf("%s?access_token=%s", addTemplateURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	var result resSubscribeAdd
	err = util.DecodeWithError(response, &result, "AddSubscribe")
	if err != nil {
		return
	}
	templateID = result.TemplateID
	return
}

// Delete 删除私有模板
func (s *Subscribe) Delete(templateID string) (err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	var msg = struct {
		TemplateID string `json:"priTmplId"`
	}{TemplateID: templateID}
	uri := fmt.Sprintf("%s?access_token=%s", delTemplateURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "DeleteSubscribe")
}
