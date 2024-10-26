// 自动生成模板CliUser
package xiao

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// cliUser表 结构体  CliUser
type CliUser struct {
	global.GVA_MODEL
	Address   string           `json:"address" form:"address" gorm:"primarykey;column:address;comment:用户地址;size:100;"`              //用户地址
	Nickname  string           `json:"nickname" form:"nickname" gorm:"column:nickname;comment:用户昵称;size:10;"`                       //用户昵称
	Avatarurl string           `json:"avatarurl" form:"avatarurl" gorm:"column:avatarurl;comment:头像链接;size:191;"`                   //头像链接
	Parent    string           `json:"parent" form:"parent" gorm:"column:parent;comment:上级地址;size:100;"`                            //上级地址
	Pullnum   *int             `json:"pullnum" form:"pullnum" gorm:"default:0;column:pullnum;comment:直推个数;size:19;"`                //直推个数
	Teamnum   *int             `json:"teamnum" form:"teamnum" gorm:"default:0;column:teamnum;comment:团队总数;"`                        //团队总数
	Mypull    datatypes.JSON   `json:"mypull" form:"mypull" gorm:"column:mypull;comment:我的直推;type:text;"swaggertype:"array,object"` //我的直推
	MyUp      datatypes.JSON   `json:"myUp" form:"myUp" gorm:"column:my_up;comment:我的上级;type:text;"swaggertype:"array,object"`      //我的上级
	Vips      datatypes.JSON   `json:"vips" form:"vips" gorm:"column:vips;comment:团队VIP;type:text;"swaggertype:"array,object"`      //团队VIP
	Status    string           `json:"status" form:"status" gorm:"default:正常;column:status;comment:当前状态;size:10;"`                  //当前状态
	Desc      string           `json:"desc" form:"desc" gorm:"default:文本备注;column:desc;comment:文本备注;size:200;"`                     //文本备注
	Desnum    *decimal.Decimal `json:"desnum" form:"desnum" gorm:"type:decimal(14,4);default:0;column:desnum;comment:金额备注;"`        //金额备注
}

// TableName cliUser表 CliUser自定义表名 cli_user
func (CliUser) TableName() string {
	return "cli_user"
}

func NewUser(address string) *CliUser {
	return &CliUser{Address: address}
}

// 查询用户
func (cliUser *CliUser) GetCliUser(tx *gorm.DB) (*CliUser, error) {
	err := tx.Where("address = ?", cliUser.Address).First(&cliUser).Error
	return cliUser, err
}

// 查询用户是否存在
func (cliUser *CliUser) CheckAddress(tx *gorm.DB, address string) error {
	var count int64
	tx.Model(&CliUser{}).Where("address = ?", address).Count(&count)
	if count > 0 {
		return fmt.Errorf("user with address %s already exists", address)
	}
	return nil
}

// 默认注册
func RootUser() {
	//查询是否存在root用户
	tx := global.GVA_DB.Begin()
	exist, _ := NewUser("root").GetCliUser(tx)
	if exist.ID != 0 {
		//存在则回滚
		fmt.Println("root user already exists")
		tx.Rollback()
		return
	}
	var cliUser = NewUser("root")
	cliUser.Nickname = "root"
	cliUser.Avatarurl = "https://cdn.jsdelivr.net/gh/flipped-aurora/gin-vue-admin/server/resource/avatar/avatar.png"
	cliUser.Parent = "root"
	tx.Create(cliUser)
	var tree = NewCliTree("root", "root")
	tx.Create(tree)
	var cliMainorder = NewCliMainorder("root")
	tx.Create(cliMainorder)
	fmt.Println("create root user success")
	tx.Commit()
}
