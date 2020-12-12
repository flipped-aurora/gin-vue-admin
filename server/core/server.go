package core

import (
	"context"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

type mShutdown interface {
	Shutdown(ctx context.Context) error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Gin-Vue-Admin
	当前版本:V2.3.6
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`, address)
	//global.GVA_LOG.Error(s.ListenAndServe().Error())

	go func() {
		// 开启一个goroutine启动服务
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.GVA_LOG.Error(err.Error())
		}
	}()
	// 新增 优雅关闭
	exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	sigChan := make(chan os.Signal, 1)
	// 注册信号
	signal.Notify(sigChan, exitSignals...)
	<-sigChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	ns, ok := s.(mShutdown)
	if !ok {
		return
	}
	if err := ns.Shutdown(ctx); err != nil {
		global.GVA_LOG.Error("Server Shutdown: " + err.Error())
		return
	}
	global.GVA_LOG.Info("Server Exitz")
}
