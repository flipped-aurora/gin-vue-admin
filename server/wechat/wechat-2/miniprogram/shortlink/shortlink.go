package shortlink

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	generateShortLinkURL = "https://api.weixin.qq.com/wxa/genwxashortlink?access_token=%s"
)

// ShortLink 短链接
type ShortLink struct {
	*context.Context
}

// NewShortLink 实例
func NewShortLink(ctx *context.Context) *ShortLink {
	return &ShortLink{ctx}
}

// ShortLinker 请求结构体
type ShortLinker struct {

	// pageUrl 通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符
	PageURL string `json:"page_url"`

	// pageTitle 页面标题，不能包含违法信息，超过20字符会用... 截断代替
	PageTitle string `json:"page_title"`

	// isPermanent 生成的 Short Link 类型，短期有效：false，永久有效：true
	IsPermanent bool `json:"is_permanent,omitempty"`
}

// resShortLinker 返回结构体
type resShortLinker struct {
	// 通用错误
	util.CommonError

	// 返回的 shortLink
	Link string `json:"link"`
}

// Generate 生成 shortLink
func (shortLink *ShortLink) generate(shortLinkParams ShortLinker) (string, error) {
	var accessToken string
	accessToken, err := shortLink.GetAccessToken()
	if err != nil {
		return "", err
	}

	urlStr := fmt.Sprintf(generateShortLinkURL, accessToken)
	response, err := util.PostJSON(urlStr, shortLinkParams)
	if err != nil {
		return "", err
	}

	// 使用通用方法返回错误
	var res resShortLinker
	err = util.DecodeWithError(response, &res, "GenerateShortLink")
	if err != nil {
		return "", err
	}

	return res.Link, nil
}

// GenerateShortLinkPermanent 生成永久shortLink
func (shortLink *ShortLink) GenerateShortLinkPermanent(PageURL, pageTitle string) (string, error) {
	return shortLink.generate(ShortLinker{
		PageURL:     PageURL,
		PageTitle:   pageTitle,
		IsPermanent: true,
	})
}

// GenerateShortLinkTemp 生成临时shortLink
func (shortLink *ShortLink) GenerateShortLinkTemp(PageURL, pageTitle string) (string, error) {
	return shortLink.generate(ShortLinker{
		PageURL:     PageURL,
		PageTitle:   pageTitle,
		IsPermanent: false,
	})
}
