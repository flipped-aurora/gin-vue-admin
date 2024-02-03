package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type UserInfoService struct {
}

// CreateUserInfo 创建用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) CreateUserInfo(userData *user.UserInfo) (err error) {
	err = global.GVA_DB.Create(userData).Error
	return err
}

// DeleteUserInfo 删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService)DeleteUserInfo(id string) (err error) {
	err = global.GVA_DB.Delete(&user.UserInfo{},"id = ?",id).Error
	return err
}

// DeleteUserInfoByIds 批量删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService)DeleteUserInfoByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]user.UserInfo{},"id in ?",ids).Error
	return err
}

// UpdateUserInfo 更新用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService)UpdateUserInfo(userData user.UserInfo) (err error) {
	err = global.GVA_DB.Save(&userData).Error
	return err
}

// GetUserInfo 根据id获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService)GetUserInfo(id string) (userData user.UserInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userData).Error
	return
}

// GetUserInfoInfoList 分页获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService)GetUserInfoInfoList(info userReq.UserInfoSearch) (list []user.UserInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.UserInfo{})
    var userDatas []user.UserInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.UserNickname != "" {
        db = db.Where("user_nickname LIKE ?","%"+ info.UserNickname+"%")
    }
    if info.UserGender != nil {
        db = db.Where("user_gender = ?",info.UserGender)
    }
    if info.UserGrade != nil {
        db = db.Where("user_grade = ?",info.UserGrade)
    }
    if info.UserProfession != "" {
        db = db.Where("user_profession LIKE ?","%"+ info.UserProfession+"%")
    }
    if info.UserCity != "" {
        db = db.Where("user_city LIKE ?","%"+ info.UserCity+"%")
    }
    if info.UserModel != nil {
        db = db.Where("user_model = ?",info.UserModel)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&userDatas).Error
	return  userDatas, total, err
}
