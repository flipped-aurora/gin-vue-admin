package home

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
)

type ShowService struct {
}

func (s *ShowService) GetInfoById(modeType int, showId uint) (model any, err error) {
	model, _, _ = createStruct(modeType)
	err = global.GVA_DB.Where("id = ?", showId).First(&model).Error
	return
}

// 获取上一页  下一页
func (s *ShowService) GetInfoNextPreviousById(modeType int, catID, showID int, ordertype string) (next, previous webcms.Class) {
	model, _, _ := createStruct(modeType)
	db1 := global.GVA_DB.Model(&model).Select("id,title").Where("enable", 1).Where("cate_id = ?", catID)
	db2 := global.GVA_DB.Model(&model).Select("id,title").Where("enable", 1).Where("cate_id = ?", catID)
	// 升序
	if ordertype == "1" {
		// 查询下一个ID
		db1.Where("id > ?", showID).Order("id ASC").First(&next)
		// 查询上一个ID
		db2.Where("id < ?", showID).Order("id ASC").First(&previous)
	} else {
		// 查询上一个ID
		db1.Where("id < ?", showID).Order("id DESC").First(&previous)
		// 查询下一个ID
		db2.Where("id > ?", showID).Order("id DESC").First(&next)
	}
	return
}
