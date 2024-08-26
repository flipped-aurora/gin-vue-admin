package jsonx

import (
	"encoding/json"
	"os"
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
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
