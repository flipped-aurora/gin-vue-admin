package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Score = new(score)

type score struct{}

// SubmitScore 提交一次五笔练习成绩
// @Tags     Wubi
// @Summary  提交五笔练习成绩
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.SubmitScore true "成绩数据"
// @Success  200  {object} response.Response{msg=string} "提交成功"
// @Router   /wubi/submitScore [post]
func (a *score) SubmitScore(c *gin.Context) {
	var in request.SubmitScore
	if err := c.ShouldBindJSON(&in); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	nick := utils.GetUserInfo(c).NickName
	rec := &model.Score{
		UserID:         uid,
		NickName:       nick,
		Mode:           in.Mode,
		TotalCount:     in.TotalCount,
		CorrectCount:   in.CorrectCount,
		ErrorCount:     in.ErrorCount,
		Accuracy:       in.Accuracy,
		DurationSecond: in.DurationSecond,
		WPM:            in.WPM,
	}
	if err := serviceScore.CreateScore(rec); err != nil {
		global.GVA_LOG.Error("成绩入库失败!", zap.Error(err))
		response.FailWithMessage("提交失败", c)
		return
	}
	response.OkWithDetailed(rec, "提交成功", c)
}

// DeleteScore 删除自己的成绩
// @Tags     Wubi
// @Summary  删除自己的一条成绩
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    ID query string true "成绩ID"
// @Success  200 {object} response.Response{msg=string} "删除成功"
// @Router   /wubi/deleteScore [delete]
func (a *score) DeleteScore(c *gin.Context) {
	id := c.Query("ID")
	if id == "" {
		response.FailWithMessage("ID必填", c)
		return
	}
	if err := serviceScore.DeleteScore(id, utils.GetUserID(c)); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetMyScores 分页获取我的成绩
// @Tags     Wubi
// @Summary  分页获取当前用户的成绩
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query request.ScoreSearch true "分页查询参数"
// @Success  200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /wubi/getMyScores [get]
func (a *score) GetMyScores(c *gin.Context) {
	var q request.ScoreSearch
	if err := c.ShouldBindQuery(&q); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceScore.GetMyScores(utils.GetUserID(c), q)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     q.Page,
		PageSize: q.PageSize,
	}, "获取成功", c)
}

// GetLeaderboard 获取五笔打字排行榜
// @Tags     Wubi
// @Summary  获取五笔打字排行榜(按WPM降序,取每用户最高记录)
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query request.LeaderboardRequest true "排行参数"
// @Success  200 {object} response.Response{data=[]model.Score,msg=string} "获取成功"
// @Router   /wubi/getLeaderboard [get]
func (a *score) GetLeaderboard(c *gin.Context) {
	var q request.LeaderboardRequest
	if err := c.ShouldBindQuery(&q); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := serviceScore.GetLeaderboard(q)
	if err != nil {
		global.GVA_LOG.Error("排行榜查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(list, c)
}
