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
				Path: "/Users/sliverhorn/Go/src/github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/initialize/api.go",
				Apis: []Api{
					{Path: `"/shop/goods/create"`, Group: `"ShopGoods"`, Method: `"POST"`, Description: `"创建商城商品"`},
					{Path: `"/shop/goods/first"`, Group: `"ShopGoods"`, Method: `"POST"`, Description: `"获取商城商品"`},
					{Path: `"/shop/goods/update"`, Group: `"ShopGoods"`, Method: `"PUT"`, Description: `"更新商城商品"`},
					{Path: `"/shop/goods/delete"`, Group: `"ShopGoods"`, Method: `"DELETE"`, Description: `"删除商城商品"`},
					{Path: `"/shop/goods/deletes"`, Group: `"ShopGoods"`, Method: `"DELETE"`, Description: `"批量删除商城商品"`},
					{Path: `"/shop/goods/list"`, Group: `"ShopGoods"`, Method: `"POST"`, Description: `"分页获取商城商品"`},
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
				Path: "/Users/sliverhorn/Go/src/github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/initialize/api.go",
				Apis: []Api{
					{Path: `"/shop/goods/create"`, Group: `"ShopGoods"`, Method: `"POST"`, Description: `"创建商城商品"`},
					{Path: `"/shop/goods/first"`, Group: `"ShopGoods"`, Method: `"POST"`, Description: `"获取商城商品"`},
					{Path: `"/shop/goods/update"`, Group: `"ShopGoods"`, Method: `"PUT"`, Description: `"更新商城商品"`},
					{Path: `"/shop/goods/delete"`, Group: `"ShopGoods"`, Method: `"DELETE"`, Description: `"删除商城商品"`},
					{Path: `"/shop/goods/deletes"`, Group: `"ShopGoods"`, Method: `"DELETE"`, Description: `"批量删除商城商品"`},
					{Path: `"/shop/goods/list"`, Group: `"ShopGoods"`, Method: `"POST"`, Description: `"分页获取商城商品"`},
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
