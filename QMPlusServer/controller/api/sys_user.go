package api

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"time"
)

var (
	USER_HEADER_IMG_PATH string = "http://qmplusimg.henrongyi.top"
	USER_HEADER_BUCKET   string = "qm-plus-img"
)

type RegistAndLoginStuct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegestStuct struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body sysModel.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/regist [post]
func Regist(c *gin.Context) {
	var R RegestStuct
	_ = c.BindJSON(&R)
	user := &sysModel.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, user := user.Regist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{
			"user": user,
		})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"user": user,
		})
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.BindJSON(&L)
	U := &sysModel.SysUser{Username: L.Username, Password: L.Password}
	if err, user := U.Login(); err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("用户名密码错误或%v", err), gin.H{})
	} else {
		tokenNext(c, *user)
	}
}

//登录以后签发jwt
func tokenNext(c *gin.Context, user sysModel.SysUser) {
	j := &middleware.JWT{
		[]byte(config.GinVueAdminconfig.JWT.SigningKey), // 唯一签名
	}
	clams := middleware.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "qmPlus",                              //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		if config.GinVueAdminconfig.System.UseMultipoint {
			var loginJwt sysModel.JwtBlacklist
			loginJwt.Jwt = token
			err, jwtStr := loginJwt.GetRedisJWT(user.Username)
			if err == redis.Nil {
				err2 := loginJwt.SetRedisJWT(user.Username)
				if err2 != nil {
					servers.ReportFormat(c, false, "设置登录状态失败", gin.H{})
				} else {
					servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
				}
			} else if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
			} else {
				var blackjWT sysModel.JwtBlacklist
				blackjWT.Jwt = jwtStr
				err3 := blackjWT.JsonInBlacklist()
				if err3 != nil {
					servers.ReportFormat(c, false, "jwt作废失败", gin.H{})
				} else {
					err2 := loginJwt.SetRedisJWT(user.Username)
					if err2 != nil {
						servers.ReportFormat(c, false, "设置登录状态失败", gin.H{})
					} else {
						servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
					}
				}
			}
		} else {
			servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
		}
	}
}

type ChangePasswordStutrc struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.ChangePasswordStutrc true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [post]
func ChangePassword(c *gin.Context) {
	var params ChangePasswordStutrc
	_ = c.BindJSON(&params)
	U := &sysModel.SysUser{Username: params.Username, Password: params.Password}
	if err, _ := U.ChangePassword(params.NewPassword); err != nil {
		servers.ReportFormat(c, false, "修改失败，请检查用户名密码", gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// @Tags SysUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	//获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*middleware.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	//便于找到用户 以后从jwt中取
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err, filePath, _ := servers.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			err, user := new(sysModel.SysUser).UploadHeaderImg(uuid, filePath)
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改数据库链接失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{"user": user})
			}
		}
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(sysModel.SysUser).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(c *gin.Context) {
	var sua SetUserAuth
	_ = c.BindJSON(&sua)
	err := new(sysModel.SysUser).SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}
