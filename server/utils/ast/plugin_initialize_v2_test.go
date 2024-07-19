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
			name: "测试 Gva插件 注册注入",
			fields: fields{
				Type:       TypePluginInitializeV2,
				Path:       filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "plugin_biz_v2.go"),
				PluginPath: filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "plugin.go"),
				ImportPath: `"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := PluginInitializeV2{
				Type:       tt.fields.Type,
				Path:       tt.fields.Path,
				PluginPath: tt.fields.PluginPath,
				ImportPath: tt.fields.ImportPath,
			}
			file, err := a.Parse(a.Path, nil)
			if err != nil {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			a.Injection(file)
			err = a.Format(a.Path, nil, file)
			if (err != nil) != tt.wantErr {
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
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Gva插件 回滚",
			fields: fields{
				Type:       TypePluginInitializeV2,
				Path:       filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "plugin_biz_v2.go"),
				PluginPath: filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "plugin.go"),
				ImportPath: `"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := PluginInitializeV2{
				Type:        tt.fields.Type,
				Path:        tt.fields.Path,
				PluginPath:  tt.fields.PluginPath,
				ImportPath:  tt.fields.ImportPath,
				StructName:  "Plugin",
				PackageName: "gva",
			}
			file, err := a.Parse(a.Path, nil)
			if err != nil {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			a.Rollback(file)
			err = a.Format(a.Path, nil, file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
