package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPluginInitializeMenu_Injection(t *testing.T) {
	type fields struct {
		Type          Type
		Path          string
		MenuSort      string
		MenuPath      string
		MenuName      string
		MenuMetaIcon  string
		MenuMetaTitle string
		MenuComponent string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Gva用户菜单 注入",
			fields: fields{
				Type:          TypePluginInitializeMenu,
				Path:          filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "initialize", "menu.go"),
				MenuSort:      "9999",
				MenuPath:      `"gvaUser"`,
				MenuName:      `"Gva用户管理"`,
				MenuMetaIcon:  `"aim"`,
				MenuMetaTitle: `"Gva用户"`,
				MenuComponent: `"plugin/gva/view/user.vue"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginInitializeMenu{
				Type:          tt.fields.Type,
				Path:          tt.fields.Path,
				MenuSort:      tt.fields.MenuSort,
				MenuPath:      tt.fields.MenuPath,
				MenuName:      tt.fields.MenuName,
				MenuMetaIcon:  tt.fields.MenuMetaIcon,
				MenuMetaTitle: tt.fields.MenuMetaTitle,
				MenuComponent: tt.fields.MenuComponent,
			}
			if err := a.Injection(); (err != nil) != tt.wantErr {
				t.Errorf("Injection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPluginInitializeMenu_Rollback(t *testing.T) {
	type fields struct {
		Type          Type
		Path          string
		MenuSort      string
		MenuPath      string
		MenuName      string
		MenuMetaIcon  string
		MenuMetaTitle string
		MenuComponent string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Gva用户菜单 注入",
			fields: fields{
				Type:          TypePluginInitializeMenu,
				Path:          filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "initialize", "menu.go"),
				MenuSort:      "9999",
				MenuPath:      `"gvaUser"`,
				MenuName:      `"Gva用户管理"`,
				MenuMetaIcon:  `"aim"`,
				MenuMetaTitle: `"Gva用户"`,
				MenuComponent: `"plugin/gva/view/user.vue"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginInitializeMenu{
				Type:          tt.fields.Type,
				Path:          tt.fields.Path,
				MenuSort:      tt.fields.MenuSort,
				MenuPath:      tt.fields.MenuPath,
				MenuName:      tt.fields.MenuName,
				MenuMetaIcon:  tt.fields.MenuMetaIcon,
				MenuMetaTitle: tt.fields.MenuMetaTitle,
				MenuComponent: tt.fields.MenuComponent,
			}
			if err := a.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
