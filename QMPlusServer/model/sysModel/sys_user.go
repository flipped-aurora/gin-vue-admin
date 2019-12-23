package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/tools"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid"`
	Username    string       `json:"userName"`
	Password    string       `json:"-"`
	NickName    string       `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string       `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityId string       `json:"authorityId" gorm:"default:888"`
}

//type Propertie struct {
//	gorm.Model
//}

//注册接口model方法
func (u *SysUser) Regist() (err error, userInter *SysUser) {
	var user SysUser
	//判断用户名是否注册
	findErr := qmsql.DEFAULTDB.Where("username = ?", u.Username).First(&user).Error
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
func (u *SysUser) ChangePassword(newPassword string) (err error, userInter *SysUser) {
	var user SysUser
	//后期修改jwt+password模式
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", tools.MD5V(newPassword)).Error
	return err, u
}

//用户更新接口
func (u *SysUser) SetUserAuthority(uuid uuid.UUID, AuthorityId string) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&SysUser{}).Update("authority_id", AuthorityId).Error
	return err
}

//用户登录
func (u *SysUser) Login() (err error, userInter *SysUser) {
	var user SysUser
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return err, &user
	}
	err = qmsql.DEFAULTDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}

// 用户头像上传更新地址
func (u *SysUser) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *SysUser) {
	var user SysUser
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// 分页获取数据  需要分页实现这个接口即可
func (u *SysUser) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(u, info)
	if err != nil {
		return
	} else {
		var userList []SysUser
		err = db.Preload("Authority").Find(&userList).Error
		return err, userList, total
	}
}
