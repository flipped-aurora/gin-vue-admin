package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type JobApplyService struct {
}

// CreateJobApply 创建JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService) CreateJobApply(jobApply *clothing.JobApply) (err error) {
	err = global.GVA_DB.Create(jobApply).Error
	return err
}

// DeleteJobApply 删除JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService)DeleteJobApply(jobApply clothing.JobApply) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.JobApply{}).Where("id = ?", jobApply.ID).Update("deleted_by", jobApply.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&jobApply).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteJobApplyByIds 批量删除JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService)DeleteJobApplyByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.JobApply{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.JobApply{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateJobApply 更新JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService)UpdateJobApply(jobApply clothing.JobApply) (err error) {
	err = global.GVA_DB.Save(&jobApply).Error
	return err
}

// GetJobApply 根据id获取JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService)GetJobApply(id uint) (jobApply clothing.JobApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&jobApply).Error
	return
}

// GetJobApplyInfoList 分页获取JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService)GetJobApplyInfoList(info clothingReq.JobApplySearch) (list []clothing.JobApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.JobApply{})
    var jobApplys []clothing.JobApply
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CroppingID != nil {
        db = db.Where("cropping_id = ?",info.CroppingID)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&jobApplys).Error
	return  jobApplys, total, err
}
