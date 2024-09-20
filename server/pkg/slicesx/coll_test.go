package slicesx

import (
	"fmt"
	"testing"
)

func TestGetUniqueElements(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}
	elements := GetUniqueElementsFromB(a, b)
	fmt.Println(elements) //[4 5]
	elements = GetUniqueElementsFromB(b, a)
	fmt.Println(elements) //[1 2]

	c := []string{"a", "b", "c"}
	d := []string{"c", "d", "e"}
	elements1 := GetUniqueElementsFromB(c, d)
	fmt.Println(elements1) //[d e]

	elements1 = GetUniqueElementsFromB(d, c)
	fmt.Println(elements1) //[a b]
}
