package stringsx

import (
	"fmt"
	"testing"
)

func TestGetSuffix(t *testing.T) {
	fmt.Println(GetSuffix("1.txt", "."))
	fmt.Println(GetSuffix("1/jpg", "."))
	fmt.Println(GetSuffix("1.png.", "."))
	fmt.Println(GetSuffix("2.1.zip", "."))
	fmt.Println(GetSuffix("你好为3131.txt", "."))
}
