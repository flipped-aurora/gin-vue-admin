package xiao

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type CliMainprofitService struct{}

// CreateCliMainprofit 创建结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) CreateCliMainprofit(climainprofit *xiao.CliMainprofit) (err error) {
	err = global.GVA_DB.Create(climainprofit).Error
	return err
}

// DeleteCliMainprofit 删除结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) DeleteCliMainprofit(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliMainprofit{}, "id = ?", ID).Error
	return err
}

// DeleteCliMainprofitByIds 批量删除结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) DeleteCliMainprofitByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliMainprofit{}, "id in ?", IDs).Error
	return err
}

// UpdateCliMainprofit 更新结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) UpdateCliMainprofit(climainprofit xiao.CliMainprofit) (err error) {
	err = global.GVA_DB.Model(&xiao.CliMainprofit{}).Where("id = ?", climainprofit.ID).Updates(&climainprofit).Error
	return err
}

// GetCliMainprofit 根据ID获取结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) GetCliMainprofit(ID string) (climainprofit xiao.CliMainprofit, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&climainprofit).Error
	return
}

// GetCliMainprofitInfoList 分页获取结算总表记录
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) GetCliMainprofitInfoList(info xiaoReq.CliMainprofitSearch) (list []xiao.CliMainprofit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliMainprofit{})
	var climainprofits []xiao.CliMainprofit
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
	orderMap["static"] = true
	orderMap["indirect"] = true
	orderMap["team"] = true
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

	err = db.Find(&climainprofits).Error
	return climainprofits, total, err
}
func (climainprofitService *CliMainprofitService) GetCliMainprofitPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// PullProfit 推荐结算
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) PullProfit() (err error) {
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

	// 查询设置信息
	var clisetting xiao.CliSet
	if err = tx.Where("id = ?", 1).First(&clisetting).Error; err != nil {
		return errors.Join(err, errors.New("查询设置失败"))
	}

	// 查询订单列表
	var cliorders []*xiao.CliOrder
	if err = tx.Where("status = ?", "正常").Find(&cliorders).Error; err != nil {
		return errors.Join(err, errors.New("查询订单失败"))
	}

	for _, cliorder := range cliorders {
		if err = processOrder(tx, cliorder, clisetting); err != nil {
			continue
		}
	}

	return nil
}

func processOrder(tx *gorm.DB, cliorder *xiao.CliOrder, clisetting xiao.CliSet) error {
	// 计算直推收益
	pullprofit := cliorder.Todaynum.Mul(*clisetting.Straight).Div(decimal.NewFromInt(100))

	// 查找上级用户
	var user xiao.CliUser
	if err := tx.Where("address = ?", cliorder.Address).First(&user).Error; err != nil {
		return errors.Join(err, errors.New("查询用户失败"))
	}
	if user.Parent == "" || user.Parent == "root" {
		return errors.Join(errors.New("查询上级用户失败"))
	}
	//查询订单总表
	mainorder, err := xiao.NewCliMainorder(user.Parent).GetCliMainorder(tx)
	if err != nil {
		return err
	}
	if *mainorder.Num == 0 {
		return nil
	}

	// 创建直推收益结算记录
	if err := createProfitRecord(tx, user.Parent, pullprofit, "直推收益"); err != nil {
		return err
	}

	// 更新结算总表
	if err := updateMainprofit(tx, user.Parent, pullprofit, "Pull"); err != nil {
		return err
	}

	// 更新提币总表
	if err := updateMainwith(tx, user.Parent, pullprofit); err != nil {
		return err
	}

	// 查询上级的上级
	var userup xiao.CliUser
	if err := tx.Where("address = ?", user.Parent).First(&userup).Error; err != nil {
		return errors.Join(err, errors.New("查询用户失败"))
	}
	if userup.Parent == "" || userup.Parent == "root" {
		return errors.Join(errors.New("查询上上级用户失败"))
	}

	// 计算间推收益
	indirectprofit := cliorder.Todaynum.Mul(*clisetting.Inderict).Div(decimal.NewFromInt(100))

	// 创建间推收益结算记录
	if err := createProfitRecord(tx, userup.Parent, indirectprofit, "间推收益"); err != nil {
		return err
	}

	// 更新结算总表
	if err := updateMainprofit(tx, userup.Parent, indirectprofit, "Indirect"); err != nil {
		return err
	}

	// 更新提币总表
	if err := updateMainwith(tx, userup.Parent, indirectprofit); err != nil {
		return err
	}

	return nil
}

func createProfitRecord(tx *gorm.DB, address string, amount decimal.Decimal, text string) error {
	cliprofit := xiao.CliProfit{
		Address: address,
		Amount:  &amount,
		Text:    text,
	}
	if err := tx.Create(&cliprofit).Error; err != nil {
		return errors.Join(err, errors.New("创建结算记录失败"))
	}
	return nil
}

