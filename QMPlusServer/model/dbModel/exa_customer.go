package dbModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"github.com/jinzhu/gorm"
)

type ExaCustomer struct {
	gorm.Model
	CustomerName string `json:"customerName"`
	CustomerPhoneData string `json:"customerPhoneData"`
	SysUserID uint `json:"sysUserId"`
	SysUserAuthorityID string `json:"sysUserAuthorityID"`
	SysUser sysModel.SysUser `json:"sysUser"`
}

//创建用户
func (e *ExaCustomer)CreateExaCustomer()(err error){
	err = qmsql.DEFAULTDB.Create(e).Error
	return err
}

//删除用户
func (e *ExaCustomer)DeleteExaCustomer()(err error){
	err = qmsql.DEFAULTDB.Delete(e).Error
	return err
}

//更新用户
func (e *ExaCustomer)UpdateExaCustomer()(err error){
	err = qmsql.DEFAULTDB.Save(e).Error
	return err
}

//获取用户信息
func (e *ExaCustomer)GetExaCustomer()(err error,customer ExaCustomer){
	err= qmsql.DEFAULTDB.Where("id = ?",e.ID).First(&customer).Error
	return
}

//获取用户列表
// 分页获取数据  需要分页实现这个接口即可
func (e *ExaCustomer) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(e, info)
	if err != nil {
		return
	} else {
		var a sysModel.SysAuthority
		a.AuthorityId = e.SysUserAuthorityID
		err,auth := a.GetAuthorityInfo()
		var dataId []string
		for _,v := range auth.DataAuthorityId{
			dataId = append(dataId, v.AuthorityId)
		}
		var CustomerList []ExaCustomer
		err = db.Where("sys_user_authority_id in (?)",dataId).Find(&CustomerList).Count(&total).Error
		if err!=nil{
			return err, CustomerList, total
		}else{
			err = db.Preload("SysUser").Where("sys_user_authority_id in (?)",dataId).Find(&CustomerList).Error
		}
		return err, CustomerList, total
	}
}