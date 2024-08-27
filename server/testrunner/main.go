package main

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"os/exec"
	"time"
)

func main() {
	r := runner.New()
	r.AddCmd("run_go_version", func(ctx *runner.Context) {
		request := ctx.Request
		now := time.Now()
		cmd := exec.Command("go", "env")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Output:", out.String())
		fmt.Println("cost:", time.Since(now))
		ctx.Response(request + "-->" + "response")
	})

	r.Run()
}
