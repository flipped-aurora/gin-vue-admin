package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPluginInitializeRouter_Injection(t *testing.T) {
	type fields struct {
		Type                 Type
		Path                 string
		ImportPath           string
		AppName              string
		GroupName            string
		PackageName          string
		FunctionName         string
		LeftRouterGroupName  string
		RightRouterGroupName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Gva插件User 注入",
			fields: fields{
				Type:                 TypePluginInitializeRouter,
				Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "initialize", "router.go"),
				ImportPath:           `"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva/router"`,
				AppName:              "Router",
				GroupName:            "User",
				PackageName:          "router",
				FunctionName:         "Init",
				LeftRouterGroupName:  "public",
				RightRouterGroupName: "private",
			},
			wantErr: false,
		},
		{
			name: "测试 中文 注入",
			fields: fields{
				Type:                 TypePluginInitializeRouter,
				Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "initialize", "router.go"),
				ImportPath:           `"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva/router"`,
				AppName:              "Router",
				GroupName:            "U中文",
				PackageName:          "router",
				FunctionName:         "Init",
				LeftRouterGroupName:  "public",
				RightRouterGroupName: "private",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginInitializeRouter{
				Type:                 tt.fields.Type,
				Path:                 tt.fields.Path,
				ImportPath:           tt.fields.ImportPath,
				AppName:              tt.fields.AppName,
				GroupName:            tt.fields.GroupName,
				PackageName:          tt.fields.PackageName,
				FunctionName:         tt.fields.FunctionName,
				LeftRouterGroupName:  tt.fields.LeftRouterGroupName,
				RightRouterGroupName: tt.fields.RightRouterGroupName,
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

func TestPluginInitializeRouter_Rollback(t *testing.T) {
	type fields struct {
		Type                 Type
		Path                 string
		ImportPath           string
		AppName              string
		GroupName            string
		PackageName          string
		FunctionName         string
		LeftRouterGroupName  string
		RightRouterGroupName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 Gva插件User 回滚",
			fields: fields{
				Type:                 TypePluginInitializeRouter,
				Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "initialize", "router.go"),
				ImportPath:           `"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva/router"`,
				AppName:              "Router",
				GroupName:            "User",
				PackageName:          "router",
				FunctionName:         "Init",
				LeftRouterGroupName:  "public",
				RightRouterGroupName: "private",
			},
			wantErr: false,
		},
		{
			name: "测试 中文 注入",
			fields: fields{
				Type:                 TypePluginInitializeRouter,
				Path:                 filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "gva", "initialize", "router.go"),
				ImportPath:           `"github.com/flipped-aurora/gin-vue-admin/server/plugin/gva/router"`,
				AppName:              "Router",
				GroupName:            "U中文",
				PackageName:          "router",
				FunctionName:         "Init",
				LeftRouterGroupName:  "public",
				RightRouterGroupName: "private",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginInitializeRouter{
				Type:                 tt.fields.Type,
				Path:                 tt.fields.Path,
				ImportPath:           tt.fields.ImportPath,
				AppName:              tt.fields.AppName,
				GroupName:            tt.fields.GroupName,
				PackageName:          tt.fields.PackageName,
				FunctionName:         tt.fields.FunctionName,
				LeftRouterGroupName:  tt.fields.LeftRouterGroupName,
				RightRouterGroupName: tt.fields.RightRouterGroupName,
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
