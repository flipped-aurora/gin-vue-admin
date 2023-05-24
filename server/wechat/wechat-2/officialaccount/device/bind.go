package device

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// ReqBind 设备绑定解绑共通实体
type ReqBind struct {
	Ticket   string `json:"ticket,omitempty"`
	DeviceID string `json:"device_id"`
	OpenID   string `json:"openid"`
}
type resBind struct {
	BaseResp util.CommonError `json:"base_resp"`
}

// Bind 设备绑定
func (d *Device) Bind(req ReqBind) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriBind, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}

// Unbind 设备解绑
func (d *Device) Unbind(req ReqBind) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriUnbind, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}

// CompelBind 强制绑定用户和设备
func (d *Device) CompelBind(req ReqBind) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriCompelBind, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}

// CompelUnbind 强制解绑用户和设备
func (d *Device) CompelUnbind(req ReqBind) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriCompelUnbind, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}
