package NestAirlinePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg"
	NestAirlinePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/NestInfo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NestAirlineService struct {
}

// CreateNestAirline 创建NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) CreateNestAirline(NtAirline *NestAirlinePkg.NestAirline) (err error) {
	err = global.GVA_DB.Create(NtAirline).Error
	return err
}

// DeleteNestAirline 删除NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) DeleteNestAirline(NtAirline NestAirlinePkg.NestAirline) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestAirlinePkg.NestAirline{}).Where("id = ?", NtAirline.ID).Update("deleted_by", NtAirline.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&NtAirline).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteNestAirlineByIds 批量删除NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) DeleteNestAirlineByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&NestAirlinePkg.NestAirline{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&NestAirlinePkg.NestAirline{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateNestAirline 更新NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) UpdateNestAirline(NtAirline NestAirlinePkg.NestAirline) (err error) {
	err = global.GVA_DB.Save(&NtAirline).Error
	return err
}

// GetNestAirline 根据id获取NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) GetNestAirline(id uint) (NtAirline NestAirlinePkg.NestAirline, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&NtAirline).Error
	return
}

// GetNestAirlineInfoList 分页获取NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
//func (NtAirlineService *NestAirlineService) GetNestAirlineInfoList(info NestAirlinePkgReq.NestAirlineSearch, c *gin.Context) (list []NestAirlinePkg.NestAirline, total int64, err error) {
//	nestInfoService := new(NestInfo.NestInfoService)
//	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
//	if err != nil {
//		return
//	}
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.GVA_DB.Model(&NestAirlinePkg.NestAirline{})
//	var NtAirlines []NestAirlinePkg.NestAirline
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
//		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
//	}
//	db.Where("nest_id in ?", nestIDList)
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//
//	err = db.Limit(limit).Offset(offset).Find(&NtAirlines).Error
//	return NtAirlines, total, err
//}
// GetNestAirlineInfoList 分页获取NestAirline记录并计算执行完的作业记录数
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) GetNestAirlineInfoList(info NestAirlinePkgReq.NestAirlineSearch, c *gin.Context) (list []NestAirlinePkg.NestAirline, total int64, err error) {
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&NestAirlinePkg.NestAirline{})
	var NtAirlines []NestAirlinePkg.NestAirline
	var copyNtAirlines []NestAirlinePkg.NestAirline
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db.Where("nest_id in ?", nestIDList)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&NtAirlines).Error
	for _, airline := range NtAirlines {
		db2 := global.GVA_DB.Model(&NestExecRecordPkg.NestExecRecord{})
		var Count int
		scanErr := db2.Raw("select count(1) Count from nest_exec_record  where missionid = ?", airline.Missionid).Scan(&Count)
		if scanErr.Error != nil {
			global.GVA_LOG.Error(scanErr.Error.Error())
		}
		airline.RecordCount = &Count
		copyNtAirlines = append(copyNtAirlines, airline)
	}
	NtAirlines = nil
	return copyNtAirlines, total, err
}

// NoPageGetNestAirlineInfoList 不分页获取NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) NoPageGetNestAirlineInfoList(nestId string, c *gin.Context) (list []map[string]interface{}, err error) {
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}
	// 创建db
	db := global.GVA_DB.Model(&NestAirlinePkg.NestAirline{})
	db.Where("nest_id in ?", nestIDList)
	//var NtAirlines []NestAirlinePkg.NestAirline
	NtAirlines := make([]map[string]interface{}, 0, 0)
	if nestId != "" {
		err = db.Where("nest_id = ?", nestId).Order("created_at desc").Find(&NtAirlines).Error
	} else {
		err = db.Order("created_at desc").Find(&NtAirlines).Error
	}

	return NtAirlines, err
}

// GetNestAirlineByMIssionId 根据missionid获取NestAirline记录
// Author [piexlmax](https://github.com/piexlmax)
func (NtAirlineService *NestAirlineService) GetNestAirlineByMIssionId(missionId string, c *gin.Context) (NtAirline NestAirlinePkg.NestAirline, err error) {
	nestInfoService := new(NestInfo.NestInfoService)
	nestIDList, err := nestInfoService.GetNestIDListByUser(c)
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("missionid = ? and nest_id in ?", missionId, nestIDList).First(&NtAirline).Error
	return
}
