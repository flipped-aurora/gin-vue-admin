package system

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	err = global.GVA_DB.First(&entity, "id = ?", api.ID).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	if !CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method) {
		return errors.New("ClearCasbin 失败")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: list interface{}, total int64, err error

func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	}

	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	if order != "" {
		orderMap := make(map[string]bool, 5)
		orderMap["id"] = true
		orderMap["path"] = true
		orderMap["api_group"] = true
		orderMap["description"] = true
		orderMap["method"] = true
		if !orderMap[order] {
			err = fmt.Errorf("非法的排序字段: %v", order)
			return apiList, total, err
		}
		OrderStr = order
		if desc {
			OrderStr = order + " desc"
		}
	}
	err = db.Order(OrderStr).Find(&apiList).Error
	return apiList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return:  apis []model.SysApi, err error

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: api model.SysApi, err error

func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	err = global.GVA_DB.First(&api, "id = ?", id).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GVA_DB.First(&oldA, "id = ?", api.ID).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		var duplicateApi system.SysApi
		if ferr := global.GVA_DB.First(&duplicateApi, "path = ? AND method = ?", api.Path, api.Method).Error; ferr != nil {
			if !errors.Is(ferr, gorm.ErrRecordNotFound) {
				return ferr
			}
		} else {
			if duplicateApi.ID != api.ID {
				return errors.New("存在相同api路径")
			}
		}

	}
	if err != nil {
		return err
	}

	err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
	if err != nil {
		return err
	}

	return global.GVA_DB.Save(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApisByIds
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []system.SysApi
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		for _, sysApi := range apis {
			CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
		}
		return err
	})
}
