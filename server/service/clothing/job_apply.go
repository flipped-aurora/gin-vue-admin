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
func (jobApplyService *JobApplyService) DeleteJobApply(jobApply clothing.JobApply) (err error) {
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
func (jobApplyService *JobApplyService) DeleteJobApplyByIds(ids request.IdsReq, deleted_by uint) (err error) {
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
func (jobApplyService *JobApplyService) UpdateJobApply(jobApply clothing.JobApply) (err error) {
	err = global.GVA_DB.Save(&jobApply).Error
	return err
}

// GetJobApply 根据id获取JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService) GetJobApply(id uint) (jobApply clothing.JobApply, err error) {
	err = global.GVA_DB.Preload("Team").Preload("Process.Style").Preload("User").Preload("CroppingRecord.Style").Where("id = ?", id).First(&jobApply).Error
	return
}

// GetJobApplyInfoList 分页获取JobApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobApplyService *JobApplyService) GetJobApplyInfoList(info clothingReq.JobApplySearch) (list []clothing.JobApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.JobApply{})
	var jobApplys []clothing.JobApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CroppingID != 0 {
		db = db.Where("cropping_id = ?", info.CroppingID)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("Team").Preload("Process.Style").Preload("User").Preload("CroppingRecord.Style").Limit(limit).Offset(offset).Find(&jobApplys).Error
	return jobApplys, total, err
}

func (jobApplyService *JobApplyService) OptApply(apply clothing.JobApply, status int) (err error) {
	rejectF := func(croppingRecord *clothing.CroppingRecord, process *clothing.Process, apply *clothing.JobApply) {
		global.GVA_DB.Model(&apply).Update("status", enum.ApplyReject)
		switch apply.JobType {
		case enum.Process:
			global.GVA_DB.Model(&clothing.Inventory{}).Where("cropping_record_id = ? and process_id = ? and size = ?", croppingRecord.ID, process.ID, apply.Size).Update("margin", gorm.Expr("margin + ?", apply.Quantity))
		case enum.Whole:
			global.GVA_DB.Model(&clothing.SizeList{}).Where("cropping_record_id = ? and size = ?", croppingRecord.ID, apply.Size).Update("margin", gorm.Expr("margin + ?", apply.Quantity))
		}
	}
	if apply.Status != nil && *apply.Status != enum.ApplyPending {
		return errors.New("已处理")
	}
	var cropping clothing.CroppingRecord
	if err := global.GVA_DB.Preload("Style").First(&cropping, apply.CroppingID).Error; err != nil {
		return errors.New("裁剪单不存在")
	}
	var job clothing.Job
	switch apply.JobType {
	case enum.Whole:
		// 整件
		if cropping.Step != enum.CroppingPending || status == enum.ApplyReject {
			rejectF(&cropping, nil, &apply)
			return errors.New("裁剪单已处理中")
		}
		job.ProcessID = 0
		job.ProcessName = "成衣"
		job.Price = cropping.Style.Price
	case enum.Process:
		var process clothing.Process
		if err := global.GVA_DB.First(&process, apply.ProcessID).Error; err != nil {
			return errors.New("工序不存在")
		}
		// 工序
		if cropping.Step == enum.CroppingComplete || status == enum.ApplyReject {
			rejectF(&cropping, &process, &apply)
			return errors.New("裁剪单已完成")
		}
		job.ProcessID = process.ID
		job.ProcessName = process.Name
		job.Price = process.Price
	}
	if err := global.GVA_DB.Model(&apply).Update("status", status).Error; err != nil {
		return err
	}
	if err := global.GVA_DB.Model(&cropping).Update("step", enum.CroppingHandling).Error; err != nil {
		return err
	}
	job.JobType = apply.JobType
	job.CroppingID = apply.CroppingID
	job.UserID = apply.UserID
	job.CompanyID = cropping.CompanyID
	job.TeamID = apply.TeamID
	job.Step = enum.CroppingHandling
	job.Quantity = int(cropping.Quantity)
	job.Income = job.Price * float64(job.Quantity)
	return global.GVA_DB.Create(&job).Error
}
