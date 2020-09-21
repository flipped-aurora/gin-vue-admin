package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHrpTriagedQueue
// @description   create a HrpTriagedQueue
// @param     hrpTriagedQueue               model.HrpTriagedQueue
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHrpTriagedQueue(hrpTriagedQueue model.HrpTriagedQueue) (err error) {
	err = global.GVA_DB.Create(&hrpTriagedQueue).Error
	return err
}

// @title    DeleteHrpTriagedQueue
// @description   delete a HrpTriagedQueue
// @auth                     （2020/04/05  20:22）
// @param     hrpTriagedQueue               model.HrpTriagedQueue
// @return                    error

func DeleteHrpTriagedQueue(hrpTriagedQueue model.HrpTriagedQueue) (err error) {
	err = global.GVA_DB.Delete(hrpTriagedQueue).Error
	return err
}

// @title    DeleteHrpTriagedQueueByIds
// @description   delete HrpTriagedQueues
// @auth                     （2020/04/05  20:22）
// @param     hrpTriagedQueue               model.HrpTriagedQueue
// @return                    error

func DeleteHrpTriagedQueueByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HrpTriagedQueue{},"id in ?",ids.Ids).Error
	return err
}

// @title    UpdateHrpTriagedQueue
// @description   update a HrpTriagedQueue
// @param     hrpTriagedQueue          *model.HrpTriagedQueue
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHrpTriagedQueue(hrpTriagedQueue *model.HrpTriagedQueue) (err error) {
	err = global.GVA_DB.Save(hrpTriagedQueue).Error
	return err
}

// @title    GetHrpTriagedQueue
// @description   get the info of a HrpTriagedQueue
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    HrpTriagedQueue        HrpTriagedQueue

func GetHrpTriagedQueue(id int32) (err error, hrpTriagedQueue model.HrpTriagedQueue) {
	err = global.GVA_DB.Where("id = ?", id).First(&hrpTriagedQueue).Error
	return
}

// @title    GetHrpTriagedQueueInfoList
// @description   get HrpTriagedQueue list by pagination, 分页获取HrpTriagedQueue
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHrpTriagedQueueInfoList(info request.HrpTriagedQueueSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.HrpTriagedQueue{})
    var hrpTriagedQueues []model.HrpTriagedQueue
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hrpTriagedQueues).Error
	return err, hrpTriagedQueues, total
}