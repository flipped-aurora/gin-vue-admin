package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	clothingRes "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type AppUserApi struct {
}

var appUserService = service.ServiceGroupApp.ClothingServiceGroup.AppUserService
var jwtService = service.ServiceGroupApp.ClothingServiceGroup.JwtService

func (appUserApi *AppUserApi) CreateAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	appUser.CreatedBy = utils.GetUserID(c)
	if err := appUserService.CreateAppUser(&appUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (appUserApi *AppUserApi) DeleteAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	appUser.DeletedBy = utils.GetUserID(c)
	if err := appUserService.DeleteAppUser(appUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (appUserApi *AppUserApi) DeleteAppUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := appUserService.DeleteAppUserByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (appUserApi *AppUserApi) UpdateAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	appUser.UpdatedBy = utils.GetUserID(c)
	if err := appUserService.UpdateAppUser(appUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (appUserApi *AppUserApi) FindAppUser(c *gin.Context) {
	var appUser clothing.AppUser
	err := c.ShouldBindQuery(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reappUser, err := appUserService.GetAppUser(appUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappUser": reappUser}, c)
	}
}

func (appUserApi *AppUserApi) GetAppUserList(c *gin.Context) {
	var pageInfo clothingReq.AppUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUserService.GetAppUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CheckExist 检测用户名或手机号存在
// @Tags AppUser
// @Summary 检测用户名或手机号存在
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.CheckExist true "检测用户名或手机号存在"
// @Success 200 {string} clothingRes.CheckExistResponse
// @Router /appUser/checkExist [get]
func (AppUserApi *AppUserApi) CheckExist(c *gin.Context) {
	var l clothingReq.CheckExist
	err := c.ShouldBindQuery(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var r clothingRes.CheckExistResponse
	if err := appUserService.CheckExist(l.Content, l.Type); err != nil {
		r.IsExist = true
		r.Msg = err.Error()
	}
	response.OkWithDetailed(r, "获取成功", c)
}

// Register 注册
// @Tags AppUser
// @Summary 注册
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.Register true "注册"
// @Success 200 {string} string "{"success":true,"msg":"创建成功"}"
// @Router /appUser/register [post]
func (appUserApi *AppUserApi) Register(c *gin.Context) {
	var l clothingReq.Register
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var appUser clothing.AppUser
	appUser.PhoneNum = l.PhoneNum
	appUser.Username = l.Username
	appUser.Password = l.Password
	if user, err := appUserService.Register(appUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		appUserApi.TokenNext(c, user)
	}
}

// Login 登录
// @Tags AppUser
// @Summary 登录
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.Login true "登录"
// @Success 200 {string} clothingRes.LoginResponse
// @Router /appUser/login [post]
func (appUserApi *AppUserApi) Login(c *gin.Context) {
	var l clothingReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(l, utils.LoginVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	u := &clothing.AppUser{PhoneNum: l.PhoneNum, Username: l.Username, Password: l.Password}
	user, err := appUserService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Status != nil && !*user.Status {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	appUserApi.TokenNext(c, *user)
}

// TokenNext 登录以后签发jwt
func (appUserApi *AppUserApi) TokenNext(c *gin.Context, user clothing.AppUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:       user.ID,
		NickName: user.Nickname,
		Username: user.Username,
	})
	userInfo := clothingRes.UserInfo{
		ID:       user.ID,
		PhoneNum: user.PhoneNum,
		Username: user.Username,
		Roles:    user.GetRoles(),
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(clothingRes.LoginResponse{
			User:      userInfo,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(clothingRes.LoginResponse{
			User:      userInfo,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT clothing.AppJwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(clothingRes.LoginResponse{
			User:      userInfo,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// GetUserInfo 获取用户详情
// @Tags AppUser
// @Summary 获取用户详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} clothing.AppUser
// @Router /appUser/getUserInfo [get]
func (appUserApi *AppUserApi) GetUserInfo(c *gin.Context) {
	userId := utils.GetUserID(c)
	if reappUser, err := appUserService.GetAppUser(userId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reappUser, c)
	}
}

// GetUserList 获取用户列表
// @Tags AppUser
// @Summary 获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.UserFilter true "获取用户列表"
// @Success 200 {string} response.PageResult
// @Router /appUser/getUserList [get]
func (appUserApi *AppUserApi) GetUserList(c *gin.Context) {
	var filter clothingReq.UserFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Sugar().Info(filter)
	userID := utils.GetUserID(c)
	if filter.CompanyID > 0 || filter.TeamID > 0 {
		companyID := filter.CompanyID
		if team, err := teamService.GetTeam(uint(filter.TeamID)); err == nil {
			companyID = int(team.CompanyID)
		}
		global.GVA_LOG.Sugar().Info(companyID)
		if !userRoleService.CheckStaff(userID, uint(companyID)) {
			response.FailWithMessage("权限不足", c)
			return
		}
	}
	if list, total, err := appUserService.GetAppUserList(filter); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		res := make([]clothingRes.UserInfo, len(list))
		for i, v := range list {
			res[i] = clothingRes.UserInfo{
				ID:       v.ID,
				Roles:    v.GetRoles(),
				PhoneNum: v.PhoneNum,
				Username: v.Username,
			}
		}
		response.OkWithDetailed(response.PageResult{
			List:     res,
			Total:    total,
			Page:     filter.Page,
			PageSize: filter.PageSize,
		}, "获取成功", c)
	}
}
