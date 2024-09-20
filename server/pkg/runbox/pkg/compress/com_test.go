package compress

import (
	"fmt"
	"testing"
)

func TestDeCompress(t *testing.T) {
	//dir := os.TempDir() +
	path, err := DeCompress("D:\\tmp\\dist.zip", "D:\\tmp")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

}
func TestDeCompressx(t *testing.T) {
	//dir := os.TempDir() +
	err := DeCompressx("D:\\tmp\\dist.zip", "D:\\tmp\\dist")
	if err != nil {
		panic(err)
	}
	fmt.Println(err)

}
