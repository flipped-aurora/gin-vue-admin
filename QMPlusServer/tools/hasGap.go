// 空值校验工具 仅用于检验空字符串 其余类型请勿使用

package tools

import (
	"errors"
	"fmt"
	"reflect"
)

func HasGap(input interface{}) error {
	getType := reflect.TypeOf(input)
	fmt.Println("获取类型 :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("所有字段", getValue)

	// 获取方法字段
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
		if value == "" {
			return errors.New(fmt.Sprintf("%s为空", field.Name))
		}
	}
	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	//for i := 0; i < getType.NumMethod(); i++ {
	//	m := getType.Method(i)
	//	fmt.Printf("%s: %v\n", m.Name, m.Type)
	//}
	return nil
}
