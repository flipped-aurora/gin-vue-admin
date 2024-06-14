package internal

import (
	"fmt"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"testing"
)

func Test_authority_CopyChildren(t *testing.T) {
	type args struct {
		ids      map[uint][]uint
		creates  map[uint]uint
		entity   *model.SysAuthority
		parent   model.SysAuthority
		entities []model.SysAuthority
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				ids: map[uint][]uint{
					1: {},
					2: {},
					3: {},
					4: {},
					5: {},
					6: {},
					7: {},
				},
				creates: map[uint]uint{},
				entity: &model.SysAuthority{
					AuthorityId: 1,
					Children: []model.SysAuthority{
						{
							AuthorityId: 2,
							Children: []model.SysAuthority{
								{AuthorityId: 5},
								{AuthorityId: 6},
							},
						},
						{AuthorityId: 3},
						{AuthorityId: 4},
					},
				},
				parent: model.SysAuthority{
					AuthorityId:   11,
					AuthorityName: "测试",
				},
				entities: []model.SysAuthority{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%p\n", tt.args.ids)
			Authority.CopyChildren(tt.args.ids, tt.args.creates, tt.args.entity, tt.args.parent, &tt.args.entities, 1)
			fmt.Println(tt.args.entities)
		})
	}
}
