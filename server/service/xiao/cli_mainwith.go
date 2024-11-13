package xiao

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/shopspring/decimal"
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
func (climainwithService *CliMainwithService) GetCliMainwithPublic() (err error) {
	// 开始事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("transaction failed: %v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	// 查询设置
	var set xiao.CliSet
	err = tx.Where("id = ?", 1).First(&set).Error
	if err != nil {
		return errors.Join(err, errors.New("查询设置失败"))
	}

	// 查询所有状态正常订单
	var orders []xiao.CliOrder
	err = tx.Where("status = ?", "正常").Find(&orders).Error
	if err != nil {
		return errors.Join(err, errors.New("查询订单失败"))
	}

	// 计算静态收益
	for _, order := range orders {
		if order.Address == "" || order.Address == "root" {
			continue
		}

		// 计算静态收益
		staticprofit := order.Amount.Mul(*set.Static).Div(decimal.NewFromInt(100))

		// 更新收益数据
		order.Todaynum = &staticprofit
		*order.Descnum = order.Descnum.Add(staticprofit)
		if err := tx.Save(&order).Error; err != nil {
			return errors.Join(err, errors.New("更新订单失败"))
		}

		// 新建收益详情
		profit := xiao.CliProfit{
			Address: order.Address,
			Amount:  &staticprofit,
			Text:    "静态收益",
		}
		if err := tx.Create(&profit).Error; err != nil {
			return errors.Join(err, errors.New("创建收益详情失败"))
		}

		// 更新收益总表
		cliMainprofit, err := xiao.NewCliMainprofit(order.Address).GetCliMainprofitAddress(tx)
		if err != nil {
			return errors.Join(err, errors.New("查询收益总表失败"))
		}
		if cliMainprofit != nil {
			*cliMainprofit.Static = cliMainprofit.Static.Add(staticprofit)
			*cliMainprofit.Amount = cliMainprofit.Amount.Add(staticprofit)
			if err := tx.Save(&cliMainprofit).Error; err != nil {
				return errors.Join(err, errors.New("更新收益总表失败"))
			}
		}

		// 更新提币总表
		cliMainwith, err := xiao.NewCliMainwith(order.Address).GetCliMainwith(tx)
		if err != nil {
			return errors.Join(err, errors.New("查询提币总表失败"))
		}
		if cliMainwith != nil {
			*cliMainwith.Withable = cliMainwith.Withable.Add(staticprofit)
			if err := tx.Save(&cliMainwith).Error; err != nil {
				return errors.Join(err, errors.New("更新提币总表失败"))
			}
		}
	}

	return nil
}
