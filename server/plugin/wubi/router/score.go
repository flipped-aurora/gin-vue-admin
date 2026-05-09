package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Score = new(score)

type score struct{}

// Init 初始化五笔打字路由
// 写操作放进 OperationRecord 中间件; 查询放在普通分组
func (r *score) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("wubi").Use(middleware.OperationRecord())
		group.POST("submitScore", apiScore.SubmitScore)    // 提交成绩
		group.DELETE("deleteScore", apiScore.DeleteScore)  // 删除成绩
	}
	{
		group := private.Group("wubi")
		group.GET("getMyScores", apiScore.GetMyScores)       // 我的成绩
		group.GET("getLeaderboard", apiScore.GetLeaderboard) // 排行榜
	}
}
