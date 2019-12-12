package main

import (
	"fmt"
	"main/config"
	"main/init/initRouter"
	"main/init/qmlog"
	"main/init/qmsql"
	"main/init/registTable"
	"net/http"
	"time"
)


// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

func main() {
	qmlog.InitLog()                            // 初始化日志
	db:=qmsql.InitMysql(config.Dbconfig.Admin) // 链接初始化数据库
	registTable.RegistTable(db)                //注册数据库表
	defer qmsql.DEFAULTDB.Close()              // 程序结束前关闭数据库链接
	Router := initRouter.InitRouter()          //注册路由
	qmlog.QMLog.Info("服务器开启")             // 日志测试代码
	//Router.RunTLS(":443","ssl.pem", "ssl.key")  // https支持 需要添加中间件
	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`欢迎使用 Gin-Vue-Admin
作者：奇淼 And Spike666
微信：shouzi_1994
默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	_ = s.ListenAndServe()
}
