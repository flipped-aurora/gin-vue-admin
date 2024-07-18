package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPackageModuleEnter_Rollback(t *testing.T) {
	type fields struct {
		Type        Type
		Path        string
		ImportPath  string
		StructName  string
		AppName     string
		GroupName   string
		ModuleName  string
		PackageName string
		ServiceName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 FileUploadAndDownloadRouter 回滚",
			fields: fields{
				Type:        TypePackageRouterModuleEnter,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", "example", "enter.go"),
				ImportPath:  `api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"`,
				StructName:  "FileUploadAndDownloadRouter",
				AppName:     "ApiGroupApp",
				GroupName:   "ExampleApiGroup",
				ModuleName:  "exaFileUploadAndDownloadApi",
				PackageName: "api",
				ServiceName: "FileUploadAndDownloadApi",
			},
			wantErr: false,
		},
		{
			name: "测试 FileUploadAndDownloadApi 回滚",
			fields: fields{
				Type:        TypePackageApiModuleEnter,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", "example", "enter.go"),
				ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/service"`,
				StructName:  "FileUploadAndDownloadApi",
				AppName:     "ServiceGroupApp",
				GroupName:   "ExampleServiceGroup",
				ModuleName:  "fileUploadAndDownloadService",
				PackageName: "service",
				ServiceName: "FileUploadAndDownloadService",
			},
			wantErr: false,
		},
		{
			name: "测试 FileUploadAndDownloadService 回滚",
			fields: fields{
				Type:        TypePackageServiceModuleEnter,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", "example", "enter.go"),
				ImportPath:  ``,
				StructName:  "FileUploadAndDownloadService",
				AppName:     "",
				GroupName:   "",
				ModuleName:  "",
				PackageName: "",
				ServiceName: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PackageModuleEnter{
				Type:        tt.fields.Type,
				Path:        tt.fields.Path,
				ImportPath:  tt.fields.ImportPath,
				StructName:  tt.fields.StructName,
				AppName:     tt.fields.AppName,
				GroupName:   tt.fields.GroupName,
				ModuleName:  tt.fields.ModuleName,
				PackageName: tt.fields.PackageName,
				ServiceName: tt.fields.ServiceName,
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

func TestPackageModuleEnter_Injection(t *testing.T) {
	type fields struct {
		Type        Type
		Path        string
		ImportPath  string
		StructName  string
		AppName     string
		GroupName   string
		ModuleName  string
		PackageName string
		ServiceName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试 FileUploadAndDownloadRouter 注入",
			fields: fields{
				Type:        TypePackageRouterModuleEnter,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", "example", "enter.go"),
				ImportPath:  `api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"`,
				StructName:  "FileUploadAndDownloadRouter",
				AppName:     "ApiGroupApp",
				GroupName:   "ExampleApiGroup",
				ModuleName:  "exaFileUploadAndDownloadApi",
				PackageName: "api",
				ServiceName: "FileUploadAndDownloadApi",
			},
			wantErr: false,
		},
		{
			name: "测试 FileUploadAndDownloadApi 注入",
			fields: fields{
				Type:        TypePackageApiModuleEnter,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", "example", "enter.go"),
				ImportPath:  `"github.com/flipped-aurora/gin-vue-admin/server/service"`,
				StructName:  "FileUploadAndDownloadApi",
				AppName:     "ServiceGroupApp",
				GroupName:   "ExampleServiceGroup",
				ModuleName:  "fileUploadAndDownloadService",
				PackageName: "service",
				ServiceName: "FileUploadAndDownloadService",
			},
			wantErr: false,
		},
		{
			name: "测试 FileUploadAndDownloadService 注入",
			fields: fields{
				Type:        TypePackageServiceModuleEnter,
				Path:        filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", "example", "enter.go"),
				ImportPath:  ``,
				StructName:  "FileUploadAndDownloadService",
				AppName:     "",
				GroupName:   "",
				ModuleName:  "",
				PackageName: "",
				ServiceName: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PackageModuleEnter{
				Type:        tt.fields.Type,
				Path:        tt.fields.Path,
				ImportPath:  tt.fields.ImportPath,
				StructName:  tt.fields.StructName,
				AppName:     tt.fields.AppName,
				GroupName:   tt.fields.GroupName,
				ModuleName:  tt.fields.ModuleName,
				PackageName: tt.fields.PackageName,
				ServiceName: tt.fields.ServiceName,
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
