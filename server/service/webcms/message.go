package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
)

type MessageService struct {
}

// CreateMessage 创建Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) CreateMessage(message webcms.Message) (err error) {
	err = global.GVA_DB.Create(&message).Error
	return err
}

// DeleteMessage 删除Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) DeleteMessage(message webcms.Message) (err error) {
	err = global.GVA_DB.Delete(&message).Error
	return err
}

// DeleteMessageByIds 批量删除Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) DeleteMessageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]webcms.Message{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMessage 更新Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) UpdateMessage(message webcms.Message) (err error) {
	err = global.GVA_DB.Save(&message).Error
	return err
}

// GetMessage 根据id获取Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) GetMessage(id uint) (message webcms.Message, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&message).Error
	return
}

// GetMessageInfoList 分页获取Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) GetMessageInfoList(info webcmsReq.MessageSearch) (list []webcms.Message, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webcms.Message{})
	var messages []webcms.Message
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&messages).Error
	return messages, total, err
}
