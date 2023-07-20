package utils

func UintArrContains(arr []uint, target uint) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func StringArrContains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
