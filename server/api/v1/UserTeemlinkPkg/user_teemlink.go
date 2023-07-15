package UserTeemlinkPkg

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/UserTeemlinkPkg"
	UserTeemlinkPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/UserTeemlinkPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserTeemlinkApi struct {
}

type TlTokenParam struct {
	Tltoken string `json:"tltoken" form:"tltoken"`
}

var usertlService = service.ServiceGroupApp.UserTeemlinkPkgServiceGroup.UserTeemlinkService

// CreateUserTeemlink 创建UserTeemlink
// @Tags UserTeemlink
// @Summary 创建UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body UserTeemlinkPkg.UserTeemlink true "创建UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usertl/createUserTeemlink [post]
func (usertlApi *UserTeemlinkApi) CreateUserTeemlink(c *gin.Context) {
	var usertl UserTeemlinkPkg.UserTeemlink
	err := c.ShouldBindJSON(&usertl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	usertl.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserID":   {utils.NotEmpty()},
		"TlUserID": {utils.NotEmpty()},
	}
	if err := utils.Verify(usertl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := usertlService.CreateUserTeemlink(&usertl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserTeemlink 删除UserTeemlink
// @Tags UserTeemlink
// @Summary 删除UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body UserTeemlinkPkg.UserTeemlink true "删除UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /usertl/deleteUserTeemlink [delete]
func (usertlApi *UserTeemlinkApi) DeleteUserTeemlink(c *gin.Context) {
	var usertl UserTeemlinkPkg.UserTeemlink
	err := c.ShouldBindJSON(&usertl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	usertl.DeletedBy = utils.GetUserID(c)
	if err := usertlService.DeleteUserTeemlink(usertl); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserTeemlinkByIds 批量删除UserTeemlink
// @Tags UserTeemlink
// @Summary 批量删除UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /usertl/deleteUserTeemlinkByIds [delete]
func (usertlApi *UserTeemlinkApi) DeleteUserTeemlinkByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := usertlService.DeleteUserTeemlinkByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserTeemlink 更新UserTeemlink
// @Tags UserTeemlink
// @Summary 更新UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body UserTeemlinkPkg.UserTeemlink true "更新UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /usertl/updateUserTeemlink [put]
func (usertlApi *UserTeemlinkApi) UpdateUserTeemlink(c *gin.Context) {
	var usertl UserTeemlinkPkg.UserTeemlink
	err := c.ShouldBindJSON(&usertl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	usertl.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserID":   {utils.NotEmpty()},
		"TlUserID": {utils.NotEmpty()},
	}
	if err := utils.Verify(usertl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := usertlService.UpdateUserTeemlink(usertl); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserTeemlink 用id查询UserTeemlink
// @Tags UserTeemlink
// @Summary 用id查询UserTeemlink
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query UserTeemlinkPkg.UserTeemlink true "用id查询UserTeemlink"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /usertl/findUserTeemlink [get]
func (usertlApi *UserTeemlinkApi) FindUserTeemlink(c *gin.Context) {
	var usertl UserTeemlinkPkg.UserTeemlink
	err := c.ShouldBindQuery(&usertl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reusertl, err := usertlService.GetUserTeemlink(usertl.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reusertl": reusertl}, c)
	}
}

// GetUserTeemlinkList 分页获取UserTeemlink列表
// @Tags UserTeemlink
// @Summary 分页获取UserTeemlink列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query UserTeemlinkPkgReq.UserTeemlinkSearch true "分页获取UserTeemlink列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usertl/getUserTeemlinkList [get]
func (usertlApi *UserTeemlinkApi) GetUserTeemlinkList(c *gin.Context) {
	var pageInfo UserTeemlinkPkgReq.UserTeemlinkSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := usertlService.GetUserTeemlinkInfoList(pageInfo); err != nil {
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

func (usertlApi *UserTeemlinkApi) GetGvaToken(c *gin.Context) {
	// var param map[string]string
	var tltokenparam TlTokenParam
	err := c.ShouldBind(&tltokenparam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// tltoken, ok := c.GetQuery("tltoken")
	// if !ok {
	// 	global.GVA_LOG.Error("获取tltoken失败!")
	// 	response.FailWithMessage("获取tltoken失败", c)
	// 	c.Abort()
	// 	return
	// }
	tltoken := tltokenparam.Tltoken
	// tltoken := c.PostForm("tltoken")
	tlJWT := utils.NewTLJWT()
	tlClaims, err := tlJWT.ParseToken(tltoken)
	if err != nil {
		if errors.Is(err, utils.TokenExpired) {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	//Stephen：找到天翎用户ID对应的本地用户
	usertl, err := service.ServiceGroupApp.UserTeemlinkPkgServiceGroup.GetUserTeemlinkByTlID(tlClaims.Username)
	if err != nil {
		response.FailWithDetailed(gin.H{"reload": true}, "TL用户未绑定", c)
		c.Abort()
		return
	}
	user, err := service.ServiceGroupApp.SystemServiceGroup.FindUserById(*usertl.UserID)
	if err != nil {
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	gvaJWT := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	gvaClaims := gvaJWT.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := gvaJWT.CreateToken(gvaClaims)
	if err != nil {
		global.GVA_LOG.Error("天翎用户新建token失败!", zap.Error(err))
		response.FailWithMessage("天翎用户新建token失败", c)
		return
	}
	response.OkWithData(gin.H{"gvatoken": token}, c)
}
