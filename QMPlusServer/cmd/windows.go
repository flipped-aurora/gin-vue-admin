package cmd

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/init/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RunWindowsServer(Router *gin.Engine, logger log.Logger) {
	address := fmt.Sprintf(":%d", config.GinVueAdminconfig.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	logger.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用 Gin-Vue-Admin
	作者：奇淼 And Spike666
	微信：shouzi_1994
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	_ = s.ListenAndServe()
}
