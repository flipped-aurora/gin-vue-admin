package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONMap map[string]interface{}

// 为了兼容无json类型的数据库实现JSONMap类型的Scan和Value方法
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Scan makes the JSONMap implement the sql.Scanner interface.
// This method decodes a JSON-encoded value into the map structure.
func (m *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*m = make(map[string]interface{})
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}
	err := json.Unmarshal(bytes, m)
	if err != nil {
		return err
	}
	return nil
}
