package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"os"
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
	r.AddCmd("echo", func(ctx *runner.Context) {
		var req EchoRequest
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.ResponseJSON(&Response{Code: -1, Message: err.Error()})
			return
		}
		defer os.Remove(ctx.Request)

		s := struct {
			Label   string `json:"label"`
			Content string `json:"content"`
			Value   int    `json:"value"`
		}{
			Label:   "json",
			Content: "hello world",
			Value:   req.PageNum * req.PageSize}

		err = ctx.ResponseJSON(s)
		if err != nil {
			panic(err)
		}
	})

	r.Run()
}
