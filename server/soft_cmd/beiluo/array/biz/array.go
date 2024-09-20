package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/slicesx"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"strconv"
	"strings"
)

func Diff(ctx *runner.Context) {

	req := struct {
		Base    string `json:"base"`
		Current string `json:"current"`
		Stp     string `json:"stp"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	bases := strings.Split(req.Base, req.Stp)
	current := strings.Split(req.Current, req.Stp)
	add, remove := slicesx.Diff(bases, current)
	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"add":    strings.Join(add, req.Stp),
			"remove": strings.Join(remove, req.Stp),
		},
	})
}
func Split(ctx *runner.Context) {
	req := struct {
		List string `json:"list"`
		Stp  string `json:"stp"`
		Size string `json:"size"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	bases := strings.Split(req.List, req.Stp)
	size, err := strconv.ParseInt(req.Size, 10, 64)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	//current := strings.Split(req.Current, req.Stp)
	slice := slicesx.SplitSlice(bases, int(size))
	s := ""
	for _, v := range slice {
		s += strings.Join(v, req.Stp) + "\n\n"
	}
	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"splitList": s,
		},
	})
}
