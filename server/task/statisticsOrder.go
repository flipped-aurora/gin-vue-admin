package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
	"gorm.io/gorm"
	"time"
)

//@author: lsk
//@function: StatisticsOrder
//@description: 每日统计昨日数据
//@return: error

func StatisticsOrder(db *gorm.DB) error {
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	var order cinema.CinemaOrder
	var statistics cinema.CinemaStatistics
	// 统计昨日订单数据
	err := db.Model(&order).Select(
		"COALESCE(SUM(CASE WHEN status = 1 THEN film_price ELSE 0 END), 0) AS price,"+
			"COALESCE(SUM(CASE WHEN status = 1 THEN 1 ELSE 0 END), 0)  AS total,"+
			"COALESCE(SUM(CASE WHEN status = 2 THEN 1 ELSE 0 END), 0)  AS refund_total,"+
			"COALESCE(SUM(CASE WHEN status = 2 THEN film_price ELSE 0 END), 0)  AS refund_price").Where("created_at BETWEEN ? AND ?", yesterday, today).Find(&statistics).Error
	if err != nil {
		return err
	}
	statistics.Date = yesterday
	err = db.Model(statistics).Create(&statistics).Error
	if err != nil {
		return err
	}
	return nil
}
