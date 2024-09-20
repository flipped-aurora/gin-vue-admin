package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/pkg/store"
)

func main() {
	fileStore := store.NewDefaultQiNiu()
	var err error
	//save, err := store.NewDefaultQiNiu().FileSave("./main.go", "runbox/beiluo/test/main.go")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(fmt.Sprintf("%+v", save))
	//file, err := fileStore.GetFile("/runbox/beiluo/test/main.go")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(fmt.Sprintf("%+v", file))
	err = fileStore.DeleteFile("runbox/beiluo/test/main.go")
	if err != nil {
		panic(err)
	}
}
