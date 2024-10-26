package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
)

type CliTreeService struct{}

// CreateCliTree 创建用户关系表记录
// Author [yourname](https://github.com/yourname)
func (cliTreeService *CliTreeService) CreateCliTree(cliTree *xiao.CliTree) (err error) {
	err = global.GVA_DB.Create(cliTree).Error
	return err
}

// DeleteCliTree 删除用户关系表记录
// Author [yourname](https://github.com/yourname)
func (cliTreeService *CliTreeService) DeleteCliTree(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliTree{}, "id = ?", ID).Error
	return err
}

// DeleteCliTreeByIds 批量删除用户关系表记录
// Author [yourname](https://github.com/yourname)
func (cliTreeService *CliTreeService) DeleteCliTreeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliTree{}, "id in ?", IDs).Error
	return err
}

// UpdateCliTree 更新用户关系表记录
// Author [yourname](https://github.com/yourname)
func (cliTreeService *CliTreeService) UpdateCliTree(cliTree xiao.CliTree) (err error) {
	err = global.GVA_DB.Model(&xiao.CliTree{}).Where("id = ?", cliTree.ID).Updates(&cliTree).Error
	return err
}

// GetCliTree 根据ID获取用户关系表记录
// Author [yourname](https://github.com/yourname)
func (cliTreeService *CliTreeService) GetCliTree(ID string) (cliTree xiao.CliTree, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliTree).Error
	return
}

// GetCliTreeInfoList 分页获取用户关系表记录
// Author [yourname](https://github.com/yourname)
func (cliTreeService *CliTreeService) GetCliTreeInfoList(info xiaoReq.CliTreeSearch) (list []xiao.CliTree, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliTree{})
	var cliTrees []xiao.CliTree
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Father != "" {
		db = db.Where("father LIKE ?", "%"+info.Father+"%")
	}
	if info.Invite != "" {
		db = db.Where("invite LIKE ?", "%"+info.Invite+"%")
	}
	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cliTrees).Error
	return cliTrees, total, err
}
func (cliTreeService *CliTreeService) GetCliTreePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
