package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	qrCreateURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	getQRImgURL = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)

const (
	actionID  = "QR_SCENE"
	actionStr = "QR_STR_SCENE"

	actionLimitID  = "QR_LIMIT_SCENE"
	actionLimitStr = "QR_LIMIT_STR_SCENE"
)

// Request 临时二维码
type Request struct {
	ExpireSeconds int64  `json:"expire_seconds,omitempty"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneStr string `json:"scene_str,omitempty"`
			SceneID  int    `json:"scene_id,omitempty"`
		} `json:"scene"`
	} `json:"action_info"`
}

// Ticket 二维码ticket
type Ticket struct {
	util.CommonError `json:",inline"`
	Ticket           string `json:"ticket"`
	ExpireSeconds    int64  `json:"expire_seconds"`
	URL              string `json:"url"`
}

// GetQRTicket 获取二维码 Ticket
func (basic *Basic) GetQRTicket(tq *Request) (t *Ticket, err error) {
	accessToken, err := basic.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(qrCreateURL, accessToken)
	response, err := util.PostJSON(uri, tq)
	if err != nil {
		err = fmt.Errorf("get qr ticket failed, %s", err)
		return
	}

	t = new(Ticket)
	err = json.Unmarshal(response, &t)
	if err != nil {
		return
	}

	if t.ErrMsg != "" {
		err = fmt.Errorf("get qr_ticket error : errcode=%v , errormsg=%v", t.ErrCode, t.ErrMsg)
		return
	}

	return
}

// ShowQRCode 通过ticket换取二维码
func ShowQRCode(tk *Ticket) string {
	return fmt.Sprintf(getQRImgURL, tk.Ticket)
}

// NewTmpQrRequest 新建临时二维码请求实例
func NewTmpQrRequest(exp time.Duration, scene interface{}) *Request {
	tq := &Request{
		ExpireSeconds: int64(exp.Seconds()),
	}
	switch reflect.ValueOf(scene).Kind() {
	case reflect.String:
		tq.ActionName = actionStr
		tq.ActionInfo.Scene.SceneStr = scene.(string)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		tq.ActionName = actionID
		tq.ActionInfo.Scene.SceneID = scene.(int)
	}

	return tq
}

// NewLimitQrRequest 新建永久二维码请求实例
func NewLimitQrRequest(scene interface{}) *Request {
	tq := &Request{}
	switch reflect.ValueOf(scene).Kind() {
	case reflect.String:
		tq.ActionName = actionLimitStr
		tq.ActionInfo.Scene.SceneStr = scene.(string)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		tq.ActionName = actionLimitID
		tq.ActionInfo.Scene.SceneID = scene.(int)
	}

	return tq
}
