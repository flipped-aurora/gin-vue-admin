package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"strings"
)

func main() {

	r := runner.New()
	r.Post("split", func(ctx *runner.Context) {
		str := ctx.ReqMap()["str"].(string)
		separator := ctx.ReqMap()["separator"].(string)
		res := strings.Split(str, separator)
		join := strings.Join(res, "\n")
		ctx.ResponseOkWithJSON(map[string]interface{}{
			"code": 0,
			"msg":  "ok",
			"data": map[string]interface{}{
				"splitString": join,
			},
		})
	})
	r.Post("ReplaceAll", func(ctx *runner.Context) {

		str := ctx.ReqMap()["str"].(string)
		newStr := ctx.ReqMap()["new_str"].(string)
		oldStr := ctx.ReqMap()["old_str"].(string)
		all := strings.ReplaceAll(str, oldStr, newStr)

		ctx.ResponseOkWithJSON(map[string]interface{}{
			"code": 0,
			"msg":  "ok",
			"data": map[string]interface{}{
				"res": all,
			},
		})
	})
	r.Run()
}
