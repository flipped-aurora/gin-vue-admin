package main

import "github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"

func main() {
	r := runner.New()
	r.AddCmd("run_go_version", func(ctx *runner.Context) {
		request := ctx.Request
		ctx.Response(request + "-->" + "response")
	})

	r.Run()
}
