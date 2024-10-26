package myutil

import (
	"encoding/json"
	"gorm.io/datatypes"
	"strconv"
)

func ArryjsonAdd(listField datatypes.JSON, addr string) (datatypes.JSON, error) {
	var existingList []string
	if listField != nil {
		if err := json.Unmarshal([]byte(listField), &existingList); err != nil {
			return nil, err
		}
	}

	// 将 address 加入到 existingList 中
	existingList = append(existingList, addr)

	// 将 addressList 转换为 JSON 字符串
	marshal, err := json.Marshal(existingList)
	if err != nil {
		return nil, err
	}

	return datatypes.JSON(marshal), nil
}

// 将字符串索引值转换为数字，加1后再次转换为字符串
func IncrementIndex(index string, add int) (string, error) {
	// 将字符串转换为整数
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		return "", err
	}

	// 加1
	indexInt = indexInt + add

	// 再次转换为字符串
	newIndexStr := strconv.Itoa(indexInt)

	return newIndexStr, nil
}
