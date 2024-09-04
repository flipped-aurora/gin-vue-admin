package main

import "github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"

func PublicConfig(desc string) *runner.Config {
	return runner.NewConfig(runner.WithIsPublicApi(true), runner.WithApiDesc(desc))
}
func main() {

	r := runner.New()
	r.Get("helloworld", func(ctx *runner.Context) {
		ctx.ResponseOkWithText("hello world")
	}, PublicConfig("helloworld 测试"))

	r.Get("json", func(ctx *runner.Context) {
		ctx.ResponseOkWithText("hello world")
	}, PublicConfig("json get 测试"))

	r.Get("text", func(ctx *runner.Context) {
		ctx.ResponseOkWithText("text hello world")
	}, PublicConfig("text get 测试"))

	r.Post("text", func(ctx *runner.Context) {
		ctx.ResponseOkWithText("text hello world")
	}, PublicConfig("text  post 测试"))
	r.Run()
}
