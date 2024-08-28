package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"github.com/sirupsen/logrus"
)

type EchoRequest struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := runner.New()
	r.Post("echo", func(ctx *runner.Context) {
		//var req EchoRequest
		//err := ctx.ShouldBindJSON(&req)
		//if err != nil {
		//	ctx.ResponseJSON(&Response{Code: -1, Message: err.Error()})
		//	return
		//}
		logrus.Infof("echo" + jsonx.JSONString(ctx.ReqMap()))

		s := struct {
			Label   string                 `json:"label"`
			Content string                 `json:"content"`
			Value   int                    `json:"value"`
			Data    map[string]interface{} `json:"data"`
		}{
			Label: "json",
			//Content: jsonx.JSONString(ctx.ReqMap()),
			Data:  ctx.ReqMap(),
			Value: 0,
		}

		err := ctx.ResponseOkWithJSON(s)
		if err != nil {
			panic(err)
		}
	})

	r.Get("file", func(ctx *runner.Context) {
		logrus.Infof("file" + jsonx.JSONString(ctx.ReqMap()))
		//ctx.ReqMap()
		jsonx.SaveFile("./request.json", ctx.ReqMap())
		ctx.ResponseOkWithFile("./request.json", true)
	})

	r.Get("helloWorld", func(ctx *runner.Context) {
		//ctx.ReqMap()
		logrus.Infof("helloWorld" + jsonx.JSONString(ctx.ReqMap()))
		ctx.ResponseOkWithText("hello world")
	})
	r.Get("ping", func(ctx *runner.Context) {
		//ctx.ReqMap()
		//logrus.Infof("helloWorld" + jsonx.JSONString(ctx.ReqMap()))
		ctx.ResponseOkWithJSON(map[string]interface{}{
			"msg": "ok",
		})
	})

	r.Run()
}
