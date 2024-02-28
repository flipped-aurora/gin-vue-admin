// 自动生成模板CinemaStatistics
package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// cinemaStatistics表 结构体  CinemaStatistics
type CinemaStatistics struct {
 global.GVA_MODEL 
      Date  string `json:"date" form:"date" gorm:"column:date;comment:时间;"`  //时间 
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:总票数;size:22;"`  //总票数 
      Total  *int `json:"total" form:"total" gorm:"column:total;comment:总票数;size:10;"`  //总票数 
}


// TableName cinemaStatistics表 CinemaStatistics自定义表名 cinema_statistics
func (CinemaStatistics) TableName() string {
  return "cinema_statistics"
}

