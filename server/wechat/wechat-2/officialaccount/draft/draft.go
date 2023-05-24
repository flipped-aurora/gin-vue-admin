package draft

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	addURL      = "https://api.weixin.qq.com/cgi-bin/draft/add"      // 新建草稿
	getURL      = "https://api.weixin.qq.com/cgi-bin/draft/get"      // 获取草稿
	deleteURL   = "https://api.weixin.qq.com/cgi-bin/draft/delete"   // 删除草稿
	updateURL   = "https://api.weixin.qq.com/cgi-bin/draft/update"   // 修改草稿
	countURL    = "https://api.weixin.qq.com/cgi-bin/draft/count"    // 获取草稿总数
	paginateURL = "https://api.weixin.qq.com/cgi-bin/draft/batchget" // 获取草稿列表
)

// Draft 草稿箱
type Draft struct {
	*context.Context
}

// NewDraft init
func NewDraft(ctx *context.Context) *Draft {
	return &Draft{
		Context: ctx,
	}
}

// Article 草稿
type Article struct {
	Title              string `json:"title"`                 // 标题
	Author             string `json:"author"`                // 作者
	Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。
	Content            string `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且去除JS
	ContentSourceURL   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久MediaID）
	ShowCoverPic       uint   `json:"show_cover_pic"`        // 是否显示封面，0为false，即不显示，1为true，即显示(默认)
	NeedOpenComment    uint   `json:"need_open_comment"`     // 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment uint   `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
}

// AddDraft 新建草稿
func (draft *Draft) AddDraft(articles []*Article) (mediaID string, err error) {
	accessToken, err := draft.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		Articles []*Article `json:"articles"`
	}
	req.Articles = articles

	uri := fmt.Sprintf("%s?access_token=%s", addURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		MediaID string `json:"media_id"`
	}
	err = util.DecodeWithError(response, &res, "AddDraft")
	if err != nil {
		return
	}
	mediaID = res.MediaID
	return
}

// GetDraft 获取草稿
func (draft *Draft) GetDraft(mediaID string) (articles []*Article, err error) {
	accessToken, err := draft.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		MediaID string `json:"media_id"`
	}
	req.MediaID = mediaID

	uri := fmt.Sprintf("%s?access_token=%s", getURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		NewsItem []*Article `json:"news_item"`
	}
	err = util.DecodeWithError(response, &res, "GetDraft")
	if err != nil {
		return
	}

	articles = res.NewsItem
	return
}

// DeleteDraft 删除草稿
func (draft *Draft) DeleteDraft(mediaID string) (err error) {
	accessToken, err := draft.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		MediaID string `json:"media_id"`
	}
	req.MediaID = mediaID

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", deleteURL, accessToken)
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "DeleteDraft")
	return
}

// UpdateDraft 修改草稿
// index 要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
func (draft *Draft) UpdateDraft(article *Article, mediaID string, index uint) (err error) {
	accessToken, err := draft.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		MediaID string   `json:"media_id"`
		Index   uint     `json:"index"`
		Article *Article `json:"articles"`
	}
	req.MediaID = mediaID
	req.Index = index
	req.Article = article

	uri := fmt.Sprintf("%s?access_token=%s", updateURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "UpdateDraft")
	return
}

// CountDraft 获取草稿总数
func (draft *Draft) CountDraft() (total uint, err error) {
	accessToken, err := draft.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", countURL, accessToken)
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		Total uint `json:"total_count"`
	}
	err = util.DecodeWithError(response, &res, "CountDraft")
	if nil != err {
		return
	}

	total = res.Total
	return
}

// ArticleList 草稿列表
type ArticleList struct {
	util.CommonError
	TotalCount int64             `json:"total_count"` // 草稿素材的总数
	ItemCount  int64             `json:"item_count"`  // 本次调用获取的素材的数量
	Item       []ArticleListItem `json:"item"`
}

// ArticleListItem 用于 ArticleList 的 item 节点
type ArticleListItem struct {
	MediaID    string             `json:"media_id"`    // 图文消息的id
	Content    ArticleListContent `json:"content"`     // 内容
	UpdateTime int64              `json:"update_time"` // 这篇图文消息素材的最后更新时间
}

// ArticleListContent 用于 ArticleListItem 的 content 节点
type ArticleListContent struct {
	NewsItem []Article `json:"news_item"` // 这篇图文消息素材的内容
}

// PaginateDraft 获取草稿列表
func (draft *Draft) PaginateDraft(offset, count int64, noReturnContent bool) (list ArticleList, err error) {
	accessToken, err := draft.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		Count           int64 `json:"count"`
		Offset          int64 `json:"offset"`
		NoReturnContent bool  `json:"no_content"`
	}
	req.Count = count
	req.Offset = offset
	req.NoReturnContent = noReturnContent

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", paginateURL, accessToken)
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &list, "PaginateDraft")
	return
}
