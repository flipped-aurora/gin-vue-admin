package security

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	mediaCheckAsyncURL = "https://api.weixin.qq.com/wxa/media_check_async?access_token=%s"
	imageCheckURL      = "https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s"
	msgCheckURL        = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
)

// Security 内容安全
type Security struct {
	*context.Context
}

// NewSecurity init
func NewSecurity(ctx *context.Context) *Security {
	return &Security{ctx}
}

// MediaCheckAsyncV1Request 图片/音频异步校验请求参数
type MediaCheckAsyncV1Request struct {
	MediaURL  string `json:"media_url"`  // 要检测的图片或音频的url，支持图片格式包括jpg, jepg, png, bmp, gif（取首帧），支持的音频格式包括mp3, aac, ac3, wma, flac, vorbis, opus, wav
	MediaType uint8  `json:"media_type"` // 1:音频;2:图片
}

// MediaCheckAsyncV1 异步校验图片/音频是否含有违法违规内容
// Deprecated
// 在2021年9月1日停止更新，请尽快更新至 2.0 接口。建议使用 MediaCheckAsync
func (security *Security) MediaCheckAsyncV1(in *MediaCheckAsyncV1Request) (traceID string, err error) {
	accessToken, err := security.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(mediaCheckAsyncURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	var res struct {
		util.CommonError
		TraceID string `json:"trace_id"`
	}
	err = util.DecodeWithError(response, &res, "MediaCheckAsyncV1")
	if err != nil {
		return
	}

	traceID = res.TraceID
	return
}

// MediaCheckAsyncRequest 图片/音频异步校验请求参数
type MediaCheckAsyncRequest struct {
	MediaURL  string `json:"media_url"`  // 要检测的图片或音频的url，支持图片格式包括jpg, jepg, png, bmp, gif（取首帧），支持的音频格式包括mp3, aac, ac3, wma, flac, vorbis, opus, wav
	MediaType uint8  `json:"media_type"` // 1:音频;2:图片
	OpenID    string `json:"openid"`     // 用户的openid（用户需在近两小时访问过小程序）
	Scene     uint8  `json:"scene"`      // 场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func (security *Security) MediaCheckAsync(in *MediaCheckAsyncRequest) (traceID string, err error) {
	accessToken, err := security.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		MediaCheckAsyncRequest
		Version uint `json:"version"` // 接口版本号，2.0版本为固定值2
	}
	req.MediaCheckAsyncRequest = *in
	req.Version = 2

	uri := fmt.Sprintf(mediaCheckAsyncURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	var res struct {
		util.CommonError
		TraceID string `json:"trace_id"`
	}
	err = util.DecodeWithError(response, &res, "MediaCheckAsync")
	if err != nil {
		return
	}

	traceID = res.TraceID
	return
}

// ImageCheckV1 校验一张图片是否含有违法违规内容（同步）
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.imgSecCheck.html
// Deprecated
// 在2021年9月1日停止更新。建议使用 MediaCheckAsync
func (security *Security) ImageCheckV1(filename string) (err error) {
	accessToken, err := security.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(imageCheckURL, accessToken)
	response, err := util.PostFile("media", filename, uri)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	return util.DecodeWithCommonError(response, "ImageCheckV1")
}

// CheckSuggest 检查建议
type CheckSuggest string

const (
	// CheckSuggestRisky 违规风险建议
	CheckSuggestRisky CheckSuggest = "risky"
	// CheckSuggestPass 安全
	CheckSuggestPass CheckSuggest = "pass"
	// CheckSuggestReview 需要审查
	CheckSuggestReview CheckSuggest = "review"
)

// MsgScene 文本场景
type MsgScene uint8

const (
	// MsgSceneMaterial 资料文件检查场景
	MsgSceneMaterial MsgScene = iota + 1
	// MsgSceneComment 评论
	MsgSceneComment
	// MsgSceneForum 论坛
	MsgSceneForum
	// MsgSceneSocialLog 社交日志
	MsgSceneSocialLog
)

// CheckLabel 检查命中标签
type CheckLabel int

func (cl CheckLabel) String() string {
	switch cl {
	case 100:
		return "正常"
	case 10001:
		return "广告"
	case 20001:
		return "时政"
	case 20002:
		return "色情"
	case 20003:
		return "辱骂"
	case 20006:
		return "违法犯罪"
	case 20008:
		return "欺诈"
	case 20012:
		return "低俗"
	case 20013:
		return "版权"
	case 21000:
		return "其他"
	default:
		return strconv.Itoa(int(cl))
	}
}

// MsgCheckRequest 文本检查请求
type MsgCheckRequest struct {
	OpenID    string   `json:"openid"`    // 用户的openid（用户需在近两小时访问过小程序）
	Scene     MsgScene `json:"scene"`     // 场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
	Content   string   `json:"content"`   // 需检测的文本内容，文本字数的上限为 2500 字，需使用 UTF-8 编码
	Nickname  string   `json:"nickname"`  // （非必填）用户昵称，需使用UTF-8编码
	Title     string   `json:"title"`     // （非必填）文本标题，需使用UTF-8编码
	Signature string   `json:"signature"` // （非必填）个性签名，该参数仅在资料类场景有效(scene=1)，需使用UTF-8编码
}

// MsgCheckResponse 文本检查响应
type MsgCheckResponse struct {
	util.CommonError
	TraceID string `json:"trace_id"` // 唯一请求标识
	Result  struct {
		Suggest CheckSuggest `json:"suggest"` // 建议
		Label   CheckLabel   `json:"label"`   // 命中标签
	} `json:"result"` // 综合结果
	Detail []struct {
		ErrCode  int64      `json:"errcode"`  // 错误码，仅当该值为0时，该项结果有效
		Strategy string     `json:"strategy"` // 策略类型
		Suggest  string     `json:"suggest"`  // 建议
		Label    CheckLabel `json:"label"`    // 命中标签
		Prob     uint       `json:"prob"`     // 置信度。0-100，越高代表越有可能属于当前返回的标签（label）
		Keyword  string     `json:"keyword"`  // 命中的自定义关键词
	} `json:"detail"` // 详细检测结果
}

// MsgCheckV1 检查一段文本是否含有违法违规内容
// Deprecated
// 在2021年9月1日停止更新，请尽快更新至 2.0 接口。建议使用 MsgCheck
func (security *Security) MsgCheckV1(content string) (res MsgCheckResponse, err error) {
	accessToken, err := security.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		Content string `json:"content"`
	}
	req.Content = content

	uri := fmt.Sprintf(msgCheckURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "security.MsgCheckV1")
	return
}

// MsgCheck 检查一段文本是否含有违法违规内容
func (security *Security) MsgCheck(in *MsgCheckRequest) (res MsgCheckResponse, err error) {
	accessToken, err := security.GetAccessToken()
	if err != nil {
		return
	}

	var req struct {
		MsgCheckRequest
		Version uint `json:"version"`
	}
	req.MsgCheckRequest = *in
	req.Version = 2

	uri := fmt.Sprintf(msgCheckURL, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "security.MsgCheck")
	return
}
