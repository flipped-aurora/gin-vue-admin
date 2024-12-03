package xiao

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/shopspring/decimal"
)

type CliMainorderService struct{}

// CreateCliMainorder 创建订单总表记录
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) CreateCliMainorder(climainorder *xiao.CliMainorder) (err error) {
	err = global.GVA_DB.Create(climainorder).Error
	return err
}

// DeleteCliMainorder 删除订单总表记录
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) DeleteCliMainorder(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliMainorder{}, "id = ?", ID).Error
	return err
}

// DeleteCliMainorderByIds 批量删除订单总表记录
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) DeleteCliMainorderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliMainorder{}, "id in ?", IDs).Error
	return err
}

// UpdateCliMainorder 更新订单总表记录
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) UpdateCliMainorder(climainorder xiao.CliMainorder) (err error) {
	err = global.GVA_DB.Model(&xiao.CliMainorder{}).Where("id = ?", climainorder.ID).Updates(&climainorder).Error
	return err
}

// GetCliMainorder 根据ID获取订单总表记录
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) GetCliMainorder(ID string) (climainorder xiao.CliMainorder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&climainorder).Error
	return
}

// GetCliMainorderInfoList 分页获取订单总表记录
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) GetCliMainorderInfoList(info xiaoReq.CliMainorderSearch) (list []xiao.CliMainorder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliMainorder{})
	var climainorders []xiao.CliMainorder
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
	orderMap["amount"] = true
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

	err = db.Find(&climainorders).Error
	return climainorders, total, err
}
func (climainorderService *CliMainorderService) GetCliMainorderPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Buy 客户端购买方法
// Author [yourname](https://github.com/yourname)
func (climainorderService *CliMainorderService) Buy(cliorder *xiao.CliOrder) (err error) {
	// 请在这里实现自己的业务逻辑
	tx := global.GVA_DB.Begin()
	// 查询所有上级
	upnodes, err := xiao.NewCliTree(cliorder.Address, "").GetAllParentNode(tx)
	if err != nil {
		return err
	}
	padiannum := decimal.NewFromInt(10)
	for _, upnode := range upnodes {
		// 获取上级的订单主表
		mainorder, err := xiao.NewCliMainorder(upnode.Address).GetCliMainorder(tx)
		if err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("获取上级订单失败"))
		}
		*mainorder.Num++
		*mainorder.Amount = mainorder.Amount.Add(*cliorder.Amount)
		tx.Save(&mainorder)

		//查询上级是否可以结算怕点
		if mainorder.Padian.Cmp(decimal.NewFromInt(0)) > 0 {
			if padiannum.Cmp(*mainorder.Padian) < 0 {
				continue
			}
			padiannum = padiannum.Sub(*mainorder.Padian)
			pronum := mainorder.Padian.Div(decimal.NewFromInt(100)).Mul(*cliorder.Amount)
			//创建个人收益结算记录
			profit := xiao.CliProfit{
				Address: upnode.Address,
				Amount:  &pronum,
				Text:    "团队入金奖励",
			}
			tx.Create(&profit)
			//更新个人总表
			cliMainprofit, err := xiao.NewCliMainprofit(upnode.Address).GetCliMainprofitAddress(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			if cliMainprofit != nil {
				*cliMainprofit.Static = cliMainprofit.Static.Add(pronum)
				*cliMainprofit.Amount = cliMainprofit.Amount.Add(pronum)
				tx.Save(&cliMainprofit)
			}
			//更新个人提币总表
			cliMainwith, err := xiao.NewCliMainwith(upnode.Address).GetCliMainwith(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			if cliMainwith != nil {
				*cliMainwith.Withable = cliMainwith.Withable.Add(pronum)
				tx.Save(&cliMainwith)
			}

		}

	}
	//更新个人订单总表
	mymainorder, err := xiao.NewCliMainorder(cliorder.Address).GetCliMainorder(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	*mymainorder.Descnum = mymainorder.Descnum.Add(*cliorder.Amount)
	tx.Save(&mymainorder)
	//创建个人订单详情
	tx.Create(&cliorder)

	return tx.Commit().Error
}
