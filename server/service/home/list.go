package home

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"go.uber.org/zap"
)

type ListService struct {
}

func (l *ListService) GetModelListByCatId(info request.PageInfo, catid int, modetype int) (modelslice any, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	model, modelslice, modelinfo := createStruct(modetype)
	db := global.GVA_DB.Model(model)
	db.Preload("CateMenus").Where("enable", 1)
	// 判断是否有子栏目
	var catids []uint
	global.GVA_DB.Model(&webcms.CateMenus{}).Where("parent_id = ?", catid).Select("id").Find(&catids)

	if len(catids) > 0 {
		db.Where("cate_id in ?", catids)
	} else {
		db.Where("cate_id = ?", catid)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order(info.OrderType).Find(&modelslice).Error
	switch modelinfo.StructName {
	case "Class":
		temp := modelslice.([]webcms.Class)
		for k, v := range temp {
			temp[k].Url = fmt.Sprint("/show/", v.CateId, "/", v.ID)
		}
		modelslice = temp
	case "Recruit":
		temp := modelslice.([]webcms.Recruit)
		for k, v := range temp {
			temp[k].Url = fmt.Sprint("/show/", v.CateId, "/", v.ID)
		}
		modelslice = temp
	}
	return
}

// 结构体映射
func createStruct(modetype int) (model any, modelslice any, modelinfo system.SysAutoCodeHistory) {
	// 通过 modeltype 获取模型信息
	err := global.GVA_DB.Where("id = ?", modetype).First(&modelinfo).Error
	if err != nil {
		global.GVA_LOG.Error("获取模型信息,err:", zap.Error(err))
	}
	switch modelinfo.StructName {
	case "Class":
		model = new(webcms.Class)
		modelslice = make([]webcms.Class, 0)
	case "Recruit":
		model = new(webcms.Recruit)
		modelslice = make([]webcms.Recruit, 0)
	}
	return
}
