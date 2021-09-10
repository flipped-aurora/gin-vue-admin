package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/kardianos/service"
	"log"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func SerStart() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

// 当前应用程序的结构体
type program struct{}

// Start 服务启动
func (p *program) Start(s service.Service) error {
	log.Println("开始服务")
	go p.run()
	return nil
}

// Stop 停止程序
func (p *program) Stop(s service.Service) error {
	log.Println("停止服务")
	return nil
}

// 启动程序
func (p *program) run() {
	//API初始化
	SerStart()
}

func main() {
	//服务的配置信息
	cfg := &service.Config{
		Name:        "gin-vue-admin",
		DisplayName: "GIN-VUE-ADMIN",
		Description: "一个基于vue和gin开发的全栈前后端分离的后台管理系统",
	}
	// 程序的接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	// 如果有参数的话我们按照对应的参数执行
	if len(os.Args) == 2 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// 如果没有那么我们就直接启动
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}
