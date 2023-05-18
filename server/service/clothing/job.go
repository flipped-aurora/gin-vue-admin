package clothing

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
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
func (jobService *JobService) DeleteJob(job clothing.Job) (err error) {
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
func (jobService *JobService) DeleteJobByIds(ids request.IdsReq, deleted_by uint) (err error) {
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
func (jobService *JobService) UpdateJob(job clothing.Job) (err error) {
	err = global.GVA_DB.Save(&job).Error
	return err
}

func (jobService *JobService) AuditApply(job clothing.Job, realQuantity uint) (err error) {
	err = global.GVA_DB.Model(&job).Updates(map[string]interface{}{
		"real_quantity": realQuantity,
		"real_income":   job.Price * float64(realQuantity),
		"step":          enum.CroppingAudit,
		"updated_by":    job.UpdatedBy,
	}).Error
	return err
}

// GetJob 根据id获取Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService) GetJob(id uint) (job clothing.Job, err error) {
	err = global.GVA_DB.Preload("Team").Where("id = ?", id).First(&job).Error
	return
}

// GetJobInfoList 分页获取Job记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobService *JobService) GetJobInfoList(info clothingReq.JobSearch) (list []clothing.Job, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.Job{})
	var jobs []clothing.Job
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CroppingID > 0 {
		db = db.Where("cropping_id = ?", info.CroppingID)
	}
	if info.TeamID > 0 {
		db = db.Where("team_id = ?", info.CroppingID)
	}
	if info.Step != 0 {
		db = db.Where("step = ?", info.Step)
	}
	if info.UserID > 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&jobs).Error
	return jobs, total, err
}

func (jobService *JobService) PostJob(cropping clothing.CroppingRecord, data clothingReq.JobList) error {
	for _, j := range data.Jobs {
		var job clothing.Job
		if err := global.GVA_DB.Where("cropping_id = ? and process_id = ?", data.CroppingID, j.ProcessID).First(&job).Error; err == nil {
			global.GVA_LOG.Sugar().Error("job 已存在:", job)
			continue
		}
		if j.ProcessID == 0 {
			if cropping.Step != enum.CroppingPending {
				return errors.New("裁剪单已处理中")
			}
			job.JobType = enum.Whole
			job.ProcessName = "成衣"
			job.Price = cropping.Style.Price
		} else {
			if cropping.Step == enum.CroppingComplete {
				return errors.New("裁剪单已完成")
			}
			job.JobType = enum.Process
			var process clothing.Process
			if err := global.GVA_DB.First(&process, j.ProcessID).Error; err != nil {
				global.GVA_LOG.Sugar().Error(err)
				continue
			}
			job.ProcessName = process.Name
			job.Price = process.Price
		}
		job.Quantity = int(cropping.Quantity)
		job.CroppingID = data.CroppingID
		job.UserID = j.UserID
		job.Income = float64(job.Quantity) * job.Price
		job.Step = enum.CroppingHandling
		job.TeamID = data.TeamID
		global.GVA_DB.Create(&job)
	}
	if err := global.GVA_DB.Model(&cropping).Update("step", enum.CroppingHandling).Error; err != nil {
		return err
	}
	return nil
}

func (jobService *JobService) JobAuditOpt(job clothing.Job, status bool) (err error) {
	if status {
		err = global.GVA_DB.Model(&job).Updates(map[string]interface{}{"step": enum.CroppingComplete, "updated_by": job.UpdatedBy}).Error
		if err != nil {
			return err
		}
		err = global.GVA_DB.Model(&clothing.UserWallet{}).Where("user_id = ? and company_id = ?", job.UserID, job.Team.CompanyID).
			Updates(map[string]interface{}{
				"wages":        gorm.Expr("wages + ?", job.RealIncome),
				"pendingWages": gorm.Expr("pendingWages + ?", job.RealIncome),
			}).Error
	} else {
		err = global.GVA_DB.Model(&job).Updates(map[string]interface{}{"step": enum.CroppingHandling, "updated_by": job.UpdatedBy}).Error
	}
	return
}
