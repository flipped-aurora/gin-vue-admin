package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysOperationRecord
// @description   create a SysOperationRecord
// @param     sysOperationRecord               model.SysOperationRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

// @title    DeleteSysOperationRecord
// @description   delete SysOperationRecords
// @auth                     （2020/04/05  20:22）
// @param     sysOperationRecord               request.IdsReq
// @return                    error

func DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysOperationRecord{},"id in (?)",ids.Ids).Error
	return err
}

// @title    DeleteSysOperationRecord
// @description   delete a SysOperationRecord
// @auth                     （2020/04/05  20:22）
// @param     sysOperationRecord               model.SysOperationRecord
// @return                    error

func DeleteSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GVA_DB.Delete(sysOperationRecord).Error
	return err
}

// @title    UpdateSysOperationRecord
// @description   update a SysOperationRecord
// @param     sysOperationRecord          *model.SysOperationRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSysOperationRecord(sysOperationRecord *model.SysOperationRecord) (err error) {
	err = global.GVA_DB.Save(sysOperationRecord).Error
	return err
}

// @title    GetSysOperationRecord
// @description   get the info of a SysOperationRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysOperationRecord        SysOperationRecord

func GetSysOperationRecord(id uint) (err error, sysOperationRecord model.SysOperationRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

// @title    GetSysOperationRecordInfoList
// @description   get SysOperationRecord list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysOperationRecord{})
	var sysOperationRecords []model.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
