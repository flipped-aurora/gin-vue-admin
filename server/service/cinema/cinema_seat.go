package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
    cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
)

type CinemaSeatService struct {
}

// CreateCinemaSeat 创建cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService) CreateCinemaSeat(cinemaSeat *cinema.CinemaSeat) (err error) {
	err = global.GVA_DB.Create(cinemaSeat).Error
	return err
}

// DeleteCinemaSeat 删除cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService)DeleteCinemaSeat(ID string) (err error) {
	err = global.GVA_DB.Delete(&cinema.CinemaSeat{},"id = ?",ID).Error
	return err
}

// DeleteCinemaSeatByIds 批量删除cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService)DeleteCinemaSeatByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]cinema.CinemaSeat{},"id in ?",IDs).Error
	return err
}

// UpdateCinemaSeat 更新cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService)UpdateCinemaSeat(cinemaSeat cinema.CinemaSeat) (err error) {
	err = global.GVA_DB.Save(&cinemaSeat).Error
	return err
}

// GetCinemaSeat 根据ID获取cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService)GetCinemaSeat(ID string) (cinemaSeat cinema.CinemaSeat, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cinemaSeat).Error
	return
}

// GetCinemaSeatInfoList 分页获取cinemaSeat表记录
// Author [piexlmax](https://github.com/piexlmax)
func (cinemaSeatService *CinemaSeatService)GetCinemaSeatInfoList(info cinemaReq.CinemaSeatSearch) (list []cinema.CinemaSeat, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cinema.CinemaSeat{})
    var cinemaSeats []cinema.CinemaSeat
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&cinemaSeats).Error
	return  cinemaSeats, total, err
}
