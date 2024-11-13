package xiao

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao/xiaores"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/online"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/wallet"
	"github.com/shopspring/decimal"
)

type CliWithdrawService struct{}

// CreateCliWithdraw 创建提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) CreateCliWithdraw(cliwithdraw *xiao.CliWithdraw) (err error) {
	err = global.GVA_DB.Create(cliwithdraw).Error
	return err
}

// DeleteCliWithdraw 删除提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) DeleteCliWithdraw(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliWithdraw{}, "id = ?", ID).Error
	return err
}

// DeleteCliWithdrawByIds 批量删除提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) DeleteCliWithdrawByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliWithdraw{}, "id in ?", IDs).Error
	return err
}

// UpdateCliWithdraw 更新提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) UpdateCliWithdraw(cliwithdraw xiao.CliWithdraw) (err error) {
	err = global.GVA_DB.Model(&xiao.CliWithdraw{}).Where("id = ?", cliwithdraw.ID).Updates(&cliwithdraw).Error
	return err
}

// GetCliWithdraw 根据ID获取提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) GetCliWithdraw(ID string) (cliwithdraw xiao.CliWithdraw, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliwithdraw).Error
	return
}

// GetCliWithdrawInfoList 分页获取提币详情记录
// Author [yourname](https://github.com/yourname)
func (cliwithdrawService *CliWithdrawService) GetCliWithdrawInfoList(info xiaoReq.CliWithdrawSearch) (list []xiao.CliWithdraw, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliWithdraw{})
	var cliwithdraws []xiao.CliWithdraw
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Text != "" {
		db = db.Where("text LIKE ?", "%"+info.Text+"%")
	}
	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
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
	orderMap["amount"] = true
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

	err = db.Find(&cliwithdraws).Error
	return cliwithdraws, total, err
}
func (cliwithdrawService *CliWithdrawService) GetCliWithdrawPublic(pageInfo *xiaoReq.CliWithdrawSearch) (res xiaores.CliWithdrawRes, err error) {
	// 此方法为获取数据源定义的数据
	// 开始事务
	tx := global.GVA_DB.Begin()

	if tx == nil {
		return res, errors.New("failed to start transaction")
	}

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

	// 查询提币总表
	maininfo, err := xiao.NewCliMainwith(pageInfo.Address).GetCliMainwith(tx)
	if err != nil {
		return res, errors.Join(err, errors.New("查询提币总表失败"))
	}

	if maininfo != nil {
		res.CliMainwiths = maininfo
	}

	// 查询提币记录
	withinfo, err := xiao.NewCliWithdraw(pageInfo.Address).GetCliWithdraw(tx)
	if err != nil {
		return res, errors.Join(err, errors.New("查询提币记录失败"))
	}
	if withinfo != nil {
		res.CliWithdraws = withinfo
	}

	return res, nil
}

