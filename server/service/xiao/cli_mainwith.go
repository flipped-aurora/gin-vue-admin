package xiao

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao/xiaores"
)

type CliMainwithService struct{}

// CreateCliMainwith 创建提币总表记录
// Author [yourname](https://github.com/yourname)
func (climainwithService *CliMainwithService) CreateCliMainwith(climainwith *xiao.CliMainwith) (err error) {
	err = global.GVA_DB.Create(climainwith).Error
	return err
}

// DeleteCliMainwith 删除提币总表记录
// Author [yourname](https://github.com/yourname)
func (climainwithService *CliMainwithService) DeleteCliMainwith(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliMainwith{}, "id = ?", ID).Error
	return err
}

// DeleteCliMainwithByIds 批量删除提币总表记录
// Author [yourname](https://github.com/yourname)
func (climainwithService *CliMainwithService) DeleteCliMainwithByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliMainwith{}, "id in ?", IDs).Error
	return err
}

// UpdateCliMainwith 更新提币总表记录
// Author [yourname](https://github.com/yourname)
func (climainwithService *CliMainwithService) UpdateCliMainwith(climainwith xiao.CliMainwith) (err error) {
	err = global.GVA_DB.Model(&xiao.CliMainwith{}).Where("id = ?", climainwith.ID).Updates(&climainwith).Error
	return err
}

// GetCliMainwith 根据ID获取提币总表记录
// Author [yourname](https://github.com/yourname)
func (climainwithService *CliMainwithService) GetCliMainwith(ID string) (climainwith xiao.CliMainwith, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&climainwith).Error
	return
}

// GetCliMainwithInfoList 分页获取提币总表记录
// Author [yourname](https://github.com/yourname)
func (climainwithService *CliMainwithService) GetCliMainwithInfoList(info xiaoReq.CliMainwithSearch) (list []xiao.CliMainwith, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliMainwith{})
	var climainwiths []xiao.CliMainwith
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["num"] = true
	orderMap["withable"] = true
	orderMap["withed"] = true
	orderMap["total"] = true
	orderMap["descnum"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&climainwiths).Error
	return climainwiths, total, err
}
func (climainwithService *CliMainwithService) GetCliMainwithPublic(pageInfo *xiaoReq.CliMainwithSearch) (res xiaores.CliWithdrawRes, err error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
	tx := global.GVA_DB
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
	//查询提币总表
	res.CliMainwiths, err = xiao.NewCliMainwith(pageInfo.Address).GetCliMainwith(tx)
	if err != nil {
		return res, err
	}
	//查询提币记录

	return res, nil
}
