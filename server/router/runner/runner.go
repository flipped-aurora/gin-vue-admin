package runner

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/httpx"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/ElasticCache"
	"strings"
	"time"
)

type Runner struct{}

func Proxy(cache ElasticCache.ElasticCache) gin.HandlerFunc {
	return func(c *gin.Context) {

		proxy, err := httpx.NewHttpProxy("http://127.0.0.1:17777")
		if err != nil {
			panic(err)
		}

		path := c.Param("path")
		// 修改请求路径，去掉/proxy前缀
		c.Request.URL.Path = "/runner/" + strings.TrimPrefix(path, "/")
		split := strings.Split(strings.TrimPrefix(path, "/"), "/")
		user := split[1]
		soft := split[2]

		key := fmt.Sprintf("%s:%s", user, soft)
		data := cache.GetAndSet(key, time.Second*5, func(key string) (data interface{}, whetherCache bool) {
			var info biz_apphub.BizToolCmdSrvApi
			err := global.GVA_DB.Debug().Model(&biz_apphub.BizToolCmdSrvApi{}).
				Where("app_code = ? and tenant_user = ?", soft, user).First(&info).Error
			if err != nil {
				return err, false
			}
			return &info, true
		})
		if v, ok := data.(error); ok {
			response.FailWithMessage(v.Error(), c)
			return
		}
		info, ok := data.(*biz_apphub.BizToolCmdSrvApi)
		if !ok || info.ID == 0 {
			response.FailWithMessage("请输入正确地址", c)
			return
		}

		query := c.Request.URL.Query()
		query.Add("_version", info.Version)
		query.Add("_type", enum.RunnerTypeMap[info.ToolType])
		c.Request.URL.RawQuery = query.Encode()
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// InitRunnerRouter 初始化 Runner 路由信息
func (s *Runner) InitRunnerRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	//newProxy, err := httpx.NewHttpProxy("http://127.0.0.1:17777")
	//if err != nil {
	//	panic(err)
	//}
	//这里应该调用关闭
	//cache := ElasticCache.New(time.Second * 5)

	{
		//Proxy
		PublicRouter.GET("/runner/*path", Proxy(cache.ProxyCache))
		PublicRouter.POST("/runner/*path", Proxy(cache.ProxyCache))
	}
}
