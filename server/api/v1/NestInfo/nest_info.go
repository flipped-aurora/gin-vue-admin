package NestInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestInfo"
	NestInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/NestInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NestInfoApi struct {
}

var nestinfoService = service.ServiceGroupApp.NestInfoServiceGroup.NestInfoService

// CreateNestInfo 创建NestInfo
// @Tags NestInfo
// @Summary 创建NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestInfo.NestInfo true "创建NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestinfo/createNestInfo [post]
func (nestinfoApi *NestInfoApi) CreateNestInfo(c *gin.Context) {
	var nestinfo NestInfo.NestInfo
	err := c.ShouldBindJSON(&nestinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nestinfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Nestid":       {utils.NotEmpty()},
		"NestLocation": {utils.NotEmpty()},
	}
	if err := utils.Verify(nestinfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := nestinfoService.CreateNestInfo(&nestinfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNestInfo 删除NestInfo
// @Tags NestInfo
// @Summary 删除NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestInfo.NestInfo true "删除NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nestinfo/deleteNestInfo [delete]
func (nestinfoApi *NestInfoApi) DeleteNestInfo(c *gin.Context) {
	var nestinfo NestInfo.NestInfo
	err := c.ShouldBindJSON(&nestinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nestinfo.DeletedBy = utils.GetUserID(c)
	if err := nestinfoService.DeleteNestInfo(nestinfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNestInfoByIds 批量删除NestInfo
// @Tags NestInfo
// @Summary 批量删除NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /nestinfo/deleteNestInfoByIds [delete]
func (nestinfoApi *NestInfoApi) DeleteNestInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := nestinfoService.DeleteNestInfoByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNestInfo 更新NestInfo
// @Tags NestInfo
// @Summary 更新NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body NestInfo.NestInfo true "更新NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /nestinfo/updateNestInfo [put]
func (nestinfoApi *NestInfoApi) UpdateNestInfo(c *gin.Context) {
	var nestinfo NestInfo.NestInfo
	err := c.ShouldBindJSON(&nestinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	nestinfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Nestid":       {utils.NotEmpty()},
		"NestLocation": {utils.NotEmpty()},
	}
	if err := utils.Verify(nestinfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := nestinfoService.UpdateNestInfo(nestinfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNestInfo 用id查询NestInfo
// @Tags NestInfo
// @Summary 用id查询NestInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NestInfo.NestInfo true "用id查询NestInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /nestinfo/findNestInfo [get]
func (nestinfoApi *NestInfoApi) FindNestInfo(c *gin.Context) {
	var nestinfo NestInfo.NestInfo
	err := c.ShouldBindQuery(&nestinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if renestinfo, err := nestinfoService.GetNestInfo(nestinfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renestinfo": renestinfo}, c)
	}
}

// GetNestInfoList 分页获取NestInfo列表
// @Tags NestInfo
// @Summary 分页获取NestInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query NestInfoReq.NestInfoSearch true "分页获取NestInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nestinfo/getNestInfoList [get]
func (nestinfoApi *NestInfoApi) GetNestInfoList(c *gin.Context) {
	var pageInfo NestInfoReq.NestInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := nestinfoService.GetNestInfoInfoList(pageInfo); err != nil {
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

func (NestInfoApi *NestInfoApi) GetAllUserList(c *gin.Context) {
	if list, total, err := nestinfoService.GetAllUserList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}

}

func (NestInfoApi *NestInfoApi) GetNestInfoInfoListWithUser(c *gin.Context) {
	if list, total, err := nestinfoService.GetNestInfoInfoListWithUser(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}

}
