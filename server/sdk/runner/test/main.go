package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {

	now := time.Now()
	cmd := exec.Command("D:\\code\\github.com\\apphub\\server\\testrunner\\testrunner.exe", "run_go_version")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", out.String())
	fmt.Println("cost:", time.Since(now))
}
