package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type AgentApi struct {
}

var agentService = service.ServiceGroupApp.ClothingServiceGroup.AgentService

// CreateAgent 创建Agent
// @Tags Agent
// @Summary 创建Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Agent true "创建Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agent/createAgent [post]
func (agentApi *AgentApi) CreateAgent(c *gin.Context) {
	var agent clothing.Agent
	err := c.ShouldBindJSON(&agent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	agent.CreatedBy = utils.GetUserID(c)
	if err := agentService.CreateAgent(&agent); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		if strings.Contains(err.Error(), "unique") {
			response.FailWithMessage("代理人已存在", c)
			return
		}
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAgent 删除Agent
// @Tags Agent
// @Summary 删除Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Agent true "删除Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agent/deleteAgent [delete]
func (agentApi *AgentApi) DeleteAgent(c *gin.Context) {
	var agent clothing.Agent
	err := c.ShouldBindJSON(&agent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	agent.DeletedBy = utils.GetUserID(c)
	if err := agentService.DeleteAgent(agent); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAgentByIds 批量删除Agent
// @Tags Agent
// @Summary 批量删除Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /agent/deleteAgentByIds [delete]
func (agentApi *AgentApi) DeleteAgentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := agentService.DeleteAgentByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAgent 更新Agent
// @Tags Agent
// @Summary 更新Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.Agent true "更新Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /agent/updateAgent [put]
func (agentApi *AgentApi) UpdateAgent(c *gin.Context) {
	var agent clothing.Agent
	err := c.ShouldBindJSON(&agent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	agent.UpdatedBy = utils.GetUserID(c)
	if err := agentService.UpdateAgent(agent); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAgent 用id查询Agent
// @Tags Agent
// @Summary 用id查询Agent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.Agent true "用id查询Agent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /agent/findAgent [get]
func (agentApi *AgentApi) FindAgent(c *gin.Context) {
	var agent clothing.Agent
	err := c.ShouldBindQuery(&agent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reagent, err := agentService.GetAgent(agent.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reagent": reagent}, c)
	}
}

// GetAgentList 分页获取Agent列表
// @Tags Agent
// @Summary 分页获取Agent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.AgentSearch true "分页获取Agent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agent/getAgentList [get]
func (agentApi *AgentApi) GetAgentList(c *gin.Context) {
	var pageInfo clothingReq.AgentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := agentService.GetAgentInfoList(pageInfo); err != nil {
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
