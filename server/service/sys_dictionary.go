package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysDictionary
// @description   create a SysDictionary
// @param     sysDictionary               model.SysDictionary
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSysDictionary(sysDictionary model.SysDictionary) (err error) {
	if (!global.GVA_DB.First(&model.SysDictionary{}, "type = ?", sysDictionary.Type).RecordNotFound()) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}

// @title    DeleteSysDictionary
// @description   delete a SysDictionary
// @auth                     （2020/04/05  20:22）
// @param     sysDictionary               model.SysDictionary
// @return                    error

func DeleteSysDictionary(sysDictionary model.SysDictionary) (err error) {
	err = global.GVA_DB.Delete(sysDictionary).Related(&sysDictionary.SysDictionaryDetails).Delete(&sysDictionary.SysDictionaryDetails).Error
	return err
}

// @title    UpdateSysDictionary
// @description   update a SysDictionary
// @param     sysDictionary          *model.SysDictionary
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSysDictionary(sysDictionary *model.SysDictionary) (err error) {
	var dict model.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := global.GVA_DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Type == sysDictionary.Type {
		err = db.Updates(sysDictionaryMap).Error
	} else {
		if (!global.GVA_DB.First(&model.SysDictionary{}, "type = ?", sysDictionary.Type).RecordNotFound()) {
			return errors.New("存在相同的type，不允许创建")
		} else {
			err = db.Updates(sysDictionaryMap).Error
		}
	}
	return err
}

// @title    GetSysDictionary
// @description   get the info of a SysDictionary
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysDictionary        SysDictionary

func GetSysDictionary(Type string, Id uint) (err error, sysDictionary model.SysDictionary) {
	err = global.GVA_DB.Where("type = ? OR id = ?", Type, Id).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	return
}

// @title    GetSysDictionaryInfoList
// @description   get SysDictionary list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysDictionaryInfoList(info request.SysDictionarySearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysDictionary{})
	var sysDictionarys []model.SysDictionary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return err, sysDictionarys, total
}
