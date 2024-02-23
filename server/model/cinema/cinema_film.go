// 自动生成模板CinemaFilm
package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// cinemaFilm表 结构体  CinemaFilm
type CinemaFilm struct {
 global.GVA_MODEL 
      Hall  *int `json:"hall" form:"hall" gorm:"column:hall;comment:影厅;"binding:"required"`  //影厅 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:电影名字;size:255;"binding:"required"`  //电影名字 
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:价格;"binding:"required"`  //价格 
      PlayTime  *time.Time `json:"playTime" form:"playTime" gorm:"column:play_time;comment:播放时间;"binding:"required"`  //播放时间 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:电影票类型;size:255;"`  //电影票类型 
}


// TableName cinemaFilm表 CinemaFilm自定义表名 cinema_film
func (CinemaFilm) TableName() string {
  return "cinema_film"
}

