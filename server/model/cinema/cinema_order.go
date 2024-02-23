// 自动生成模板CinemaOrder
package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// cinemaOrder表 结构体  CinemaOrder
type CinemaOrder struct {
 global.GVA_MODEL 
      SeatId  *int `json:"seatId" form:"seatId" gorm:"column:seat_id;comment:座位ID;size:10;"`  //座位ID 
      FilmId  *int `json:"filmId" form:"filmId" gorm:"column:film_id;comment:电影院;size:10;"`  //电影院 
      FilmHall  *bool `json:"filmHall" form:"filmHall" gorm:"column:film_hall;comment:影厅;"`  //影厅 
      FilmSeat  string `json:"filmSeat" form:"filmSeat" gorm:"column:film_seat;comment:选座;size:512;"`  //选座 
      FilmName  string `json:"filmName" form:"filmName" gorm:"column:film_name;comment:电影价格;size:255;"`  //电影价格 
      FilmType  string `json:"filmType" form:"filmType" gorm:"column:film_type;comment:电影类型;size:255;"`  //电影类型 
      PlayTime  *time.Time `json:"playTime" form:"playTime" gorm:"column:play_time;comment:播放时间;"`  //播放时间 
      FilmPrice  *float64 `json:"filmPrice" form:"filmPrice" gorm:"column:film_price;comment:价格;size:10;"`  //价格 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:状态 1. 正常 0.退款;"`  //状态 1. 正常 0.退款 
}


// TableName cinemaOrder表 CinemaOrder自定义表名 cinema_order
func (CinemaOrder) TableName() string {
  return "cinema_order"
}

