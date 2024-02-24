package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CinemaFilmSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Hall *int   `json:"hall" form:"hall" `
	Name string `json:"name" form:"name" `
	request.PageInfo
}

type CinemaFilmCreate struct {
	Hall     *int     `json:"hall" form:"hall" gorm:"column:hall;comment:影厅;"binding:"required"`
	Name     string   `json:"name" form:"name" gorm:"column:name;comment:电影名字;size:255;"binding:"required"`
	Price    *float64 `json:"price" form:"price" gorm:"column:price;comment:价格;"binding:"required"`
	PlayTime string   `json:"playTime" form:"playTime" gorm:"column:play_time;comment:播放时间;"binding:"required"`
	Type     string   `json:"type" form:"type" gorm:"column:type;comment:电影票类型;size:255;"`
}
