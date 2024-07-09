package system

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"reflect"
	"testing"
)

func Test_autoCodePackage_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		info *request.SysAutoCodePackageCreate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试1",
			args: args{
				ctx: context.Background(),
				info: &request.SysAutoCodePackageCreate{
					Desc:        "描述",
					Label:       "展示名",
					Template:    "package",
					PackageName: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &autoCodePackage{}
			if err := a.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_autoCodePackage_templates(t *testing.T) {
	type args struct {
		ctx    context.Context
		entity model.SysAutoCodePackage
		info   request.AutoCode
	}
	tests := []struct {
		name      string
		args      args
		wantCode  map[string]string
		wantEnter map[string]map[string]string
		wantErr   bool
	}{
		{
			name: "测试1",
			args: args{
				ctx: context.Background(),
				entity: model.SysAutoCodePackage{
					Desc:        "描述",
					Label:       "展示名",
					Template:    "package",
					PackageName: "preview",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &autoCodePackage{}
			gotCode, gotEnter, err := s.templates(tt.args.ctx, tt.args.entity, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("templates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCode, tt.wantCode) {
				t.Errorf("templates() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if !reflect.DeepEqual(gotEnter, tt.wantEnter) {
				t.Errorf("templates() gotEnter = %v, want %v", gotEnter, tt.wantEnter)
			}
		})
	}
}
