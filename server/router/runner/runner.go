package runner

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/httpx"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/ElasticCache"
	"github.com/nats-io/nats.go"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Runner struct{}

func urlQueryToMap(values url.Values) map[string]string {
	mp := make(map[string]string)
	for k, v := range values {
		if len(v) == 1 {
			mp[k] = v[0]
		}
		if len(v) == 0 {
			continue
		}
		if len(v) > 1 {
			mp[k] = strings.Join(v, ",")
		}
	}
	return mp
}

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

func HandMsg(c *gin.Context) {
	path := c.Param("path")
	split := strings.Split(strings.TrimPrefix(path, "/"), "/")
	user := split[1]
	soft := split[2]
	runnerCmd := strings.Join(split[3:], "/")
	key := fmt.Sprintf("%s:%s", user, soft)
	data := cache.ProxyCache.GetAndSet(key, time.Second*5, func(key string) (data interface{}, whetherCache bool) {
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

	sub := fmt.Sprintf("runner.run.%s.%s.%s", user, soft, runnerCmd)
	var body []byte
	if c.Request.Method == "POST" {
		body, _ = io.ReadAll(c.Request.Body)
	} else if c.Request.Method == "GET" {
		toMap := urlQueryToMap(c.Request.URL.Query())
		body, _ = json.Marshal(toMap)
	}

	msg := nats.NewMsg(sub)
	msg.Header.Set("soft", soft)
	msg.Header.Set("user", user)
	msg.Header.Set("cmd", runnerCmd)
	msg.Header.Set("version", info.Version)
	msg.Header.Set("method", c.Request.Method)
	msg.Header.Set("type", info.ToolType)
	msg.Header.Set("fs_path", info.OssPath)
	msg.Data = body
	resp, err := global.NatsClient.RequestMsg(msg, time.Second*200)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	errStr := resp.Header.Get("error")
	if errStr != "" { //系统级错误需要告知管理员和用户
		c.Data(500, c.Request.Header.Get("Content-Type"), nil)
		return
	}

	code := resp.Header.Get("Status-Code")
	fmt.Println(code)
	fmt.Println(resp.Header)
	status, err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		c.Data(500, c.Request.Header.Get("Content-Type"), nil)
		return
	}
	c.Data(int(status), resp.Header.Get("Content-Type"), resp.Data)
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
		//PublicRouter.GET("/runner/*path", Proxy(cache.ProxyCache))
		//PublicRouter.POST("/runner/*path", Proxy(cache.ProxyCache))

		PublicRouter.GET("/runner/*path", HandMsg)
		PublicRouter.POST("/runner/*path", HandMsg)
	}
}
