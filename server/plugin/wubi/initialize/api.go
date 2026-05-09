package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/wubi/submitScore", Description: "提交五笔成绩", ApiGroup: "五笔打字", Method: "POST"},
		{Path: "/wubi/deleteScore", Description: "删除五笔成绩", ApiGroup: "五笔打字", Method: "DELETE"},
		{Path: "/wubi/getMyScores", Description: "获取我的五笔成绩", ApiGroup: "五笔打字", Method: "GET"},
		{Path: "/wubi/getLeaderboard", Description: "获取五笔排行榜", ApiGroup: "五笔打字", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
