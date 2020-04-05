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

type RegisterAndLoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type RegisterStruct struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22 ）
// @return    err             error
// @return    userInter       *SysUser
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

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22 ）
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser
func (u *SysUser) ChangePassword(newPassword string) (err error, userInter *SysUser) {
	var user SysUser
	//后期修改jwt+password模式
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22 ）
// @param     uuid             UUID
// @param     authorityId      string
// @return    err              error
func (u *SysUser) SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22 ）
// @return    err             error
// @return    userInter       *SysUser
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

// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22 ）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser
func (u *SysUser) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *SysUser) {
	var user SysUser
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22 ）
// @param     PageInfo         int
// @return    err              error
// @return    list             interface{}
// @return    total            int
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
