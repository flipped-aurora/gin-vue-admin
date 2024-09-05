package main

import "github.com/flipped-aurora/gin-vue-admin/server/pkg/osx"

func main() {
	err := osx.CopyDirectory("D:/code/github.com/apphub/server/test/v1", "D:/code/github.com/apphub/server/test/v2")
	if err != nil {
		panic(err)
	}
}
