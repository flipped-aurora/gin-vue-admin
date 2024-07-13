package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPluginInitialize_Injection(t *testing.T) {
	type fields struct {
		Type       Type
		Path       string
		PluginPath string
		ImportPath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Shop插件 注册注入",
			fields: fields{
				Type:       TypePluginInitialize,
				Path:       filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "plugin_biz_v2.go"),
				PluginPath: filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "plugin.go"),
				ImportPath: `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewPluginInitialize(tt.fields.Type, tt.fields.Path, tt.fields.PluginPath, tt.fields.ImportPath)
			if err := a.Injection(); (err != nil) != tt.wantErr {
				t.Errorf("Injection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPluginInitialize_Rollback(t *testing.T) {
	type fields struct {
		Type        Type
		Path        string
		PluginPath  string
		ImportPath  string
		PluginName  string
		StructName  string
		PackageName string
		version     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Shop插件 回滚",
			fields: fields{
				Type:       TypePluginInitialize,
				Path:       filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "plugin_biz_v2.go"),
				PluginPath: filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "plugin.go"),
				ImportPath: `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := PluginInitialize{
				Type:        tt.fields.Type,
				Path:        tt.fields.Path,
				PluginPath:  tt.fields.PluginPath,
				ImportPath:  tt.fields.ImportPath,
				StructName:  "Plugin",
				PackageName: "shop",
				Version:     "bizPluginV2",
			}
			if err := a.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
