package slicesx

type Eq interface {
	string | int | int64 | int32 | int16 | int8 | float32 | float64 | uint | uint8 | uint32 | uint64
}

// GetUniqueElementsFromB 取B集合的差集（B−A），返回B集合存在但是A集合不存在的元素,A={1,2,3},B={3,4,5}       output  B−A={4,5}
func GetUniqueElementsFromB[T Eq](A, B []T) []T {
	var uniqueElements []T

	// 创建一个map来存储a中的元素，以便快速查找
	aMap := make(map[interface{}]bool)
	for _, elem := range A {
		aMap[elem] = true
	}

	// 遍历b，找出不在a中的元素
	for _, elem := range B {
		if !aMap[elem] {
			uniqueElements = append(uniqueElements, elem)
		}
	}

	return uniqueElements
}
