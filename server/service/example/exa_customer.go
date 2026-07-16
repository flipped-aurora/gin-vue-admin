package example

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type CustomerService struct{}

var CustomerServiceApp = new(CustomerService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) CreateExaCustomer(ctx context.Context, e example.ExaCustomer) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) DeleteExaCustomer(ctx context.Context, e example.ExaCustomer) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (exa *CustomerService) UpdateExaCustomer(ctx context.Context, e *example.ExaCustomer) (err error) {
	// 归属列(dept_id/created_by)创建时盖章后即不可变, 更新时忽略, 防止被表单未回传的零值覆盖
	err = global.GVA_DB.WithContext(ctx).Omit("dept_id", "created_by").Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: customer model.ExaCustomer, err error

func (exa *CustomerService) GetExaCustomer(ctx context.Context, id uint) (customer example.ExaCustomer, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *CustomerService) GetCustomerInfoList(ctx context.Context, sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	_ = sysUserAuthorityID // 数据范围已由数据权限引擎接管, 旧的按角色 DataAuthorityId 手写过滤已移除
	limit, offset := info.LimitOffset()
	// 数据范围(本部门及子级/仅本人)由数据权限引擎的 GORM 回调按当前角色 data_scope
	// 自动追加 dept_id/created_by 过滤, 此处只写常规分页查询即可。
	db := global.GVA_DB.WithContext(ctx).Model(&example.ExaCustomer{})
	var CustomerList []example.ExaCustomer
	if err = db.Count(&total).Error; err != nil {
		return CustomerList, total, err
	}
	err = db.Limit(limit).Offset(offset).Preload("SysUser").Find(&CustomerList).Error
	return CustomerList, total, err
}
