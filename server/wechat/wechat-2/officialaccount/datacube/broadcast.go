package datacube

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	getArticleSummary = "https://api.weixin.qq.com/datacube/getarticlesummary"
	getArticleTotal   = "https://api.weixin.qq.com/datacube/getarticletotal"
	getUserRead       = "https://api.weixin.qq.com/datacube/getuserread"
	getUserReadHour   = "https://api.weixin.qq.com/datacube/getuserreadhour"
	getUserShare      = "https://api.weixin.qq.com/datacube/getusershare"
	getUserShareHour  = "https://api.weixin.qq.com/datacube/getusersharehour"
)

// ResArticleSummary 获取图文群发每日数据响应
type ResArticleSummary struct {
	util.CommonError

	List []struct {
		RefDate          string `json:"ref_date"`
		MsgID            string `json:"msgid"`
		Title            string `json:"title"`
		IntPageReadUser  int    `json:"int_page_read_user"`
		IntPageReadCount int    `json:"int_page_read_count"`
		OriPageReadUser  int    `json:"ori_page_read_user"`
		OriPageReadCount int    `json:"ori_page_read_count"`
		ShareUser        int    `json:"share_user"`
		ShareCount       int    `json:"share_count"`
		AddToFavUser     int    `json:"add_to_fav_user"`
		AddToFavCount    int    `json:"add_to_fav_count"`
	} `json:"list"`
}

// ResArticleTotal 获取图文群发总数据响应
type ResArticleTotal struct {
	util.CommonError

	List []struct {
		RefDate string                `json:"ref_date"`
		MsgID   string                `json:"msgid"`
		Title   string                `json:"title"`
		Details []ArticleTotalDetails `json:"details"`
	} `json:"list"`
}

// ArticleTotalDetails 获取图文群发总数据响应文字详情
type ArticleTotalDetails struct {
	StatDate                    string `json:"stat_date"`
	TargetUser                  int    `json:"target_user"`
	IntPageReadUser             int    `json:"int_page_read_user"`
	IntPageReadCount            int    `json:"int_page_read_count"`
	OriPageReadUser             int    `json:"ori_page_read_user"`
	OriPageReadCount            int    `json:"ori_page_read_count"`
	ShareUser                   int    `json:"share_user"`
	ShareCount                  int    `json:"share_count"`
	AddToFavUser                int    `json:"add_to_fav_user"`
	AddToFavCount               int    `json:"add_to_fav_count"`
	IntPageFromSessionReadUser  int    `json:"int_page_from_session_read_user"`
	IntPageFromSessionReadCount int    `json:"int_page_from_session_read_count"`
	IntPageFromHistMsgReadUser  int    `json:"int_page_from_hist_msg_read_user"`
	IntPageFromHistMsgReadCount int    `json:"int_page_from_hist_msg_read_count"`
	IntPageFromFeedReadUser     int    `json:"int_page_from_feed_read_user"`
	IntPageFromFeedReadCount    int    `json:"int_page_from_feed_read_count"`
	IntPageFromFriendsReadUser  int    `json:"int_page_from_friends_read_user"`
	IntPageFromFriendsReadCount int    `json:"int_page_from_friends_read_count"`
	IntPageFromOtherReadUser    int    `json:"int_page_from_other_read_user"`
	IntPageFromOtherReadCount   int    `json:"int_page_from_other_read_count"`
	FeedShareFromSessionUser    int    `json:"feed_share_from_session_user"`
	FeedShareFromSessionCnt     int    `json:"feed_share_from_session_cnt"`
	FeedShareFromFeedUser       int    `json:"feed_share_from_feed_user"`
	FeedShareFromFeedCnt        int    `json:"feed_share_from_feed_cnt"`
	FeedShareFromOtherUser      int    `json:"feed_share_from_other_user"`
	FeedShareFromOtherCnt       int    `json:"feed_share_from_other_cnt"`
}

// ResUserRead 获取图文统计数据响应
type ResUserRead struct {
	util.CommonError

	List []struct {
		RefDate          string `json:"ref_date"`
		UserSource       int    `json:"user_source"`
		IntPageReadUser  int    `json:"int_page_read_user"`
		IntPageReadCount int    `json:"int_page_read_count"`
		OriPageReadUser  int    `json:"ori_page_read_user"`
		OriPageReadCount int    `json:"ori_page_read_count"`
		ShareUser        int    `json:"share_user"`
		ShareCount       int    `json:"share_count"`
		AddToFavUser     int    `json:"add_to_fav_user"`
		AddToFavCount    int    `json:"add_to_fav_count"`
	} `json:"list"`
}

// ResUserReadHour 获取图文统计分时数据
type ResUserReadHour struct {
	util.CommonError

	List []struct {
		RefDate          string `json:"ref_date"`
		RefHour          int    `json:"ref_hour"`
		UserSource       int    `json:"user_source"`
		IntPageReadUser  int    `json:"int_page_read_user"`
		IntPageReadCount int    `json:"int_page_read_count"`
		OriPageReadUser  int    `json:"ori_page_read_user"`
		OriPageReadCount int    `json:"ori_page_read_count"`
		ShareUser        int    `json:"share_user"`
		ShareCount       int    `json:"share_count"`
		AddToFavUser     int    `json:"add_to_fav_user"`
		AddToFavCount    int    `json:"add_to_fav_count"`
	} `json:"list"`
}

// ResUserShare 获取图文分享转发数据
type ResUserShare struct {
	util.CommonError

	List []struct {
		RefDate    string `json:"ref_date"`
		ShareScene int    `json:"share_scene"`
		ShareCount int    `json:"share_count"`
		ShareUser  int    `json:"share_user"`
	} `json:"list"`
}

// ResUserShareHour  获取图文分享转发分时数据
type ResUserShareHour struct {
	util.CommonError

	List []struct {
		RefDate    string `json:"ref_date"`
		RefHour    int    `json:"ref_hour"`
		ShareScene int    `json:"share_scene"`
		ShareCount int    `json:"share_count"`
		ShareUser  int    `json:"share_user"`
	} `json:"list"`
}

// GetArticleSummary 获取图文群发每日数据
func (cube *DataCube) GetArticleSummary(s string, e string) (resArticleSummary ResArticleSummary, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getArticleSummary, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resArticleSummary, "GetArticleSummary")
	return
}

// GetArticleTotal 获取图文群发总数据
func (cube *DataCube) GetArticleTotal(s string, e string) (resArticleTotal ResArticleTotal, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getArticleTotal, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resArticleTotal, "GetArticleTotal")
	return
}

// GetUserRead 获取图文统计数据
func (cube *DataCube) GetUserRead(s string, e string) (resUserRead ResUserRead, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUserRead, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resUserRead, "GetUserRead")
	return
}

// GetUserReadHour 获取图文统计分时数据
func (cube *DataCube) GetUserReadHour(s string, e string) (resUserReadHour ResUserReadHour, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUserReadHour, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resUserReadHour, "GetUserReadHour")
	return
}

// GetUserShare 获取图文分享转发数据
func (cube *DataCube) GetUserShare(s string, e string) (resUserShare ResUserShare, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUserShare, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resUserShare, "GetUserShare")
	return
}

// GetUserShareHour 获取图文分享转发分时数据
func (cube *DataCube) GetUserShareHour(s string, e string) (resUserShareHour ResUserShareHour, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUserShareHour, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resUserShareHour, "GetUserShareHour")
	return
}
