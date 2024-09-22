package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"time"
)

func main() {
	r := runner.New()
	r.Get("ts", func(ctx *runner.Context) {
		ctx.ResponseOkWithJSON(map[string]interface{}{
			"code": 0,
			"msg":  "ok",
			"data": map[string]interface{}{
				"ts": time.Now().Unix(),
			},
		})
	})
	r.Run()
}
