package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// SysDepartment 部门(组织架构树)
// 照 SysAuthority 的自关联树范式实现,额外维护 Ancestors 祖级链,便于后续数据权限"本部门及子级"快速取子树
type SysDepartment struct {
	global.GVA_MODEL
	Name      string          `json:"name" gorm:"index;comment:部门名称"`                  // 部门名称
	ParentId  uint            `json:"parentId" gorm:"default:0;comment:父部门ID"`         // 父部门ID(0为顶级)
	Ancestors string          `json:"ancestors" gorm:"comment:祖级链,逗号分隔如 0,1,5"`        // 祖级链
	Sort      int             `json:"sort" gorm:"default:0;comment:排序"`                // 排序
	LeaderId  uint            `json:"leaderId" gorm:"comment:负责人用户ID"`                 // 负责人(关联 sys_users)
	Leader    *SysUser        `json:"leader" form:"-" gorm:"foreignKey:LeaderId;references:ID"` // 负责人用户(联系电话/邮箱从该用户带出);form:"-" 阻断 gin 绑定递归(与 SysUser.Dept 成环会栈溢出)
	Status    *bool           `json:"status" gorm:"default:true;comment:是否启用"`         // 是否启用
	Children  []SysDepartment `json:"children" gorm:"-"`                               // 子部门(内存组装,不建列)
	NamePath  string          `json:"namePath" gorm:"-"`                               // 公司/部门全路径名(内存组装,不建列)
}

func (SysDepartment) TableName() string {
	return "sys_departments"
}
