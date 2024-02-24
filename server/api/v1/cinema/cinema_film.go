package cinema

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cinema"
	cinemaReq "github.com/flipped-aurora/gin-vue-admin/server/model/cinema/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CinemaFilmApi struct {
}

var cinemaFilmService = service.ServiceGroupApp.CinemaServiceGroup.CinemaFilmService

// CreateCinemaFilm 创建cinemaFilm表
// @Tags CinemaFilm
// @Summary 创建cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinemaReq.CinemaFilmCreate true "创建cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cinemaFilm/createCinemaFilm [post]
func (cinemaFilmApi *CinemaFilmApi) CreateCinemaFilm(c *gin.Context) {
	var cinemaFilm cinemaReq.CinemaFilmCreate
	err := c.ShouldBindJSON(&cinemaFilm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaFilmService.CreateCinemaFilm(&cinemaFilm); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCinemaFilm 删除cinemaFilm表
// @Tags CinemaFilm
// @Summary 删除cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaFilm true "删除cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cinemaFilm/deleteCinemaFilm [delete]
func (cinemaFilmApi *CinemaFilmApi) DeleteCinemaFilm(c *gin.Context) {
	ID := c.Query("ID")
	if err := cinemaFilmService.DeleteCinemaFilm(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCinemaFilmByIds 批量删除cinemaFilm表
// @Tags CinemaFilm
// @Summary 批量删除cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cinemaFilm/deleteCinemaFilmByIds [delete]
func (cinemaFilmApi *CinemaFilmApi) DeleteCinemaFilmByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := cinemaFilmService.DeleteCinemaFilmByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCinemaFilm 更新cinemaFilm表
// @Tags CinemaFilm
// @Summary 更新cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cinema.CinemaFilm true "更新cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cinemaFilm/updateCinemaFilm [put]
func (cinemaFilmApi *CinemaFilmApi) UpdateCinemaFilm(c *gin.Context) {
	var cinemaFilm cinema.CinemaFilm
	err := c.ShouldBindJSON(&cinemaFilm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cinemaFilmService.UpdateCinemaFilm(cinemaFilm); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCinemaFilm 用id查询cinemaFilm表
// @Tags CinemaFilm
// @Summary 用id查询cinemaFilm表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinema.CinemaFilm true "用id查询cinemaFilm表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cinemaFilm/findCinemaFilm [get]
func (cinemaFilmApi *CinemaFilmApi) FindCinemaFilm(c *gin.Context) {
	ID := c.Query("ID")
	if recinemaFilm, err := cinemaFilmService.GetCinemaFilm(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recinemaFilm": recinemaFilm}, c)
	}
}

// GetCinemaFilmList 分页获取cinemaFilm表列表
// @Tags CinemaFilm
// @Summary 分页获取cinemaFilm表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cinemaReq.CinemaFilmSearch true "分页获取cinemaFilm表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cinemaFilm/getCinemaFilmList [get]
func (cinemaFilmApi *CinemaFilmApi) GetCinemaFilmList(c *gin.Context) {
	var pageInfo cinemaReq.CinemaFilmSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cinemaFilmService.GetCinemaFilmInfoList(pageInfo); err != nil {
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
