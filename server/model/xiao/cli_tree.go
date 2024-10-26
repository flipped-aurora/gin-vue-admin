// 自动生成模板CliTree
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// 用户关系表 结构体  CliTree
type CliTree struct {
	global.GVA_MODEL
	Address     string `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"` //用户地址
	Father      string `json:"father" form:"father" gorm:"column:father;comment:父级节点;size:100;"`               //父级节点
	Invite      string `json:"invite" form:"invite" gorm:"column:invite;comment:邀请节点;size:100;"`               //邀请节点
	Leftval     *int   `json:"leftval" form:"leftval" gorm:"default:1;column:leftval;comment:左值;"`             //左值
	Rightval    *int   `json:"rightval" form:"rightval" gorm:"default:2;column:rightval;comment:右值;"`          //右值
	Level       *int   `json:"level" form:"level" gorm:"default:1;column:level;comment:层级;"`                   //层级
	Vip         *int   `json:"vip" form:"vip" gorm:"default:0;column:vip;comment:用户等级;"`                       //用户等级
	Description string `json:"description" form:"description" gorm:"column:description;comment:备注;"`           //备注
}

// TableName 用户关系表 CliTree自定义表名 cli_tree
func (CliTree) TableName() string {
	return "cli_tree"
}

// NewCliTree
func NewCliTree(address, invite string) *CliTree {
	leftval := 1
	rightval := 2
	level := 1
	vip := 0

	return &CliTree{
		Address:  address,
		Invite:   invite,
		Leftval:  &leftval,
		Rightval: &rightval,
		Level:    &level,
		Vip:      &vip,
	}
}

// 获取自己的节点
func (clitree *CliTree) GetSelfNode(tx *gorm.DB) error {
	// 假设我们想在地址的开头匹配
	return tx.Where("address LIKE ?", "%"+clitree.Address+"%").First(clitree).Error
}

// GetAllTree 获取整棵树
func (tree *CliTree) GetAllTree(tx *gorm.DB) ([]*CliTree, error) {
	var trees []*CliTree                          // 定义一个Tree切片来存储查询结果
	if err := tx.Find(&trees).Error; err != nil { // 使用Find方法查询所有记录
		return nil, err // 如果查询出错，则返回错误
	}
	return trees, nil // 成功查询后返回Tree切片和nil错误
}

// GetFatherNode 获取父节点
func (tree *CliTree) GetFatherNode(tx *gorm.DB) (*CliTree, error) {
	//先查询自己节点
	var exists = &CliTree{
		Address: tree.Address,
	}

	if err := exists.GetSelfNode(tx); err != nil { // 调用GetSelfNode方法获取自己节点
		return nil, err // 如果获取自己节点出错，则返回错误
	}
	var father CliTree                                                                // 定义一个Tree变量来存储查询结果
	if err := tx.Where("address = ?", tree.Father).First(&father).Error; err != nil { // 使用Where方法查询指定条件的记录
		return nil, err // 如果查询出错，则返回错误
	}
	return &father, nil // 成功查询后返回Tree指针和nil错误
}

// 获取伞下所有节点
func (tree *CliTree) GetAllChilds(tx *gorm.DB) ([]*CliTree, error) {
	var exists = &CliTree{
		Address: tree.Address,
	}

	if err := exists.GetSelfNode(tx); err != nil { // 调用GetSelfNode方法获取自己节点
		return nil, err // 如果获取自己节点出错，则返回错误
	}
	var childs []*CliTree                                                                                                   // 定义一个Tree切片来存储查询结果
	if err := tx.Where("leftval > ? and rightval < ?", *exists.Leftval, *exists.Rightval).Find(&childs).Error; err != nil { // 使用Where方法查询指定条件的记录
		return nil, err // 如果查询出错，则返回错误
	}
	return childs, nil // 成功查询后返回Tree切片和nil错误
}

// 获取指定层级的节点
func (tree *CliTree) GetChildByLevel(tx *gorm.DB, add int) ([]*CliTree, error) {
	var exists = &CliTree{
		Address: tree.Address,
	}
	if err := exists.GetSelfNode(tx); err != nil { // 调用GetSelfNode方法获取自己节点
		return nil, err // 如果获取自己节点出错，则返回错误
	}
	var childs []*CliTree                                                                                                                              // 定义一个Tree切片来存储查询结果
	if err := tx.Where("leftval > ? and rightval < ? and level = ?", *tree.Leftval, *tree.Rightval, *tree.Level+add).Find(&childs).Error; err != nil { // 使用Where方法查询指定条件的记录
	}
	return childs, nil
}

// 获取我邀请的节点
func (tree *CliTree) GetMyInviteNode(tx *gorm.DB) (*CliTree, error) {
	exists := &CliTree{
		Address: tree.Address,
	}
	if err := exists.GetSelfNode(tx); err != nil { // 调用GetSelfNode方法获取自己节点
		return nil, err // 如果获取自己节点出错，则返回错误
	}
	var invitetree CliTree
	tx.Where("invite = ?", exists.Address).First(&invitetree)
	return &invitetree, nil
}

// 获取所有上级节点
func (tree *CliTree) GetAllParentNode(tx *gorm.DB) ([]*CliTree, error) {
	exists := &CliTree{
		Address: tree.Address,
	}
	if err := exists.GetSelfNode(tx); err != nil { // 调用GetSelfNode方法获取自己节点
		return nil, err // 如果获取自己节点出错，则返回错误
	}
	var childs []*CliTree // 定义一个Tree切片来存储查询结
	if err := tx.Where("leftval < ? and rightval > ?", *exists.Leftval, *exists.Rightval).Order("id DESC").Find(&childs).Error; err != nil {
		return nil, err
	}
	return childs, nil
}

// 创建节点
func (tree *CliTree) CreateNode(tx *gorm.DB) (err error) {
	//1.查询父节点是否存在
	var parenttree CliTree
	err = tx.Where("address = ?", tree.Invite).First(&parenttree).Error
	if err != nil {
		return err
	}
	//2.查询当前地址是否已经存在
	var exists = &CliTree{
		Address: tree.Address,
	}
	if err = exists.GetSelfNode(tx); err == nil { // 调用GetSelfNode方法获取自己节点
		return err // 如果获取自己节点出错，则返回错误
	}

	//3.如果不存在，创建新的节点

	tree.Leftval = parenttree.Rightval
	*tree.Rightval = *parenttree.Rightval + 1
	*tree.Level = *parenttree.Level + 1

	//5.将所有左值或右值大于等于新增节点父节点的右值的节点的左值或右值加2
	err = tx.Model(&CliTree{}).Where("leftval >= ?", *parenttree.Rightval).Update("leftval", gorm.Expr("leftval + ?", 2)).Error
	if err != nil {
		return err
	}
	err = tx.Model(&CliTree{}).Where("rightval >= ?", *parenttree.Rightval).Update("rightval", gorm.Expr("rightval + ?", 2)).Error
	if err != nil {
		return err
	}
	//4.插入节点
	err = tx.Create(&tree).Error
	if err != nil {
		return err
	}
	global.GVA_LOG.Info("创建节点成功")
	return nil
}
