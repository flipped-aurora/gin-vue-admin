package utils

import (
	"encoding/json"
	"strings"
)

func GetJSONKeys(jsonStr string) (keys []string, err error) {
	// 使用json.Decoder，以便在解析过程中记录键的顺序
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	// 确保数据是一个对象
	if t != json.Delim('{') {
		return nil, err
	}
	for dec.More() {
		t, err = dec.Token()
		if err != nil {
			return nil, err
		}
		keys = append(keys, t.(string))

		// 解析值
		var value interface{}
		err = dec.Decode(&value)
		if err != nil {
			return nil, err
		}
	}
	return keys, nil
}
