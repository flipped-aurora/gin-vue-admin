// Package slicesx ...
package slicesx

import (
	"reflect"
)

// CollectionEquationStrings 两个集合的元素相等，不考虑排序问题，a,b,c ~ c,b,a
func CollectionEquationStrings(list1 []string, list2 []string) bool {
	if list1 == nil && list2 == nil {
		return true
	}
	if len(list1) != len(list2) {
		return false
	}
	mp := make(map[string]struct{})
	for _, data := range list1 {
		mp[data] = struct{}{}
	}

	for _, data := range list2 {
		if _, ok := mp[data]; !ok {
			return false
		}
	}
	return true
}

// In ...
func In[T any](list []T, e T) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}

func MapIntGroup[T any](
	list []T, selector func(T) int, filters ...func(key interface{}, list []T) []T) map[int][]T {
	mp := make(map[int][]T)
	for _, data := range list {
		key := selector(data)
		if _, ok := mp[key]; ok {
			mp[key] = append(mp[key], data)
		} else {
			mp[key] = make([]T, 0, 1)
			mp[key] = append(mp[key], data)
		}
	}
	if filters != nil {
		for _, filter := range filters {
			for k, v := range mp {
				res := filter(k, v)
				if res == nil {
					delete(mp, k)
				} else {
					mp[k] = res
				}
			}
		}
	}
	return mp
}
func ContainsString(list []string, el string) bool {
	for _, e := range list {
		if e == el {
			return true
		}
	}
	return false
}
func RemoveString(list []string, el string) []string {
	res := []string{}
	for _, e := range list {
		if e != el {
			res = append(res, e)
		}
	}
	return res
}

func Range[T any](list []T, fn func(T) T) []T {
	var res []T
	for _, t := range list {
		res = append(res, fn(t))
	}
	return res
}

func RemoveBy[T any](list []T, fn func(T) bool) []T {
	var res []T
	for _, t := range list {
		del := fn(t)
		if !del {
			res = append(res, t)
		}
	}
	return res
}

func AddFirst[T any](list []T, el T) []T {
	var res []T
	res = append(res, el)
	res = append(res, list...)

	return res
}

// RemoveExistElement
// source=a;b;c;g
// elements=a;c;d;e;f
// return d;e;f
func RemoveExistElement(elements []string, source []string) []string {
	return Select(source, func(s string) bool {
		return !ContainsString(elements, s)
	})
}
