package analysis

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 获取用户访问小程序日留存
	getAnalysisDailyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=%s"
	// 获取用户访问小程序月留存
	getAnalysisMonthlyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=%s"
	// 获取用户访问小程序周留存
	getAnalysisWeeklyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=%s"
	// 获取用户访问小程序数据概况
	getAnalysisDailySummaryURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=%s"
	// 获取用户访问小程序数据日趋势
	getAnalysisDailyVisitTrendURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=%s"
	// 获取用户访问小程序数据月趋势
	getAnalysisMonthlyVisitTrendURL = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=%s"
	// 获取用户访问小程序数据周趋势
	getAnalysisWeeklyVisitTrendURL = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=%s"
	// 获取小程序新增或活跃用户的画像分布数据
	getAnalysisUserPortraitURL = "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=%s"
	// 获取用户小程序访问分布数据
	getAnalysisVisitDistributionURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=%s"
	// 访问页面
	getAnalysisVisitPageURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=%s"
)

// Analysis analyis 数据分析
type Analysis struct {
	*context.Context
}

// NewAnalysis new
func NewAnalysis(ctx *context.Context) *Analysis {
	return &Analysis{ctx}
}

// fetchData 拉取统计数据
func (analysis *Analysis) fetchData(urlStr string, body interface{}) (response []byte, err error) {
	var accessToken string
	accessToken, err = analysis.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)
	response, err = util.PostJSON(urlStr, body)
	return
}

