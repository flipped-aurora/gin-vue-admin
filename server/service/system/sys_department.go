package system

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type SysDepartmentService struct{}

var SysDepartmentServiceApp = new(SysDepartmentService)

// buildAncestors 依据父部门推算祖级链: 顶级为 "0", 子级为 父.Ancestors + "," + 父ID
func (s *SysDepartmentService) buildAncestors(ctx context.Context, parentId uint) (string, error) {
	if parentId == 0 {
		return "0", nil
	}
	var parent system.SysDepartment
	if err := global.GVA_DB.WithContext(ctx).First(&parent, parentId).Error; err != nil {
		return "", errors.New("父部门不存在")
	}
	return parent.Ancestors + "," + strconv.Itoa(int(parent.ID)), nil
}

// buildDepartmentNamePath 依据部门自身与 id→name 映射, 拼出 "公司/部门" 全路径名(纯函数)
// dept.Ancestors 为祖级 ID 链(顶级 "0"、不含自身, 如 "0,1,5"); "0" 与映射中缺失的祖级会被跳过
func buildDepartmentNamePath(dept system.SysDepartment, nameByID map[uint]string) string {
	var parts []string
	for _, seg := range strings.Split(dept.Ancestors, ",") {
		seg = strings.TrimSpace(seg)
		if seg == "" || seg == "0" {
			continue
		}
		id, err := strconv.ParseUint(seg, 10, 64)
		if err != nil {
			continue
		}
		if name := nameByID[uint(id)]; name != "" {
			parts = append(parts, name)
		}
	}
	if dept.Name != "" {
		parts = append(parts, dept.Name)
	}
	return strings.Join(parts, "/")
}

// FillNamePaths 依据各部门 Ancestors 解析出 "公司/部门" 全路径名, 按索引写回每个部门的 NamePath。
// 只做一次 id→name 查询; 查询失败返回 error, 调用方可容错忽略(前端会降级为裸部门名)。
func (s *SysDepartmentService) FillNamePaths(ctx context.Context, depts []system.SysDepartment) error {
	if len(depts) == 0 {
		return nil
	}
	// 收集所有祖级 ID(去掉虚拟根 "0")
	idSet := make(map[uint]struct{})
	for _, d := range depts {
		for _, seg := range strings.Split(d.Ancestors, ",") {
			seg = strings.TrimSpace(seg)
			if seg == "" || seg == "0" {
				continue
			}
			if id, err := strconv.ParseUint(seg, 10, 64); err == nil {
				idSet[uint(id)] = struct{}{}
			}
		}
	}
	nameByID := make(map[uint]string, len(idSet))
	if len(idSet) > 0 {
		ids := make([]uint, 0, len(idSet))
		for id := range idSet {
			ids = append(ids, id)
		}
		var rows []system.SysDepartment
		if err := global.GVA_DB.WithContext(ctx).Model(&system.SysDepartment{}).
			Select("id", "name").Where("id IN ?", ids).Find(&rows).Error; err != nil {
			return err
		}
		for _, r := range rows {
			nameByID[r.ID] = r.Name
		}
	}
	// 值切片须按索引写回, 否则改的是副本
	for i := range depts {
		depts[i].NamePath = buildDepartmentNamePath(depts[i], nameByID)
	}
	return nil
}

// CreateSysDepartment 创建部门并自动维护祖级链
func (s *SysDepartmentService) CreateSysDepartment(ctx context.Context, dept *system.SysDepartment) (err error) {
	ancestors, err := s.buildAncestors(ctx, dept.ParentId)
	if err != nil {
		return err
	}
	dept.Ancestors = ancestors
	dept.Children = nil
	return global.GVA_DB.WithContext(ctx).Create(dept).Error
}

// UpdateSysDepartment 更新部门, 若父部门变更则重算本节点祖级链
// 注: 子孙节点的祖级链重算(部门移动)属后续阶段, 此处仅保证本节点正确
func (s *SysDepartmentService) UpdateSysDepartment(ctx context.Context, dept *system.SysDepartment) (err error) {
	if dept.ParentId == dept.ID {
		return errors.New("父部门不能是自己")
	}
	ancestors, err := s.buildAncestors(ctx, dept.ParentId)
	if err != nil {
		return err
	}
	return global.GVA_DB.WithContext(ctx).Model(&system.SysDepartment{}).Where("id = ?", dept.ID).Updates(map[string]interface{}{
		"name":      dept.Name,
		"parent_id": dept.ParentId,
		"ancestors": ancestors,
		"sort":      dept.Sort,
		"leader_id": dept.LeaderId,
		"status":    dept.Status,
	}).Error
}

