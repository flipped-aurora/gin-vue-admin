package dbModel

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"main/init/mysql"
	"main/model/modelInterface"
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
func NewUser(user User) *User {
	return &User{UserName: user.UserName, PassWord: user.PassWord, NickName: user.NickName, HeaderImg: user.HeaderImg}
}

func (u *User) Create() (err error, user modelInterface.CURD) {
	err = mysql.DEFAULTDB.Create(u).Error
	return err, u
}

func (u *User) Delete() (err error, user modelInterface.CURD) {
	err = mysql.DEFAULTDB.Create(u).Error
	return err, u
}

func (u *User) Updata() (err error, user modelInterface.CURD) {
	err = mysql.DEFAULTDB.Create(u).Error
	return err, u
}

func (u *User) Read() (err error, user modelInterface.CURD) {
	err = mysql.DEFAULTDB.Create(u).Error
	return err, u
}
