package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
)

type CinemaStatisticsService struct {
}

// GetCinemaStatisticsInfoList 分页获取cinemaStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaStatisticsService *CinemaStatisticsService) GetCinemaStatisticsInfoList(info cinemaReq.CinemaStatisticsSearch) (list []cinema.CinemaStatistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cinema.CinemaStatistics{})
	var cinemaStatistics []cinema.CinemaStatistics
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("date BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cinemaStatistics).Error
	return cinemaStatistics, total, err
}
