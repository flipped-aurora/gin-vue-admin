package slicesx

import (
	"fmt"
	"testing"
)

func TestDiff(t *testing.T) {
	add, remove := Diff([]string{"1", "2", "3", "4"}, []string{"3", "4", "5", "6"})
	fmt.Println(add)
	fmt.Println(remove)
}
