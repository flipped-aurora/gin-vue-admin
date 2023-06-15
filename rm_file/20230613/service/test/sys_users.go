package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
)

type SysUsersService struct {
}

// CreateSysUsers 创建SysUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) CreateSysUsers(sysUsers *test.SysUsers) (err error) {
	err = global.GVA_DB.Create(sysUsers).Error
	return err
}

// DeleteSysUsers 删除SysUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService)DeleteSysUsers(sysUsers test.SysUsers) (err error) {
	err = global.GVA_DB.Delete(&sysUsers).Error
	return err
}

// DeleteSysUsersByIds 批量删除SysUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService)DeleteSysUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]test.SysUsers{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysUsers 更新SysUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService)UpdateSysUsers(sysUsers test.SysUsers) (err error) {
	err = global.GVA_DB.Save(&sysUsers).Error
	return err
}

// GetSysUsers 根据id获取SysUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService)GetSysUsers(id uint) (sysUsers test.SysUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysUsers).Error
	return
}

// GetSysUsersInfoList 分页获取SysUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService)GetSysUsersInfoList(info testReq.SysUsersSearch) (list []test.SysUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&test.SysUsers{})
    var sysUserss []test.SysUsers
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&sysUserss).Error
	return  sysUserss, total, err
}
