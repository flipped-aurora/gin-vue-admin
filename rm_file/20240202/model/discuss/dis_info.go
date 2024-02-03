// 自动生成模板DisInfo
package discuss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	"gorm.io/datatypes"
)

// 帖子信息 结构体  DisInfo
type DisInfo struct {
      global.GVA_MODEL
      DisTitle  string `json:"disTitle" form:"disTitle" gorm:"column:dis_title;comment:帖子标题;size:300;"binding:"required"`  //帖子标题 
      DisContent  string `json:"disContent" form:"disContent" gorm:"column:dis_content;comment:帖子内容;size:2000;type:text;"`  //帖子内容 
      DisComId  *int `json:"disComId" form:"disComId" gorm:"column:dis_com_id;comment:帖子关联比赛id;"`  //帖子关联比赛id 
      DisUserId  *int `json:"disUserId" form:"disUserId" gorm:"column:dis_user_id;comment:发帖人id;"`  //发帖人id 
      DisLoveNumber  *int `json:"disLoveNumber" form:"disLoveNumber" gorm:"column:dis_love_number;comment:帖子点赞人数;"`  //帖子点赞人数 
      DisCollectNumber  *int `json:"disCollectNumber" form:"disCollectNumber" gorm:"column:dis_collect_number;comment:帖子收藏人数;"`  //帖子收藏人数 
      DisModel  *int `json:"disModel" form:"disModel" gorm:"column:dis_model;comment:帖子类型;"`  //帖子类型 
      DisPicture  datatypes.JSON `json:"disPicture" form:"disPicture" gorm:"column:dis_picture;comment:帖子图片;"`  //帖子图片 
      DisLookNumber  *int `json:"disLookNumber" form:"disLookNumber" gorm:"column:dis_look_number;comment:帖子阅读量;"`  //帖子阅读量 
      DisStatus  *int `json:"disStatus" form:"disStatus" gorm:"column:dis_status;comment:帖子状态;"`  //帖子状态 
}


// TableName 帖子信息 DisInfo自定义表名 disInfo
func (DisInfo) TableName() string {
  return "disInfo"
}

