package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type JobService struct {
}

// CreateJob 创建Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService) CreateJob(job *clothing.Job) (err error) {
	err = global.GVA_DB.Create(job).Error
	return err
}

// DeleteJob 删除Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)DeleteJob(job clothing.Job) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Job{}).Where("id = ?", job.ID).Update("deleted_by", job.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&job).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteJobByIds 批量删除Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)DeleteJobByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Job{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Job{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateJob 更新Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)UpdateJob(job clothing.Job) (err error) {
	err = global.GVA_DB.Save(&job).Error
	return err
}

// GetJob 根据id获取Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)GetJob(id uint) (job clothing.Job, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&job).Error
	return
}

// GetJobInfoList 分页获取Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService)GetJobInfoList(info clothingReq.JobSearch) (list []clothing.Job, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.Job{})
    var jobs []clothing.Job
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Step != nil {
        db = db.Where("step = ?",info.Step)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&jobs).Error
	return  jobs, total, err
}
