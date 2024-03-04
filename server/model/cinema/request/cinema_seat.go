package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CinemaSeatSearch struct {
	FilmId int        `json:"filmId" form:"filmId" gorm:"column:film_id;comment:电影院;size:10;"` // 电影院
	Date   *time.Time `json:"date" form:"date" gorm:"column:date;comment:打印日期;"`

	request.PageInfo
}

type CinemaSeatCreate struct {
	FilmId    int        `json:"filmId" form:"filmId" gorm:"column:film_id;comment:电影院;size:10;"`                           // 电影院
	Date      *time.Time `json:"date" form:"date" gorm:"column:date;comment:打印日期;"`                                         // 打印日期
	Positions []string   `json:"positions" form:"position" binding:"required" gorm:"column:position;comment:几排几座;size:10;"` // 几排几座
}
