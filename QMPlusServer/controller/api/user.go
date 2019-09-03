package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model/dbModel"
	"main/model/modelInterface"
)

type RegistStuct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

// @Summary 创建用户
// @Produce  application/x-www-form-urlencoded
// @Param data body api.RegistStuct true "用户注册接口"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/regist [post]
func Regist(c *gin.Context) {
	var R RegistStuct
	_ = c.BindJSON(&R)
	U := dbModel.NewUser(dbModel.User{UserName: R.UserName, PassWord: R.PassWord})
	var curd modelInterface.CURD
	curd = U
	err, user := curd.Create()
	fmt.Println(err, user)
}