// CliWithdraw
func (cliwithdrawService *CliWithdrawService) CliWithdraw(cliwithdraw *xiao.CliWithdraw) (err error) {
	//开启事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//计算手续费
	var setinfo *xiao.CliSet
	err = tx.First(&setinfo).Error
	if err != nil {
		return errors.Join(err, errors.New("查询设置表失败"))
	}
	fee := cliwithdraw.Amount.Mul(*setinfo.Fee).Div(decimal.NewFromInt(100))
	if cliwithdraw.Desc == "赎回" {
		//计算提币金额=投入资金-提币总额
		//查询投入资金
		var maininfo xiao.CliMainorder
		if err := tx.Raw(`SELECT * FROM cli_mainorder WHERE address = ? FOR UPDATE`, cliwithdraw.Address).Scan(&maininfo).Error; err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("查询投入资金失败"))
		}
		//查询提币总额
		var withedinfo xiao.CliMainwith
		if err := tx.Raw(`SELECT * FROM cli_mainwith WHERE address = ? FOR UPDATE`, cliwithdraw.Address).Scan(&withedinfo).Error; err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("查询提币总额失败"))
		}
		nowtiqu := maininfo.Descnum.Mul(decimal.NewFromFloat(0.8)).Sub(*withedinfo.Withed)
		if nowtiqu.Cmp(decimal.NewFromInt(0)) < 0 {
			tx.Rollback()
			return errors.New("可提金额为0")
		}
		//更新个人业绩
		myselfnum := decimal.NewFromInt(0)
		maininfo.Descnum = &myselfnum
		if err := tx.Save(maininfo).Error; err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("更新投入资金失败"))
		}
		//更新个人收益
		withedinfo.Withable = &myselfnum
		*withedinfo.Withed = withedinfo.Withed.Add(nowtiqu)
		if err := tx.Save(withedinfo).Error; err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("更新个人收益失败"))
		}
		//提币方法
		truenum := nowtiqu.Sub(fee)
		clienteth, err := online.NewClient()
		if err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("连接以太坊节点失败"))
		}
		//0x79D954564b77C9550327B3e11cFe31472bc1e0d0
		store, err := wallet.ImportKeyStore("0xa4416da064bbee02d0f71130cf1414279a32280a")
		if err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("导入私钥失败"))
		}
		transfer, err := online.Bsc20Transfer(clienteth, store, "0x55d398326f99059fF775485246999027B3197955", cliwithdraw.Address, &truenum)
		if err != nil {
			tx.Rollback()
			return errors.Join(err, errors.New("转账失败"))
		}

		//创建提币记录
		cliwithdraw.Amount = &nowtiqu
		cliwithdraw.Status = "赎回本金"
		cliwithdraw.Desc = transfer
		cliwithdraw.Descnum = &fee
		err = tx.Create(&cliwithdraw).Error
		if err != nil {
			// 回滚事务
			return errors.Join(err, errors.New("创建提币记录失败"))
		}
		err = tx.Commit().Error
		if err != nil {
			return errors.Join(err, errors.New("提交事务失败"))
		}
		return nil

	}
	// 查询提币总表并加锁
	var maininfo xiao.CliMainwith
	if err := tx.Raw(`SELECT * FROM cli_mainwith WHERE address = ? FOR UPDATE`, cliwithdraw.Address).Scan(&maininfo).Error; err != nil {
		return errors.Join(err, errors.New("查询提币总表失败"))
	}
	//查询提币总表
	//maininfo, err := xiao.NewCliMainwith(cliwithdraw.Address).GetCliMainwith(tx)
	//if err != nil {
	//	return errors.Join(err, errors.New("查询提币总表失败"))
	//}
	if maininfo.Withable.Cmp(*cliwithdraw.Amount) < 0 {
		return errors.New("提币金额大于可提金额")
	}
	*maininfo.Withable = maininfo.Withable.Sub(*cliwithdraw.Amount)
	*maininfo.Withed = maininfo.Withed.Add(*cliwithdraw.Amount)
	if err := tx.Save(maininfo).Error; err != nil {
		return errors.Join(err, errors.New("更新提币总表失败"))
	}

	truenum := cliwithdraw.Amount.Sub(fee)

	//提币方法
	clienteth, err := online.NewClient()
	if err != nil {
		return errors.Join(err, errors.New("连接以太坊节点失败"))
	}
	//0x79D954564b77C9550327B3e11cFe31472bc1e0d0
	store, err := wallet.ImportKeyStore("0xa4416da064bbee02d0f71130cf1414279a32280a")
	if err != nil {
		return errors.Join(err, errors.New("导入私钥失败"))
	}
	transfer, err := online.Bsc20Transfer(clienteth, store, "0x55d398326f99059fF775485246999027B3197955", cliwithdraw.Address, &truenum)
	if err != nil {
		return errors.Join(err, errors.New("转账失败"))
	}
	//创建提币记录
	cliwithdraw.Status = "自动发放"
	cliwithdraw.Desc = transfer
	cliwithdraw.Descnum = &fee
	err = tx.Create(&cliwithdraw).Error
	if err != nil {
		// 回滚事务
		return errors.Join(err, errors.New("创建提币记录失败"))
	}
	err = tx.Commit().Error
	if err != nil {
		return errors.Join(err, errors.New("提交事务失败"))
	}
	return nil
}
