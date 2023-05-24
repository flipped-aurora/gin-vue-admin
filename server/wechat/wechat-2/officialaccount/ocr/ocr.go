package ocr

import (
	"fmt"
	"net/url"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/officialaccount/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	ocrIDCardURL         = "https://api.weixin.qq.com/cv/ocr/idcard"
	ocrBankCardURL       = "https://api.weixin.qq.com/cv/ocr/bankcard"
	ocrDrivingURL        = "https://api.weixin.qq.com/cv/ocr/driving"
	ocrDrivingLicenseURL = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	ocrBizLicenseURL     = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	ocrCommonURL         = "https://api.weixin.qq.com/cv/ocr/comm"
	ocrPlateNumberURL    = "https://api.weixin.qq.com/cv/ocr/platenum"
)

// OCR struct
type OCR struct {
	*context.Context
}

// coordinate 坐标
type coordinate struct {
	X int64 `json:"x,omitempty"`
	Y int64 `json:"y,omitempty"`
}

// position 位置
type position struct {
	LeftTop     coordinate `json:"left_top"`
	RightTop    coordinate `json:"right_top"`
	RightBottom coordinate `json:"right_bottom"`
	LeftBottom  coordinate `json:"left_bottom"`
}

// imageSize 图片尺寸
type imageSize struct {
	Width  int64 `json:"w,omitempty"`
	Height int64 `json:"h,omitempty"`
}

// ResDriving 行驶证返回结果
type ResDriving struct {
	util.CommonError

	PlateNumber       string              `json:"plate_num,omitempty"`
	VehicleType       string              `json:"vehicle_type,omitempty"`
	Owner             string              `json:"owner,omitempty"`
	Address           string              `json:"addr,omitempty"`
	UseCharacter      string              `json:"use_character,omitempty"`
	Model             string              `json:"model,omitempty"`
	Vin               string              `json:"vin,omitempty"`
	EngineNumber      string              `json:"engine_num,omitempty"`
	RegisterDate      string              `json:"register_date,omitempty"`
	IssueDate         string              `json:"issue_date,omitempty"`
	PlateNumberB      string              `json:"plate_num_b,omitempty"`
	Record            string              `json:"record,omitempty"`
	PassengersNumber  string              `json:"passengers_num,omitempty"`
	TotalQuality      string              `json:"total_quality,omitempty"`
	PrepareQuality    string              `json:"prepare_quality,omitempty"`
	OverallSize       string              `json:"overall_size,omitempty"`
	CardPositionFront map[string]position `json:"card_position_front,omitempty"`
	CardPositionBack  map[string]position `json:"card_position_back,omitempty"`
	ImageSize         imageSize           `json:"img_size,omitempty"`
}

// ResIDCard 身份证返回结果
type ResIDCard struct {
	util.CommonError

	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	Address     string `json:"addr,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	ValidDate   string `json:"valid_date,omitempty"`
}

// ResBankCard 银行卡返回结果
type ResBankCard struct {
	util.CommonError

	Number string `json:"number,omitempty"`
}

// ResDrivingLicense 驾驶证返回结果
type ResDrivingLicense struct {
	util.CommonError

	IDNumber     string `json:"id_num,omitempty"`
	Name         string `json:"name,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Nationality  string `json:"nationality,omitempty"`
	Address      string `json:"address,omitempty"`
	Birthday     string `json:"birth_date,omitempty"`
	IssueDate    string `json:"issue_date,omitempty"`
	CarClass     string `json:"car_class,omitempty"`
	ValidFrom    string `json:"valid_from,omitempty"`
	ValidTo      string `json:"valid_to,omitempty"`
	OfficialSeal string `json:"official_seal,omitempty"`
}

// ResBizLicense 营业执照返回结果
type ResBizLicense struct {
	util.CommonError

	RegisterNumber      string              `json:"reg_num,omitempty"`
	Serial              string              `json:"serial,omitempty"`
	LegalRepresentative string              `json:"legal_representative,omitempty"`
	EnterpriseName      string              `json:"enterprise_name,omitempty"`
	TypeOfOrganization  string              `json:"type_of_organization,omitempty"`
	Address             string              `json:"address,omitempty"`
	TypeOfEnterprise    string              `json:"type_of_enterprise,omitempty"`
	BusinessScope       string              `json:"business_scope,omitempty"`
	RegisteredCapital   string              `json:"registered_capital,omitempty"`
	PaidInCapital       string              `json:"paid_in_capital,omitempty"`
	ValidPeriod         string              `json:"valid_period,omitempty"`
	RegisterDate        string              `json:"registered_date,omitempty"`
	CertPosition        map[string]position `json:"cert_position,omitempty"`
	ImageSize           imageSize           `json:"img_size,omitempty"`
}

// ResCommon 公共印刷品返回结果
type ResCommon struct {
	util.CommonError

	Items     []commonItem `json:"items,omitempty"`
	ImageSize imageSize    `json:"img_size,omitempty"`
}

// commonItem 公共元素
type commonItem struct {
	Position position `json:"pos"`
	Text     string   `json:"text"`
}

// ResPlateNumber 车牌号返回结果
type ResPlateNumber struct {
	util.CommonError

	Number string `json:"number"`
}

// NewOCR 实例
func NewOCR(c *context.Context) *OCR {
	ocr := new(OCR)
	ocr.Context = c
	return ocr
}

// IDCard 身份证OCR识别接口
func (ocr *OCR) IDCard(path string) (resIDCard ResIDCard, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrIDCardURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resIDCard, "OCRIDCard")

	return
}

// BankCard 银行卡OCR识别接口
func (ocr *OCR) BankCard(path string) (resBankCard ResBankCard, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrBankCardURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resBankCard, "OCRBankCard")

	return
}

// Driving 行驶证OCR识别接口
func (ocr *OCR) Driving(path string) (resDriving ResDriving, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrDrivingURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resDriving, "OCRDriving")

	return
}

// DrivingLicense 驾驶证OCR识别接口
func (ocr *OCR) DrivingLicense(path string) (resDrivingLicense ResDrivingLicense, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrDrivingLicenseURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resDrivingLicense, "OCRDrivingLicense")

	return
}

// BizLicense 营业执照OCR识别接口
func (ocr *OCR) BizLicense(path string) (resBizLicense ResBizLicense, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrBizLicenseURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resBizLicense, "OCRBizLicense")

	return
}

// Common 通用印刷体OCR识别接口
func (ocr *OCR) Common(path string) (resCommon ResCommon, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrCommonURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resCommon, "OCRCommon")

	return
}

// PlateNumber 车牌OCR识别接口
func (ocr *OCR) PlateNumber(path string) (resPlateNumber ResPlateNumber, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.HTTPPost(fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrPlateNumberURL, url.QueryEscape(path), accessToken), "")
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resPlateNumber, "OCRPlateNumber")

	return
}
