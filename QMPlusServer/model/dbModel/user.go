package dbModel

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"main/init/qmsql"
	"main/tools"
)

type User struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"uuid"`
	UserName   string    `json:"userName"`
	PassWord   string    `json:"passWord"`
	NickName   string    `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg  string    `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	//Propertie                //	多余属性自行添加
	//PropertieId uint  // 自动关联 Propertie 的Id 附加属性过多 建议创建一对一关系
}

//type Propertie struct {
//	gorm.Model
//}

//注册接口model方法
func (u *User) Regist() (err error, userInter *User) {
	var user User
	//判断用户名是否注册
	findErr := qmsql.DEFAULTDB.Where("user_name = ?", u.UserName).First(&user).Error
	//err为nil表明读取到了 不能注册
	if findErr == nil {
		return errors.New("用户名已注册"), nil
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.PassWord = tools.MD5V(u.PassWord)
		u.UUID = uuid.NewV4()
		err = qmsql.DEFAULTDB.Create(u).Error
	}
	return err, u
}

//修改用户密码
func (u *User) ChangePassWord(newPassWord string) (err error, userInter *User) {
	var user User
	//后期修改jwt+password模式
	u.PassWord = tools.MD5V(u.PassWord)
	err = qmsql.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.UserName, u.PassWord).First(&user).Update("pass_word", tools.MD5V(newPassWord)).Error
	return err, u
}

//用户更新接口
func (u *User) UpdataUser() (err error, userInter *User) {
	err = qmsql.DEFAULTDB.Create(u).Error
	return err, u
}

//用户登录
func (u *User) Login() (err error, userInter *User) {
	var user User
	u.PassWord = tools.MD5V(u.PassWord)
	err = qmsql.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.UserName, u.PassWord).First(&user).Error
	return err, &user
}
