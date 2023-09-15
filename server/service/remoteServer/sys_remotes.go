package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    remoteServerReq "github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer/request"
    "gorm.io/gorm"
)

type SysRemotesService struct {
}

// CreateSysRemotes 创建远程服务器配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysRemotesService *SysRemotesService) CreateSysRemotes(sysRemotes *remoteServer.SysRemotes) (err error) {
	err = global.GVA_DB.Create(sysRemotes).Error
	return err
}

// DeleteSysRemotes 删除远程服务器配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysRemotesService *SysRemotesService)DeleteSysRemotes(sysRemotes remoteServer.SysRemotes) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&remoteServer.SysRemotes{}).Where("id = ?", sysRemotes.ID).Update("deleted_by", sysRemotes.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&sysRemotes).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteSysRemotesByIds 批量删除远程服务器配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysRemotesService *SysRemotesService)DeleteSysRemotesByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&remoteServer.SysRemotes{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&remoteServer.SysRemotes{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateSysRemotes 更新远程服务器配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysRemotesService *SysRemotesService)UpdateSysRemotes(sysRemotes remoteServer.SysRemotes) (err error) {
	err = global.GVA_DB.Save(&sysRemotes).Error
	return err
}

// GetSysRemotes 根据id获取远程服务器配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysRemotesService *SysRemotesService)GetSysRemotes(id uint) (sysRemotes remoteServer.SysRemotes, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysRemotes).Error
	return
}

// GetSysRemotesInfoList 分页获取远程服务器配置表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysRemotesService *SysRemotesService)GetSysRemotesInfoList(info remoteServerReq.SysRemotesSearch) (list []remoteServer.SysRemotes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&remoteServer.SysRemotes{})
    var sysRemotess []remoteServer.SysRemotes
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Code != "" {
        db = db.Where("code = ?",info.Code)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&sysRemotess).Error
	return  sysRemotess, total, err
}
