package slicesx

// Diff 函数，第一个参数是原始数据，第二个参数是当前数据，分析对比得出需要删除和需要添加的数据
// base原始数据 1,2,3,4
// current当前数据 3,4,5,6
// add需要添加的数据:  1,2
// remove需要删除的数据: 5,6
func Diff(base []string, current []string) (add []string, remove []string) {
	m := make(map[string]struct{}, len(base))
	m1 := make(map[string]struct{}, len(current))
	for i := 0; i < len(base); i++ {
		m[base[i]] = struct{}{}
	}

	for i := 0; i < len(current); i++ {
		m1[current[i]] = struct{}{}
	}
	for s, _ := range m {
		_, ok := m1[s]
		if !ok {
			add = append(add, s)
		}
	}

	for s, _ := range m1 {
		_, ok := m[s]
		if !ok {
			remove = append(remove, s)
		}
	}
	return
}
