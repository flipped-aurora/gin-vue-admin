package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPackageInitializeRouter_Injection(t *testing.T) {
	type fields struct {
		Type            Type
		Path            string
		ImportPath      string
		AppName         string
		GroupName       string
		ModuleName      string
		PackageName     string
		FunctionName    string
		RouterGroupName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 InitCustomerRouter 注入",
			fields: fields{
				Type:            TypePackageInitializeRouter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router_biz.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/router"`,
				AppName:         "RouterGroupApp",
				GroupName:       "Example",
				ModuleName:      "exampleRouter",
				PackageName:     "router",
				FunctionName:    "InitCustomerRouter",
				RouterGroupName: "privateGroup",
			},
			wantErr: false,
		},
		{
			name: "测试 InitFileUploadAndDownloadRouter 注入",
			fields: fields{
				Type:            TypePackageInitializeRouter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router_biz.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/router"`,
				AppName:         "RouterGroupApp",
				GroupName:       "Example",
				ModuleName:      "exampleRouter",
				PackageName:     "router",
				FunctionName:    "InitFileUploadAndDownloadRouter",
				RouterGroupName: "privateGroup",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PackageInitializeRouter{
				Type:                 tt.fields.Type,
				Path:                 tt.fields.Path,
				ImportPath:           tt.fields.ImportPath,
				AppName:              tt.fields.AppName,
				GroupName:            tt.fields.GroupName,
				ModuleName:           tt.fields.ModuleName,
				PackageName:          tt.fields.PackageName,
				FunctionName:         tt.fields.FunctionName,
				RouterGroupName:      tt.fields.RouterGroupName,
				LeftRouterGroupName:  "privateGroup",
				RightRouterGroupName: "publicGroup",
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

func TestPackageInitializeRouter_Rollback(t *testing.T) {
	type fields struct {
		Type            Type
		Path            string
		ImportPath      string
		AppName         string
		GroupName       string
		ModuleName      string
		PackageName     string
		FunctionName    string
		RouterGroupName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "测试 InitCustomerRouter 回滚",
			fields: fields{
				Type:            TypePackageInitializeRouter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router_biz.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/router"`,
				AppName:         "RouterGroupApp",
				GroupName:       "Example",
				ModuleName:      "exampleRouter",
				PackageName:     "router",
				FunctionName:    "InitCustomerRouter",
				RouterGroupName: "privateGroup",
			},
			wantErr: false,
		},
		{
			name: "测试 InitFileUploadAndDownloadRouter 回滚",
			fields: fields{
				Type:            TypePackageInitializeRouter,
				Path:            filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "initialize", "router_biz.go"),
				ImportPath:      `"github.com/flipped-aurora/gin-vue-admin/server/router"`,
				AppName:         "RouterGroupApp",
				GroupName:       "Example",
				ModuleName:      "exampleRouter",
				PackageName:     "router",
				FunctionName:    "InitFileUploadAndDownloadRouter",
				RouterGroupName: "privateGroup",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PackageInitializeRouter{
				Type:            tt.fields.Type,
				Path:            tt.fields.Path,
				ImportPath:      tt.fields.ImportPath,
				AppName:         tt.fields.AppName,
				GroupName:       tt.fields.GroupName,
				ModuleName:      tt.fields.ModuleName,
				PackageName:     tt.fields.PackageName,
				FunctionName:    tt.fields.FunctionName,
				RouterGroupName: tt.fields.RouterGroupName,
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
