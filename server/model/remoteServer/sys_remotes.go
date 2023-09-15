// 自动生成模板SysRemotes
package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 远程服务器配置表 结构体  SysRemotes
type SysRemotes struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）;size:191;"`  //机器简称 
      Code  string `json:"code" form:"code" gorm:"column:code;comment:字典名（英）;size:191;"`  //项目编号 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:状态;"`  //状态 
      IsRemove  *bool `json:"isRemoved" form:"isRemoved" gorm:"column:isRemoved;comment:描述;"`  //删除状态 
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:远程服务器ip;size:191;"`  //ip地址 
      Port  string `json:"port" form:"port" gorm:"column:port;comment:远程服务器端口号;size:191;"`  //端口号 
      User  string `json:"userName" form:"userName" gorm:"column:userName;comment:账号;size:191;"`  //账号 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;size:191;"`  //密码 
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`  //备注 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 远程服务器配置表 SysRemotes自定义表名 sys_remotes
func (SysRemotes) TableName() string {
  return "sys_remotes"
}

