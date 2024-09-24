package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"github.com/flipped-aurora/gin-vue-admin/server/soft_cmd/beiluo/array/biz"
)

func main() {
	r := runner.New()
	r.Post("diff", biz.Diff)
	r.Post("Split", biz.Split)
	r.Post("ComputeIntersection", biz.ComputeIntersection)
	r.Run()
}
