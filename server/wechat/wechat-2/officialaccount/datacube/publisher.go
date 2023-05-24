package datacube

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

// AdSlot 广告位类型
type AdSlot string

const (
	// SlotIDBizBottom 公众号底部广告
	SlotIDBizBottom AdSlot = "SLOT_ID_BIZ_BOTTOM"
	// SlotIDBizMidContext 公众号文中广告
	SlotIDBizMidContext AdSlot = "SLOT_ID_BIZ_MID_CONTEXT"
	// SlotIDBizVideoEnd 公众号视频后贴
	SlotIDBizVideoEnd AdSlot = "SLOT_ID_BIZ_VIDEO_END"
	// SlotIDBizSponsor 公众号互选广告
	SlotIDBizSponsor AdSlot = "SLOT_ID_BIZ_SPONSOR"
	// SlotIDBizCps 公众号返佣商品
	SlotIDBizCps AdSlot = "SLOT_ID_BIZ_CPS"
	// SlotIDWeappBanner 小程序banner
	SlotIDWeappBanner AdSlot = "SLOT_ID_WEAPP_BANNER"
	// SlotIDWeappRewardVideo 小程序激励视频
	SlotIDWeappRewardVideo AdSlot = "SLOT_ID_WEAPP_REWARD_VIDEO"
	// SlotIDWeappInterstitial 小程序插屏广告
	SlotIDWeappInterstitial AdSlot = "SLOT_ID_WEAPP_INTERSTITIAL"
	// SlotIDWeappVideoFeeds 小程序视频广告
	SlotIDWeappVideoFeeds AdSlot = "SLOT_ID_WEAPP_VIDEO_FEEDS"
	// SlotIDWeappVideoBegin 小程序视频前贴
	SlotIDWeappVideoBegin AdSlot = "SLOT_ID_WEAPP_VIDEO_BEGIN"
	// SlotIDWeappBox 小程序格子广告
	SlotIDWeappBox AdSlot = "SLOT_ID_WEAPP_BOX"
)

const (
	publisherURL = "https://api.weixin.qq.com/publisher/stat"
)

const (
	actionPublisherAdPosGeneral = "publisher_adpos_general"
	actionPublisherCpsGeneral   = "publisher_cps_general"
	actionPublisherSettlement   = "publisher_settlement"
)

// BaseResp 错误信息
type BaseResp struct {
	ErrMsg string `json:"err_msg"`
	Ret    int    `json:"ret"`
}

// ResPublisherAdPos 公众号分广告位数据响应
type ResPublisherAdPos struct {
	util.CommonError

	BaseResp BaseResp        `json:"base_resp"`
	List     []ResAdPosList  `json:"list"`
	Summary  ResAdPosSummary `json:"summary"`
	TotalNum int             `json:"total_num"`
}

// ResAdPosList 公众号分广告位列表
type ResAdPosList struct {
	SlotID        int64   `json:"slot_id"`
	AdSlot        string  `json:"ad_slot"`
	Date          string  `json:"date"`
	ReqSuccCount  int     `json:"req_succ_count"`
	ExposureCount int     `json:"exposure_count"`
	ExposureRate  float64 `json:"exposure_rate"`
	ClickCount    int     `json:"click_count"`
	ClickRate     float64 `json:"click_rate"`
	Income        int     `json:"income"`
	Ecpm          float64 `json:"ecpm"`
}

// ResAdPosSummary 公众号分广告位概览
type ResAdPosSummary struct {
	ReqSuccCount  int     `json:"req_succ_count"`
	ExposureCount int     `json:"exposure_count"`
	ExposureRate  float64 `json:"exposure_rate"`
	ClickCount    int     `json:"click_count"`
	ClickRate     float64 `json:"click_rate"`
	Income        int     `json:"income"`
	Ecpm          float64 `json:"ecpm"`
}

// ResPublisherCps 公众号返佣商品数据响应
type ResPublisherCps struct {
	util.CommonError

	BaseResp BaseResp      `json:"base_resp"`
	List     []ResCpsList  `json:"list"`
	Summary  ResCpsSummary `json:"summary"`
	TotalNum int           `json:"total_num"`
}

// ResCpsList 公众号返佣商品列表
type ResCpsList struct {
	Date            string  `json:"date"`
	ExposureCount   int     `json:"exposure_count"`
	ClickCount      int     `json:"click_count"`
	ClickRate       float64 `json:"click_rate"`
	OrderCount      int     `json:"order_count"`
	OrderRate       float64 `json:"order_rate"`
	TotalFee        int     `json:"total_fee"`
	TotalCommission int     `json:"total_commission"`
}

