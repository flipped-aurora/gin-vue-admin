package biz_apphub

// biz_apphub 结构体  BizAppHubRecord
type BizAppHubRecord struct {
	AppId  uint   `json:"appId" form:"appId" gorm:"column:app_id;comment:应用id;"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;type:text;comment:备注;"`
	//OperateUser string `json:"operateUser" gorm:"column:operate_user;comment:更新者"` //介绍视频
	BizAppHub
}

// TableName biz_apphub BizAppHub自定义表名 biz_apphub
func (BizAppHubRecord) TableName() string {
	return "biz_apphub_record"
}
