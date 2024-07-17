package system

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"reflect"
	"testing"
)

func Test_autoCodeTemplate_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		info request.AutoCode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &autoCodeTemplate{}
			if err := s.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_autoCodeTemplate_Preview(t *testing.T) {
	type args struct {
		ctx  context.Context
		info request.AutoCode
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "测试 package",
			args: args{
				ctx:  context.Background(),
				info: request.AutoCode{},
			},
			wantErr: false,
		},
		{
			name: "测试 plugin",
			args: args{
				ctx:  context.Background(),
				info: request.AutoCode{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testJson := `{"structName":"SysUser","tableName":"sys_users","packageName":"sysUsers","package":"gva","abbreviation":"sysUsers","description":"sysUsers表","businessDB":"","autoCreateApiToSql":true,"autoCreateMenuToSql":true,"autoMigrate":true,"gvaModel":true,"autoCreateResource":false,"fields":[{"fieldName":"Uuid","fieldDesc":"用户UUID","fieldType":"string","dataType":"varchar","fieldJson":"uuid","primaryKey":false,"dataTypeLong":"191","columnName":"uuid","comment":"用户UUID","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Username","fieldDesc":"用户登录名","fieldType":"string","dataType":"varchar","fieldJson":"username","primaryKey":false,"dataTypeLong":"191","columnName":"username","comment":"用户登录名","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Password","fieldDesc":"用户登录密码","fieldType":"string","dataType":"varchar","fieldJson":"password","primaryKey":false,"dataTypeLong":"191","columnName":"password","comment":"用户登录密码","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"NickName","fieldDesc":"用户昵称","fieldType":"string","dataType":"varchar","fieldJson":"nickName","primaryKey":false,"dataTypeLong":"191","columnName":"nick_name","comment":"用户昵称","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"SideMode","fieldDesc":"用户侧边主题","fieldType":"string","dataType":"varchar","fieldJson":"sideMode","primaryKey":false,"dataTypeLong":"191","columnName":"side_mode","comment":"用户侧边主题","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"HeaderImg","fieldDesc":"用户头像","fieldType":"string","dataType":"varchar","fieldJson":"headerImg","primaryKey":false,"dataTypeLong":"191","columnName":"header_img","comment":"用户头像","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"BaseColor","fieldDesc":"基础颜色","fieldType":"string","dataType":"varchar","fieldJson":"baseColor","primaryKey":false,"dataTypeLong":"191","columnName":"base_color","comment":"基础颜色","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"AuthorityId","fieldDesc":"用户角色ID","fieldType":"int","dataType":"bigint","fieldJson":"authorityId","primaryKey":false,"dataTypeLong":"20","columnName":"authority_id","comment":"用户角色ID","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Phone","fieldDesc":"用户手机号","fieldType":"string","dataType":"varchar","fieldJson":"phone","primaryKey":false,"dataTypeLong":"191","columnName":"phone","comment":"用户手机号","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Email","fieldDesc":"用户邮箱","fieldType":"string","dataType":"varchar","fieldJson":"email","primaryKey":false,"dataTypeLong":"191","columnName":"email","comment":"用户邮箱","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}},{"fieldName":"Enable","fieldDesc":"用户是否被冻结 1正常 2冻结","fieldType":"int","dataType":"bigint","fieldJson":"enable","primaryKey":false,"dataTypeLong":"19","columnName":"enable","comment":"用户是否被冻结 1正常 2冻结","require":false,"errorText":"","clearable":true,"fieldSearchType":"","fieldIndexType":"","dictType":"","front":true,"dataSource":{"association":1,"table":"","label":"","value":""}}],"humpPackageName":"sys_users"}`
			err := json.Unmarshal([]byte(testJson), &tt.args.info)
			if err != nil {
				t.Error(err)
				return
			}
			err = tt.args.info.Pretreatment()
			if err != nil {
				t.Error(err)
				return
			}
			got, err := AutoCodeTemplate.Preview(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("Preview() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Preview() got = %v, want %v", got, tt.want)
			}
		})
	}
}
