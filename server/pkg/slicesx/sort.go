package slicesx

// Sort a > b 降序
func Sort[T any](arr []T, less func(a, b T) bool) []T {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if less(arr[j], arr[i]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
