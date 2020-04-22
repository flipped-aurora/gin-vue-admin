package api

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model/customerModel"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"gin-vue-admin/tools"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	CUS_HEADER_IMG_PATH string = "http://q8g3778mf.bkt.clouddn.com"
	CUS_HEADER_BUCKET   string = "lqkimage"
)

type CusStruct struct {
	Username  string `json:"userName"`
	Password  string `json:"passWord"`
	NickName  string `json:"nickName"`
	HeaderImg string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
}

type CustomerId struct {
	UUID uuid.UUID
}

func CusRegist(c *gin.Context) {
	var R RegestStuct
	_ = c.ShouldBindJSON(&R)
	user := &customerModel.Customers{Username: R.Username, Nickname: R.NickName, Password: R.Password, Image: R.HeaderImg}
	err, user := user.Register()
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

func CusLogin(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.ShouldBindJSON(&L)
	U := &customerModel.Customers{Username: L.Username, Password: L.Password}
	if err, user := U.Login(); err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("用户名密码错误或%v", err), gin.H{})
	} else {
		custokenNext(c, *user)
	}
}

//登录以后签发jwt
func custokenNext(c *gin.Context, user customerModel.Customers) {
	j := &middleware.JWT{
		[]byte(config.GinVueAdminconfig.JWT.SigningKey), // 唯一签名
	}
	clams := middleware.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.Nickname,
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

func ChangeCusPassword(c *gin.Context) {
	var params ChangePasswordStutrc
	_ = c.ShouldBindJSON(&params)
	U := &customerModel.Customers{Username: params.Username, Password: params.Password}
	if err, _ := U.ChangePassword(params.NewPassword); err != nil {
		servers.ReportFormat(c, false, "修改失败，请检查用户名密码", gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

func UploadCusHeaderImg(c *gin.Context) {
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
		err, filePath, _ := servers.Upload(header, CUS_HEADER_BUCKET, CUS_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			err, user := new(customerModel.Customers).UploadHeaderImg(uuid, filePath)
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改数据库链接失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{"user": user})
			}
		}
	}
}

func AddCustomer(c *gin.Context) {
	var cus customerModel.Customers
	_ = c.ShouldBindJSON(&cus)

	err := cus.AddCustomer()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败, %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "添加成功", gin.H{})
	}
}

func UpdateCustomer(c *gin.Context) {
	var cus customerModel.Customers
	_ = c.ShouldBindJSON(&cus)

	err := cus.UpdateCustomer()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改失败, %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

func DeleteCustomer(c *gin.Context) {
	var customerId CustomerId
	var cus customerModel.Customers
	_ = c.ShouldBindJSON(&customerId)
	err := cus.DeleteCustomer(customerId.UUID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败, %v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "删除成功", gin.H{})
	}
}

func GetCustomerList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(customerModel.Customers).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"cusList":  list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

func GetCustomerById(c *gin.Context) {
	var customerId CustomerId
	var cus customerModel.Customers

	_ = c.ShouldBindJSON(&customerId)
	err := cus.GetCustomerById(customerId.UUID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败,%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{"customer": cus})
	}
}

type Phone struct {
	Phone string `json:"phone"`
}

func SendPhoneCaptcha(c *gin.Context) {
	var phone Phone
	_ = c.ShouldBindJSON(&phone)
	code := tools.SMSSend(phone.Phone)
	servers.ReportFormat(c, true, "发送成功", gin.H{"captcha": code})
}
