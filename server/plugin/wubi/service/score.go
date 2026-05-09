package service

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/model/request"
)

var Score = new(score)

type score struct{}

// CreateScore 写入一条练习成绩
func (s *score) CreateScore(rec *model.Score) error {
	return global.GVA_DB.Create(rec).Error
}

// DeleteScore 删除自己的一条成绩 仅本人可删
func (s *score) DeleteScore(id string, userID uint) error {
	return global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Score{}).Error
}

// GetMyScores 查询当前用户的成绩 分页
func (s *score) GetMyScores(userID uint, info request.ScoreSearch) (list []model.Score, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&model.Score{}).Where("user_id = ?", userID)
	if info.Mode != "" {
		db = db.Where("mode = ?", info.Mode)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at DESC").Find(&list).Error
	return
}

// GetLeaderboard 排行榜 按 WPM 降序 取每个用户的最高成绩
func (s *score) GetLeaderboard(req request.LeaderboardRequest) (list []model.Score, err error) {
	limit := req.Limit
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	// 子查询：每个用户在指定模式/时间窗口内的最高 WPM 对应的记录 ID
	sub := global.GVA_DB.Model(&model.Score{}).
		Select("user_id, MAX(wpm) AS max_wpm")
	if req.Mode != "" {
		sub = sub.Where("mode = ?", req.Mode)
	}
	if req.Days > 0 {
		since := time.Now().AddDate(0, 0, -req.Days)
		sub = sub.Where("created_at >= ?", since)
	}
	sub = sub.Group("user_id")

	db := global.GVA_DB.Model(&model.Score{}).
		Joins("JOIN (?) AS top ON gva_wubi_score.user_id = top.user_id AND gva_wubi_score.wpm = top.max_wpm", sub)
	if req.Mode != "" {
		db = db.Where("gva_wubi_score.mode = ?", req.Mode)
	}
	order := "gva_wubi_score.wpm DESC"
	if req.Order == "accuracy" {
		order = "gva_wubi_score.accuracy DESC, gva_wubi_score.wpm DESC"
	}
	err = db.Order(order).Limit(limit).Find(&list).Error
	return
}
