package xiao

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/myutil"
)

type CliUserService struct{}

// CreateCliUser 创建cliUser表记录
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) CreateCliUser(cliUser *xiao.CliUser) (err error) {
	err = global.GVA_DB.Create(cliUser).Error
	return err
}

// DeleteCliUser 删除cliUser表记录
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) DeleteCliUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&xiao.CliUser{}, "id = ?", ID).Error
	return err
}

// DeleteCliUserByIds 批量删除cliUser表记录
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) DeleteCliUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]xiao.CliUser{}, "id in ?", IDs).Error
	return err
}

// UpdateCliUser 更新cliUser表记录
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) UpdateCliUser(cliUser xiao.CliUser) (err error) {
	err = global.GVA_DB.Model(&xiao.CliUser{}).Where("id = ?", cliUser.ID).Updates(&cliUser).Error
	return err
}

// GetCliUser 根据ID获取cliUser表记录
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) GetCliUser(ID string) (cliUser xiao.CliUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cliUser).Error
	return
}

// GetCliUserInfoList 分页获取cliUser表记录
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) GetCliUserInfoList(info xiaoReq.CliUserSearch) (list []xiao.CliUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&xiao.CliUser{})
	var cliUsers []xiao.CliUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Parent != "" {
		db = db.Where("parent LIKE ?", "%"+info.Parent+"%")
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
	orderMap["pullnum"] = true
	orderMap["teamnum"] = true
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

	err = db.Find(&cliUsers).Error
	return cliUsers, total, err
}
func (cliUserService *CliUserService) GetCliUserPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Rigester 客户端注册方法
// Author [yourname](https://github.com/yourname)
func (cliUserService *CliUserService) Rigester(cliUser *xiao.CliUser) (err error) {

	// 请在这里实现自己的业务逻辑
	tx := global.GVA_DB.Begin()
	//1.创建注册表 2.创建用户关系表 3.创建订单总表 4.创建结算总表 5.创建提币总表
	//更新登录表
	loadinfo, err := xiao.Newload(cliUser.Address).CheckAddress(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if loadinfo.Desc == "注册成功" {
		tx.Rollback()
		return errors.New("用户已注册")
	}
	loadinfo.Usdt = cliUser.Desnum
	*loadinfo.Loadtimes++
	loadinfo.Desc = "注册成功"
	err = tx.Save(&loadinfo).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.Join(err, errors.New("更新登录信息失败"))
	}
	//创建注册表
	err = tx.Create(&cliUser).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.Join(err, errors.New("创建注册表失败"))
	}
	//创建订单总表
	cliOrder := xiao.NewCliMainorder(cliUser.Address)
	err = tx.Create(&cliOrder).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.Join(err, errors.New("创建订单总表失败"))
	}
	//创建结算总表
	climainpro := xiao.NewCliMainprofit(cliUser.Address)
	err = tx.Create(&climainpro).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.Join(err, errors.New("创建结算总表失败"))
	}
	//创建提币总表
	cliWithdraw := xiao.NewCliMainwith(cliUser.Address)
	err = tx.Create(&cliWithdraw).Error
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return errors.Join(err, errors.New("创建提币总表失败"))
	}
	//更新上级
	parentuser, err := xiao.NewUser(cliUser.Parent).GetCliUser(tx)
	if err != nil {
		tx.Rollback()
		return errors.Join(err, errors.New("上级不存在"))
	}
	*parentuser.Pullnum++
	parentuser.Mypull, err = myutil.ArryjsonAdd(parentuser.Mypull, cliUser.Address)
	if err != nil {
		tx.Rollback()
		return errors.Join(err, errors.New("更新上级失败"))
	}
	tx.Save(parentuser)
	//创建用户关系表
	err = xiao.NewCliTree(cliUser.Address, cliUser.Parent).CreateNode(tx)
	if err != nil {
		tx.Rollback()
		return errors.Join(err, errors.New("创建用户关系表失败"))
	}
	//查询所有上级
	upnodes, err := xiao.NewCliTree(cliUser.Address, "").GetAllParentNode(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, upnode := range upnodes {
		// 更新团队成员
		userinfo, err := xiao.NewUser(upnode.Address).GetCliUser(tx)
		if err != nil {
			return err
		}
		*userinfo.Teamnum++
		tx.Save(&userinfo)
	}

	return tx.Commit().Error
}
