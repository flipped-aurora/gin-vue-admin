package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"time"
)

type CinemaOrderService struct {
}

// CreateCinemaOrder 创建cinemaOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaOrderService *CinemaOrderService) CreateCinemaOrder(cinemaOrder *cinema.CinemaOrder) (err error) {
	err = global.GVA_DB.Create(cinemaOrder).Error
	return err
}

// DeleteCinemaOrder 删除cinemaOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaOrderService *CinemaOrderService) DeleteCinemaOrder(ID string) (err error) {
	err = global.GVA_DB.Delete(&cinema.CinemaOrder{}, "id = ?", ID).Error
	return err
}

// DeleteCinemaOrderByIds 批量删除cinemaOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaOrderService *CinemaOrderService) DeleteCinemaOrderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cinema.CinemaOrder{}, "id in ?", IDs).Error
	return err
}

// UpdateCinemaOrder 更新cinemaOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaOrderService *CinemaOrderService) UpdateCinemaOrder(cinemaOrder cinema.CinemaOrder) (err error) {
	err = global.GVA_DB.Save(&cinemaOrder).Error
	return err
}

// GetCinemaOrder 根据ID获取cinemaOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaOrderService *CinemaOrderService) GetCinemaOrder(ID string) (cinemaOrder cinema.CinemaOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cinemaOrder).Error
	return
}

// GetCinemaOrderInfoList 分页获取cinemaOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaOrderService *CinemaOrderService) GetCinemaOrderInfoList(info cinemaReq.CinemaOrderSearch) (list []cinema.CinemaOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cinema.CinemaOrder{})
	var cinemaOrders []cinema.CinemaOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.OrderStatus != 0 {
		db = db.Where("status = ?", info.OrderStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&cinemaOrders).Error
	return cinemaOrders, total, err
}

// GetCinemaOrderStatistics 获取cinemaOrder
func (cinemaOrderService *CinemaOrderService) GetCinemaOrderStatistics() (cinemaOrderStatistics []cinema.StatisticsOrder, err error) {
	db := global.GVA_DB.Model(&cinema.CinemaOrder{})
	var cinemaOrderStats []cinema.StatisticsOrder // 最近一个月的订单统计
	// 查询最近一个月的订单统计
	err = db.Select("date(created_at) as date, sum(film_price) as totalPrice, count(*) as totalCount").
		Where("created_at >= ?", time.Now().AddDate(0, 0, -30)).
		Group("date(created_at)").Find(&cinemaOrderStats).Error

	if err != nil {
		return nil, err
	}
	return cinemaOrderStats, nil
}
