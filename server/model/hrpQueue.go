// 自动生成模板HrpQueue
package model

import (
	//"gorm.io/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type HrpQueue struct {
      ID int32 `json:"id" form:"id" gorm:"column:id;comment:'ID';type:integer;"`
      CreateUid int32  `json:"createUid" form:"createUid" gorm:"column:create_uid;comment:'Created by';type:integer;"`
      CreateDate time.Time  `json:"createDate" form:"createDate" gorm:"column:create_date;comment:'Created on';type:timestamp without time zone;"`
      WriteUid  int32 `json:"writeUid" form:"writeUid" gorm:"column:write_uid;comment:'Last Updated by';type:integer;"`
      WriteDate time.Time  `json:"writeDate" form:"writeDate" gorm:"column:write_date;comment:'Last Updated on';type:timestamp without time zone;"`
      OutpatientNo  int64 `json:"outpatientNo" form:"outpatientNo" gorm:"column:outpatient_no;comment:'门诊号';type:integer;"`
      PartnerId  int32  `json:"partnerId" form:"partnerId" gorm:"column:partner_id;comment:'患者';type:integer;"`
      Spell  string `json:"spell" form:"spell" gorm:"column:spell;comment:'姓名拼音';type:character varying;"`
      EmployeeId  int32 `json:"employeeId" form:"employeeId" gorm:"column:employee_id;comment:'挂号医生';type:integer;"`
      RoomId  int32 `json:"roomId" form:"roomId" gorm:"column:room_id;comment:'诊室';type:integer;"`
      DepartmentId  int32 `json:"departmentId" form:"departmentId" gorm:"column:department_id;comment:'科室';type:integer;"`
      RegisterTypeId  int32 `json:"registerTypeId" form:"registerTypeId" gorm:"column:register_type_id;comment:'号类';type:character varying;"`
      ReconsultationTrue  bool `json:"reconsultationTrue" form:"reconsultationTrue" gorm:"column:reconsultation_true;comment:'复诊';type:boolean;"`
      EmergTrue bool  `json:"emergTrue" form:"emergTrue" gorm:"column:emerg_true;comment:'急诊';type:boolean;"`
      AppointmentTrue  bool `json:"appointmentTrue" form:"appointmentTrue" gorm:"column:appointment_true;comment:'是否预约';type:boolean;"`
      OrderNo  int32 `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:'';type:integer;"`
      OperatorCode string  `json:"operatorCode" form:"operatorCode" gorm:"column:operator_code;comment:'收费员';type:character varying;"`
      HappenDatetime int64   `json:"happenDatetime" form:"happenDatetime" gorm:"column:happen_datetime;comment:'发生时间';type:timestamp without time zone;"`
      RegisteDatetime  time.Time `json:"registeDatetime" form:"registeDatetime" gorm:"column:registe_datetime;comment:'登记时间';type:timestamp without time zone;"`
      TriagedTrue  bool  `json:"triagedTrue" form:"triagedTrue" gorm:"column:triaged_true;comment:'是否分诊';type:boolean;"`
      RegisterId int32  `json:"registerId" form:"registerId" gorm:"column:register_id;comment:'来源表';type:integer;"`
      TriagedQueueId   int32`json:"triagedQueueId" form:"triagedQueueId" gorm:"column:triaged_queue_id;comment:'';type:integer;"`
      QueueNameId int32  `json:"queueNameId" form:"queueNameId" gorm:"column:queue_name_id;comment:'队列名称ID';type:integer;"`
      Active  bool  `json:"active" form:"active" gorm:"column:active;comment:'时效';type:boolean;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:'';type:character varying;"`
      Old  string `json:"old" form:"old" gorm:"column:old;comment:'';type:character varying;"`
      Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:'';type:character varying;"`
      RegisterType  string `json:"registerType" form:"registerType" gorm:"column:register_type;comment:'';type:character varying;"`
      RecordState int  `json:"recordState" form:"recordState" gorm:"column:record_state;comment:'';type:integer;"`
      Origin  int `json:"origin" form:"origin" gorm:"column:origin;comment:'';type:integer;"`
      ShowOrderNo  string `json:"showOrderNo" form:"showOrderNo" gorm:"column:show_order_no;comment:'';type:character varying;"`
      BillNos string  `json:"billNos" form:"billNos" gorm:"column:bill_nos;comment:'';type:character varying;"`
      AppointmentKind  string `json:"appointmentKind" form:"appointmentKind" gorm:"column:appointment_kind;comment:'';type:character varying;"`
      //DisposeBillNos   `json:"disposeBillNos" form:"disposeBillNos" gorm:"column:dispose_bill_nos;comment:'';type:character varying[];"`
      ExecutionTrue  bool  `json:"executionTrue" form:"executionTrue" gorm:"column:execution_true;comment:'';type:boolean;"`
      ExecutionAutoTriageTrue  bool `json:"executionAutoTriageTrue" form:"executionAutoTriageTrue" gorm:"column:execution_auto_triage_true;comment:'';type:boolean;"`
      State int  `json:"state" form:"state" gorm:"column:state;comment:'';type:integer;"`
      QueuePriorityLevel  int `json:"queuePriorityLevel" form:"queuePriorityLevel" gorm:"column:queue_priority_level;comment:'';type:integer;"`
      QueueGroupId  int32 `json:"queueGroupId" form:"queueGroupId" gorm:"column:queue_group_id;comment:'';type:integer;"`
}


func (HrpQueue) TableName() string {
  return "hrp_queue"
}
