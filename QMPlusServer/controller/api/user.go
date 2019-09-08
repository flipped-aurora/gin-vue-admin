package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/controller/servers"
	"main/middleware"
	"main/model/dbModel"
	"main/model/modelInterface"
	"mime/multipart"
	"time"
)

var (
	USER_HEADER_IMG_PATH string = "http://qmplusimg.henrongyi.top"
	USER_HEADER_BUCKET   string = "qm-plus-img"
)

type RegistAndLoginStuct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "用户注册接口"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/regist [post]
func Regist(c *gin.Context) {
	var R RegistAndLoginStuct
	_ = c.BindJSON(&R)

	U := &dbModel.User{UserName: R.UserName, PassWord: R.PassWord}
	err, user := U.Regist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{
			"user": user,
		})
	} else {
		servers.ReportFormat(c, false, "创建成功", gin.H{
			"user": user,
		})
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "用户登录接口"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.BindJSON(&L)
	U := &dbModel.User{UserName: L.UserName, PassWord: L.PassWord}
	if err, user := U.Login(); err != nil {
		servers.ReportFormat(c, false, "用户名密码错误", gin.H{"user": user})
	} else {
		tokenNext(c, *user)
	}
}

//登录以后签发jwt
func tokenNext(c *gin.Context, user dbModel.User) {
	j := &middleware.JWT{
		[]byte("qmPlus"), // 唯一签名
	}
	clams := middleware.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),   // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*7), // 过期时间 一周
			Issuer:    "qmPlus",                          //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "token": token})
	}
}

type ChangePassWordStutrc struct {
	UserName    string `json:"userName"`
	PassWord    string `json:"passWord"`
	NewPassWord string `json:"newPassWord"`
}

// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.ChangePassWordStutrc true "用户修改密码"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassWord [post]
func ChangePassWord(c *gin.Context) {
	var params ChangePassWordStutrc
	_ = c.BindJSON(&params)
	U := &dbModel.User{UserName: params.UserName, PassWord: params.PassWord}
	if err, _ := U.ChangePassWord(params.NewPassWord); err != nil {
		servers.ReportFormat(c, false, "修改失败，请检查用户名密码", gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// @Tags User
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param userName formData string true "用户上传头像"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	//获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*middleware.CustomClaims)
	fmt.Println(waitUse.NickName)
	_, header, err := c.Request.FormFile("headerImg")
	//便于找到用户 以后从jwt中取
	userName := c.PostForm("userName")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err, filePath := servers.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			err, user := new(dbModel.User).UploadHeaderImg(userName, filePath)

			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改数据库链接失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{"user": user})
			}
		}
	}
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getInfoList [post]
func GetInfoList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(dbModel.User).GetInfoList(pageInfo)
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
