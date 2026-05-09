package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Score 五笔打字成绩
type Score struct {
	global.GVA_MODEL
	UserID         uint    `json:"userID" form:"userID" gorm:"column:user_id;index;comment:练习用户ID;"`
	NickName       string  `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称(冗余,便于排行榜展示);"`
	Mode           string  `json:"mode" form:"mode" gorm:"column:mode;size:32;default:char;comment:练习模式 char/word/article;"`
	TotalCount     int     `json:"totalCount" form:"totalCount" gorm:"column:total_count;comment:本轮题目总数;"`
	CorrectCount   int     `json:"correctCount" form:"correctCount" gorm:"column:correct_count;comment:正确数;"`
	ErrorCount     int     `json:"errorCount" form:"errorCount" gorm:"column:error_count;comment:错误数;"`
	Accuracy       float64 `json:"accuracy" form:"accuracy" gorm:"column:accuracy;comment:准确率(0-100);"`
	DurationSecond int     `json:"durationSecond" form:"durationSecond" gorm:"column:duration_second;comment:耗时(秒);"`
	WPM            float64 `json:"wpm" form:"wpm" gorm:"column:wpm;comment:每分钟字数;"`
}

// TableName Score 自定义表名 gva_wubi_score
func (Score) TableName() string {
	return "gva_wubi_score"
}
