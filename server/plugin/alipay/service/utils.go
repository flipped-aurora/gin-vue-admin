package service

import (
	"strconv"
)

// 将传过来的字符串转成int64  例如：0.02 to 2
func StrToInt64(s string) (*int64, error) {
	s1, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	s2 := int64(s1 * 100)
	return &s2, nil
}
