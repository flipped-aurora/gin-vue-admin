package home

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

type HomeApi struct{}

var (
	webconfigService = service.ServiceGroupApp.WebcmsServiceGroup.WebconfigService
	linksService     = service.ServiceGroupApp.WebcmsServiceGroup.LinksService
	homeService      = service.ServiceGroupApp.HomeServiceGroup.HomeService
	swiperService    = service.ServiceGroupApp.WebcmsServiceGroup.SwiperService
)

func (h *HomeApi) Home(c *gin.Context) {
	// 获取站点id
	siteid := c.GetInt("siteid")
	links, err := linksService.GetLinksinfo()
	if err != nil {
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
	}
	webconfig := getwebcommon(siteid)
	webconfig["title"] = fmt.Sprint("首页", "-", webconfig["company"])
	webconfig["links"] = links
	c.HTML(http.StatusOK, "index.html", webconfig)
}

// 搜索
func (h *HomeApi) Search(c *gin.Context) {
	p := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(p)
	if err != nil {
		response.Fail500Error(c)
		return
	}
	keyword := c.Query("keyword")
	var pageInfo request.PageInfo
	pageInfo.Page = page
	pageInfo.PageSize = 10
	pageInfo.Keyword = keyword
	searchlist, err := homeService.SearchInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
	}

	c.HTML(http.StatusOK, "search.html", gin.H{
		"title":      "Search Result",
		"searchlist": searchlist,
	})
}
