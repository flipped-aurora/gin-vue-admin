package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "gorm.io/gorm"
)

type AgentService struct {
}

// CreateAgent 创建Agent记录
// Author [piexlmax](https://github.com/piexlmax)
func (agentService *AgentService) CreateAgent(agent *clothing.Agent) (err error) {
	err = global.GVA_DB.Create(agent).Error
	return err
}

// DeleteAgent 删除Agent记录
// Author [piexlmax](https://github.com/piexlmax)
func (agentService *AgentService)DeleteAgent(agent clothing.Agent) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Agent{}).Where("id = ?", agent.ID).Update("deleted_by", agent.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&agent).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAgentByIds 批量删除Agent记录
// Author [piexlmax](https://github.com/piexlmax)
func (agentService *AgentService)DeleteAgentByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&clothing.Agent{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.Agent{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAgent 更新Agent记录
// Author [piexlmax](https://github.com/piexlmax)
func (agentService *AgentService)UpdateAgent(agent clothing.Agent) (err error) {
	err = global.GVA_DB.Save(&agent).Error
	return err
}

// GetAgent 根据id获取Agent记录
// Author [piexlmax](https://github.com/piexlmax)
func (agentService *AgentService)GetAgent(id uint) (agent clothing.Agent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&agent).Error
	return
}

// GetAgentInfoList 分页获取Agent记录
// Author [piexlmax](https://github.com/piexlmax)
func (agentService *AgentService)GetAgentInfoList(info clothingReq.AgentSearch) (list []clothing.Agent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&clothing.Agent{})
    var agents []clothing.Agent
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&agents).Error
	return  agents, total, err
}
