package ast

import "testing"

func TestPluginInitializeApi_Injection(t *testing.T) {
	type fields struct {
		Type Type
		Path string
		Apis []Api
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Goods 注入",
			fields: fields{
				Type: TypePluginInitializeApi,
				Path: "/Users/sliverhorn/Go/src/github.com/flipped-aurora/gin-vue-admin/server/plugin/gva/initialize/api.go",
				Apis: []Api{
					{Path: `"/gva/goods/create"`, Group: `"GvaGoods"`, Method: `"POST"`, Description: `"创建Gva商品"`},
					{Path: `"/gva/goods/first"`, Group: `"GvaGoods"`, Method: `"POST"`, Description: `"获取Gva商品"`},
					{Path: `"/gva/goods/update"`, Group: `"GvaGoods"`, Method: `"PUT"`, Description: `"更新Gva商品"`},
					{Path: `"/gva/goods/delete"`, Group: `"GvaGoods"`, Method: `"DELETE"`, Description: `"删除Gva商品"`},
					{Path: `"/gva/goods/deletes"`, Group: `"GvaGoods"`, Method: `"DELETE"`, Description: `"批量删除Gva商品"`},
					{Path: `"/gva/goods/list"`, Group: `"GvaGoods"`, Method: `"POST"`, Description: `"分页获取Gva商品"`},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := PluginInitializeApi{
				Type: tt.fields.Type,
				Path: tt.fields.Path,
				Apis: tt.fields.Apis,
			}
			if err := a.Injection(); (err != nil) != tt.wantErr {
				t.Errorf("Injection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPluginInitializeApi_Rollback(t *testing.T) {
	type fields struct {
		Type Type
		Path string
		Apis []Api
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Goods 回滚",
			fields: fields{
				Type: TypePluginInitializeApi,
				Path: "/Users/sliverhorn/Go/src/github.com/flipped-aurora/gin-vue-admin/server/plugin/gva/initialize/api.go",
				Apis: []Api{
					{Path: `"/gva/goods/create"`, Group: `"GvaGoods"`, Method: `"POST"`, Description: `"创建Gva商品"`},
					{Path: `"/gva/goods/first"`, Group: `"GvaGoods"`, Method: `"POST"`, Description: `"获取Gva商品"`},
					{Path: `"/gva/goods/update"`, Group: `"GvaGoods"`, Method: `"PUT"`, Description: `"更新Gva商品"`},
					{Path: `"/gva/goods/delete"`, Group: `"GvaGoods"`, Method: `"DELETE"`, Description: `"删除Gva商品"`},
					{Path: `"/gva/goods/deletes"`, Group: `"GvaGoods"`, Method: `"DELETE"`, Description: `"批量删除Gva商品"`},
					{Path: `"/gva/goods/list"`, Group: `"GvaGoods"`, Method: `"POST"`, Description: `"分页获取Gva商品"`},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := PluginInitializeApi{
				Type: tt.fields.Type,
				Path: tt.fields.Path,
				Apis: tt.fields.Apis,
			}
			if err := a.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
