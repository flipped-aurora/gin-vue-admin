package dbModel

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"main/controller/servers"
	"main/init/qmsql"
	"main/model/modelInterface"
	"main/tools"
)

type User struct {
	gorm.Model  `json:"-"`
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"-"`
	Password    string    `json:"-"`
	NickName    string    `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string    `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	Authority   Authority `json:"authority" form:"ForeignKey:authority_id;AssociationForeignKey:authority_id"`
	AuthorityId float64   `json:"-" gorm:"default:888"`
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
	findErr := qmsql.DEFAULTDB.Where("user_name = ?", u.Username).First(&user).Error
	//err为nil表明读取到了 不能注册
	if findErr == nil {
		return errors.New("用户名已注册"), nil
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = tools.MD5V(u.Password)
		u.UUID = uuid.NewV4()
		err = qmsql.DEFAULTDB.Create(u).Error
	}
	return err, u
}

//修改用户密码
func (u *User) ChangePassword(newPassword string) (err error, userInter *User) {
	var user User
	//后期修改jwt+password模式
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.Username, u.Password).First(&user).Update("pass_word", tools.MD5V(newPassword)).Error
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
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("user_name = ? AND pass_word = ?", u.Username, u.Password).First(&user).Error
	err = qmsql.DEFAULTDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}

// 用户头像上传更新地址
func (u *User) UploadHeaderImg(username string, filePath string) (err error, userInter *User) {
	var user User
	err = qmsql.DEFAULTDB.Where("user_name = ?", username).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// 分页获取数据  需要分页实现这个接口即可
func (u *User) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(u, info)
	if err != nil {
		return
	} else {
		var userList []User
		err = db.Find(&userList).Error
		return err, userList, total
	}
}
