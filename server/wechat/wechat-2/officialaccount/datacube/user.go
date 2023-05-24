package datacube

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	getUserSummary    = "https://api.weixin.qq.com/datacube/getusersummary"
	getUserAccumulate = "https://api.weixin.qq.com/datacube/getusercumulate"
)

// ResUserSummary 获取用户增减数据响应
type ResUserSummary struct {
	util.CommonError

	List []struct {
		RefDate    string `json:"ref_date"`
		UserSource int    `json:"user_source"`
		NewUser    int    `json:"new_user"`
		CancelUser int    `json:"cancel_user"`
	} `json:"list"`
}

// ResUserAccumulate 获取累计用户数据响应
type ResUserAccumulate struct {
	util.CommonError

	List []struct {
		RefDate      string `json:"ref_date"`
		CumulateUser int    `json:"cumulate_user"`
	} `json:"list"`
}

// GetUserSummary 获取用户增减数据
func (cube *DataCube) GetUserSummary(s string, e string) (resUserSummary ResUserSummary, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUserSummary, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resUserSummary, "GetUserSummary")
	return
}

// GetUserAccumulate 获取累计用户数据
func (cube *DataCube) GetUserAccumulate(s string, e string) (resUserAccumulate ResUserAccumulate, err error) {
	accessToken, err := cube.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", getUserAccumulate, accessToken)
	reqDate := &reqDate{
		BeginDate: s,
		EndDate:   e,
	}

	response, err := util.PostJSON(uri, reqDate)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &resUserAccumulate, "GetUserAccumulate")
	return
}
