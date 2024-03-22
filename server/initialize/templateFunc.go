package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"math"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

var (
	cateMenusService = service.ServiceGroupApp.WebcmsServiceGroup.CateMenusService
	courseService    = service.ServiceGroupApp.WebcmsServiceGroup.CourseService
	classService     = service.ServiceGroupApp.WebcmsServiceGroup.ClassService
	webconfigService = service.ServiceGroupApp.WebcmsServiceGroup.WebconfigService
	swiperService    = service.ServiceGroupApp.WebcmsServiceGroup.SwiperService
	listApi          = v1.ApiGroupApp.HomeApiGroup.ListApi
)

func TemplateFunc() template.FuncMap {
	return template.FuncMap{
		"category":          category,
		"safe":              safe,
		"list":              list,
		"json2map":          json2map,
		"paginator":         paginator,
		"getswiper":         getswiper,
		"getallInfobycatid": getAllInfoByCatid,
		"catpos":            catpos,
		"getwebconfig":      getwebconfig,
		"uinttoint":         uinttoint,
		"getcateinfo":       getcateinfo,
	}
}

// 获取栏目信息
func getcateinfo(catid int) (info webcms.CateMenus) {
	info, err := cateMenusService.GetCateMenus(uint(catid))
	if err != nil {
		global.GVA_LOG.Error("获取信息失败!", zap.Error(err))
	}
	return
}

// 获取banner 图
func getswiper(siteid int) (list []webcms.Swiper) {
	list = swiperService.GetSwiperList(siteid)
	return
}

// 获取getwebconfig 参数
func getwebconfig(siteid int) map[string]string {
	res := global.GVA_REDIS.HGetAll(context.Background(), fmt.Sprint("webconfig:", siteid))
	if res.Err() != nil {
		global.GVA_LOG.Error("获取参数失败!", zap.Error(res.Err()))
		return nil
	} else {
		return res.Val()
	}
}

// json 转为 map
func json2map(str string) (mapdata []map[string]interface{}) {
	mapdata = make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(str), &mapdata)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func safe(s string) template.HTML {
	return template.HTML(s)
}

func uinttoint(param uint) int {
	return int(param)
}

// 获取栏目列表
func category(pid uint, siteid int) (list []webcms.CateMenus) {
	if pid == 0 {
		list, _ = cateMenusService.GetCateMenusInfoList2(siteid)
	} else {
		list, _ = cateMenusService.GetCateMenusListByPid(pid)
	}
	return
}

// 获取产品列表
func list(catid, page, num int) (list any) {
	_, list, _ = listApi.ContentLists(catid, page, num)
	return
}

// 获取当前栏目下所有产品
func getAllInfoByCatid(catid uint) (list []webcms.Course) {
	list, _ = courseService.GetAllInfoByCatid(catid)
	return
}

// 分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page

	return paginatorMap
}

/**
 * 当前路径
 * 返回指定栏目路径层级
 * @param catid 栏目id
 * @param symbol 栏目间隔符
 */
func catpos(catid uint, symbol string) string {
	// <a href="">首页</a> > 关于禹冰
	info, err := cateMenusService.GetCateMenus(catid)
	html := " > "
	if err != nil {
		return ""
	}
	if info.ParentId == 0 {
		return html + " <a href=" + info.Url + ">" + info.Name + "</a> "
	} else {
		html = html + " <a href=" + info.Url + ">" + info.Name + "</a> "
		return catpos(info.ParentId, symbol) + html
	}

}
