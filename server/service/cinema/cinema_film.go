package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
    cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
)

type CinemaFilmService struct {
}

// CreateCinemaFilm 创建cinemaFilm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaFilmService *CinemaFilmService) CreateCinemaFilm(cinemaFilm *cinema.CinemaFilm) (err error) {
	err = global.GVA_DB.Create(cinemaFilm).Error
	return err
}

// DeleteCinemaFilm 删除cinemaFilm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaFilmService *CinemaFilmService)DeleteCinemaFilm(ID string) (err error) {
	err = global.GVA_DB.Delete(&cinema.CinemaFilm{},"id = ?",ID).Error
	return err
}

// DeleteCinemaFilmByIds 批量删除cinemaFilm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaFilmService *CinemaFilmService)DeleteCinemaFilmByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cinema.CinemaFilm{},"id in ?",IDs).Error
	return err
}

// UpdateCinemaFilm 更新cinemaFilm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaFilmService *CinemaFilmService)UpdateCinemaFilm(cinemaFilm cinema.CinemaFilm) (err error) {
	err = global.GVA_DB.Save(&cinemaFilm).Error
	return err
}

// GetCinemaFilm 根据ID获取cinemaFilm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaFilmService *CinemaFilmService)GetCinemaFilm(ID string) (cinemaFilm cinema.CinemaFilm, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cinemaFilm).Error
	return
}

// GetCinemaFilmInfoList 分页获取cinemaFilm表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaFilmService *CinemaFilmService)GetCinemaFilmInfoList(info cinemaReq.CinemaFilmSearch) (list []cinema.CinemaFilm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cinema.CinemaFilm{})
    var cinemaFilms []cinema.CinemaFilm
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Hall != nil {
        db = db.Where("hall = ?",info.Hall)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&cinemaFilms).Error
	return  cinemaFilms, total, err
}
