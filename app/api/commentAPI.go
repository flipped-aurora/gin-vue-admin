package api

import (
	"app/global"
	"app/model/comment"
	"app/mysqlDB"
	"app/response"
	"app/service"
	"github.com/gin-gonic/gin"
)

type CommentInfoApi struct {
}

var commentDataService = service.ServiceApi.CommentService

// CreateCommentInfo 创建评论信息
// @Tags CommentInfo
// @Summary 创建评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.CommentInfo true "创建评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /commentData/createCommentInfo [post]
func (commentDataApi *CommentInfoApi) CreateCommentInfo(c *gin.Context) {
	var commentData comment.CommentInfo
	err := c.ShouldBindJSON(&commentData)
	userId, ok := c.Get("userID")
	if !ok {
		response.FailWithMessage("用户身份验证失败", c)
		return
	}
	commentData.CommentUserId = userId.(uint)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentDataService.CreateCommentInfo(&commentData); err != nil {
		//global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// CreateCommentSonInfo 创建评论信息
// @Tags CommentInfo
// @Summary 创建评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.CommentInfo true "创建评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /commentData/createCommentInfo [post]
func (commentDataApi *CommentInfoApi) CreateCommentSonInfo(c *gin.Context) {
	var commentData comment.CommentSonInfo
	err := c.ShouldBindJSON(&commentData)

	userId, ok := c.Get("userID")
	if !ok {
		response.FailWithMessage("用户身份验证失败", c)
		return
	}
	commentData.CommentUserId = userId.(uint)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentDataService.CreateCommentSonInfo(&commentData); err != nil {
		//global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommentInfo 删除评论信息
// @Tags CommentInfo
// @Summary 删除评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.CommentInfo true "删除评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentData/deleteCommentInfo [delete]
func (commentDataApi *CommentInfoApi) DeleteCommentInfo(c *gin.Context) {
	id := c.Query("ID")
	userId, ok := c.Get("userID")
	if !ok {
		response.FailWithMessage("身份验证失败", c)
	}
	useridd := userId.(uint)
	user, err := mysqlDB.FindUserByID(useridd)
	if err != nil {
		response.FailWithMessage("无法查找到该用户", c)
		return
	}
	comm, err := mysqlDB.FindCommentByID(id)
	if comm.ID == 0 {
		response.FailWithMessage("该评论已被删除", c)
		return
	}
	if comm.CommentUserId != user.ID && user.UserModel != global.UserModelGuanLiYuan {
		response.FailWithMessage("删除权限不足", c)
	}

	if err := commentDataService.DeleteCommentInfo(id); err != nil {
		//global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentSonInfo 删除评论信息
// @Tags CommentInfo
// @Summary 删除评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.CommentInfo true "删除评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentData/deleteCommentInfo [delete]
func (commentDataApi *CommentInfoApi) DeleteCommentSonInfo(c *gin.Context) {
	id := c.Query("ID")
	userId, ok := c.Get("userID")
	if !ok {
		response.FailWithMessage("身份验证失败", c)
	}
	useridd := userId.(uint)
	user, err := mysqlDB.FindUserByID(useridd)
	if err != nil {
		response.FailWithMessage("无法查找到该用户", c)
		return
	}
	comm, err := mysqlDB.FindCommentSonByID(id)
	if comm.ID == 0 {
		response.FailWithMessage("该评论已被删除", c)
		return
	}
	if comm.CommentUserId != user.ID && user.UserModel != global.UserModelGuanLiYuan {
		response.FailWithMessage("删除权限不足", c)
	}

	if err := commentDataService.DeleteCommentSonInfo(id); err != nil {
		//global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentInfoByIds 批量删除评论信息
// @Tags CommentInfo
// @Summary 批量删除评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /commentData/deleteCommentInfoByIds [delete]
func (commentDataApi *CommentInfoApi) DeleteCommentInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := commentDataService.DeleteCommentInfoByIds(ids); err != nil {
		//global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCommentInfo 更新评论信息
// @Tags CommentInfo
// @Summary 更新评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.CommentInfo true "更新评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commentData/updateCommentInfo [put]
func (commentDataApi *CommentInfoApi) UpdateCommentInfo(c *gin.Context) {
	var commentData comment.CommentInfo
	err := c.ShouldBindJSON(&commentData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentDataService.UpdateCommentInfo(commentData); err != nil {
		//global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateCommentSonInfo 更新评论信息
// @Tags CommentInfo
// @Summary 更新评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body comment.CommentInfo true "更新评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commentData/updateCommentInfo [put]
func (commentDataApi *CommentInfoApi) UpdateCommentSonInfo(c *gin.Context) {
	var commentData comment.CommentSonInfo
	err := c.ShouldBindJSON(&commentData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentDataService.UpdateCommentSonInfo(commentData); err != nil {
		//global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCommentInfo 用id查询评论信息
// @Tags CommentInfo
// @Summary 用id查询评论信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query comment.CommentInfo true "用id查询评论信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commentData/findCommentInfo [get]
func (commentDataApi *CommentInfoApi) FindCommentInfo(c *gin.Context) {
	id := c.Query("ID")

	if recommentData, err := commentDataService.GetCommentInfo(id); err != nil {
		//global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recommentData": recommentData}, c)
	}
}

// GetCommentInfoList 分页获取评论信息列表
// @Tags CommentInfo
// @Summary 分页获取评论信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query commentReq.CommentInfoSearch true "分页获取评论信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commentData/getCommentInfoList [get]
func (commentDataApi *CommentInfoApi) GetCommentInfoList(c *gin.Context) {
	var pageInfo comment.CommentInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := commentDataService.GetCommentInfoInfoList(pageInfo); err != nil {
		//global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
