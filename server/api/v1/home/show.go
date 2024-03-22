package home

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ShowService            = service.ServiceGroupApp.HomeServiceGroup.ShowService
	temp                   gin.H
	nextInfo, previousInfo webcms.Class
)

type ShowApi struct{}

func (l *ShowApi) Show(c *gin.Context) {
	catid := c.Param("catid")
	catId, err := strconv.Atoi(catid)
	if err != nil {
		global.GVA_LOG.Error("栏目id参数错误!", zap.Error(err))
		response.Fail404Error(c)
		return
	}
	showid := c.Param("id")
	showId, err := strconv.Atoi(showid)
	if err != nil {
		global.GVA_LOG.Error("id参数转换错误!", zap.Error(err))
		response.Fail404Error(c)
		return
	}
	// 获取站点id
	siteid := c.GetInt("siteid")

	// 获取栏目信息
	cateinfo, err := catemenusService.GetCateMenus(uint(catId))
	if err != nil {
		global.GVA_LOG.Error("获取栏目信息错误!", zap.Error(err))
		response.Fail500Error(c)
		return
	}
	// 判断是否选择了模板文件
	if cateinfo.ShowTemplate == "" {
		global.GVA_LOG.Error(fmt.Sprintf("栏目详情页模板 %s 暂未设置!", cateinfo.Name))
		response.Fail500Error(c)
		return
	}
	parentinfo := webcms.CateMenus{}
	if cateinfo.ParentId != 0 {
		// 获取父级栏目信息
		parentinfo, err = catemenusService.GetCateMenus(cateinfo.ParentId)
		if err != nil {
			response.Fail500Error(c)
			return
		}
	}
	res := getwebcommon(siteid)
	// 通过栏目模型进行数据获取
	modetype := cateinfo.ModeType
	showinfo, err := ShowService.GetInfoById(modetype, uint(showId))
	if err != nil {
		response.Fail500Error(c)
		return
	}
	res["showinfo"] = showinfo
	nextInfo, previousInfo = ShowService.GetInfoNextPreviousById(modetype, catId, showId, cateinfo.OrderType)
	// 生成上一页或者 下一页
	res["next_url"] = "#"
	if nextInfo.ID != 0 {
		res["next"] = gin.H{
			"url":   fmt.Sprintf("/show/%d/%d", catId, nextInfo.ID),
			"title": nextInfo.Title,
		}
	}
	res["previous_url"] = "#"
	if previousInfo.ID != 0 {
		res["previous"] = gin.H{
			"url":   fmt.Sprintf("/show/%d/%d", catId, previousInfo.ID),
			"title": previousInfo.Title,
		}
	}
	res["cateinfo"] = cateinfo
	res["parentinfo"] = parentinfo
	c.HTML(http.StatusOK, cateinfo.ShowTemplate, res)
}
