// 自动生成模板BizToolCmdSrvApi
package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 后端工具指令api 结构体  BizToolCmdSrvApi
type BizToolCmdSrvApi struct {
	global.GVA_MODEL
	AppName     string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称（中文）;" binding:"required"`     //应用名称（中文）
	AppCode     string `json:"appCode" form:"appCode" gorm:"column:app_code;comment:;" binding:"required"`             //应用名称（英文标识）
	ToolType    string `json:"tool_type" form:"tool_type" gorm:"column:tool_type;comment:工具类型;" binding:"required"`    //工具类型
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;" binding:"required"`                    //标题
	Desc        string `json:"desc" form:"desc" gorm:"column:desc;comment:;" binding:"required"`                       //应用介绍
	Classify    string `json:"classify" form:"classify" gorm:"column:classify;comment:;" binding:"required"`           //分类
	Version     string `json:"version" form:"version" gorm:"column:version;comment:;" binding:"required"`              //应用版本
	Mode        string `json:"mode" form:"mode" gorm:"column:mode;comment:;" binding:"required"`                       //收费模式
	DevelopMode string `json:"developMode" form:"developMode" gorm:"column:develop_mode;comment:;" binding:"required"` //后续迭代
	OssPath     string `json:"ossPath" form:"ossPath" gorm:"column:oss_path;comment:;"`                                //文件地址
	Cover       string `json:"cover" form:"cover" gorm:"column:cover;comment:;"`                                       //封面地址
	Tags        string `json:"tags" form:"tags" gorm:"column:tags;comment:;"`                                          //应用标签
	Video       string `json:"video" form:"video" gorm:"column:video;comment:;"`                                       //介绍视频
	TenantID    *int   `json:"tenant_id" form:"tenant_id" gorm:"column:tenant_id;comment:所属租户id;"`                     //所属租户id
	TenantUser  string `json:"tenant_user" form:"tenant_user" gorm:"column:tenant_user;comment:所属租户;"`                 //所属租户
	Zone        string `json:"zone" form:"zone" gorm:"column:zone;comment:应用存储的服务器地址;"`                                //应用存储的服务器地址
	Status      string `json:"status" form:"status" gorm:"column:status;comment:状态，1启用，2停用;"`                          //状态
	ApiPath     string `json:"api_path" form:"api_path" gorm:"column:api_path;comment:前端调用的接口地址;"`                     //接口地址
	IsPublic    string `json:"is_public" form:"is_public" gorm:"column:is_public;comment:是否公开项目;"`                     //是否公开
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
	OperateUser string `json:"operateUser" gorm:"column:operate_user;comment:更新者"` //介绍视频
	Remark      string `json:"remark" gorm:"-"`
	ApiHost     string `json:"api_host" form:"api_host" gorm:"-"` //接口地址前缀
}

// TableName 后端工具指令api BizToolCmdSrvApi自定义表名 biz_tool_cmd_srv_api
func (BizToolCmdSrvApi) TableName() string {
	return "biz_tool_cmd_srv_api"
}

// biz_apphub 结构体  BizAppHubRecord
type BizToolCmdSrvApiRecord struct {
	AppId  uint   `json:"appId" form:"appId" gorm:"column:app_id;comment:应用id;"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;type:text;comment:备注;"`
	//OperateUser string `json:"operateUser" gorm:"column:operate_user;comment:更新者"` //介绍视频
	BizToolCmdSrvApi
}

// TableName biz_apphub BizAppHub自定义表名 biz_apphub
func (BizToolCmdSrvApiRecord) TableName() string {
	return "biz_tool_cmd_srv_api_record"
}
