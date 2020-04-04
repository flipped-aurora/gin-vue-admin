package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"github.com/dchest/captcha"
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

type RegisterAndLoginStuct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
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
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R RegestStuct
	_ = c.ShouldBindJSON(&R)
	user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, user := user.Register()
	if err != nil {
		response.Result(response.ERROR, gin.H{
			"user": user,
		}, fmt.Sprintf("%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"user": user,
		}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.RegisterAndLoginStuct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L RegisterAndLoginStuct
	_ = c.ShouldBindJSON(&L)
	if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := U.Login(); err != nil {
			response.Result(response.ERROR, gin.H{}, fmt.Sprintf("用户名密码错误或%v", err), c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.Result(response.ERROR, gin.H{}, "验证码错误", c)
	}

}

//登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
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
		response.Result(response.ERROR, gin.H{}, "获取token失败", c)
	} else {
		if global.GVA_CONFIG.System.UseMultipoint {
			var loginJwt model.JwtBlacklist
			loginJwt.Jwt = token
			err, jwtStr := loginJwt.GetRedisJWT(user.Username)
			if err == redis.Nil {
				err2 := loginJwt.SetRedisJWT(user.Username)
				if err2 != nil {
					response.Result(response.ERROR, gin.H{}, "设置登录状态失败", c)
				} else {
					response.Result(response.SUCCESS, gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000}, "登录成功", c)
				}
			} else if err != nil {
				response.Result(response.ERROR, gin.H{}, fmt.Sprintf("%v", err), c)
			} else {
				var blackjWT model.JwtBlacklist
				blackjWT.Jwt = jwtStr
				err3 := blackjWT.JsonInBlacklist()
				if err3 != nil {
					response.Result(response.ERROR, gin.H{}, "jwt作废失败", c)
				} else {
					err2 := loginJwt.SetRedisJWT(user.Username)
					if err2 != nil {
						response.Result(response.ERROR, gin.H{}, "设置登录状态失败", c)
					} else {
						response.Result(response.SUCCESS, gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000}, "登录成功", c)
					}
				}
			}
		} else {
			response.Result(response.SUCCESS, gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000}, "登录成功", c)
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
// @Router /user/changePassword [put]
func ChangePassword(c *gin.Context) {
	var params ChangePasswordStutrc
	_ = c.ShouldBindJSON(&params)
	U := &model.SysUser{Username: params.Username, Password: params.Password}
	if err, _ := U.ChangePassword(params.NewPassword); err != nil {
		response.Result(response.ERROR, gin.H{}, "修改失败，请检查用户名密码", c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "修改成功", c)
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
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		//文件上传后拿到文件路径
		err, filePath, _ := utils.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			response.Result(response.ERROR, gin.H{}, fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			err, user := new(model.SysUser).UploadHeaderImg(uuid, filePath)
			if err != nil {
				response.Result(response.ERROR, gin.H{}, fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				response.Result(response.SUCCESS, gin.H{"user": user}, "上传成功", c)
			}
		}
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo model.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(model.SysUser).GetInfoList(pageInfo)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{
			"userList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取数据成功", c)
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
	_ = c.ShouldBindJSON(&sua)
	err := new(model.SysUser).SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		response.Result(response.ERROR, gin.H{}, fmt.Sprintf("修改失败，%v", err), c)
	} else {
		response.Result(response.SUCCESS, gin.H{}, "修改成功", c)
	}
}
