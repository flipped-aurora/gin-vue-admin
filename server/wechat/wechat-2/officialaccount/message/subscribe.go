package message

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	subscribeSendURL                      = "https://api.weixin.qq.com/cgi-bin/message/subscribe/bizsend"
	subscribeTemplateListURL              = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"
	subscribeTemplateAddURL               = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate"
	subscribeTemplateDelURL               = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate"
	subscribeTemplateGetCategoryURL       = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory"
	subscribeTemplateGetPubTplKeyWorksURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords"
	subscribeTemplateGetPubTplTitles      = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles"
)

// Subscribe 订阅消息
type Subscribe struct {
	*context.Context
}

// NewSubscribe 实例化
func NewSubscribe(context *context.Context) *Subscribe {
	tpl := new(Subscribe)
	tpl.Context = context
	return tpl
}

// SubscribeMessage 发送的订阅消息内容
type SubscribeMessage struct {
	ToUser      string                        `json:"touser"`         // 必须, 接受者OpenID
	TemplateID  string                        `json:"template_id"`    // 必须, 模版ID
	Page        string                        `json:"page,omitempty"` // 可选, 跳转网页时填写
	Data        map[string]*SubscribeDataItem `json:"data"`           // 必须, 模板数据
	MiniProgram struct {
		AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` // 可选,跳转至小程序地址
}

// SubscribeDataItem 模版内某个 .DATA 的值
type SubscribeDataItem struct {
	Value string `json:"value"`
}

// Send 发送订阅消息
func (tpl *Subscribe) Send(msg *SubscribeMessage) (err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "SendSubscribeMessage")
}

// PrivateSubscribeItem 私有订阅消息模板
type PrivateSubscribeItem struct {
	PriTmplID string `json:"priTmplId"` //	添加至帐号下的模板 id，发送订阅通知时所需
	Title     string `json:"title"`     // 模版标题
	Content   string `json:"content"`   // 模版内容
	Example   string `json:"example"`   // 模板内容示例
	SubType   int    `json:"type"`      // 模版类型，2 为一次性订阅，3 为长期订阅
}

type resPrivateSubscribeList struct {
	util.CommonError
	SubscriptionList []*PrivateSubscribeItem `json:"data"`
}

// List 获取私有订阅消息模板列表
func (tpl *Subscribe) List() (templateList []*PrivateSubscribeItem, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeTemplateListURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resPrivateSubscribeList
	err = util.DecodeWithError(response, &res, "ListSubscribe")
	if err != nil {
		return
	}
	templateList = res.SubscriptionList
	return
}

type resSubscribeAdd struct {
	util.CommonError

	TemplateID string `json:"priTmplId"`
}

// Add 添加订阅消息模板
func (tpl *Subscribe) Add(ShortID string, kidList []int, sceneDesc string) (templateID string, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	var msg = struct {
		TemplateIDShort string `json:"tid"`
		SceneDesc       string `json:"sceneDesc"`
		KidList         []int  `json:"kidList"`
	}{TemplateIDShort: ShortID, SceneDesc: sceneDesc, KidList: kidList}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeTemplateAddURL, accessToken)
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
func (tpl *Subscribe) Delete(templateID string) (err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	var msg = struct {
		TemplateID string `json:"priTmplId"`
	}{TemplateID: templateID}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeTemplateDelURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "DeleteSubscribe")
}

// PublicTemplateCategory 公众号类目
type PublicTemplateCategory struct {
	ID   int    `json:"id"`   // 类目ID
	Name string `json:"name"` // 类目的中文名
}

type resSubscribeCategoryList struct {
	util.CommonError
	CategoryList []*PublicTemplateCategory `json:"data"`
}

// GetCategory 获取公众号类目
func (tpl *Subscribe) GetCategory() (categoryList []*PublicTemplateCategory, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeTemplateGetCategoryURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result resSubscribeCategoryList
	err = util.DecodeWithError(response, &result, "GetCategory")
	if err != nil {
		return
	}
	categoryList = result.CategoryList
	return
}

// PublicTemplateKeyWords 模板中的关键词
type PublicTemplateKeyWords struct {
	KeyWordsID int    `json:"kid"`     // 关键词 id
	Name       string `json:"name"`    // 关键词内容
	Example    string `json:"example"` // 关键词内容对应的示例
	Rule       string `json:"rule"`    // 参数类型
}

type resPublicTemplateKeyWordsList struct {
	util.CommonError
	KeyWordsList []*PublicTemplateKeyWords `json:"data"` // 关键词列表
}

// GetPubTplKeyWordsByID 获取模板中的关键词
func (tpl *Subscribe) GetPubTplKeyWordsByID(titleID string) (keyWordsList []*PublicTemplateKeyWords, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&tid=%s", subscribeTemplateGetPubTplKeyWorksURL, accessToken, titleID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result resPublicTemplateKeyWordsList
	err = util.DecodeWithError(response, &result, "GetPublicTemplateKeyWords")
	if err != nil {
		return
	}
	keyWordsList = result.KeyWordsList
	return
}

// PublicTemplateTitle 类目下的公共模板
type PublicTemplateTitle struct {
	TitleID    int    `json:"tid"`        // 模版标题 id
	Title      string `json:"title"`      // 模版标题
	Type       int    `json:"type"`       // 模版类型，2 为一次性订阅，3 为长期订阅
	CategoryID string `json:"categoryId"` // 模版所属类目 id
}

type resPublicTemplateTitleList struct {
	util.CommonError
	Count             int                    `json:"count"` // 公共模板列表总数
	TemplateTitleList []*PublicTemplateTitle `json:"data"`  // 模板标题列表
}

// GetPublicTemplateTitleList 获取类目下的公共模板
func (tpl *Subscribe) GetPublicTemplateTitleList(ids string, start int, limit int) (count int, templateTitleList []*PublicTemplateTitle, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&ids=%s&start=%d&limit=%d", subscribeTemplateGetPubTplTitles, accessToken, ids, start, limit)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result resPublicTemplateTitleList
	err = util.DecodeWithError(response, &result, "GetPublicTemplateTitle")
	if err != nil {
		return
	}
	count = result.Count
	templateTitleList = result.TemplateTitleList
	return
}
