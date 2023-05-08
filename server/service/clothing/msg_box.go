package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/enum"
	"gorm.io/gorm"
)

type MsgBoxService struct {
}

// CreateMsgBox 创建MsgBox记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgBoxService *MsgBoxService) CreateMsgBox(msgBox *clothing.MsgBox) (err error) {
	err = global.GVA_DB.Create(msgBox).Error
	return err
}

// DeleteMsgBox 删除MsgBox记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgBoxService *MsgBoxService) DeleteMsgBox(msgBox clothing.MsgBox) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.MsgBox{}).Where("id = ?", msgBox.ID).Update("deleted_by", msgBox.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&msgBox).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMsgBoxByIds 批量删除MsgBox记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgBoxService *MsgBoxService) DeleteMsgBoxByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&clothing.MsgBox{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&clothing.MsgBox{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMsgBox 更新MsgBox记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgBoxService *MsgBoxService) UpdateMsgBox(msgBox clothing.MsgBox) (err error) {
	err = global.GVA_DB.Save(&msgBox).Error
	return err
}

// GetMsgBox 根据id获取MsgBox记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgBoxService *MsgBoxService) GetMsgBox(id uint) (msgBox clothing.MsgBox, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&msgBox).Error
	return
}

func (msgBoxService *MsgBoxService) SetRead(id uint) (err error) {
	err = global.GVA_DB.Model(&clothing.MsgBox{}).Where("id = ?", id).Update("status", enum.MsgRead).Error
	return
}

// GetMsgBoxInfoList 分页获取MsgBox记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgBoxService *MsgBoxService) GetMsgBoxInfoList(info clothingReq.MsgBoxSearch) (list []clothing.MsgBox, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&clothing.MsgBox{}).Order("status asc,id desc")
	var msgBoxs []clothing.MsgBox
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MsgType != 0 {
		db = db.Where("msg_type = ?", info.MsgType)
	}
	if info.To != 0 {
		db = db.Where("to_user = ?", info.To)
	}
	if info.From != 0 {
		db = db.Where("from_user = ?", info.From)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&msgBoxs).Error
	return msgBoxs, total, err
}

func (msgBoxService *MsgBoxService) SendMsg(from, to uint, objType int, objID uint) error {
	var msg clothing.MsgBox
	status := new(bool)
	*status = enum.MsgPending
	msg.From = from
	msg.To = to
	msg.MsgType = uint(objType)
	msg.MsgID = objID
	msg.Status = status
	return global.GVA_DB.Create(&msg).Error

}
