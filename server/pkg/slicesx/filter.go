// Package slicesx ...
package slicesx

// Select 选择指定条件元素
func Select[T any](list []T, fn func(T) bool) []T {
	var res []T
	for _, t := range list {
		if fn(t) {
			res = append(res, t)
		}
	}
	return res
}

// Find 查找元素
func Find[T any](list []T, fn func(T) bool) T {
	var n T
	for _, t := range list {
		if fn(t) {
			return t
		}
	}
	return n
}

// Filter 根据指定字段去重
func Filter[T, V any](list []T, selectField func(T) V) []T {
	mp := make(map[interface{}]struct{})
	res := make([]T, 0, len(list))
	for _, t := range list {
		key := selectField(t)
		if _, ok := mp[key]; !ok {
			res = append(res, t)
			mp[key] = struct{}{}
		}
	}
	return res
}

// Transfer 类型转换
func Transfer[T, V any](list []T, fn func(T) V) []V {
	var res []V
	for _, t := range list {
		res = append(res, fn(t))
	}
	return res
}

// SplitSlice 拆分切片
func SplitSlice[T any](slice []T, chunkSize int) [][]T {
	var result [][]T
	length := len(slice)
	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize
		if end > length {
			end = length
		}
		result = append(result, slice[i:end])
	}
	return result
}

// StringMap 切片根据指定字段转换成map
func StringMap[T any](list []T, selector func(T) string) map[string]T {
	mp := make(map[string]T)
	for _, data := range list {
		mp[selector(data)] = data
	}
	return mp
}

// MapValueCount 根据字段分组统计
func MapValueCount[T any](list []T, selector func(T) string) map[interface{}]int {
	mp := make(map[interface{}]int)
	for _, data := range list {
		key := selector(data)
		if _, ok := mp[key]; ok {
			mp[key]++
		} else {
			mp[key] = 1
		}
	}
	return mp
}

// MapValueGroup 根据指定字段分组，filter过滤
func MapValueGroup[T any](
	list []T, selector func(T) interface{}, filters ...func(key interface{}, list []T) []T) map[interface{}][]T {
	mp := make(map[interface{}][]T)
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

// GroupByField 根据指定字段分组，filter过滤
func GroupByField[T any](
	list []T, selector func(T) interface{}, filters ...func(key interface{}, list []T) []T) map[interface{}][]T {
	mp := make(map[interface{}][]T)
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

// Map 指定字段为key转换为map
func Map[T any](list []T, selector func(T) interface{}) map[interface{}]T {
	mp := make(map[interface{}]T)
	for _, data := range list {
		mp[selector(data)] = data
	}
	return mp
}

// Merge 将多个切片合并
func Merge[T any](lists ...[]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}
