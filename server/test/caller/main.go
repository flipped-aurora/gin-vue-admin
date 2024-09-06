package main

import (
	"fmt"
	"runtime"
)

func Test() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println("pc:", pc, "file:", file, "line:", line, "ok:", ok)
}

func main() {
	Test()
}