// DeleteSysDepartment 删除部门, 存在子部门或已有用户归属时禁止删除
func (s *SysDepartmentService) DeleteSysDepartment(ctx context.Context, id uint) (err error) {
	if id == 0 {
		return errors.New("部门ID不能为空")
	}
	if !errors.Is(global.GVA_DB.WithContext(ctx).Where("parent_id = ?", id).First(&system.SysDepartment{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在子部门,不允许删除")
	}
	var userCount int64
	if err = global.GVA_DB.WithContext(ctx).Model(&system.SysUser{}).Where("dept_id = ?", id).Count(&userCount).Error; err != nil {
		return err
	}
	if userCount > 0 {
		return errors.New("该部门下存在用户,不允许删除")
	}
	var joinCount int64
	if err = global.GVA_DB.WithContext(ctx).Model(&system.SysUserDepartment{}).Where("sys_department_id = ?", id).Count(&joinCount).Error; err != nil {
		return err
	}
	if joinCount > 0 {
		return errors.New("该部门下存在用户,不允许删除")
	}
	return global.GVA_DB.WithContext(ctx).Delete(&system.SysDepartment{}, id).Error
}

// GetSysDepartment 获取单个部门
func (s *SysDepartmentService) GetSysDepartment(ctx context.Context, id uint) (dept system.SysDepartment, err error) {
	err = global.GVA_DB.WithContext(ctx).Preload("Leader").First(&dept, id).Error
	return
}

// GetSysDepartmentTree 获取部门树; 带名称搜索时平铺返回匹配项, 否则返回整棵树
func (s *SysDepartmentService) GetSysDepartmentTree(ctx context.Context, name string) (list []system.SysDepartment, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysDepartment{}).Preload("Leader").Order("sort")
	if name != "" {
		err = db.Where("name LIKE ?", "%"+name+"%").Find(&list).Error
		return list, err
	}
	err = db.Where("parent_id = ?", 0).Find(&list).Error
	if err != nil {
		return nil, err
	}
	for i := range list {
		if err = s.findChildren(ctx, &list[i]); err != nil {
			return nil, err
		}
	}
	return list, err
}

// findChildren 递归组装子部门(照 AuthorityService.findChildrenAuthority)
func (s *SysDepartmentService) findChildren(ctx context.Context, dept *system.SysDepartment) (err error) {
	err = global.GVA_DB.WithContext(ctx).Preload("Leader").Where("parent_id = ?", dept.ID).Order("sort").Find(&dept.Children).Error
	if len(dept.Children) > 0 {
		for i := range dept.Children {
			if err = s.findChildren(ctx, &dept.Children[i]); err != nil {
				return err
			}
		}
	}
	return err
}

// GetDepartmentUserIds 获取部门下的用户ID(多部门归属, 反向查询)
func (s *SysDepartmentService) GetDepartmentUserIds(ctx context.Context, deptId uint) (ids []uint, err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&system.SysUserDepartment{}).
		Where("sys_department_id = ?", deptId).Pluck("sys_user_id", &ids).Error
	return
}

// SetDepartmentUsers 全量覆盖部门的用户成员(反向分配), 并维护主部门 dept_id 一致性
func (s *SysDepartmentService) SetDepartmentUsers(ctx context.Context, deptId uint, userIds []uint) error {
	if deptId == 0 {
		return errors.New("部门ID不能为空")
	}
	return global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var current []uint
		if err := tx.Model(&system.SysUserDepartment{}).Where("sys_department_id = ?", deptId).
			Pluck("sys_user_id", &current).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserDepartment{}, "sys_department_id = ?", deptId).Error; err != nil {
			return err
		}
		if len(userIds) > 0 {
			records := make([]system.SysUserDepartment, 0, len(userIds))
			for _, uid := range userIds {
				records = append(records, system.SysUserDepartment{SysUserId: uid, SysDepartmentId: deptId})
			}
			if err := tx.Create(&records).Error; err != nil {
				return err
			}
		}
		// 主部门一致性: 被移除者若主部门是本部门则清空; 归属本部门且无主部门者补设为本部门
		target := make(map[uint]struct{}, len(userIds))
		for _, uid := range userIds {
			target[uid] = struct{}{}
		}
		var removed []uint
		for _, uid := range current {
			if _, ok := target[uid]; !ok {
				removed = append(removed, uid)
			}
		}
		if len(removed) > 0 {
			if err := tx.Model(&system.SysUser{}).Where("dept_id = ? AND id IN ?", deptId, removed).
				Update("dept_id", 0).Error; err != nil {
				return err
			}
		}
		if len(userIds) > 0 {
			if err := tx.Model(&system.SysUser{}).Where("dept_id = ? AND id IN ?", 0, userIds).
				Update("dept_id", deptId).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
