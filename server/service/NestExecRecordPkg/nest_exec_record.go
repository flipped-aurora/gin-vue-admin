package NestExecRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg"
	NestExecRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestInfo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NestExecRecordService struct {
}

// CreateNestExecRecord 创建NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) CreateNestExecRecord(NtERecord *NestExecRecordPkg.NestExecRecord) (err error) {
	err = global.GVA_DB.Create(NtERecord).Error
	return err
}

// DeleteNestExecRecord 删除NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) DeleteNestExecRecord(NtERecord NestExecRecordPkg.NestExecRecord) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestExecRecordPkg.NestExecRecord{}).Where("id = ?", NtERecord.ID).Update("deleted_by", NtERecord.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&NtERecord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteNestExecRecordByIds 批量删除NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) DeleteNestExecRecordByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestExecRecordPkg.NestExecRecord{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&NestExecRecordPkg.NestExecRecord{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateNestExecRecord 更新NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) UpdateNestExecRecord(NtERecord NestExecRecordPkg.NestExecRecord) (err error) {
	err = global.GVA_DB.Save(&NtERecord).Error
	return err
}

// GetNestExecRecord 根据id获取NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) GetNestExecRecord(id uint) (NtERecord NestExecRecordPkg.NestExecRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&NtERecord).Error
	return
}

// GetNestExecRecordInfoList 分页获取NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) GetNestExecRecordInfoList(info NestExecRecordPkgReq.NestExecRecordSearch, c *gin.Context) (list []NestExecRecordPkg.NestExecRecord, total int64, err error) {
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&NestExecRecordPkg.NestExecRecord{})
	var NtERecords []NestExecRecordPkg.NestExecRecord
	db.Where("nest_id in ?", nestIDList)
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.NestId != "" {
		db = db.Where("nest_id = ?", info.NestId)
	}
	if info.ExecuteId != "" {
		db = db.Where("execute_id = ?", info.ExecuteId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
		if *info.Status == 3 {
			db = db.Order("execute_at desc")
		}
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&NtERecords).Error
	return NtERecords, total, err
}

// NoPageGetNestExecRecordInfoList 不分页获取NestExecRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtERecordService *NestExecRecordService) NoPageGetNestExecRecordInfoList(c *gin.Context) (list []map[string]interface{}, err error) {
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}
	// 创建db
	db := global.GVA_DB.Model(&NestExecRecordPkg.NestExecRecord{})
	db.Where("nest_id in ?", nestIDList)
	NtERecords := make([]map[string]interface{}, 0, 0)
	err = db.Order("execute_at desc").Where("execute_at <> ''").Find(&NtERecords).Error
	return NtERecords, err
}
