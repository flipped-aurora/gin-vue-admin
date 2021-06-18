package core

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"time"

	"github.com/eyotang/game-proxy/server/global"
	"github.com/eyotang/game-proxy/server/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 功能面
	fcRouter := initialize.FuncRouters()
	fcAddr := fmt.Sprintf(":%d", global.GVA_CONFIG.System.FuncAddr)
	fs := initServer(fcAddr, fcRouter)
	binding.EnableDecoderUseNumber = true
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("function server run success on ", zap.String("address", fcAddr))
	go fs.ListenAndServe()

	// 管理面
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Game-Proxy
	当前版本:V0.1.0
    加群方式:微信号: eyotang2016
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
