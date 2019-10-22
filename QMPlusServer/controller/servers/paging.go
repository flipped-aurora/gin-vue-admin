package servers

import (
	"github.com/jinzhu/gorm"
	"main/init/qmsql"
	"main/model/modelInterface"
)

//获取分页功能 接收实现了分页接口的结构体 返回搜索完成的结果 许需要自行scan 或者find
func PagingServer(paging modelInterface.Paging, info modelInterface.PageInfo) (err error, db *gorm.DB, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = qmsql.DEFAULTDB.Model(paging).Count(&total).Error
	db = qmsql.DEFAULTDB.Limit(limit).Offset(offset).Order("id desc")
	return err, db, total
}
