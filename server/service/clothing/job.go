package clothing

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"gorm.io/gorm"
	"time"
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

func (jobService *JobService) ChangeWorker(job clothing.Job) (err error) {
	err = global.GVA_DB.Model(&job).Update("user_id", job.UserID).Error
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
	err = global.GVA_DB.Preload("Team").Preload("Process.Style").Preload("User").Preload("CroppingRecord.Style").Where("id = ?", id).First(&job).Error
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
		db = db.Where("team_id = ?", info.TeamID)
	}
	if info.Step != 0 {
		db = db.Where("step = ?", info.Step)
	}
	if info.UserID > 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	if len(info.Size) > 0 {
		db = db.Where("size = ?", info.Size)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("Team").Preload("Process.Style").Preload("User").Preload("CroppingRecord.Style").Limit(limit).Offset(offset).Find(&jobs).Error
	return jobs, total, err
}

func (jobService *JobService) PostJob(cropping clothing.CroppingRecord, data clothingReq.JobList) error {
	var temp bool
	for _, j := range data.Jobs {
		var job clothing.Job
		if j.ProcessID == 0 {
			job.JobType = enum.Whole
			job.ProcessName = "成衣"
			job.Price = cropping.Style.Price
			var s clothing.SizeList
			global.GVA_DB.Where("cropping_record_id = ? and size = ?", cropping.ID, j.Size).First(&s)
			if s.Margin < j.Quantity {
				temp = true
				continue
			} else {
				global.GVA_DB.Model(&s).Update("margin", gorm.Expr("margin - ?", j.Quantity))
			}
		} else {
			job.JobType = enum.Process
			var process clothing.Process
			if err := global.GVA_DB.First(&process, j.ProcessID).Error; err != nil {
				temp = true
				continue
			}
			var i clothing.Inventory
			global.GVA_DB.Where("cropping_record_id = ? and process_id = ? and size = ?", cropping.ID, process.ID, j.Size).First(&i)
			if int(i.Margin) < job.Quantity {
				temp = true
				continue
			} else {
				global.GVA_DB.Model(&i).Update("margin", gorm.Expr("margin - ?", j.Quantity))
			}
			job.ProcessName = process.Name
			job.Price = process.Price
		}
		job.Quantity = int(j.Quantity)
		job.CroppingID = data.CroppingID
		job.CompanyID = cropping.CompanyID
		job.UserID = j.UserID
		job.Size = j.Size
		job.Income = float64(job.Quantity) * job.Price
		job.Step = enum.CroppingHandling
		job.TeamID = data.TeamID
		global.GVA_DB.Create(&job)
	}
	if err := global.GVA_DB.Model(&cropping).Update("step", enum.CroppingHandling).Error; err != nil {
		return err
	}
	if temp {
		return errors.New("部分分配失败")
	}
	return nil
}

func (jobService *JobService) JobAuditOpt(job clothing.Job, status bool) (err error) {
	if status {
		err = global.GVA_DB.Model(&job).Updates(map[string]interface{}{"step": enum.CroppingComplete, "updated_by": job.UpdatedBy, "complete_at": time.Now()}).Error
		if err != nil {
			return err
		}
		var wallet clothing.UserWallet
		global.GVA_DB.Where("user_id = ? and company_id = ?", job.UserID, job.Team.CompanyID).FirstOrCreate(&wallet)
		err = global.GVA_DB.Model(&wallet).
			Updates(map[string]interface{}{
				"wages":         gorm.Expr("wages + ?", job.RealIncome),
				"pending_wages": gorm.Expr("pending_wages + ?", job.RealIncome),
			}).Error
	} else {
		err = global.GVA_DB.Model(&job).Updates(map[string]interface{}{"step": enum.CroppingHandling, "updated_by": job.UpdatedBy}).Error
	}
	return
}

func (jobService *JobService) GetWagesDetail(info clothingReq.JobSearch) (list []clothing.Job, total int64, err error) {
	if info.CompanyID == 0 && info.UserID == 0 {
		err = errors.New("缺少参数")
		return
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&clothing.Job{}).Joins("join app_user on app_user.id = job.user_id").Where("job.step = ? AND job.is_cancel = ?", enum.CroppingComplete, false)
	if info.Keyword != "" {
		db = db.Where("app_user.username like ? or app_user.nickname like ? app_user.phone_num like ?", fmt.Sprintf("%%%s%%", info.Keyword), fmt.Sprintf("%%%s%%", info.Keyword), fmt.Sprintf("%%%s%%", info.Keyword))
	}
	if info.StartCreatedAt != nil {
		db = db.Where("job.complete_at >= ?", *info.StartCreatedAt)
	}
	if info.EndCreatedAt != nil {
		db = db.Where("job.complete_at <= ?", *info.EndCreatedAt)
	}
	if info.CompanyID > 0 {
		db = db.Where("job.company_id = ?", info.CompanyID)
	}
	if info.TeamID > 0 {
		db = db.Where("job.team_id = ?", info.TeamID)
	}
	if info.UserID > 0 {
		db = db.Where("job.user_id = ?", info.UserID)
	}

	db = db.Select("job.user_id,job.company_id,sum(job.real_income) as real_income").Group("user_id,company_id").Having("sum(job.real_income) >= ?", info.Coin)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("User").Preload("Company").Limit(limit).Offset(offset).Find(&list).Error
	return
}
