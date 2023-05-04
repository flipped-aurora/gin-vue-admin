package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type JobQuestionService struct {
}

// CreateJobQuestion 创建JobQuestion记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobQuestionService *JobQuestionService) CreateJobQuestion(jobQuestion *clothing.JobQuestion) (err error) {
	err = global.GVA_DB.Create(jobQuestion).Error
	return err
}

// DeleteJobQuestion 删除JobQuestion记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobQuestionService *JobQuestionService) DeleteJobQuestion(jobQuestion clothing.JobQuestion) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.JobQuestion{}).Where("id = ?", jobQuestion.ID).Update("deleted_by", jobQuestion.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&jobQuestion).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteJobQuestionByIds 批量删除JobQuestion记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobQuestionService *JobQuestionService) DeleteJobQuestionByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.JobQuestion{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.JobQuestion{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateJobQuestion 更新JobQuestion记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobQuestionService *JobQuestionService) UpdateJobQuestion(jobQuestion clothing.JobQuestion) (err error) {
	err = global.GVA_DB.Save(&jobQuestion).Error
	return err
}

// GetJobQuestion 根据id获取JobQuestion记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobQuestionService *JobQuestionService) GetJobQuestion(id uint) (jobQuestion clothing.JobQuestion, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&jobQuestion).Error
	return
}

// GetJobQuestionInfoList 分页获取JobQuestion记录
// Author [piexlmax](https://github.com/piexlmax)
func (jobQuestionService *JobQuestionService) GetJobQuestionInfoList(info clothingReq.JobQuestionSearch) (list []clothing.JobQuestion, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.JobQuestion{})
	var jobQuestions []clothing.JobQuestion
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.JobID != 0 {
		db = db.Where("job_id = ?", info.JobID)
	}
	if info.UserID != 0 {
		db = db.Where("user_id = ?", info.UserID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&jobQuestions).Error
	return jobQuestions, total, err
}