// ResCpsSummary 公众号返佣概览
type ResCpsSummary struct {
	ExposureCount   int     `json:"exposure_count"`
	ClickCount      int     `json:"click_count"`
	ClickRate       float64 `json:"click_rate"`
	OrderCount      int     `json:"order_count"`
	OrderRate       float64 `json:"order_rate"`
	TotalFee        int     `json:"total_fee"`
	TotalCommission int     `json:"total_commission"`
}

// ResPublisherSettlement 公众号结算收入数据及结算主体信息响应
type ResPublisherSettlement struct {
	util.CommonError

	BaseResp          BaseResp         `json:"base_resp"`
	Body              string           `json:"body"`
	PenaltyAll        int              `json:"penalty_all"`
	RevenueAll        int64            `json:"revenue_all"`
	SettledRevenueAll int64            `json:"settled_revenue_all"`
	SettlementList    []SettlementList `json:"settlement_list"`
	TotalNum          int              `json:"total_num"`
}

// SettlementList 结算单列表
type SettlementList struct {
	Date           string        `json:"date"`
	Zone           string        `json:"zone"`
	Month          string        `json:"month"`
	Order          int           `json:"order"`
	SettStatus     int           `json:"sett_status"`
	SettledRevenue int           `json:"settled_revenue"`
	SettNo         string        `json:"sett_no"`
	MailSendCnt    string        `json:"mail_send_cnt"`
	SlotRevenue    []SlotRevenue `json:"slot_revenue"`
}

// SlotRevenue 产生收入的广告
type SlotRevenue struct {
	SlotID             string `json:"slot_id"`
	SlotSettledRevenue int    `json:"slot_settled_revenue"`
}

// ParamsPublisher 拉取数据参数
type ParamsPublisher struct {
	Action    string `json:"action"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	AdSlot    AdSlot `json:"ad_slot"`
}

// fetchData 拉取统计数据
func (cube *DataCube) fetchData(params ParamsPublisher) (response []byte, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("action", params.Action)
	v.Add("access_token", accessToken)
	v.Add("page", strconv.Itoa(params.Page))
	v.Add("page_size", strconv.Itoa(params.PageSize))
	v.Add("start_date", params.StartDate)
	v.Add("end_date", params.EndDate)
	if params.AdSlot != "" {
		v.Add("ad_slot", string(params.AdSlot))
	}

	uri := fmt.Sprintf("%s?%s", publisherURL, v.Encode())

	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	return
}

// GetPublisherAdPosGeneral 获取公众号分广告位数据
func (cube *DataCube) GetPublisherAdPosGeneral(startDate, endDate string, page, pageSize int, adSlot AdSlot) (resPublisherAdPos ResPublisherAdPos, err error) {
	params := ParamsPublisher{
		Action:    actionPublisherAdPosGeneral,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  pageSize,
		AdSlot:    adSlot,
	}

	response, err := cube.fetchData(params)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resPublisherAdPos, "GetPublisherAdPosGeneral")
	if err != nil {
		return
	}

	if resPublisherAdPos.BaseResp.Ret != 0 {
		err = fmt.Errorf("GetPublisherAdPosGeneral Error , errcode=%d , errmsg=%s", resPublisherAdPos.BaseResp.Ret, resPublisherAdPos.BaseResp.ErrMsg)
		return
	}
	return
}

// GetPublisherCpsGeneral 获取公众号返佣商品数据
func (cube *DataCube) GetPublisherCpsGeneral(startDate, endDate string, page, pageSize int) (resPublisherCps ResPublisherCps, err error) {
	params := ParamsPublisher{
		Action:    actionPublisherCpsGeneral,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  pageSize,
	}

	response, err := cube.fetchData(params)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resPublisherCps, "GetPublisherCpsGeneral")
	if err != nil {
		return
	}

	if resPublisherCps.BaseResp.Ret != 0 {
		err = fmt.Errorf("GetPublisherCpsGeneral Error , errcode=%d , errmsg=%s", resPublisherCps.BaseResp.Ret, resPublisherCps.BaseResp.ErrMsg)
		return
	}
	return
}

// GetPublisherSettlement 获取公众号结算收入数据及结算主体信息
func (cube *DataCube) GetPublisherSettlement(startDate, endDate string, page, pageSize int) (resPublisherSettlement ResPublisherSettlement, err error) {
	params := ParamsPublisher{
		Action:    actionPublisherSettlement,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  pageSize,
	}

	response, err := cube.fetchData(params)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resPublisherSettlement, "GetPublisherSettlement")
	if err != nil {
		return
	}

	if resPublisherSettlement.BaseResp.Ret != 0 {
		err = fmt.Errorf("GetPublisherSettlement Error , errcode=%d , errmsg=%s", resPublisherSettlement.BaseResp.Ret, resPublisherSettlement.BaseResp.ErrMsg)
		return
	}
	return
}
