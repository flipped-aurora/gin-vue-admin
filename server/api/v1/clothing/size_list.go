package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    clothingReq "github.com/flipped-aurora/gin-vue-admin/server/model/clothing/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SizeListApi struct {
}

var sizeListService = service.ServiceGroupApp.ClothingServiceGroup.SizeListService


// CreateSizeList 创建SizeList
// @Tags SizeList
// @Summary 创建SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.SizeList true "创建SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sizeList/createSizeList [post]
func (sizeListApi *SizeListApi) CreateSizeList(c *gin.Context) {
	var sizeList clothing.SizeList
	err := c.ShouldBindJSON(&sizeList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    sizeList.CreatedBy = utils.GetUserID(c)
	if err := sizeListService.CreateSizeList(&sizeList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSizeList 删除SizeList
// @Tags SizeList
// @Summary 删除SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.SizeList true "删除SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sizeList/deleteSizeList [delete]
func (sizeListApi *SizeListApi) DeleteSizeList(c *gin.Context) {
	var sizeList clothing.SizeList
	err := c.ShouldBindJSON(&sizeList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    sizeList.DeletedBy = utils.GetUserID(c)
	if err := sizeListService.DeleteSizeList(sizeList); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSizeListByIds 批量删除SizeList
// @Tags SizeList
// @Summary 批量删除SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sizeList/deleteSizeListByIds [delete]
func (sizeListApi *SizeListApi) DeleteSizeListByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := sizeListService.DeleteSizeListByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSizeList 更新SizeList
// @Tags SizeList
// @Summary 更新SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body clothing.SizeList true "更新SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sizeList/updateSizeList [put]
func (sizeListApi *SizeListApi) UpdateSizeList(c *gin.Context) {
	var sizeList clothing.SizeList
	err := c.ShouldBindJSON(&sizeList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    sizeList.UpdatedBy = utils.GetUserID(c)
	if err := sizeListService.UpdateSizeList(sizeList); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSizeList 用id查询SizeList
// @Tags SizeList
// @Summary 用id查询SizeList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothing.SizeList true "用id查询SizeList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sizeList/findSizeList [get]
func (sizeListApi *SizeListApi) FindSizeList(c *gin.Context) {
	var sizeList clothing.SizeList
	err := c.ShouldBindQuery(&sizeList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resizeList, err := sizeListService.GetSizeList(sizeList.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resizeList": resizeList}, c)
	}
}

// GetSizeListList 分页获取SizeList列表
// @Tags SizeList
// @Summary 分页获取SizeList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query clothingReq.SizeListSearch true "分页获取SizeList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sizeList/getSizeListList [get]
func (sizeListApi *SizeListApi) GetSizeListList(c *gin.Context) {
	var pageInfo clothingReq.SizeListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sizeListService.GetSizeListInfoList(pageInfo); err != nil {
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