// RetainItem 留存项结构
type RetainItem struct {
	Key   int `json:"key"`   // 标识，0开始表示当天，1表示1甜后，以此类推
	Value int `json:"value"` // key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// ResAnalysisRetain 小程序留存数据返回
type ResAnalysisRetain struct {
	util.CommonError
	RefDate    string       `json:"ref_date"`     // 日期
	VisitUVNew []RetainItem `json:"visit_uv_new"` // 新增用户留存
	VisitUV    []RetainItem `json:"visit_uv"`     // 活跃用户留存
}

// getAnalysisRetain 获取用户访问小程序留存数据(日、月、周)
func (analysis *Analysis) getAnalysisRetain(urlStr string, beginDate, endDate string) (result ResAnalysisRetain, err error) {
	body := map[string]string{
		"begin_date": beginDate,
		"end_date":   endDate,
	}
	response, err := analysis.fetchData(urlStr, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("getAnalysisRetain error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// GetAnalysisDailyRetain 获取用户访问小程序日留存
func (analysis *Analysis) GetAnalysisDailyRetain(beginDate, endDate string) (result ResAnalysisRetain, err error) {
	return analysis.getAnalysisRetain(getAnalysisDailyRetainURL, beginDate, endDate)
}

// GetAnalysisMonthlyRetain 获取用户访问小程序月留存
func (analysis *Analysis) GetAnalysisMonthlyRetain(beginDate, endDate string) (result ResAnalysisRetain, err error) {
	return analysis.getAnalysisRetain(getAnalysisMonthlyRetainURL, beginDate, endDate)
}

// GetAnalysisWeeklyRetain 获取用户访问小程序周留存
func (analysis *Analysis) GetAnalysisWeeklyRetain(beginDate, endDate string) (result ResAnalysisRetain, err error) {
	return analysis.getAnalysisRetain(getAnalysisWeeklyRetainURL, beginDate, endDate)
}

// ResAnalysisDailySummary 小程序访问数据概况
type ResAnalysisDailySummary struct {
	util.CommonError
	List []struct {
		RefDate    string `json:"ref_date"`    // 日期
		VisitTotal int    `json:"visit_total"` // 累计用户数
		SharePV    int    `json:"share_pv"`    // 转发次数
		ShareUV    int    `json:"share_uv"`    // 转发人数
	} `json:"list"`
}

// GetAnalysisDailySummary 获取用户访问小程序数据概况
func (analysis *Analysis) GetAnalysisDailySummary(beginDate, endDate string) (result ResAnalysisDailySummary, err error) {
	body := map[string]string{
		"begin_date": beginDate,
		"end_date":   endDate,
	}
	response, err := analysis.fetchData(getAnalysisDailySummaryURL, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetAnalysisDailySummary error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// ResAnalysisVisitTrend 小程序访问数据趋势(日、月、周)
type ResAnalysisVisitTrend struct {
	util.CommonError
	List []struct {
		RefDate         string  `json:"ref_date"`          // 日期
		SessionCnt      int     `json:"session_cnt"`       // 打开次数
		VisitPV         int     `json:"visit_pv"`          // 访问次数
		VisitUV         int     `json:"visit_uv"`          // 访问人数
		VisitUVNew      int     `json:"visit_uv_new"`      // 新用户数
		StayTimeUV      float64 `json:"stay_time_uv"`      // 人均停留时长
		StayTimeSession float64 `json:"stay_time_session"` // 次均停留时常
		VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度
	} `json:"list"`
}

// getAnalysisRetain 获取小程序访问数据趋势(日、月、周)
func (analysis *Analysis) getAnalysisVisitTrend(urlStr string, beginDate, endDate string) (result ResAnalysisVisitTrend, err error) {
	body := map[string]string{
		"begin_date": beginDate,
		"end_date":   endDate,
	}
	response, err := analysis.fetchData(urlStr, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("getAnalysisVisitTrend error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// GetAnalysisDailyVisitTrend 获取用户访问小程序数据日趋势
func (analysis *Analysis) GetAnalysisDailyVisitTrend(beginDate, endDate string) (result ResAnalysisVisitTrend, err error) {
	return analysis.getAnalysisVisitTrend(getAnalysisDailyVisitTrendURL, beginDate, endDate)
}

// GetAnalysisMonthlyVisitTrend 获取用户访问小程序数据月趋势
func (analysis *Analysis) GetAnalysisMonthlyVisitTrend(beginDate, endDate string) (result ResAnalysisVisitTrend, err error) {
	return analysis.getAnalysisVisitTrend(getAnalysisMonthlyVisitTrendURL, beginDate, endDate)
}

// GetAnalysisWeeklyVisitTrend 获取用户访问小程序数据周趋势
func (analysis *Analysis) GetAnalysisWeeklyVisitTrend(beginDate, endDate string) (result ResAnalysisVisitTrend, err error) {
	return analysis.getAnalysisVisitTrend(getAnalysisWeeklyVisitTrendURL, beginDate, endDate)
}

// UserPortraitItem 用户画像项目
type UserPortraitItem struct {
	ID    int    `json:"id"`    // 属性值id
	Name  string `json:"name"`  // 属性值名称
	Value int    `json:"value"` // 该场景访问uv
}

// UserPortrait 用户画像
type UserPortrait struct {
	Index     int                `json:"index"`     // 分布类型
	Province  []UserPortraitItem `json:"province"`  // 省份，如北京、广东等
	City      []UserPortraitItem `json:"city"`      // 城市，如北京、广州等
	Genders   []UserPortraitItem `json:"genders"`   // 性别，包括男、女、未知
	Platforms []UserPortraitItem `json:"platforms"` // 终端类型，包括iPhone, android, 其他
	Devices   []UserPortraitItem `json:"devices"`   // 机型，如苹果iPhone 6, OPPO R9等
	Ages      []UserPortraitItem `json:"ages"`      // 年龄，包括17岁以下、18-24对等区间
}

// ResAnalysisUserPortrait 小程序新增或活跃用户的画像分布数据返回
type ResAnalysisUserPortrait struct {
	util.CommonError
	RefDate    string       `json:"ref_date"`     // 日期
	VisitUVNew UserPortrait `json:"visit_uv_new"` // 新用户画像
	VisitUV    UserPortrait `json:"visit_uv"`     // 活跃用户画像
}

// GetAnalysisUserPortrait 获取小程序新增或活跃用户的画像分布数据
func (analysis *Analysis) GetAnalysisUserPortrait(beginDate, endDate string) (result ResAnalysisUserPortrait, err error) {
	body := map[string]string{
		"begin_date": beginDate,
		"end_date":   endDate,
	}
	response, err := analysis.fetchData(getAnalysisUserPortraitURL, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetAnalysisUserPortrait error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// VisitDistributionIndexItem 访问分数数据结构
type VisitDistributionIndexItem struct {
	Key                 int `json:"key"`                    // 场景id
	Value               int `json:"value"`                  // 该场景id访问pv
	AccessSourceVisitUV int `json:"access_source_visit_uv"` // 该场景id访问uv
}

// VisitDistributionIndex 访问分布单分布类型数据
type VisitDistributionIndex struct {
	Index    string                       `json:"index"`     // 分布类型
	ItemList []VisitDistributionIndexItem `json:"item_list"` // 分布数据列表
}

// ResAnalysisVisitDistribution 小程序访问分布数据返回
type ResAnalysisVisitDistribution struct {
	util.CommonError
	RefDate string                   `json:"ref_date"` // 日期
	List    []VisitDistributionIndex `json:"list"`     // 数据列表
}

// GetAnalysisVisitDistribution 获取用户小程序访问分布数据
func (analysis *Analysis) GetAnalysisVisitDistribution(beginDate, endDate string) (result ResAnalysisVisitDistribution, err error) {
	body := map[string]string{
		"begin_date": beginDate,
		"end_date":   endDate,
	}
	response, err := analysis.fetchData(getAnalysisVisitDistributionURL, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetAnalysisVisitDistribution error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

// VisitPageItem 访问单个页面的数据结构
type VisitPageItem struct {
	PagePath       string  `json:"page_path"`        // 页面路径
	PageVisitPV    int     `json:"page_visit_pv"`    // 访问次数
	PageVisitUV    int     `json:"page_visit_uv"`    // 访问人数
	PageStaytimePV float64 `json:"page_staytime_pv"` // 次均停留时常
	EntrypagePV    int     `json:"entrypage_pv"`     // 进入页次数
	ExitpagePV     int     `json:"exitpage_pv"`      // 退出页次数
	PageSharePV    int     `json:"page_share_pv"`    // 转发次数
	PageShareUV    int     `json:"page_share_uv"`    // 转发人数
}

// ResAnalysisVisitPage 访问小程序页面访问数据返回
type ResAnalysisVisitPage struct {
	util.CommonError
	RefDate string          `json:"ref_date"` // 日期
	List    []VisitPageItem `json:"list"`     // 数据列表
}

// GetAnalysisVisitPage 获取小程序页面访问数据
func (analysis *Analysis) GetAnalysisVisitPage(beginDate, endDate string) (result ResAnalysisVisitPage, err error) {
	body := map[string]string{
		"begin_date": beginDate,
		"end_date":   endDate,
	}
	response, err := analysis.fetchData(getAnalysisVisitPageURL, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetAnalysisVisitPage error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
