package material

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	addNewsURL          = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	updateNewsURL       = "https://api.weixin.qq.com/cgi-bin/material/update_news"
	addMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	delMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	getMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/get_material"
	getMaterialCountURL = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount"
	batchGetMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/batchget_material"
)

// PermanentMaterialType 永久素材类型
type PermanentMaterialType string

const (
	// PermanentMaterialTypeImage 永久素材图片类型（image）
	PermanentMaterialTypeImage PermanentMaterialType = "image"
	// PermanentMaterialTypeVideo 永久素材视频类型（video）
	PermanentMaterialTypeVideo PermanentMaterialType = "video"
	// PermanentMaterialTypeVoice 永久素材语音类型 （voice）
	PermanentMaterialTypeVoice PermanentMaterialType = "voice"
	// PermanentMaterialTypeNews 永久素材图文类型（news）
	PermanentMaterialTypeNews PermanentMaterialType = "news"
)

// Material 素材管理
type Material struct {
	*context.Context
}

// NewMaterial init
func NewMaterial(context *context.Context) *Material {
	material := new(Material)
	material.Context = context
	return material
}

// Article 永久图文素材
type Article struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ThumbURL         string `json:"thumb_url"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
	URL              string `json:"url"`
	DownURL          string `json:"down_url"`
}

// GetNews 获取/下载永久素材
func (material *Material) GetNews(id string) ([]*Article, error) {
	accessToken, err := material.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", getMaterialURL, accessToken)

	var req struct {
		MediaID string `json:"media_id"`
	}
	req.MediaID = id
	responseBytes, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	var res struct {
		NewsItem []*Article `json:"news_item"`
	}
	err = json.Unmarshal(responseBytes, &res)
	if err != nil {
		return nil, err
	}

	return res.NewsItem, nil
}

// reqArticles 永久性图文素材请求信息
type reqArticles struct {
	Articles []*Article `json:"articles"`
}

// resArticles 永久性图文素材返回结果
type resArticles struct {
	util.CommonError

	MediaID string `json:"media_id"`
}

// AddNews 新增永久图文素材
func (material *Material) AddNews(articles []*Article) (mediaID string, err error) {
	req := &reqArticles{articles}

	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", addNewsURL, accessToken)
	responseBytes, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var res resArticles
	err = json.Unmarshal(responseBytes, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		return "", fmt.Errorf("errcode=%d,errmsg=%s", res.ErrCode, res.ErrMsg)
	}
	mediaID = res.MediaID
	return
}

// reqUpdateArticle 更新永久性图文素材请求信息
type reqUpdateArticle struct {
	MediaID  string   `json:"media_id"`
	Index    int64    `json:"index"`
	Articles *Article `json:"articles"`
}

// UpdateNews 更新永久图文素材
func (material *Material) UpdateNews(article *Article, mediaID string, index int64) (err error) {
	req := &reqUpdateArticle{mediaID, index, article}

	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", updateNewsURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "UpdateNews")
}

// resAddMaterial 永久性素材上传返回的结果
type resAddMaterial struct {
	util.CommonError

	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// AddMaterial 上传永久性素材（处理视频需要单独上传）
func (material *Material) AddMaterial(mediaType MediaType, filename string) (mediaID string, url string, err error) {
	if mediaType == MediaTypeVideo {
		err = errors.New("永久视频素材上传使用 AddVideo 方法")
		return
	}
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=%s", addMaterialURL, accessToken, mediaType)
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	mediaID = resMaterial.MediaID
	url = resMaterial.URL
	return
}

type reqVideo struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

// AddVideo 永久视频素材文件上传
func (material *Material) AddVideo(filename, title, introduction string) (mediaID string, url string, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=video", addMaterialURL, accessToken)

	videoDesc := &reqVideo{
		Title:        title,
		Introduction: introduction,
	}
	var fieldValue []byte
	fieldValue, err = json.Marshal(videoDesc)
	if err != nil {
		return
	}

	fields := []util.MultipartFormField{
		{
			IsFile:    true,
			Fieldname: "media",
			Filename:  filename,
		},
		{
			IsFile:    false,
			Fieldname: "description",
			Value:     fieldValue,
		},
	}

	var response []byte
	response, err = util.PostMultipartForm(fields, uri)
	if err != nil {
		return
	}

	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	mediaID = resMaterial.MediaID
	url = resMaterial.URL
	return
}

type reqDeleteMaterial struct {
	MediaID string `json:"media_id"`
}

// DeleteMaterial 删除永久素材
func (material *Material) DeleteMaterial(mediaID string) error {
	accessToken, err := material.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", delMaterialURL, accessToken)
	response, err := util.PostJSON(uri, reqDeleteMaterial{mediaID})
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteMaterial")
}

// ArticleList 永久素材列表
type ArticleList struct {
	util.CommonError
	TotalCount int64             `json:"total_count"`
	ItemCount  int64             `json:"item_count"`
	Item       []ArticleListItem `json:"item"`
}

// ArticleListItem 用于ArticleList的item节点
type ArticleListItem struct {
	MediaID    string             `json:"media_id"`
	Content    ArticleListContent `json:"content"`
	Name       string             `json:"name"`
	URL        string             `json:"url"`
	UpdateTime int64              `json:"update_time"`
}

// ArticleListContent 用于ArticleListItem的content节点
type ArticleListContent struct {
	NewsItem   []Article `json:"news_item"`
	UpdateTime int64     `json:"update_time"`
	CreateTime int64     `json:"create_time"`
}

// reqBatchGetMaterial BatchGetMaterial请求参数
type reqBatchGetMaterial struct {
	Type   PermanentMaterialType `json:"type"`
	Count  int64                 `json:"count"`
	Offset int64                 `json:"offset"`
}

// BatchGetMaterial 批量获取永久素材
//
//reference:https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (material *Material) BatchGetMaterial(permanentMaterialType PermanentMaterialType, offset, count int64) (list ArticleList, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", batchGetMaterialURL, accessToken)

	req := reqBatchGetMaterial{
		Type:   permanentMaterialType,
		Offset: offset,
		Count:  count,
	}

	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &list, "BatchGetMaterial")
	return
}

// ResMaterialCount 素材总数
type ResMaterialCount struct {
	util.CommonError
	VoiceCount int64 `json:"voice_count"` // 语音总数量
	VideoCount int64 `json:"video_count"` // 视频总数量
	ImageCount int64 `json:"image_count"` // 图片总数量
	NewsCount  int64 `json:"news_count"`  // 图文总数量
}

// GetMaterialCount 获取素材总数.
func (material *Material) GetMaterialCount() (res ResMaterialCount, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getMaterialCountURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &res, "GetMaterialCount")
	return
}
