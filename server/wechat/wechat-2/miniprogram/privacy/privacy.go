package privacy

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// Privacy 小程序授权隐私设置
type Privacy struct {
	*context.Context
}

// NewPrivacy 实例化小程序隐私接口
// 文档地址 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html
func NewPrivacy(context *context.Context) *Privacy {
	if context == nil {
		panic("NewPrivacy got a nil context")
	}
	return &Privacy{
		context,
	}
}

// OwnerSetting 收集方（开发者）信息配置
type OwnerSetting struct {
	ContactEmail         string `json:"contact_email"`
	ContactPhone         string `json:"contact_phone"`
	ContactQQ            string `json:"contact_qq"`
	ContactWeixin        string `json:"contact_weixin"`
	ExtFileMediaID       string `json:"ext_file_media_id"`
	NoticeMethod         string `json:"notice_method"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
}

// SettingItem 收集权限的配置
type SettingItem struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyText string `json:"privacy_text"`
}

// SetPrivacySettingRequest 设置权限的请求参数
type SetPrivacySettingRequest struct {
	PrivacyVer   int           `json:"privacy_ver"`
	OwnerSetting OwnerSetting  `json:"owner_setting"`
	SettingList  []SettingItem `json:"setting_list"`
}

const (
	setPrivacySettingURL    = "https://api.weixin.qq.com/cgi-bin/component/setprivacysetting"
	getPrivacySettingURL    = "https://api.weixin.qq.com/cgi-bin/component/getprivacysetting"
	uploadPrivacyExtFileURL = "https://api.weixin.qq.com/cgi-bin/component/uploadprivacyextfile"

	// PrivacyV1 用户隐私保护指引的版本，1表示现网版本。
	PrivacyV1 = 1
	// PrivacyV2 2表示开发版。默认是2开发版。
	PrivacyV2 = 2
)

// GetPrivacySettingResponse 获取权限配置的响应结果
type GetPrivacySettingResponse struct {
	util.CommonError
	CodeExist    int                   `json:"code_exist"`
	PrivacyList  []string              `json:"privacy_list"`
	SettingList  []SettingResponseItem `json:"setting_list"`
	UpdateTime   int64                 `json:"update_time"`
	OwnerSetting OwnerSetting          `json:"owner_setting"`
	PrivacyDesc  DescList              `json:"privacy_desc"`
}

// SettingResponseItem 获取权限设置的响应明细
type SettingResponseItem struct {
	PrivacyKey   string `json:"privacy_key"`
	PrivacyText  string `json:"privacy_text"`
	PrivacyLabel string `json:"privacy_label"`
}

// DescList 权限列表(保持与官方一致)
type DescList struct {
	PrivacyDescList []Desc `json:"privacy_desc_list"`
}

// Desc 权限列表明细(保持与官方一致)
type Desc struct {
	PrivacyDesc string `json:"privacy_desc"`
	PrivacyKey  string `json:"privacy_key"`
}

// GetPrivacySetting 获取小程序权限配置
func (s *Privacy) GetPrivacySetting(privacyVer int) (GetPrivacySettingResponse, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return GetPrivacySettingResponse{}, err
	}

	response, err := util.PostJSON(fmt.Sprintf("%s?access_token=%s", getPrivacySettingURL, accessToken), map[string]int{
		"privacy_ver": privacyVer,
	})
	if err != nil {
		return GetPrivacySettingResponse{}, err
	}
	// 返回错误信息
	var result GetPrivacySettingResponse
	if err = util.DecodeWithError(response, &result, "getprivacysetting"); err != nil {
		return GetPrivacySettingResponse{}, err
	}

	return result, nil
}

// SetPrivacySetting 更新小程序权限配置
func (s *Privacy) SetPrivacySetting(privacyVer int, ownerSetting OwnerSetting, settingList []SettingItem) error {
	if privacyVer == PrivacyV1 && len(settingList) > 0 {
		return errors.New("当privacy_ver传2或者不传时，setting_list是必填；当privacy_ver传1时，该参数不可传")
	}
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return err
	}

	response, err := util.PostJSON(fmt.Sprintf("%s?access_token=%s", setPrivacySettingURL, accessToken), SetPrivacySettingRequest{
		PrivacyVer:   privacyVer,
		OwnerSetting: ownerSetting,
		SettingList:  settingList,
	})
	if err != nil {
		return err
	}

	// 返回错误信息
	if err = util.DecodeWithCommonError(response, "setprivacysetting"); err != nil {
		return err
	}

	return err
}

// UploadPrivacyExtFileResponse 上传权限定义模板响应参数
type UploadPrivacyExtFileResponse struct {
	util.CommonError
	ExtFileMediaID string `json:"ext_file_media_id"`
}

// UploadPrivacyExtFile 上传权限定义模板
func (s *Privacy) UploadPrivacyExtFile(fileData []byte) (UploadPrivacyExtFileResponse, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return UploadPrivacyExtFileResponse{}, err
	}

	response, err := util.PostJSON(fmt.Sprintf("%s?access_token=%s", uploadPrivacyExtFileURL, accessToken), map[string][]byte{
		"file": fileData,
	})
	if err != nil {
		return UploadPrivacyExtFileResponse{}, err
	}

	// 返回错误信息
	var result UploadPrivacyExtFileResponse
	if err = util.DecodeWithError(response, &result, "setprivacysetting"); err != nil {
		return UploadPrivacyExtFileResponse{}, err
	}

	return result, err
}
