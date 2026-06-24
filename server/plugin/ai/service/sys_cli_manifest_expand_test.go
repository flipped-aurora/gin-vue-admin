package service

import (
	"encoding/json"
	"testing"
)

func TestExpandBodySchemaParametersUnfoldsModelFields(t *testing.T) {
	specJSON := `{
		"basePath": "/api",
		"definitions": {
			"request.CreateUser": {
				"properties": {
					"username": {"type": "string", "description": "用户名"},
					"password": {"type": "string", "description": "密码"},
					"phone": {"type": "string"},
					"role": {"$ref": "#/definitions/request.Role"}
				}
			}
		}
	}`
	var spec swaggerSpec
	if err := json.Unmarshal([]byte(specJSON), &spec); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	param := swaggerParameter{
		Name:   "data",
		In:     "body",
		Schema: map[string]interface{}{"$ref": "#/definitions/request.CreateUser"},
	}
	expanded := expandBodySchemaParameters(param, &spec)
	if len(expanded) != 4 {
		t.Fatalf("expanded count = %d, want 4: %+v", len(expanded), expanded)
	}
	// 字段按字母序稳定排列
	if expanded[0].Flag != "password" || expanded[1].Flag != "phone" || expanded[2].Flag != "role" || expanded[3].Flag != "username" {
		t.Fatalf("unexpected order: %+v", expanded)
	}
	// 嵌套对象 $ref 降级为 json，标量保留具体类型
	role := expanded[2]
	if role.Type != "json" {
		t.Fatalf("role type = %q, want json", role.Type)
	}
	user := expanded[3]
	if user.Type != "string" || user.Description != "用户名" {
		t.Fatalf("username param = %+v", user)
	}
}

func TestExpandBodySchemaParametersReturnsNilWhenNotRef(t *testing.T) {
	spec := &swaggerSpec{Definitions: map[string]swaggerDefinition{}}
	// 非 $ref body（无 schema）
	if got := expandBodySchemaParameters(swaggerParameter{In: "body"}, spec); got != nil {
		t.Fatalf("expected nil for empty schema, got %+v", got)
	}
	// 引用的模型不存在
	missing := swaggerParameter{In: "body", Schema: map[string]interface{}{"$ref": "#/definitions/Missing"}}
	if got := expandBodySchemaParameters(missing, spec); got != nil {
		t.Fatalf("expected nil for missing definition, got %+v", got)
	}
}
