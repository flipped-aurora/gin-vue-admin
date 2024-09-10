package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
	"strings"
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

func (apiService *ApiService) GetApiGroups() (groups []string, groupApiMap map[string]string, err error) {
	var apis []system.SysApi
	err = global.GVA_DB.Find(&apis).Error
	if err != nil {
		return
	}
	groupApiMap = make(map[string]string, 0)
	for i := range apis {
		pathArr := strings.Split(apis[i].Path, "/")
		newGroup := true
		for i2 := range groups {
			if groups[i2] == apis[i].ApiGroup {
				newGroup = false
			}
		}
		if newGroup {
			groups = append(groups, apis[i].ApiGroup)
		}
		groupApiMap[pathArr[1]] = apis[i].ApiGroup
	}
	return
}

func (apiService *ApiService) SyncApi() (newApis, deleteApis, ignoreApis []system.SysApi, err error) {
	newApis = make([]system.SysApi, 0)
	deleteApis = make([]system.SysApi, 0)
	ignoreApis = make([]system.SysApi, 0)
	var apis []system.SysApi
	err = global.GVA_DB.Find(&apis).Error
	if err != nil {
		return
	}
	var ignores []system.SysIgnoreApi
	err = global.GVA_DB.Find(&ignores).Error
	if err != nil {
		return
	}

	for i := range ignores {
		ignoreApis = append(ignoreApis, system.SysApi{
			Path:        ignores[i].Path,
			Description: "",
			ApiGroup:    "",
			Method:      ignores[i].Method,
		})
	}

	var cacheApis []system.SysApi
	for i := range global.GVA_ROUTERS {
		ignoresFlag := false
		for j := range ignores {
			if ignores[j].Path == global.GVA_ROUTERS[i].Path && ignores[j].Method == global.GVA_ROUTERS[i].Method {
				ignoresFlag = true
			}
		}
		if !ignoresFlag {
			cacheApis = append(cacheApis, system.SysApi{
				Path:   global.GVA_ROUTERS[i].Path,
				Method: global.GVA_ROUTERS[i].Method,
			})
		}
	}

	//对比数据库中的api和内存中的api，如果数据库中的api不存在于内存中，则把api放入删除数组，如果内存中的api不存在于数据库中，则把api放入新增数组
	for i := range cacheApis {
		var flag bool
		// 如果存在于内存不存在于api数组中
		for j := range apis {
			if cacheApis[i].Path == apis[j].Path && cacheApis[i].Method == apis[j].Method {
				flag = true
			}
		}
		if !flag {
			newApis = append(newApis, system.SysApi{
				Path:        cacheApis[i].Path,
				Description: "",
				ApiGroup:    "",
				Method:      cacheApis[i].Method,
			})
		}
	}

	for i := range apis {
		var flag bool
		// 如果存在于api数组不存在于内存
		for j := range cacheApis {
			if cacheApis[j].Path == apis[i].Path && cacheApis[j].Method == apis[i].Method {
				flag = true
			}
		}
		if !flag {
			deleteApis = append(deleteApis, apis[i])
		}
	}
	return
}

func (apiService *ApiService) IgnoreApi(ignoreApi system.SysIgnoreApi) (err error) {
	if ignoreApi.Flag {
		return global.GVA_DB.Create(&ignoreApi).Error
	}
	return global.GVA_DB.Unscoped().Delete(&ignoreApi, "path = ? AND method = ?", ignoreApi.Path, ignoreApi.Method).Error
}

func (apiService *ApiService) EnterSyncApi(syncApis systemRes.SysSyncApis) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		if syncApis.NewApis != nil && len(syncApis.NewApis) > 0 {
			txErr = tx.Create(&syncApis.NewApis).Error
			if txErr != nil {
				return txErr
			}
		}
		for i := range syncApis.DeleteApis {
			CasbinServiceApp.ClearCasbin(1, syncApis.DeleteApis[i].Path, syncApis.DeleteApis[i].Method)
			txErr = tx.Delete(&system.SysApi{}, "path = ? AND method = ?", syncApis.DeleteApis[i].Path, syncApis.DeleteApis[i].Method).Error
			if txErr != nil {
				return txErr
			}
		}
		return nil
	})
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
	CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
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

func (apiService *ApiService) GetAllApis(authorityID uint) (apis []system.SysApi, err error) {
	parentAuthorityID, err := AuthorityServiceApp.GetParentAuthorityID(authorityID)
	if err != nil {
		return nil, err
	}
	err = global.GVA_DB.Order("id desc").Find(&apis).Error
	if parentAuthorityID == 0 || !global.GVA_CONFIG.System.UseStrictAuth {
		return
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(authorityID)
	// 挑选 apis里面的path和method也在paths里面的api
	var authApis []system.SysApi
	for i := range apis {
		for j := range paths {
			if paths[j].Path == apis[i].Path && paths[j].Method == apis[i].Method {
				authApis = append(authApis, apis[i])
			}
		}
	}
	return authApis, err
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
