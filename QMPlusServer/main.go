package main

import (
	"main/config"
	"main/init"
	"net/http"
	"time"
)

// @Summary 打印测试功能
// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
// @Host 127.0.0.1:8080
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
// @Router / [get]
func main() {
	init.InitMysql(config.Dbconfig.Admin)
	defer init.DEFAULTDB.Close()
	init.InitRouter()
	s := &http.Server{
		Addr:           ":8888",
		Handler:        init.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
