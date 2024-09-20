// Package slicesx ...
package slicesx

import (
	"fmt"
	"testing"
)

// TestIn ...
func TestIn(t *testing.T) {

	type User struct {
		Name string
		Age  int
	}

	list := []int{1, 2, 3, 4, 5}
	xh := User{"xh", 21}
	bl := User{Name: "bl", Age: 22}
	xy := User{Name: "xy", Age: 22}
	list2 := []User{bl, xh}
	fmt.Println(In(list, 1))
	fmt.Println(In(list, 7))
	fmt.Println(In(list, 9))
	fmt.Println(In(list, 1))

	fmt.Println(In(list2, bl))
	fmt.Println(In(list2, xh))
	fmt.Println(In(list2, xy))

	//
	//true
	//false
	//false
	//true
	//true
	//true
	//false
	//--- PASS: TestIn (0.00s)
	//PASS
}

func TestRemoveExistElement(t *testing.T) {
	element := RemoveExistElement([]string{"1", "2", "3"}, []string{"3", "4", "5"})
	fmt.Println(element)

	element = RemoveExistElement([]string{"1", "2"}, []string{"3", "4", "5"})
	fmt.Println(element)

	element = RemoveExistElement([]string{"3", "4", "5"}, []string{"3", "4", "5"})
	fmt.Println(element)
}
