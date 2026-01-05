// Package common 提供通用的基础类型定义
//
// 功能说明：
// 本包定义了项目中常用的基础类型，包括：
// - JSONMap: 用于数据库JSON字段的映射类型
// - TreeNode: 树形结构的通用接口
package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONMap JSON映射类型，用于GORM的JSON字段
//
// 功能说明：
// 实现了driver.Valuer和sql.Scanner接口，使得map[string]interface{}类型
// 可以直接作为数据库的JSON字段类型使用。
//
// 使用场景：
// - 存储动态结构的JSON数据
// - 存储配置信息、扩展字段等
//
// 实现细节：
// - Value(): 将map序列化为JSON字符串存入数据库
// - Scan(): 从数据库读取JSON字符串并反序列化为map
//
// 使用示例：
//
//	type User struct {
//	    ID   uint
//	    Name string
//	    Meta common.JSONMap `gorm:"type:json"` // 使用JSONMap类型
//	}
//
//	// 写入数据
//	user := User{
//	    Name: "张三",
//	    Meta: common.JSONMap{
//	        "age": 25,
//	        "city": "北京",
//	        "tags": []string{"vip", "active"},
//	    },
//	}
//	db.Create(&user)
//
//	// 读取数据
//	var user User
//	db.First(&user, 1)
//	fmt.Println(user.Meta["age"]) // 输出: 25
type JSONMap map[string]interface{}

// Value 实现driver.Valuer接口，将JSONMap转换为数据库值
//
// 功能说明：
// 当GORM保存JSONMap到数据库时，会自动调用此方法将map序列化为JSON字符串。
//
// 返回值：
// - driver.Value: JSON字符串的字节数组或nil
// - error: 序列化错误
//
// 实现细节：
// - 如果map为nil，返回nil（数据库存储NULL）
// - 否则使用json.Marshal序列化为JSON字符串
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Scan 实现sql.Scanner接口，从数据库值扫描到JSONMap
//
// 功能说明：
// 当GORM从数据库读取JSON字段时，会自动调用此方法将JSON字符串反序列化为map。
//
// 参数说明：
// - value: 数据库中的值，可能是[]byte、string或nil
//
// 返回值：
// - error: 反序列化错误或类型不匹配错误
//
// 实现细节：
// - 如果value为nil，初始化为空map
// - 如果value是[]byte，直接反序列化
// - 如果value是string，先转换为[]byte再反序列化
// - 其他类型返回错误
func (m *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*m = make(map[string]interface{})
		return nil
	}
	var err error
	switch value.(type) {
	case []byte:
		// 数据库返回的JSON字节数组
		err = json.Unmarshal(value.([]byte), m)
	case string:
		// 某些数据库驱动可能返回字符串
		err = json.Unmarshal([]byte(value.(string)), m)
	default:
		err = errors.New("basetypes.JSONMap.Scan: invalid value type")
	}
	if err != nil {
		return err
	}
	return nil
}

// TreeNode 树形结构的通用接口
//
// 功能说明：
// 定义了树形结构节点的通用行为，任何需要构建树形结构的数据类型都可以实现此接口。
//
// 类型参数：
// - T: 节点类型，通常是实现接口的结构体本身
//
// 方法说明：
// - GetChildren(): 获取子节点列表
// - SetChildren(): 设置子节点
// - GetID(): 获取节点ID
// - GetParentID(): 获取父节点ID
//
// 使用场景：
// - 菜单树、部门树、分类树等层级结构
//
// 使用示例：
//
//	type Menu struct {
//	    ID       int
//	    ParentID int
//	    Name     string
//	    Children []Menu
//	}
//
//	func (m Menu) GetChildren() []Menu {
//	    return m.Children
//	}
//
//	func (m *Menu) SetChildren(children Menu) {
//	    m.Children = append(m.Children, children)
//	}
//
//	func (m Menu) GetID() int {
//	    return m.ID
//	}
//
//	func (m Menu) GetParentID() int {
//	    return m.ParentID
//	}
//
// 注意事项：
// - 这是一个泛型接口，T通常是实现接口的结构体类型
// - SetChildren的参数类型可能需要根据实际需求调整
type TreeNode[T any] interface {
	GetChildren() []T       // 获取子节点列表
	SetChildren(children T) // 设置子节点（注意：参数类型可能需要调整）
	GetID() int             // 获取节点ID
	GetParentID() int       // 获取父节点ID
}
