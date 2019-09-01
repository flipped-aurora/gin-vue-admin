package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"uuid"`
	UserName   string    `json:"userName"`
	PassWord   string    `json:"passWord"`
	NickName   string    `json:"nickName" gorm:"default:'galeone'"`
	HeaderImg  string    `json:"headerImg" gorm:"default:'galeone'"`
	//Propertie                //	多余属性自行添加
	//PropertieId uint  // 自动关联 Propertie 的Id 附加属性过多 建议创建一对一关系
}

//type Propertie struct {
//	gorm.Model
//}
