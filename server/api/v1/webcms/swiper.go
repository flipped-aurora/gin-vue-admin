package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	webcmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/webcms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SwiperApi struct {
}

var swiperService = service.ServiceGroupApp.WebcmsServiceGroup.SwiperService

// CreateSwiper 创建Swiper
// @Tags Swiper
// @Summary 创建Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Swiper true "创建Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /swiper/createSwiper [post]
func (swiperApi *SwiperApi) CreateSwiper(c *gin.Context) {
	var swiper webcms.Swiper
	err := c.ShouldBindJSON(&swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Thumb":  {utils.NotEmpty()},
		"Enable": {utils.NotEmpty()},
	}
	if err := utils.Verify(swiper, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	swiper.Siteid = siteinfo["id"]
	if err := swiperService.CreateSwiper(swiper); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSwiper 删除Swiper
// @Tags Swiper
// @Summary 删除Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Swiper true "删除Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /swiper/deleteSwiper [delete]
func (swiperApi *SwiperApi) DeleteSwiper(c *gin.Context) {
	var swiper webcms.Swiper
	err := c.ShouldBindJSON(&swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := swiperService.DeleteSwiper(swiper); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSwiperByIds 批量删除Swiper
// @Tags Swiper
// @Summary 批量删除Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /swiper/deleteSwiperByIds [delete]
func (swiperApi *SwiperApi) DeleteSwiperByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := swiperService.DeleteSwiperByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSwiper 更新Swiper
// @Tags Swiper
// @Summary 更新Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webcms.Swiper true "更新Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /swiper/updateSwiper [put]
func (swiperApi *SwiperApi) UpdateSwiper(c *gin.Context) {
	var swiper webcms.Swiper
	err := c.ShouldBindJSON(&swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Thumb":  {utils.NotEmpty()},
		"Enable": {utils.NotEmpty()},
	}
	if err := utils.Verify(swiper, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := swiperService.UpdateSwiper(swiper); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSwiper 用id查询Swiper
// @Tags Swiper
// @Summary 用id查询Swiper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcms.Swiper true "用id查询Swiper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /swiper/findSwiper [get]
func (swiperApi *SwiperApi) FindSwiper(c *gin.Context) {
	var swiper webcms.Swiper
	err := c.ShouldBindQuery(&swiper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reswiper, err := swiperService.GetSwiper(swiper.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reswiper": reswiper}, c)
	}
}

// GetSwiperList 分页获取Swiper列表
// @Tags Swiper
// @Summary 分页获取Swiper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webcmsReq.SwiperSearch true "分页获取Swiper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /swiper/getSwiperList [get]
func (swiperApi *SwiperApi) GetSwiperList(c *gin.Context) {
	var pageInfo webcmsReq.SwiperSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取站点id
	siteinfo := c.GetStringMapString("siteinfo")
	if list, total, err := swiperService.GetSwiperInfoList(pageInfo, siteinfo["id"]); err != nil {
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
