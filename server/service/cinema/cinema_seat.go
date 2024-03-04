package cinema

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"github.com/pkg/errors"
)

type CinemaSeatService struct {
}

// CreateCinemaSeat 创建cinemaSeat表记录
func (cinemaSeatService *CinemaSeatService) CreateCinemaSeat(req *cinemaReq.CinemaSeatCreate) (err error) {
	// 判断是否有重复数据
	tx := global.GVA_DB
	// 日期格式化
	date := req.Date.Format("2006-01-02")
	for _, v := range req.Positions {
		var count int64
		err = tx.Model(cinema.CinemaSeat{}).Where("film_id = ? AND date = ? AND position = ?", req.FilmId, date, v).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			err = errors.New(v + "座位已经售卖")
			return err
		}
	}
	var film cinema.CinemaFilm
	err = tx.First(&film, req.FilmId).Error
	if err != nil {
		return err
	}
	tx = tx.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	for _, v := range req.Positions {
		seat := cinema.CinemaSeat{
			FilmId:   &req.FilmId,
			Date:     date,
			Position: v,
		}
		err := tx.Create(&seat).Error
		if err != nil {
			return err
		}
		filmId := int(film.ID)
		seatId := int(seat.ID)
		insertOrder := cinema.CinemaOrder{
			FilmId:    &filmId,
			FilmHall:  film.Hall,
			FilmSeat:  v,
			FilmName:  film.Name,
			FilmType:  film.Type,
			PlayTime:  film.PlayTime,
			FilmPrice: film.Price,
			SeatId:    &seatId,
			Status:    1,
		}
		err = tx.Create(&insertOrder).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteCinemaSeat 删除cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService) DeleteCinemaSeat(ID string) (err error) {
	fmt.Println(ID)
	err = global.GVA_DB.Delete(&cinema.CinemaSeat{}, "id = ?", ID).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Model(&cinema.CinemaOrder{}).Where("seat_id = ? ", ID).Update("status", 2).Error
	return err
}

// DeleteCinemaSeatByIds 批量删除cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService) DeleteCinemaSeatByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cinema.CinemaSeat{}, "id in ?", IDs).Error
	return err
}

// GetCinemaSeat 根据ID获取cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService) GetCinemaSeat(ID string) (cinemaSeat cinema.CinemaSeat, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cinemaSeat).Error
	return
}

// GetCinemaSeatInfoList 分页获取cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService) GetCinemaSeatInfoList(info cinemaReq.CinemaSeatSearch) (list []cinema.CinemaSeat, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cinema.CinemaSeat{})
	var cinemaSeats []cinema.CinemaSeat
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.FilmId != 0 {
		db = db.Where("film_id = ?", info.FilmId)
	} else {
		return cinemaSeats, total, err
	}

	if info.Date != nil {
		date := info.Date.Format("2006-01-02")
		db = db.Where("date = ?", date)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&cinemaSeats).Error
	return cinemaSeats, total, err
}
