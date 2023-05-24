package component

import (
	"fmt"

	openContext "github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/openplatform/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	fastregisterweappURL = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp"
)

// Component 快速创建小程序
type Component struct {
	*openContext.Context
}

// NewComponent new
func NewComponent(opContext *openContext.Context) *Component {
	return &Component{opContext}
}

// RegisterMiniProgramParam 快速注册小程序参数
type RegisterMiniProgramParam struct {
	Name               string `json:"name"`                 // 企业名
	Code               string `json:"code"`                 // 企业代码
	CodeType           string `json:"code_type"`            // 企业代码类型 1：统一社会信用代码（18 位） 2：组织机构代码（9 位 xxxxxxxx-x） 3：营业执照注册号(15 位)
	LegalPersonaWechat string `json:"legal_persona_wechat"` // 法人微信号
	LegalPersonaName   string `json:"legal_persona_name"`   // 法人姓名（绑定银行卡）
	ComponentPhone     string `json:"component_phone"`      // 第三方联系电话（方便法人与第三方联系）
}

// RegisterMiniProgram 快速创建小程
// reference: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Fast_Registration_Interface_document.html
func (component *Component) RegisterMiniProgram(param *RegisterMiniProgramParam) error {
	componentAK, err := component.GetComponentAccessToken()
	if err != nil {
		return nil
	}
	url := fmt.Sprintf(fastregisterweappURL+"?action=create&component_access_token=%s", componentAK)
	data, err := util.PostJSON(url, param)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "component/fastregisterweapp?action=create")
}

// GetRegistrationStatusParam 查询任务创建状态
type GetRegistrationStatusParam struct {
	Name               string `json:"name"`                 // 企业名
	LegalPersonaWechat string `json:"legal_persona_wechat"` // 法人微信号
	LegalPersonaName   string `json:"legal_persona_name"`   // 法人姓名（绑定银行卡）

}

// GetRegistrationStatus 查询创建任务状态.
func (component *Component) GetRegistrationStatus(param *GetRegistrationStatusParam) error {
	componentAK, err := component.GetComponentAccessToken()
	if err != nil {
		return nil
	}
	url := fmt.Sprintf(fastregisterweappURL+"?action=search&component_access_token=%s", componentAK)
	data, err := util.PostJSON(url, param)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "component/fastregisterweapp?action=search")
}
