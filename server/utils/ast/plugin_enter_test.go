package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPluginEnter_Injection(t *testing.T) {
	type fields struct {
		Type            Type
		Path            string
		ImportPath      string
		StructName      string
		StructCamelName string
		ModuleName      string
		GroupName       string
		PackageName     string
		ServiceName     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 商城插件UserService 注入",
			fields: fields{
				Type:            TypePluginServiceEnter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "service", "enter.go"),
				ImportPath:      "",
				StructName:      "User",
				StructCamelName: "user",
				ModuleName:      "",
				GroupName:       "",
				PackageName:     "",
				ServiceName:     "",
			},
			wantErr: false,
		},
		{
			name: "测试 商城插件UserApi 注入",
			fields: fields{
				Type:            TypePluginApiEnter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "api", "enter.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/service"`,
				StructName:      "User",
				StructCamelName: "user",
				ModuleName:      "serviceUser",
				GroupName:       "Service",
				PackageName:     "service",
				ServiceName:     "User",
			},
			wantErr: false,
		},
		{
			name: "测试 商城插件UserRouter 注入",
			fields: fields{
				Type:            TypePluginRouterEnter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "router", "enter.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/api"`,
				StructName:      "User",
				StructCamelName: "user",
				ModuleName:      "userApi",
				GroupName:       "Api",
				PackageName:     "api",
				ServiceName:     "User",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginEnter{
				Type:            tt.fields.Type,
				Path:            tt.fields.Path,
				ImportPath:      tt.fields.ImportPath,
				StructName:      tt.fields.StructName,
				StructCamelName: tt.fields.StructCamelName,
				ModuleName:      tt.fields.ModuleName,
				GroupName:       tt.fields.GroupName,
				PackageName:     tt.fields.PackageName,
				ServiceName:     tt.fields.ServiceName,
			}
			if err := a.Injection(); (err != nil) != tt.wantErr {
				t.Errorf("Injection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPluginEnter_Rollback(t *testing.T) {
	type fields struct {
		Type            Type
		Path            string
		ImportPath      string
		StructName      string
		StructCamelName string
		ModuleName      string
		GroupName       string
		PackageName     string
		ServiceName     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 商城插件UserRouter 回滚",
			fields: fields{
				Type:            TypePluginRouterEnter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "router", "enter.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/api"`,
				StructName:      "User",
				StructCamelName: "user",
				ModuleName:      "userApi",
				GroupName:       "Api",
				PackageName:     "api",
				ServiceName:     "User",
			},
			wantErr: false,
		},
		{
			name: "测试 商城插件UserApi 回滚",
			fields: fields{
				Type:            TypePluginApiEnter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "api", "enter.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/plugin/shop/service"`,
				StructName:      "User",
				StructCamelName: "user",
				ModuleName:      "serviceUser",
				GroupName:       "Service",
				PackageName:     "service",
				ServiceName:     "User",
			},
			wantErr: false,
		},
		{
			name: "测试 商城插件UserService 回滚",
			fields: fields{
				Type:            TypePluginServiceEnter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "shop", "service", "enter.go"),
				ImportPath:      "",
				StructName:      "User",
				StructCamelName: "user",
				ModuleName:      "",
				GroupName:       "",
				PackageName:     "",
				ServiceName:     "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginEnter{
				Type:            tt.fields.Type,
				Path:            tt.fields.Path,
				ImportPath:      tt.fields.ImportPath,
				StructName:      tt.fields.StructName,
				StructCamelName: tt.fields.StructCamelName,
				ModuleName:      tt.fields.ModuleName,
				GroupName:       tt.fields.GroupName,
				PackageName:     tt.fields.PackageName,
				ServiceName:     tt.fields.ServiceName,
			}
			if err := a.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
