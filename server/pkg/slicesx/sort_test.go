package slicesx

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {

	list := Sort([]int{4, 2, 3, 5, 1, 6}, func(a, b int) bool {
		return a > b
	})
	fmt.Println(list) //[6 5 4 3 2 1]
	list = Sort([]int{4, 2, 3, 5, 1, 6}, func(a, b int) bool {
		return a < b
	})
	fmt.Println(list) //[1 2 3 4 5 6]
}
