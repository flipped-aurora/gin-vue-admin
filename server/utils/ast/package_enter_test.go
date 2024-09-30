package ast

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"path/filepath"
	"testing"
)

func TestPackageEnter_Rollback(t *testing.T) {
	type fields struct {
		Type              Type
		Path              string
		ImportPath        string
		StructName        string
		PackageName       string
		PackageStructName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: global.Translate("utils.exampleApiGroupRollback"),
			fields: fields{
				Type:              TypePackageApiEnter,
				Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", "enter.go"),
				ImportPath:        `"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"`,
				StructName:        "ExampleApiGroup",
				PackageName:       "example",
				PackageStructName: "ApiGroup",
			},
			wantErr: false,
		},
		{
			name: global.Translate("utils.exampleRouterGroupRollback"),
			fields: fields{
				Type:              TypePackageRouterEnter,
				Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", "enter.go"),
				ImportPath:        `"github.com/flipped-aurora/gin-vue-admin/server/router/example"`,
				StructName:        "Example",
				PackageName:       "example",
				PackageStructName: "RouterGroup",
			},
			wantErr: false,
		},
		{
			name: global.Translate("utils.exampleServiceGroupRollback"),
			fields: fields{
				Type:              TypePackageServiceEnter,
				Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", "enter.go"),
				ImportPath:        `"github.com/flipped-aurora/gin-vue-admin/server/service/example"`,
				StructName:        "ExampleServiceGroup",
				PackageName:       "example",
				PackageStructName: "ServiceGroup",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PackageEnter{
				Type:              tt.fields.Type,
				Path:              tt.fields.Path,
				ImportPath:        tt.fields.ImportPath,
				StructName:        tt.fields.StructName,
				PackageName:       tt.fields.PackageName,
				PackageStructName: tt.fields.PackageStructName,
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

func TestPackageEnter_Injection(t *testing.T) {
	type fields struct {
		Type              Type
		Path              string
		ImportPath        string
		StructName        string
		PackageName       string
		PackageStructName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: global.Translate("utils.exampleApiGroupInjection"),
			fields: fields{
				Type:              TypePackageApiEnter,
				Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "api", "v1", "enter.go"),
				ImportPath:        `"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"`,
				StructName:        "ExampleApiGroup",
				PackageName:       "example",
				PackageStructName: "ApiGroup",
			},
		},
		{
			name: global.Translate("utils.exampleRouterGroupInjection"),
			fields: fields{
				Type:              TypePackageRouterEnter,
				Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "router", "enter.go"),
				ImportPath:        `"github.com/flipped-aurora/gin-vue-admin/server/router/example"`,
				StructName:        "Example",
				PackageName:       "example",
				PackageStructName: "RouterGroup",
			},
			wantErr: false,
		},
		{
			name: global.Translate("utils.exampleServiceGroupInjection"),
			fields: fields{
				Type:              TypePackageServiceEnter,
				Path:              filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "service", "enter.go"),
				ImportPath:        `"github.com/flipped-aurora/gin-vue-admin/server/service/example"`,
				StructName:        "ExampleServiceGroup",
				PackageName:       "example",
				PackageStructName: "ServiceGroup",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PackageEnter{
				Type:              tt.fields.Type,
				Path:              tt.fields.Path,
				ImportPath:        tt.fields.ImportPath,
				StructName:        tt.fields.StructName,
				PackageName:       tt.fields.PackageName,
				PackageStructName: tt.fields.PackageStructName,
			}
			file, err := a.Parse(a.Path, nil)
			if err != nil {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			a.Injection(file)
			err = a.Format(a.Path, nil, file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
