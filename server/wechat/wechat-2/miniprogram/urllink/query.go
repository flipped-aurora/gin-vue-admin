package urllink

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const queryURL = "https://api.weixin.qq.com/wxa/query_urllink"

// ULQueryResult 返回的结果
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.query.html 返回值
type ULQueryResult struct {
	util.CommonError

	URLLinkInfo struct {
		Appid      string `json:"appid"`
		Path       string `json:"path"`
		Query      string `json:"query"`
		CreateTime int64  `json:"create_time"`
		ExpireTime int64  `json:"expire_time"`
		EnvVersion string `json:"env_version"`
		CloudBase  struct {
			Env           string `json:"env"`
			Domain        string `json:"domain"`
			Path          string `json:"path"`
			Query         string `json:"query"`
			ResourceAppid string `json:"resource_appid"`
		} `json:"cloud_base"`
	} `json:"url_link_info"`
	VisitOpenid string `json:"visit_openid"`
}

// Query 查询小程序 url_link 配置。
func (u *URLLink) Query(urlLink string) (*ULQueryResult, error) {
	accessToken, err := u.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", queryURL, accessToken)
	response, err := util.PostJSON(uri, map[string]string{"url_link": urlLink})
	if err != nil {
		return nil, err
	}
	var resp ULQueryResult
	err = util.DecodeWithError(response, &resp, "URLLink.Query")
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
