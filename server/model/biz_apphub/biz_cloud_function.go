// 自动生成模板BizCloudFunction
package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 云函数 结构体  BizCloudFunction
type BizCloudFunction struct {
    global.GVA_MODEL
    CnName  string `json:"cn_name" form:"cn_name" gorm:"column:cn_name;comment:函数中文名;" binding:"required"`  //函数中文名 
    CodeName  string `json:"code_name" form:"code_name" gorm:"column:code_name;comment:函数英文标识;" binding:"required"`  //函数英文标识 
    Classify  string `json:"classify" form:"classify" gorm:"column:classify;comment:函数分类（集合，数学，字符串处理，化学......）;" binding:"required"`  //函数分类 
    ExecMode  string `json:"exec_mode" form:"exec_mode" gorm:"column:exec_mode;comment:云函数执行方式;" binding:"required"`  //云函数执行方式 
    Title  string `json:"title" form:"title" gorm:"column:title;comment:函数标题;" binding:"required"`  //函数标题 
    Content  string `json:"content" form:"content" gorm:"column:content;comment:;"`  //函数详细介绍 
    ContentType  string `json:"content_type" form:"content_type" gorm:"column:content_type;comment:内容类型;" binding:"required"`  //内容类型 
    Param  datatypes.JSON `json:"param" form:"param" gorm:"column:param;comment:[ { "code":"A", "desc":"参数描述", "mode":"in/out(in是输入参数，out是输出参数)", "mock_data":"mock的数据", "type":"类型，string/number" } ];type:text;" binding:"required"swaggertype:"object"`  //函数参数 
    IsPublic  string `json:"is_public" form:"is_public" gorm:"column:is_public;comment:是否公开;" binding:"required"`  //是否公开 
    ApiConfig  datatypes.JSON `json:"api_config" form:"api_config" gorm:"column:api_config;comment:接口配置{     "path": "/api/",     "method": "post" };type:text;"swaggertype:"object"`  //接口配置 
    ScriptCode  string `json:"script_code" form:"script_code" gorm:"column:script_code;comment:JS代码;"`  //JS代码 
    Tags  string `json:"tags" form:"tags" gorm:"column:tags;comment:tags;"`  //标签 
    Views  *int `json:"views" form:"views" gorm:"default:0;column:views;comment:浏览量;"`  //浏览量 
    ExecCount  *int `json:"exec_count" form:"exec_count" gorm:"default:0;column:exec_count;comment:执行次数;"`  //执行次数 
    Coll  *int `json:"coll" form:"coll" gorm:"default:0;column:coll;comment:收藏数量;"`  //收藏数量 
    Like  *int `json:"like" form:"like" gorm:"default:0;column:like;comment:点赞量;"`  //点赞量 
}


// TableName 云函数 BizCloudFunction自定义表名 biz_cloud_function
func (BizCloudFunction) TableName() string {
    return "biz_cloud_function"
}

