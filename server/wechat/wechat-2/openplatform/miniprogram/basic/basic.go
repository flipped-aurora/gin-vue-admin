package basic

import (
	"fmt"

	openContext "github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/openplatform/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	getAccountBasicInfoURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"
)

// Basic 基础信息设置
type Basic struct {
	*openContext.Context
	appID string
}

// NewBasic new
func NewBasic(opContext *openContext.Context, appID string) *Basic {
	return &Basic{Context: opContext, appID: appID}
}

// AccountBasicInfo 基础信息
type AccountBasicInfo struct {
	util.CommonError
}

// GetAccountBasicInfo 获取小程序基础信息
//
//reference:https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html
func (basic *Basic) GetAccountBasicInfo() (*AccountBasicInfo, error) {
	ak, err := basic.GetAuthrAccessToken(basic.AppID)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getAccountBasicInfoURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	result := &AccountBasicInfo{}
	if err := util.DecodeWithError(data, result, "account/getaccountbasicinfo"); err != nil {
		return nil, err
	}
	return result, nil
}

// modify_domain设置服务器域名
// TODO
// func (encryptor *Basic) modifyDomain() {
// }
