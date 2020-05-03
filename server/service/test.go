package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTestTest
// @description   create a TestTest
// @param     t               model.TestTest
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateTestTest(t model.TestTest) (err error) {
	err = global.GVA_DB.Create(&t).Error
	return err
}

// @title    DeleteTestTest
// @description   delete a TestTest
// @auth                     （2020/04/05  20:22）
// @param     t               model.TestTest
// @return                    error

func DeleteTestTest(t model.TestTest) (err error) {
	err = global.GVA_DB.Delete(t).Error
	return err
}

// @title    UpdateTestTest
// @description   update a TestTest
// @param     t          *model.TestTest
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateTestTest(t *model.TestTest) (err error) {
	err = global.GVA_DB.Save(t).Error
	return err
}

// @title    GetTestTest
// @description   get the info of a TestTest
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    TestTest        TestTest

func GetTestTest(id uint) (err error, t model.TestTest) {
	err = global.GVA_DB.Where("id = ?", id).First(&t).Error
	return
}

// @title    GetTestTestInfoList
// @description   get TestTest list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetTestTestInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
    	offset := info.PageSize * (info.Page - 1)
    	db := global.GVA_DB
    	var ts []model.TestTest
    	err = db.Find(&ts).Count(&total).Error
    	err = db.Limit(limit).Offset(offset).Find(&ts).Error
    	return err,ts, total
}
