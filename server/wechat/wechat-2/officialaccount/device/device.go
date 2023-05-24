package device

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	uriAuthorize    = "https://api.weixin.qq.com/device/authorize_device"
	uriQRCode       = "https://api.weixin.qq.com/device/create_qrcode"
	uriVerifyQRCode = "https://api.weixin.qq.com/device/verify_qrcode"
	uriBind         = "https://api.weixin.qq.com/device/bind"
	uriUnbind       = "https://api.weixin.qq.com/device/unbind"
	uriCompelBind   = "https://api.weixin.qq.com/device/compel_bind"
	uriCompelUnbind = "https://api.weixin.qq.com/device/compel_unbind"
	uriState        = "https://api.weixin.qq.com/device/get_stat"
)

// Device struct
type Device struct {
	*context.Context
}

// NewDevice 实例
func NewDevice(context *context.Context) *Device {
	device := new(Device)
	device.Context = context
	return device
}

// ResDeviceState 设备状态响应实体
type ResDeviceState struct {
	util.CommonError
	Status     int    `json:"status"`
	StatusInfo string `json:"status_info"`
}

// State 设备状态查询
func (d *Device) State(device string) (res ResDeviceState, err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&device_id=%s", uriState, accessToken, device)
	var response []byte
	if response, err = util.HTTPGet(uri); err != nil {
		return
	}
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeviceState Error , errcode=%d , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
