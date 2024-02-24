// 自动生成模板CinemaSeat
package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// cinemaSeat表 结构体  CinemaSeat
type CinemaSeat struct {
	global.GVA_MODEL
	FilmId   *int   `json:"filmId" form:"filmId" gorm:"column:film_id;comment:电影院;size:10;"`       //电影院
	Date     string `json:"date" form:"date" gorm:"column:date;comment:打印日期;"`                     //打印日期
	Position string `json:"position" form:"position" gorm:"column:position;comment:几排几座;size:10;"` //几排几座
	Status   *bool  `json:"status" form:"status" gorm:"column:status;comment:状态1: 正常"`             //状态1: 正常 2: 重打 0:退款
}

// TableName cinemaSeat表 CinemaSeat自定义表名 cinema_seat
func (CinemaSeat) TableName() string {
	return "cinema_seat"
}
