package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"net/http"
	"time"
)

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	// 插件安装 暂时只是后台功能 添加model 添加路由 添加对数据库的操作  详细插件测试模板可看https://github.com/piexlmax/gvaplug  此处不建议投入生产
	//err := initialize.InstallPlug(global.GVA_DB, Router, gvaplug.GvaPlug{})
	//if err != nil {
	//	panic(fmt.Sprintf("插件安装失败： %v", err))
	//}
	// end 插件描述

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用 Gin-Vue-Admin
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	global.GVA_LOG.Error(s.ListenAndServe())
}
