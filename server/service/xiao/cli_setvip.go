package xiao

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/online"
	teleclient2 "github.com/flipped-aurora/gin-vue-admin/server/plugin/teleclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type CliSetvipService struct{}

// CreateCliSetvip 创建团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) CreateCliSetvip(clisetvip *xiao.CliSetvip) (err error) {
	err = global.GVA_DB.Create(clisetvip).Error
	return err
}

// DeleteCliSetvip 删除团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) DeleteCliSetvip(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliSetvip{}, "id = ?", ID).Error
	return err
}

// DeleteCliSetvipByIds 批量删除团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) DeleteCliSetvipByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliSetvip{}, "id in ?", IDs).Error
	return err
}

// UpdateCliSetvip 更新团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) UpdateCliSetvip(clisetvip xiao.CliSetvip) (err error) {
	err = global.GVA_DB.Model(&xiao.CliSetvip{}).Where("id = ?", clisetvip.ID).Updates(&clisetvip).Error
	return err
}

// GetCliSetvip 根据ID获取团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) GetCliSetvip(ID string) (clisetvip xiao.CliSetvip, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&clisetvip).Error
	return
}

// GetCliSetvipInfoList 分页获取团队设置记录
// Author [yourname](https://github.com/yourname)
func (clisetvipService *CliSetvipService) GetCliSetvipInfoList(info xiaoReq.CliSetvipSearch) (list []xiao.CliSetvip, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliSetvip{})
	var clisetvips []xiao.CliSetvip
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&clisetvips).Error
	return clisetvips, total, err
}
func (clisetvipService *CliSetvipService) GetCliSetvipPublic() (err error) {
	// 此方法为获取数据源定义的数据
	const usdtaddress = "0x55d398326f99059fF775485246999027B3197955"

	// 获取 BNB 客户端
	bnbclient, err := online.NewClient()
	if err != nil {
		return errors.Join(err, errors.New("获取 bnbclient 失败"))
	}

	// 获取 Telegram 客户端
	teleclient, err := teleclient2.Telecli()
	if err != nil {
		return errors.Join(err, errors.New("获取 teleclient 失败"))
	}

	// 查询所有用户表
	var allusers []xiao.CliUser
	if err := global.GVA_DB.Find(&allusers).Error; err != nil {
		return errors.Join(err, errors.New("查询用户表失败"))
	}

	// 存储符合条件的地址和金额
	var qualifiedUsers []string
	var totalAmount decimal.Decimal

	for i, user := range allusers {
		bnb20balance, err := online.Bsc20Balance(bnbclient, usdtaddress, user.Address)
		if err != nil {
			return errors.Join(err, errors.New("查询 bnb20 余额失败"))
		}
		if bnb20balance.Cmp(decimal.NewFromInt(100)) > 0 {
			// 取整数
			intBalance := bnb20balance.Floor()
			// 将符合条件的地址和金额添加到列表中
			qualifiedUsers = append(qualifiedUsers, fmt.Sprintf("%d. 地址：%s\n金额：**%s**\n", i+1, user.Address, intBalance.String()))
			// 累加总金额
			totalAmount = totalAmount.Add(intBalance)
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 构建消息
	var messageBuilder strings.Builder
	messageBuilder.WriteString(fmt.Sprintf("萧敬腾机器人发来消息\n"))
	messageBuilder.WriteString(fmt.Sprintf("符合条件的地址数量：%d\n", len(qualifiedUsers)))
	messageBuilder.WriteString(fmt.Sprintf("总金额：%s\n", totalAmount.String()))
	for _, user := range qualifiedUsers {
		messageBuilder.WriteString(user)
		messageBuilder.WriteString("\n")
	}

	// 发送消息
	msg := tgbotapi.NewMessage(-1002375968132, messageBuilder.String())
	_, err = teleclient.Send(msg)
	if err != nil {
		return errors.Join(err, errors.New("发送失败"))
	}
	return nil
}
