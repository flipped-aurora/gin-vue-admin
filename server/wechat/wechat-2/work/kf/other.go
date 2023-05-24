package kf

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 获取视频号绑定状态
	corpQualification = "https://qyapi.weixin.qq.com/cgi-bin/kf/get_corp_qualification?access_token=%s"
)

// CorpQualificationSchema 获取视频号绑定状态响应内容
type CorpQualificationSchema struct {
	util.CommonError
	WechatChannelsBinding bool `json:"wechat_channels_binding"` // 	当企业具有绑定成功的视频号时，返回true，否则返回false。 1. 企业申请绑定视频号且由视频号管理员确认后，才为绑定成功状态 2. 至少有一个绑定成功的视频号就会返回true
}

// GetCorpQualification 获取视频号绑定状态
// 微信客服可接待的客户数，和企业是否已完成主体验证、是否绑定视频号相关。
//
// 企业未完成主体验证时，微信客服仅可累计接待100位客户
// 企业已验证但未绑定视频号时，微信客服仅可累计接待10000位客户
// 企业已验证且已绑定视频号时，微信客服可接待的客户数不受限制
//
// 开发者可获取状态后，在应用等地方提示企业去完成主体验证或绑定视频号。
func (r *Client) GetCorpQualification() (info CorpQualificationSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.HTTPGet(fmt.Sprintf(corpQualification, accessToken)); err != nil {
		return info, err
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
