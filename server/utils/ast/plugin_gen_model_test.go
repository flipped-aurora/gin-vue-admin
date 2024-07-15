package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPluginGenModel_Injection(t *testing.T) {
	type fields struct {
		Type        Type
		Path        string
		ImportPath  string
		PackageName string
		StructName  string
		IsNew       bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 ShopUser 结构体注入",
			fields: fields{
				Type:        TypePluginGenModel,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "gen", "main.go"),
				ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/model"`,
				PackageName: "model",
				StructName:  "User",
				IsNew:       false,
			},
		},
		{
			name: "测试 ShopUser 结构体注入",
			fields: fields{
				Type:        TypePluginGenModel,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "gen", "main.go"),
				ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/model"`,
				PackageName: "model",
				StructName:  "User",
				IsNew:       true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginGenModel{
				Type:        tt.fields.Type,
				Path:        tt.fields.Path,
				ImportPath:  tt.fields.ImportPath,
				PackageName: tt.fields.PackageName,
				StructName:  tt.fields.StructName,
				IsNew:       tt.fields.IsNew,
			}
			if err := a.Injection(); (err != nil) != tt.wantErr {
				t.Errorf("Injection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPluginGenModel_Rollback(t *testing.T) {
	type fields struct {
		Type        Type
		Path        string
		ImportPath  string
		PackageName string
		StructName  string
		IsNew       bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 ShopUser 回滚",
			fields: fields{
				Type:        TypePluginGenModel,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "gen", "main.go"),
				ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/model"`,
				PackageName: "model",
				StructName:  "User",
				IsNew:       false,
			},
		},
		{
			name: "测试 ShopUser 回滚",
			fields: fields{
				Type:        TypePluginGenModel,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "gen", "main.go"),
				ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/model"`,
				PackageName: "model",
				StructName:  "User",
				IsNew:       true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginGenModel{
				Type:        tt.fields.Type,
				Path:        tt.fields.Path,
				ImportPath:  tt.fields.ImportPath,
				PackageName: tt.fields.PackageName,
				StructName:  tt.fields.StructName,
				IsNew:       tt.fields.IsNew,
			}
			if err := a.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
