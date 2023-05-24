package freepublish

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	publishURL      = "https://api.weixin.qq.com/cgi-bin/freepublish/submit"     // 发布接口
	selectStateURL  = "https://api.weixin.qq.com/cgi-bin/freepublish/get"        // 发布状态轮询接口
	deleteURL       = "https://api.weixin.qq.com/cgi-bin/freepublish/delete"     // 删除发布
	firstArticleURL = "https://api.weixin.qq.com/cgi-bin/freepublish/getarticle" // 通过 article_id 获取已发布文章
	paginateURL     = "https://api.weixin.qq.com/cgi-bin/freepublish/batchget"   // 获取成功发布列表
)

// PublishStatus 发布状态
type PublishStatus uint

const (
	// PublishStatusSuccess 0:成功
	PublishStatusSuccess PublishStatus = iota
	// PublishStatusPublishing 1:发布中
	PublishStatusPublishing
	// PublishStatusOriginalFail 2:原创失败
	PublishStatusOriginalFail
	// PublishStatusFail 3:常规失败
	PublishStatusFail
	// PublishStatusAuditRefused 4:平台审核不通过
	PublishStatusAuditRefused
	// PublishStatusUserDeleted 5:成功后用户删除所有文章
	PublishStatusUserDeleted
	// PublishStatusSystemBanned 6:成功后系统封禁所有文章
	PublishStatusSystemBanned
)

// FreePublish 发布能力
type FreePublish struct {
	*context.Context
}

// NewFreePublish init
func NewFreePublish(ctx *context.Context) *FreePublish {
	return &FreePublish{
		Context: ctx,
	}
}

// Publish 发布接口。需要先将图文素材以草稿的形式保存（见“草稿箱/新建草稿”，
// 如需从已保存的草稿中选择，见“草稿箱/获取草稿列表”），选择要发布的草稿 media_id 进行发布
func (freePublish *FreePublish) Publish(mediaID string) (publishID int64, err error) {
	var accessToken string
	accessToken, err = freePublish.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		MediaID string `json:"media_id"`
	}
	req.MediaID = mediaID

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", publishURL, accessToken)
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		PublishID int64 `json:"publish_id"`
	}
	err = util.DecodeWithError(response, &res, "SubmitFreePublish")
	if err != nil {
		return
	}

	publishID = res.PublishID
	return
}

// PublishStatusList 发布任务状态列表
type PublishStatusList struct {
	util.CommonError
	PublishID     int64                `json:"publish_id"`     // 发布任务id
	PublishStatus PublishStatus        `json:"publish_status"` // 发布状态
	ArticleID     string               `json:"article_id"`     // 当发布状态为0时（即成功）时，返回图文的 article_id，可用于“客服消息”场景
	ArticleDetail PublishArticleDetail `json:"article_detail"` // 发布任务文章成功状态详情
	FailIndex     []uint               `json:"fail_idx"`       // 当发布状态为2或4时，返回不通过的文章编号，第一篇为 1；其他发布状态则为空
}

// PublishArticleDetail 发布任务成功详情
type PublishArticleDetail struct {
	Count uint                 `json:"count"` // 当发布状态为0时（即成功）时，返回文章数量
	Items []PublishArticleItem `json:"item"`
}

// PublishArticleItem 发布任务成功的文章内容
type PublishArticleItem struct {
	Index      uint   `json:"idx"`         // 当发布状态为0时（即成功）时，返回文章对应的编号
	ArticleURL string `json:"article_url"` // 当发布状态为0时（即成功）时，返回图文的永久链接
}

// SelectStatus 发布状态轮询接口
func (freePublish *FreePublish) SelectStatus(publishID int64) (list PublishStatusList, err error) {
	accessToken, err := freePublish.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		PublishID int64 `json:"publish_id"`
	}
	req.PublishID = publishID

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", selectStateURL, accessToken)
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &list, "SelectStatusFreePublish")
	return
}

// Delete 删除发布。
// index 要删除的文章在图文消息中的位置，第一篇编号为1，该字段不填或填0会删除全部文章
// !!!此操作不可逆，请谨慎操作!!!删除后微信公众号后台仍然会有记录!!!
func (freePublish *FreePublish) Delete(articleID string, index uint) (err error) {
	accessToken, err := freePublish.GetAccessToken()
	if err != nil {
		return err
	}

	var req struct {
		ArticleID string `json:"article_id"`
		Index     uint   `json:"index"`
	}
	req.ArticleID = articleID
	req.Index = index

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", deleteURL, accessToken)
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteFreePublish")
}

// Article 图文信息内容
type Article struct {
	Title              string `json:"title"`                 // 标题
	Author             string `json:"author"`                // 作者
	Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
	Content            string `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
	ContentSourceURL   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（一定是永久MediaID）
	ShowCoverPic       uint   `json:"show_cover_pic"`        // 是否显示封面，0为false，即不显示，1为true，即显示(默认)
	NeedOpenComment    uint   `json:"need_open_comment"`     // 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment uint   `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
	URL                string `json:"url"`                   // 图文消息的URL
	IsDeleted          bool   `json:"is_deleted"`            // 该图文是否被删除
}

// First 通过 article_id 获取已发布文章
func (freePublish *FreePublish) First(articleID string) (list []Article, err error) {
	accessToken, err := freePublish.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		ArticleID string `json:"article_id"`
	}
	req.ArticleID = articleID

	var response []byte
	uri := fmt.Sprintf("%s?access_token=%s", firstArticleURL, accessToken)
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		NewsItem []Article `json:"news_item"`
	}
	err = util.DecodeWithError(response, &res, "FirstFreePublish")
	if err != nil {
		return
	}

	list = res.NewsItem
	return
}

// ArticleList 发布列表
type ArticleList struct {
	util.CommonError
	TotalCount int64             `json:"total_count"` // 成功发布素材的总数
	ItemCount  int64             `json:"item_count"`  // 本次调用获取的素材的数量
	Item       []ArticleListItem `json:"item"`
}

// ArticleListItem 用于 ArticleList 的 item 节点
type ArticleListItem struct {
	ArticleID  string             `json:"article_id"`  // 成功发布的图文消息id
	Content    ArticleListContent `json:"content"`     // 内容
	UpdateTime int64              `json:"update_time"` // 这篇图文消息素材的最后更新时间
}

// ArticleListContent 用于 ArticleListItem 的 content 节点
type ArticleListContent struct {
	NewsItem []Article `json:"news_item"` // 这篇图文消息素材的内容
}

// Paginate 获取成功发布列表
func (freePublish *FreePublish) Paginate(offset, count int64, noReturnContent bool) (list ArticleList, err error) {
	var accessToken string
	accessToken, err = freePublish.GetAccessToken()
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

	err = util.DecodeWithError(response, &list, "PaginateFreePublish")
	return
}
