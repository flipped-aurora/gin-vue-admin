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

type CategoryApi struct{}

var catemenusService = service.ServiceGroupApp.WebcmsServiceGroup.CateMenusService

func (h *CategoryApi) Category(c *gin.Context) {
	catid := c.Param("id")
	id, err := strconv.Atoi(catid)
	if err != nil {
		global.GVA_LOG.Error("栏目id参数错误!", zap.Error(err))
		response.Fail404Error(c)
		return
	}
	// 获取栏目信息
	info, err := catemenusService.GetCateMenus(uint(id))
	if err != nil {
		global.GVA_LOG.Error("获取栏目信息错误!", zap.Error(err))
		response.Fail500Error(c)
		return
	}
	// 跳转到子栏目
	if info.IsJump {
		childinfo, _ := catemenusService.GetCateMenusListByPid(info.ID)
		if len(childinfo) > 0 {
			c.Redirect(http.StatusMovedPermanently, childinfo[0].Url)
			return
		}
	}
	// 判断是否选择了模板文件
	if info.ListTemplate == "" {
		global.GVA_LOG.Error(fmt.Sprintf("栏目模板 %s 暂未设置!", info.Name))
		response.Fail500Error(c)
		return
	}

	parentinfo := webcms.CateMenus{}
	if info.ParentId != 0 {
		// 获取父级栏目信息
		parentinfo, err = catemenusService.GetCateMenus(info.ParentId)
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("获取父级栏目信息失败 err: %s ", info.Name))
			response.Fail500Error(c)
			return
		}
	}
	// 获取站点信息
	siteinfo := c.GetStringMap("siteinfo")
	res := getwebcommon(siteinfo["site_url"].(string))
	res["title"] = fmt.Sprint(info.Name, "-", res["company"])
	res["info"] = info
	res["parentinfo"] = parentinfo
	c.HTML(http.StatusOK, info.ListTemplate, res)
}
