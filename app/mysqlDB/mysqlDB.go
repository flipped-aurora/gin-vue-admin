package mysqlDB

import (
	"app/global"
	"app/model/comment"
	"app/model/competition"
	"app/model/discuss"
	"app/model/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	dsn := global.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("无法连接数据库！", err)
		panic("无法连接数据库！")
		return
	}
	err = db.AutoMigrate(&user.UserInfo{}, &competition.ComInfo{}, &discuss.DisInfo{}, &comment.CommentInfo{}, &comment.CommentSonInfo{})
	if err != nil {
		fmt.Println(err)
		return
	}

	DB = db
}
func FindUserByID(id uint) (user.UserInfo, error) {
	var useR user.UserInfo
	err := DB.Model(user.UserInfo{}).Where("ID = ?", id).First(&useR).Error
	return useR, err
}
func FindUserByIDString(id string) (user.UserInfo, error) {
	var useR user.UserInfo
	err := DB.Model(user.UserInfo{}).Where("ID = ?", id).First(&useR).Error
	return useR, err
}
func FindDisByID(id string) (discuss.DisInfo, error) {
	var disS discuss.DisInfo
	err := DB.Model(discuss.DisInfo{}).Where("ID = ?", id).First(&disS).Error
	return disS, err
}
func FindCommentByID(id string) (comment.CommentInfo, error) {
	var commm comment.CommentInfo
	err := DB.Model(comment.CommentInfo{}).Where("ID = ?", id).First(&commm).Error
	return commm, err
}
func FindCommentSonByID(id string) (comment.CommentSonInfo, error) {
	var commm comment.CommentSonInfo
	err := DB.Model(comment.CommentSonInfo{}).Where("ID = ?", id).First(&commm).Error
	return commm, err
}
