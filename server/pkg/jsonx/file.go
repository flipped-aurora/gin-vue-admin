package jsonx

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func UnmarshalFromFile(filePath string, v interface{}) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, v)
}

func SaveFile(filePath string, el interface{}) error {
	marshal, err := json.Marshal(el)
	if err != nil {
		return err
	}
	backCount := 0
back:
	if backCount > 1 {
		return fmt.Errorf("SaveFile backCount > 1")
	}
	file, err := os.Create(filePath)
	if err != nil {
		s := err.Error()
		if strings.Contains(s, "The system cannot find the path specified") { //如果目录不存在就篡改就目录
			split := strings.Split(filePath, "/")
			p := strings.Join(split[0:len(split)-1], "/")
			if len(split) > 2 {
				err = os.MkdirAll(p, os.ModePerm)
				if err != nil {
					return err
				}
				err = nil
				backCount++
				goto back
			}
		}

		return err
	}
	defer file.Close()
	_, err = file.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
