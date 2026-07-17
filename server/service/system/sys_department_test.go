package system

import (
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// TestBuildDepartmentNamePath 覆盖「公司/部门」全路径拼接的各类场景(纯函数, 不触碰 DB)
func TestBuildDepartmentNamePath(t *testing.T) {
	nameByID := map[uint]string{
		1: "总公司",
		5: "研发中心",
	}
	cases := []struct {
		name string
		dept system.SysDepartment
		want string
	}{
		{
			name: "顶级部门只显示自身名",
			dept: system.SysDepartment{Name: "总公司", Ancestors: "0"},
			want: "总公司",
		},
		{
			name: "二级部门拼上公司前缀",
			dept: system.SysDepartment{Name: "研发部", Ancestors: "0,1"},
			want: "总公司/研发部",
		},
		{
			name: "多级部门整条祖级链拼接",
			dept: system.SysDepartment{Name: "后端组", Ancestors: "0,1,5"},
			want: "总公司/研发中心/后端组",
		},
		{
			name: "祖级名缺失时跳过该段",
			dept: system.SysDepartment{Name: "研发部", Ancestors: "0,9"},
			want: "研发部",
		},
		{
			name: "空祖级链退化为部门名",
			dept: system.SysDepartment{Name: "游离部门", Ancestors: ""},
			want: "游离部门",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := buildDepartmentNamePath(c.dept, nameByID)
			if got != c.want {
				t.Fatalf("buildDepartmentNamePath() = %q, want %q", got, c.want)
			}
		})
	}
}
