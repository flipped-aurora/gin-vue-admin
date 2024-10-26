package xiao

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
	xiaoReq "github.com/flipped-aurora/gin-vue-admin/server/model/xiao/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao/xiaores"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

type CliLoadApi struct{}

// CreateCliLoad 创建cliLoad表
// @Tags CliLoad
// @Summary 创建cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliLoad true "创建cliLoad表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cliLoad/createCliLoad [post]
func (cliLoadApi *CliLoadApi) CreateCliLoad(c *gin.Context) {
	var cliLoad xiao.CliLoad
	err := c.ShouldBindJSON(&cliLoad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliLoadService.CreateCliLoad(&cliLoad)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCliLoad 删除cliLoad表
// @Tags CliLoad
// @Summary 删除cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliLoad true "删除cliLoad表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cliLoad/deleteCliLoad [delete]
func (cliLoadApi *CliLoadApi) DeleteCliLoad(c *gin.Context) {
	ID := c.Query("ID")
	err := cliLoadService.DeleteCliLoad(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCliLoadByIds 批量删除cliLoad表
// @Tags CliLoad
// @Summary 批量删除cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cliLoad/deleteCliLoadByIds [delete]
func (cliLoadApi *CliLoadApi) DeleteCliLoadByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cliLoadService.DeleteCliLoadByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCliLoad 更新cliLoad表
// @Tags CliLoad
// @Summary 更新cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliLoad true "更新cliLoad表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cliLoad/updateCliLoad [put]
func (cliLoadApi *CliLoadApi) UpdateCliLoad(c *gin.Context) {
	var cliLoad xiao.CliLoad
	err := c.ShouldBindJSON(&cliLoad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cliLoadService.UpdateCliLoad(cliLoad)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCliLoad 用id查询cliLoad表
// @Tags CliLoad
// @Summary 用id查询cliLoad表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiao.CliLoad true "用id查询cliLoad表"
// @Success 200 {object} response.Response{data=xiao.CliLoad,msg=string} "查询成功"
// @Router /cliLoad/findCliLoad [get]
func (cliLoadApi *CliLoadApi) FindCliLoad(c *gin.Context) {
	ID := c.Query("ID")
	recliLoad, err := cliLoadService.GetCliLoad(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recliLoad, c)
}

// GetCliLoadList 分页获取cliLoad表列表
// @Tags CliLoad
// @Summary 分页获取cliLoad表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliLoadSearch true "分页获取cliLoad表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cliLoad/getCliLoadList [get]
func (cliLoadApi *CliLoadApi) GetCliLoadList(c *gin.Context) {
	var pageInfo xiaoReq.CliLoadSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cliLoadService.GetCliLoadInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetCliLoadPublic 不需要鉴权的cliLoad表接口
// @Tags CliLoad
// @Summary 不需要鉴权的cliLoad表接口
// @accept application/json
// @Produce application/json
// @Param data query xiaoReq.CliLoadSearch true "分页获取cliLoad表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cliLoad/getCliLoadPublic [get]
func (cliLoadApi *CliLoadApi) GetCliLoadPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	cliLoadService.GetCliLoadPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的cliLoad表接口信息",
	}, "获取成功", c)
}

// Login 用户登录方法
// @Tags CliLoad
// @Summary 用户登录方法
// @accept application/json
// @Produce application/json
// @Param data body xiao.CliLoad true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /cliLoad/login [POST]
func (cliLoadApi *CliLoadApi) Login(c *gin.Context) {
	// 请添加自己的业务逻辑
	var cliLoad xiao.CliLoad
	err := c.ShouldBindJSON(&cliLoad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info, err := cliLoadService.Login(&cliLoad)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	//查询用户信息
	cliLoadApi.TokenNext(c, &info)
}

func (cliloadApi *CliLoadApi) TokenNext(c *gin.Context, loadres *xiaores.CliLoadResponse) {
	jwtHandler := utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 初始化JWT处理器

	// 创建JWT Claims
	claims := jwtHandler.CreateClaims(systemReq.BaseClaims{
		NickName: loadres.Load.Desc,    // 备注信息
		Username: loadres.Load.Address, // 用户地址
		ID:       loadres.Load.ID,
	})

	// 生成Token
	token, err := jwtHandler.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("生成Token失败", zap.Error(err))
		response.FailWithMessage("生成Token失败", c)
		return
	}

	// 构建响应数据
	respData := map[string]interface{}{
		"loadinfo":   loadres,
		"token":      token,
		"expires_at": claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}

	// 多点登录未启用，直接返回Token
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(respData, "登录成功", c)
		return
	}

	// 处理单点登录逻辑
	oldJwt, err := jwtService.GetRedisJWT(loadres.User.Address)
	switch {
	case errors.Is(err, redis.Nil): // 旧Token不存在于Redis，设置新Token
		if setErr := jwtService.SetRedisJWT(token, loadres.User.Address); setErr != nil {
			global.GVA_LOG.Error("设置新Token至Redis失败", zap.Error(setErr))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(respData, "登录成功", c)
	case err != nil: // Redis操作错误
		global.GVA_LOG.Error("从Redis获取Token失败", zap.Error(err))
		response.FailWithMessage("单点登录失败", c)
	default: // 旧Token存在，加入黑名单并设置新Token
		blacklistEntry := system.JwtBlacklist{Jwt: oldJwt}
		if blacklistErr := jwtService.JsonInBlacklist(blacklistEntry); blacklistErr != nil {
			global.GVA_LOG.Error("将旧Token加入黑名单失败", zap.Error(blacklistErr))
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if setNewTokenErr := jwtService.SetRedisJWT(token, loadres.User.Address); setNewTokenErr != nil {
			global.GVA_LOG.Error("设置新Token至Redis失败", zap.Error(setNewTokenErr))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(respData, "登录成功", c)
		return
	}
}
