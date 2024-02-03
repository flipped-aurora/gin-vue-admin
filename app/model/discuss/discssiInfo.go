// 自动生成模板DiscussInfo
package discuss

import (
	"app/model/competition"
	"app/model/request"
	"app/model/user"
	"gorm.io/gorm"
	"time"

	"gorm.io/datatypes"
)

// 帖子信息 结构体  DiscussInfo
type DisInfo struct {
	ID        uint                `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time           // 创建时间
	UpdatedAt time.Time           // 更新时间
	DeletedAt gorm.DeletedAt      `gorm:"index" json:"-"`                                                      // 删除时间
	DisComId  uint                `json:"disComId" form:"disComId" gorm:"column:dis_com_id;comment:帖子关联比赛id;"` //帖子关联比赛id
	Com       competition.ComInfo `json:"comInfo" form:"comInfo" gorm:"foreignKey:DisComId"`
	DisUserId uint                `json:"disUserId" form:"disUserId" gorm:"column:dis_user_id;comment:发帖人id;"` //发帖人id
	User      user.UserInfo       `json:"userInfo" form:"userInfo" gorm:"foreignKey:DisUserId"`
	//以上不要动了
	DisTitle         string         `json:"disTitle" form:"disTitle" gorm:"column:dis_title;comment:帖子标题;size:300;not null" binding:"required"` //帖子标题
	DisContent       string         `json:"disContent" form:"disContent" gorm:"column:dis_content;comment:帖子内容;size:2000;type:text;"`           //帖子内容
	DisLoveNumber    int            `json:"disLoveNumber" form:"disLoveNumber" gorm:"column:dis_love_number;comment:帖子点赞人数;default:0"`          //帖子点赞人数
	DisCollectNumber int            `json:"disCollectNumber" form:"disCollectNumber" gorm:"column:dis_collect_number;comment:帖子收藏人数;default:0"` //帖子收藏人数
	DisModel         int            `json:"disModel" form:"disModel" gorm:"column:dis_model;comment:帖子类型;default:0"`                            //帖子类型
	DisPicture       datatypes.JSON `json:"disPicture" form:"disPicture" gorm:"column:dis_picture;comment:帖子图片;"`                               //帖子图片
	DisLookNumber    int            `json:"disLookNumber" form:"disLookNumber" gorm:"column:dis_look_number;comment:帖子阅读量;default:0"`           //帖子阅读量
	DisStatus        int            `json:"disStatus" form:"disStatus" gorm:"column:dis_status;comment:帖子状态;default:0"`                         //帖子状态
}

// TableName 帖子信息 DisInfo自定义表名 disInfo
func (DisInfo) TableName() string {
	return "disInfo"
}

type DisInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	DisTitle       string     `json:"disTitle" form:"disTitle" `
	DisComId       *uint      `json:"disComId" form:"disComId" `
	DisUserId      *uint      `json:"disUserId" form:"disUserId" `
	DisModel       *int       `json:"disModel" form:"disModel" `
	DisStatus      *int       `json:"disStatus" form:"disStatus" `
	request.PageInfo
}
