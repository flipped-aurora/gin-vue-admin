package cmd

import (
	"fmt"
	"gin-vue-admin/config"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"syscall"
	"time"
)

func RunLinuxServer(Router *gin.Engine) {
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", config.GinVueAdminconfig.System.Addr)

	server := endless.NewServer(endPoint, Router)
	server.BeforeBegin = func(add string) {
		fmt.Printf(`欢迎使用 Gin-Vue-Admin
	作者：奇淼 And Spike666
	微信：shouzi_1994
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	Actual pid is %d
`, fmt.Sprintf(":%d", config.GinVueAdminconfig.System.Addr), syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server err: %v", err)
	}
}
