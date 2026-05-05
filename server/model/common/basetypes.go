package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JSONMap map[string]any

type JSONSlice[T any] []T

var mysqlJSONTypeCache sync.Map

func (JSONMap) GormDataType() string {
	return "json"
}

func (JSONMap) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return resolveJSONDBDataType(db)
}

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m *JSONMap) Scan(value any) error {
	if value == nil {
		*m = make(map[string]any)
		return nil
	}
	var err error
	switch typed := value.(type) {
	case []byte:
		err = json.Unmarshal(typed, m)
	case string:
		err = json.Unmarshal([]byte(typed), m)
	default:
		err = errors.New("basetypes.JSONMap.Scan: invalid value type")
	}
	if err != nil {
		return err
	}
	return nil
}

func (JSONSlice[T]) GormDataType() string {
	return "json"
}

func (JSONSlice[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return resolveJSONDBDataType(db)
}

func (s JSONSlice[T]) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *JSONSlice[T]) Scan(value any) error {
	if value == nil {
		*s = JSONSlice[T]{}
		return nil
	}
	var err error
	switch typed := value.(type) {
	case []byte:
		err = json.Unmarshal(typed, s)
	case string:
		err = json.Unmarshal([]byte(typed), s)
	default:
		err = errors.New("basetypes.JSONSlice.Scan: invalid value type")
	}
	if err != nil {
		return err
	}
	return nil
}

func resolveJSONDBDataType(db *gorm.DB) string {
	switch db.Dialector.Name() {
	case "mysql":
		if mysqlSupportsJSON(db) {
			return "JSON"
		}
		return "LONGTEXT"
	case "postgres":
		return "JSONB"
	default:
		return "JSON"
	}
}

func mysqlSupportsJSON(db *gorm.DB) bool {
	sqlDB, err := db.DB()
	if err != nil {
		return true
	}

	cacheKey := sqlDB
	if cached, ok := mysqlJSONTypeCache.Load(cacheKey); ok {
		return cached.(bool)
	}

	supports := true
	var version string
	if err := db.Raw("SELECT VERSION()").Scan(&version).Error; err == nil {
		lowerVersion := strings.ToLower(version)
		if strings.Contains(lowerVersion, "mariadb") {
			supports = false
		} else if strings.HasPrefix(version, "5.5.") || strings.HasPrefix(version, "5.6.") {
			supports = false
		}
	}

	mysqlJSONTypeCache.Store(cacheKey, supports)
	return supports
}

type TreeNode[T any] interface {
	GetChildren() []T
	SetChildren(children T)
	GetID() int
	GetParentID() int
}
