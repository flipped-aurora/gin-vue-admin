package xiao

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao/xiaores"
	"gorm.io/gorm"
)

type CliLoadService struct{}

// CreateCliLoad 创建cliLoad表记录
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) CreateCliLoad(cliLoad *xiao.CliLoad) (err error) {
	err = global.GVA_DB.Create(cliLoad).Error
	return err
}

// DeleteCliLoad 删除cliLoad表记录
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) DeleteCliLoad(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliLoad{}, "id = ?", ID).Error
	return err
}

// DeleteCliLoadByIds 批量删除cliLoad表记录
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) DeleteCliLoadByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliLoad{}, "id in ?", IDs).Error
	return err
}

// UpdateCliLoad 更新cliLoad表记录
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) UpdateCliLoad(cliLoad xiao.CliLoad) (err error) {
	err = global.GVA_DB.Model(&xiao.CliLoad{}).Where("id = ?", cliLoad.ID).Updates(&cliLoad).Error
	return err
}

// GetCliLoad 根据ID获取cliLoad表记录
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) GetCliLoad(ID string) (cliLoad xiao.CliLoad, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliLoad).Error
	return
}

// GetCliLoadInfoList 分页获取cliLoad表记录
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) GetCliLoadInfoList(info xiaoReq.CliLoadSearch) (list []xiao.CliLoad, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliLoad{})
	var cliLoads []xiao.CliLoad
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Usdt != nil {
		db = db.Where("usdt > ?", info.Usdt)
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
	orderMap["loadtimes"] = true
	orderMap["lasttime"] = true
	orderMap["usdt"] = true
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

	err = db.Find(&cliLoads).Error
	return cliLoads, total, err
}
func (cliLoadService *CliLoadService) GetCliLoadPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Login 用户登录方法
// Author [yourname](https://github.com/yourname)
func (cliLoadService *CliLoadService) Login(cliLoad *xiao.CliLoad) (resinfo xiaores.CliLoadResponse, err error) {
	if cliLoad.Address == "" {
		return resinfo, errors.New("用户地址为空")
	}
	// 开启事务
	tx := global.GVA_DB.Begin()
	if tx == nil {
		return resinfo, errors.New("failed to start transaction")
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
	// 检查地址是否已经登录
	resinfo.Load, err = cliLoad.CheckAddress(tx)
	//如果没有记录则创建记录
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resinfo.Load = xiao.Newload(cliLoad.Address)
		resinfo.Load.Usdt = cliLoad.Usdt
		resinfo.Load.Loadip = cliLoad.Loadip
		resinfo.Load.Loadaddr = cliLoad.Loadaddr
		err = tx.Create(&resinfo.Load).Error
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return resinfo, err
		}
		err = tx.Commit().Error
		// 设置用户 ID
		resinfo.Load.ID = uint(1)
		return resinfo, err
	}
	// 更新登录信息
	if resinfo.Load != nil {
		*resinfo.Load.Loadtimes++
		resinfo.Load.Usdt = cliLoad.Usdt
		resinfo.Load.Loadip = cliLoad.Loadip
		resinfo.Load.Loadaddr = cliLoad.Loadaddr
		err = tx.Save(&resinfo.Load).Error
		fmt.Printf("更新登录信息:%v\n", resinfo.Load)
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return resinfo, errors.Join(err, errors.New("更新登录信息失败"))
		}

	}
	if resinfo.Load.Desc == "未注册" {
		tx.Commit()
		return resinfo, errors.New("用户未注册")
	}
	// 查询用户信息
	resinfo.User, err = xiao.NewUser(cliLoad.Address).GetCliUser(tx)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return resinfo, errors.Join(err, errors.New("查询用户信息失败"))
	}
	//查询业绩总表
	resinfo.Mainorder, err = xiao.NewCliMainorder(cliLoad.Address).GetCliMainorder(tx)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return resinfo, errors.Join(err, errors.New("查询业绩总表失败"))
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return resinfo, err
	}

	return resinfo, nil
}
