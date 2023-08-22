// 自动生成模板OrderInfo
package orderinfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// OrderInfo 结构体
type OrderInfo struct {
      global.GVA_MODEL
      OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;size:32;"`
      TicketNumber  string `json:"ticketNumber" form:"ticketNumber" gorm:"column:ticket_number;comment:订单取票号;size:12;"`
      ContactPhone  string `json:"contactPhone" form:"contactPhone" gorm:"column:contact_phone;comment:订单联系电话;size:11;"`
      IsTransfer  *bool `json:"isTransfer" form:"isTransfer" gorm:"column:is_transfer;comment:是否换乘.0:false;1:true;"`
      IsOccupySeat  *int `json:"isOccupySeat" form:"isOccupySeat" gorm:"column:is_occupy_seat;comment:是否占座.0:false;1:true;"`
      CompleteStatus  *int `json:"completeStatus" form:"completeStatus" gorm:"column:complete_status;comment:完成状态.0:未完成;1:出票成功;2:出票失败;"`
      FailReason  string `json:"failReason" form:"failReason" gorm:"column:fail_reason;comment:出票失败原因;size:255;"`
      MachineName  string `json:"machineName" form:"machineName" gorm:"column:machine_name;comment:设备名称;size:128;"`
}


// TableName OrderInfo 表名
func (OrderInfo) TableName() string {
  return "order_info"
}

