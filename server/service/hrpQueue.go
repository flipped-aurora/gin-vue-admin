package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHrpQueue
// @description   create a HrpQueue
// @param     hrpQueue               model.HrpQueue
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHrpQueue(hrpQueue model.HrpQueue) (err error) {
	err = global.GVA_DB.Create(&hrpQueue).Error
	return err
}

// @title    DeleteHrpQueue
// @description   delete a HrpQueue
// @auth                     （2020/04/05  20:22）
// @param     hrpQueue               model.HrpQueue
// @return                    error

func DeleteHrpQueue(hrpQueue model.HrpQueue) (err error) {
	err = global.GVA_DB.Delete(hrpQueue).Error
	return err
}

// @title    DeleteHrpQueueByIds
// @description   delete HrpQueues
// @auth                     （2020/04/05  20:22）
// @param     hrpQueue               model.HrpQueue
// @return                    error

func DeleteHrpQueueByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HrpQueue{},"id in ?",ids.Ids).Error
	return err
}

// @title    UpdateHrpQueue
// @description   update a HrpQueue
// @param     hrpQueue          *model.HrpQueue
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHrpQueue(hrpQueue *model.HrpQueue) (err error) {
	err = global.GVA_DB.Save(hrpQueue).Error
	return err
}

// @title    GetHrpQueue
// @description   get the info of a HrpQueue
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    HrpQueue        HrpQueue

func GetHrpQueue(id int32) (err error, hrpQueue model.HrpQueue) {
	err = global.GVA_DB.Where("id = ?", id).First(&hrpQueue).Error
	return
}

// @title    GetHrpQueueInfoList
// @description   get HrpQueue list by pagination, 分页获取HrpQueue
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHrpQueueInfoList(info request.HrpQueueSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.HrpQueue{})
    var hrpQueues []model.HrpQueue
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hrpQueues).Error
	return err, hrpQueues, total
}