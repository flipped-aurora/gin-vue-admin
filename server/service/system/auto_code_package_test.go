package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
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
