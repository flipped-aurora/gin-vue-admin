package user

import (
	"app/model/user"
	"app/mysqlDB"
)

type UserInfoService struct {
}

// CreateUserInfo 创建用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) CreateUserInfo(userData *user.UserInfo) (err error) {
	err = mysqlDB.DB.Create(userData).Error
	return err
}

// DeleteUserInfo 删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) DeleteUserInfo(id string) (err error) {
	err = mysqlDB.DB.Model(user.UserInfo{}).Where("id = ?", id).Error
	if err != nil {
		return err
	}
	err = mysqlDB.DB.Delete(&user.UserInfo{}, "id = ?", id).Error
	return err
}

// DeleteUserInfoByIds 批量删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) DeleteUserInfoByIds(ids []string) (err error) {
	err = mysqlDB.DB.Delete(&[]user.UserInfo{}, "id in ?", ids).Error
	return err
}

// UpdateUserInfo 更新用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) UpdateUserInfo(userData user.UserInfo) (err error) {
	err = mysqlDB.DB.Updates(&userData).Error
	return err
}

// GetUserInfo 根据id获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) GetUserInfo(id string) (userData user.UserInfo, err error) {
	err = mysqlDB.DB.Where("id = ?", id).First(&userData).Error
	return
}

// GetUserInfo 根据wxOpenid获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) GetUserInfoByWxOpenid(wxopenid string) (userData user.UserInfo, err error) {
	err = mysqlDB.DB.Where("user_wxopenid = ?", wxopenid).First(&userData).Error
	return
}

// GetUserInfoInfoList 分页获取用户信息记录
// // Author [piexlmax](https://github.com/piexlmax)
func (userDataService *UserInfoService) GetUserInfoInfoList(info user.UserInfoSearch) (list []user.UserInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := mysqlDB.DB.Model(&user.UserInfo{})
	var userDatas []user.UserInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserNickname != "" {
		db = db.Where("user_nickname LIKE ?", "%"+info.UserNickname+"%")
	}
	if info.UserGender != nil {
		db = db.Where("user_gender = ?", info.UserGender)
	}
	if info.UserGrade != nil {
		db = db.Where("user_grade = ?", info.UserGrade)
	}
	if info.UserProfession != "" {
		db = db.Where("user_profession LIKE ?", "%"+info.UserProfession+"%")
	}
	if info.UserCity != "" {
		db = db.Where("user_city LIKE ?", "%"+info.UserCity+"%")
	}
	if info.UserModel != nil {
		db = db.Where("user_model = ?", info.UserModel)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userDatas).Error
	return userDatas, total, err
}
func (this UserInfoService) GetInfo(id string) (reData interface{}, err error) {
	return this.GetUserInfo(id)
}
