package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"time"
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

	err = db.Order("id DESC").Find(&cinemaStatistics).Error
	return cinemaStatistics, total, err
}

func (cinemaStatisticsService *CinemaStatisticsService) GetCinemaStatisticsToday() (today cinema.CinemaStatisticsToday, err error) {
	db := global.GVA_DB.Model(&cinema.CinemaOrder{})
	err = db.Select("COALESCE(SUM(CASE WHEN status = 1 THEN film_price ELSE 0 END), 0) AS price, COALESCE(SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END), 0) AS total").
		Where("created_at >= ?", time.Now().Format("2006-01-02")).Find(&today).Error
	return today, err
}
