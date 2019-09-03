package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model/dbModel"
)

type RegistAndLoginStuct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

// @Tags User
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "用户注册接口"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/regist [post]
func Regist(c *gin.Context) {
	var R RegistAndLoginStuct
	_ = c.BindJSON(&R)

	U := &dbModel.User{UserName: R.UserName, PassWord: R.PassWord}
	err, user := U.Regist()
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     fmt.Sprintf("%v", err),
			"data": gin.H{
				"user": user,
			},
		})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "创建成功",
			"data": gin.H{
				"user": user,
			},
		})
	}
}

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.RegistAndLoginStuct true "用户登录接口"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.BindJSON(&L)
	U := &dbModel.User{UserName: L.UserName, PassWord: L.PassWord}
	if err, user := U.Login(); err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "用户名密码错误",
			"data": gin.H{
				"user": user,
			},
		})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "登录成功",
			"data": gin.H{
				"user": user,
			},
		})
	}
}

type ChangePassWordStutrc struct {
	UserName    string `json:"userName"`
	PassWord    string `json:"passWord"`
	NewPassWord string `json:"newPassWord"`
}

// @Tags User
// @Summary 用户修改密码
// @Produce  application/json
// @Param data body api.ChangePassWordStutrc true "用户修改密码"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassWord [post]
func ChangePassWord(c *gin.Context) {
	var params ChangePassWordStutrc
	_ = c.BindJSON(&params)
	U := &dbModel.User{UserName: params.UserName, PassWord: params.PassWord}
	if err, _ := U.ChangePassWord(params.NewPassWord); err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "修改失败，请检查用户名密码",
			"data":    gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "修改成功",
			"data":    gin.H{},
		})
	}
}