func updateMainprofit(tx *gorm.DB, address string, amount decimal.Decimal, field string) error {
	var mainpro *xiao.CliMainprofit
	var err error

	switch field {
	case "Pull":
		mainpro, err = xiao.NewCliMainprofit(address).GetCliMainprofitAddress(tx)
		if err != nil {
			return errors.Join(err, errors.New("查询结算总表失败"))
		}
		*mainpro.Pull = mainpro.Pull.Add(amount)
		*mainpro.Amount = mainpro.Amount.Add(amount)
	case "Indirect":
		mainpro, err = xiao.NewCliMainprofit(address).GetCliMainprofitAddress(tx)
		if err != nil {
			return errors.Join(err, errors.New("查询结算总表失败"))
		}
		*mainpro.Indirect = mainpro.Indirect.Add(amount)
		*mainpro.Amount = mainpro.Amount.Add(amount)
	case "Team":
		mainpro, err = xiao.NewCliMainprofit(address).GetCliMainprofitAddress(tx)
		if err != nil {
			return errors.Join(err, errors.New("查询结算总表失败"))
		}
		*mainpro.Team = mainpro.Team.Add(amount)
		*mainpro.Amount = mainpro.Amount.Add(amount)
	default:
		return errors.New("未知字段")
	}

	if err := tx.Save(mainpro).Error; err != nil {
		return errors.Join(err, errors.New("更新结算总表失败"))
	}
	return nil
}

func updateMainwith(tx *gorm.DB, address string, amount decimal.Decimal) error {
	mainwith, err := xiao.NewCliMainwith(address).GetCliMainwith(tx)
	if err != nil {
		return errors.Join(err, errors.New("查询提币总表失败"))
	}
	*mainwith.Withable = mainwith.Withable.Add(amount)
	if err := tx.Save(mainwith).Error; err != nil {
		return errors.Join(err, errors.New("更新提币总表失败"))
	}
	return nil
}

// TeamProfit 团队结算
// Author [yourname](https://github.com/yourname)
func (climainprofitService *CliMainprofitService) TeamProfit() (err error) {
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("transaction failed: %v", r)
			log.Println("Transaction failed:", r) // 记录日志
		} else if err != nil {
			tx.Rollback()
		}
	}()

	// 查询设置信息
	var clisetting []xiao.CliSetvip
	if err = tx.Find(&clisetting).Error; err != nil {
		return errors.New("查询设置失败: " + err.Error())
	}

	// 查询订单表
	var orders []*xiao.CliOrder
	if err = tx.Where("status = ?", "正常").Find(&orders).Error; err != nil {
		return errors.New("查询订单失败: " + err.Error())
	}

	for _, order := range orders {
		if order.Address == "" || order.Address == "root" {
			continue
		}

		// 查询所有上级
		uptrees, err := xiao.NewCliTree(order.Address, "").GetAllParentNode(tx)
		if err != nil {
			log.Println("查询上级节点失败:", err) // 记录日志
			continue
		}

		for _, uptree := range uptrees {
			mainorder, err := xiao.NewCliMainorder(uptree.Address).GetCliMainorder(tx)
			if err != nil {
				return errors.New("获取主订单失败: " + err.Error())
			}

			var rates decimal.Decimal
			var rateIndex int
			if mainorder.Desc != "备注" {
				rateIndex, err = strconv.Atoi(mainorder.Desc)
				if err != nil {
					log.Println("转换备注为整数失败:", err) // 记录日志
					continue
				}
			}

			for i := 0; i < len(clisetting)-1; i++ {
				if mainorder.Amount.Cmp(*clisetting[i].Come) >= 0 && mainorder.Amount.Cmp(*clisetting[i+1].Come) < 0 && mainorder.Descnum.Cmp(decimal.NewFromInt(1000)) >= 0 && rates.IsZero() {
					fmt.Printf("订单ID: %d, 金额: %s, 符合设置项 %d 的条件\n", mainorder.ID, mainorder.Amount.String(), i+1)
					rates = *clisetting[i].Rate
					break
				}
			}

			// 检查最后一个设置项
			if mainorder.Amount.Cmp(*clisetting[len(clisetting)-1].Come) >= 0 {
				fmt.Printf("订单ID: %d, 金额: %s, 符合设置项 %d 的条件\n", mainorder.ID, mainorder.Amount.String(), len(clisetting))
				rates = *clisetting[len(clisetting)-1].Rate
			}

			if rateIndex > 0 {
				rates = *clisetting[rateIndex-1].Rate
			}

			if rates.IsZero() {
				continue
			}

			teamprofit := order.Todaynum.Mul(rates).Div(decimal.NewFromInt(100))

			// 成交收益详情
			profit := xiao.CliProfit{
				Address: uptree.Address,
				Amount:  &teamprofit,
				Text:    "团队收益",
			}
			if err := tx.Create(&profit).Error; err != nil {
				return errors.New("创建收益记录失败: " + err.Error())
			}

			// 更新收益总表
			cliMainprofit, err := xiao.NewCliMainprofit(uptree.Address).GetCliMainprofitAddress(tx)
			if err != nil {
				return errors.New("查询收益总表失败: " + err.Error())
			}

			*cliMainprofit.Team = cliMainprofit.Team.Add(teamprofit)
			*cliMainprofit.Amount = cliMainprofit.Amount.Add(teamprofit)
			if err := tx.Save(&cliMainprofit).Error; err != nil {
				return errors.New("更新收益总表失败: " + err.Error())
			}

			// 更新提币总表
			cliMainwith, err := xiao.NewCliMainwith(uptree.Address).GetCliMainwith(tx)
			if err != nil {
				return errors.New("查询提币总表失败: " + err.Error())
			}
			if cliMainwith != nil {
				*cliMainwith.Withable = cliMainwith.Withable.Add(teamprofit)
				if err := tx.Save(&cliMainwith).Error; err != nil {
					return errors.New("更新提币总表失败: " + err.Error())
				}
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("提交事务失败: " + err.Error())
	}

	return nil
}

// 假设的计算团队收益的函数
