package xiao

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao/xiaores"
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
func (cliTreeService *CliTreeService) GetCliTreePublic(pageInfo *xiaoReq.CliTreeSearch) (res xiaores.CliTreeRes, err error) {
	// 此方法为获取数据源定义的数据
	// 查询直推和团队数据
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("transaction failed: %v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	res.Pulltrees, err = xiao.NewCliTree(pageInfo.Address, "").GetChildByLevel(tx, 1)
	if err != nil {
		return res, err
	}
	for _, pulltree := range res.Pulltrees {
		userinfo, _ := xiao.NewUser(pulltree.Address).GetCliUser(tx)
		pulltree.Father = userinfo.Nickname
		pulltree.Description = userinfo.Avatarurl
	}
	res.Teamtrees, err = xiao.NewCliTree(pageInfo.Address, "").GetPartChilds(tx, 1)
	if err != nil {
		return res, err
	}
	for _, teamtree := range res.Teamtrees {
		userinfo, _ := xiao.NewUser(teamtree.Address).GetCliUser(tx)
		teamtree.Father = userinfo.Nickname
		teamtree.Description = userinfo.Avatarurl
	}
	res.Cliuser, err = xiao.NewUser(pageInfo.Address).GetCliUser(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}
