package model

import (
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
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
func (u *SysUser) Register() (err error, userInter *SysUser) {
	var user SysUser
	//判断用户名是否注册
	notResigt := global.GVA_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	//notResigt为false表明读取到了 不能注册
	if !notResigt {
		return errors.New("用户名已注册"), nil
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GVA_DB.Create(u).Error
	}
	return err, u
}

//修改用户密码
func (u *SysUser) ChangePassword(newPassword string) (err error, userInter *SysUser) {
	var user SysUser
	//后期修改jwt+password模式
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//用户更新接口
func (u *SysUser) SetUserAuthority(uuid uuid.UUID, AuthorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&SysUser{}).Update("authority_id", AuthorityId).Error
	return err
}

//用户登录
func (u *SysUser) Login() (err error, userInter *SysUser) {
	var user SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return err, &user
	}
	err = global.GVA_DB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}

// 用户头像上传更新地址
func (u *SysUser) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *SysUser) {
	var user SysUser
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// 分页获取数据
func (u *SysUser) GetInfoList(info PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	if err != nil {
		return
	} else {
		var userList []SysUser
		err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
		return err, userList, total
	}
}
