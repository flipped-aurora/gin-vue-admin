package tencentcloudapi

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	str, _ := Translate("番茄")
	fmt.Println(str)
}
