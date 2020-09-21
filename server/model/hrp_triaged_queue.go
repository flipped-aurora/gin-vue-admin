// 自动生成模板HrpTriagedQueue
package model

import (
	//"gorm.io/gorm"
      "time"
)

// 如果含有time.Time 请自行import time包
type HrpTriagedQueue struct {
      ID int32 `json:"id" form:"id" gorm:"column:id;comment:'ID';type:integer;"`
      CreateDate time.Time  `json:"createDate" form:"createDate" gorm:"column:create_date;comment:'Created on';type:timestamp without time zone;"`
      WriteUid int32  `json:"writeUid" form:"writeUid" gorm:"column:write_uid;comment:'Last Updated by';type:integer;"`
      CreateUid  int32 `json:"createUid" form:"createUid" gorm:"column:create_uid;comment:'Created by';type:integer;"`
      WriteDate  time.Time `json:"writeDate" form:"writeDate" gorm:"column:write_date;comment:'修改时间';type:timestamp without time zone;"`
      QueueNameId  int32 `json:"queueNameId" form:"queueNameId" gorm:"column:queue_name_id;comment:' 队列名称ID';type:integer;"`
      QueueId  int32 `json:"queueId" form:"queueId" gorm:"column:queue_id;comment:'';type:integer;"`
      EquipmentId int32  `json:"equipmentId" form:"equipmentId" gorm:"column:equipment_id;comment:'执行设备';type:integer;"`
      PartnerId   int32 `json:"partnerId" form:"partnerId" gorm:"column:partner_id;comment:'患者';type:integer;"`
      OutpatientNo int64  `json:"outpatientNo" form:"outpatientNo" gorm:"column:outpatient_no;comment:'门诊号';type:integer;"`
      Spell  string `json:"spell" form:"spell" gorm:"column:spell;comment:'姓名拼音';type:character varying;"`
      TriagedEmployeeIds  string  `json:"triagedEmployeeIds" form:"triagedEmployeeIds" gorm:"column:triaged_employee_ids;comment:'医生';type:character varying;"`
      TriagedRoomIds  string `json:"triagedRoomIds" form:"triagedRoomIds" gorm:"column:triaged_room_ids;comment:'诊室';type:character varying;"`
      OrderNo  int32 `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:'顺序号';type:integer;"`
      ShowOrderNo  string `json:"showOrderNo" form:"showOrderNo" gorm:"column:show_order_no;comment:'顺序号';type:character varying;"`
      AppointmentTrue  bool  `json:"appointmentTrue" form:"appointmentTrue" gorm:"column:appointment_true;comment:'';type:boolean;"`
      ReconsultationTrue  bool `json:"reconsultationTrue" form:"reconsultationTrue" gorm:"column:reconsultation_true;comment:'复诊';type:boolean;"`
      EmergTrue  bool `json:"emergTrue" form:"emergTrue" gorm:"column:emerg_true;comment:'加急';type:boolean;"`
      RegisterTypeId  int32 `json:"registerTypeId" form:"registerTypeId" gorm:"column:register_type_id;comment:'号类';type:integer;"`
      Active  bool `json:"active" form:"active" gorm:"column:active;comment:'时效';type:boolean;"`
      State  int `json:"state" form:"state" gorm:"column:state;comment:'就诊状态';type:integer;"`
      Stage int  `json:"stage" form:"stage" gorm:"column:stage;comment:'';type:integer;"`
      DepartmentId int32  `json:"departmentId" form:"departmentId" gorm:"column:department_id;comment:'';type:integer;"`
      Name string  `json:"name" form:"name" gorm:"column:name;comment:'';type:character varying;"`
      RegisterType  string `json:"registerType" form:"registerType" gorm:"column:register_type;comment:'';type:character varying;"`
      Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:'';type:character varying;"`
      Old  string `json:"old" form:"old" gorm:"column:old;comment:'';type:character varying;"`
      DiagnoseEmployeeId int32  `json:"diagnoseEmployeeId" form:"diagnoseEmployeeId" gorm:"column:diagnose_employee_id;comment:'';type:integer;"`
      DiagnoseRoomId  int32 `json:"diagnoseRoomId" form:"diagnoseRoomId" gorm:"column:diagnose_room_id;comment:'';type:integer;"`
      DiagnoseEquipmentId  int32 `json:"diagnoseEquipmentId" form:"diagnoseEquipmentId" gorm:"column:diagnose_equipment_id;comment:'';type:integer;"`
      DiagnoseTime time.Time  `json:"diagnoseTime" form:"diagnoseTime" gorm:"column:diagnose_time;comment:'';type:timestamp without time zone;"`
      HappenDatetime time.Time  `json:"happenDatetime" form:"happenDatetime" gorm:"column:happen_datetime;comment:'';type:timestamp without time zone;"`
      TriageDatetime  time.Time `json:"triageDatetime" form:"triageDatetime" gorm:"column:triage_datetime;comment:'';type:timestamp without time zone;"`
      InpatientNo int64  `json:"inpatientNo" form:"inpatientNo" gorm:"column:inpatient_no;comment:'';type:integer;"`
      QueueGroupId  int32 `json:"queueGroupId" form:"queueGroupId" gorm:"column:queue_group_id;comment:'';type:integer;"`
      QueuePriorityLevel int32  `json:"queuePriorityLevel" form:"queuePriorityLevel" gorm:"column:queue_priority_level;comment:'';type:integer;"`
}


func (HrpTriagedQueue) TableName() string {
  return "hrp_triaged_queue"
}
