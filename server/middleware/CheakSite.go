package middleware

import (
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CheckSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取配置文件siteid
		session := sessions.Default(c)
		siteid := session.Get("siteid")
		// 未有获取到数据 重新查询表第一条数据 并保存数据
		var webconfig map[string]any
		if siteid == nil {
			err := global.GVA_DB.Model(webcms.Webconfig{}).First(&webconfig).Error
			if err != nil {
				global.GVA_LOG.Error("获取默认站点信息失败!", zap.Error(err))
			}
			siteid = webconfig["id"]
			session.Set("siteid", siteid)
			session.Save()
		}
		err := global.GVA_DB.Model(webcms.Webconfig{}).Where("id", siteid).Find(&webconfig).Error
		if err != nil {
			global.GVA_LOG.Error("获取默认站点信息失败!", zap.Error(err))
		}
		c.Set("siteinfo", webconfig)
		c.Next()
	}
}
func HomeCheckSite() gin.HandlerFunc {
	// 通过域名进行判断站点 并获取相应站点的数据
	return func(c *gin.Context) {
		// 1.获取访问链接
		scheme := "http://"
		host := c.Request.Host
		if c.Request.TLS != nil || c.Request.Header.Get("X-Forwarded-Proto") == "https" {
			scheme = "https://"
		}
		url := strings.Join([]string{scheme, host, "/"}, "")

		// fmt.Println(url, "1111")
		//获取站点信息
		siteinfo := webcms.Webconfig{}
		// 查不到站点 跳转404
		err := global.GVA_DB.Table("webconfig").Where("site_url = ?", url).First(&siteinfo).Error
		if err != nil {
			global.GVA_LOG.Error("查不到站点 跳转404", zap.Error(err))
			c.HTML(http.StatusOK, "404.html", gin.H{
				"title": "404",
			})
			c.Abort()
		} else {
			c.Set("siteinfo", siteinfo)
			c.Next()
		}
		c.Next()
	}
}
