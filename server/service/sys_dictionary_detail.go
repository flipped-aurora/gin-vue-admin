package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysDictionaryDetail
// @description   create a SysDictionaryDetail
// @param     sysDictionaryDetail               model.SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSysDictionaryDetail(sysDictionaryDetail model.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

// @title    DeleteSysDictionaryDetail
// @description   delete a SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @param     sysDictionaryDetail               model.SysDictionaryDetail
// @return                    error

func DeleteSysDictionaryDetail(sysDictionaryDetail model.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(sysDictionaryDetail).Error
	return err
}

// @title    UpdateSysDictionaryDetail
// @description   update a SysDictionaryDetail
// @param     sysDictionaryDetail          *model.SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSysDictionaryDetail(sysDictionaryDetail *model.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	return err
}

// @title    GetSysDictionaryDetail
// @description   get the info of a SysDictionaryDetail
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysDictionaryDetail        SysDictionaryDetail

func GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail model.SysDictionaryDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

// @title    GetSysDictionaryDetailInfoList
// @description   get SysDictionaryDetail list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysDictionaryDetail{})
	var sysDictionaryDetails []model.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err, sysDictionaryDetails, total
}
