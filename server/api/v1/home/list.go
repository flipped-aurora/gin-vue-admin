package home

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ListApi struct{}

var (
	listService = service.ServiceGroupApp.HomeServiceGroup.ListService
	contentlist interface{}
	total       int64
)

func (l *ListApi) List(c *gin.Context) {
	id := c.Param("catid")
	catid, err := strconv.Atoi(id)
	if err != nil {
		global.GVA_LOG.Error("栏目id参数错误!", zap.Error(err))
		response.Fail404Error(c)
		return
	}
	p := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(p)
	if err != nil {
		global.GVA_LOG.Error("分页编号转换失败!", zap.Error(err))
		response.Fail404Error(c)
		return
	}
	// 获取站点id
	siteid := c.GetInt("siteid")
	info, contentlist, err := l.ContentLists(catid, page, 0)
	if err != nil {
		global.GVA_LOG.Error("获取模型列表 err:", zap.Error(err))
		response.Fail500Error(c)
		return
	}
	parentinfo := webcms.CateMenus{}
	if info.ParentId != 0 {
		// 获取父级栏目信息
		parentinfo, err = catemenusService.GetCateMenus(info.ParentId)
		if err != nil {
			response.Fail500Error(c)
			return
		}
	}
	res := getwebcommon(siteid)
	res["title"] = fmt.Sprint(info.Name, "-", res["company"])
	res["info"] = info
	res["parentinfo"] = parentinfo
	res["contentlist"] = contentlist
	res["page"] = int(page)
	res["total"] = total
	// 生成静态文件
	// tempath := GetGenerateHtml(data, info.ListTemplate, fmt.Sprint("list-", info.ID, "-", page, ".html"))
	// if tempath != "" {
	// 	c.HTML(http.StatusOK, tempath, nil)
	// } else {
	// 获取
	c.HTML(http.StatusOK, info.ListTemplate, res)
	// }

}

func (l *ListApi) ContentLists(catid, page, num int) (info webcms.CateMenus, contentlist any, err error) {
	var pageInfo request.PageInfo
	pageInfo.Page = page
	// 获取栏目信息
	info, err = catemenusService.GetCateMenus(uint(catid))
	if err != nil {
		global.GVA_LOG.Error("获取栏目信息错误!", zap.Error(err))
		return
	}
	// 判断是否选择了模板文件
	if info.ListTemplate == "" {
		global.GVA_LOG.Error(fmt.Sprintf("栏目列表页模板 %s 暂未设置!", info.Name))
		err = errors.New(fmt.Sprintf("栏目列表页模板 %s 暂未设置!", info.Name))
		return
	}
	// 通过栏目模型进行数据获取
	modetype := info.ModeType
	// 每页数量
	pageInfo.PageSize = info.PgSize
	if num != 0 {
		pageInfo.PageSize = num
	}

	// 排序方式
	if info.OrderType == "1" {
		// 升序
		pageInfo.OrderType = "id ASC"
	} else {
		pageInfo.OrderType = "id DESC"
	}
	contentlist, total, err = listService.GetModelListByCatId(pageInfo, catid, modetype)
	return
}
