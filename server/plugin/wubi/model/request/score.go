package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// ScoreSearch 五笔成绩搜索 用于"我的成绩"分页
type ScoreSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Mode           string     `json:"mode" form:"mode"`
	request.PageInfo
}

// LeaderboardRequest 排行榜查询
type LeaderboardRequest struct {
	Mode  string `json:"mode" form:"mode"`               // char/word/article 留空则不过滤
	Limit int    `json:"limit" form:"limit"`             // 默认 20
	Days  int    `json:"days" form:"days"`               // 仅最近 N 天 0=全部
	Order string `json:"order" form:"order" example:"wpm"` // wpm/accuracy 默认 wpm
}

// SubmitScore 客户端提交成绩用 不让客户端自由设置 user/nickname
type SubmitScore struct {
	Mode           string  `json:"mode" form:"mode" binding:"required"`
	TotalCount     int     `json:"totalCount" form:"totalCount" binding:"required,min=1"`
	CorrectCount   int     `json:"correctCount" form:"correctCount" binding:"min=0"`
	ErrorCount     int     `json:"errorCount" form:"errorCount" binding:"min=0"`
	DurationSecond int     `json:"durationSecond" form:"durationSecond" binding:"required,min=1"`
	Accuracy       float64 `json:"accuracy" form:"accuracy" binding:"min=0,max=100"`
	WPM            float64 `json:"wpm" form:"wpm" binding:"min=0"`
}
