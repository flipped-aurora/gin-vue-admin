// 自动生成模板BizAppHub
package biz_apphub

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// biz_apphub 结构体  BizAppHub
type BizAppHub struct {
	global.GVA_MODEL
	AppName string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称（中文）;" binding:"required"` //应用名称（中文）
	AppCode string `json:"appCode" form:"appCode" gorm:"column:app_code;comment:;" binding:"required"`         //应用名称（英文标识）
	Title   string `json:"title" form:"title" gorm:"column:title;comment:;" binding:"required"`                //标题
	//User        string `json:"user" gorm:"column:user;comment:;"`                                                      //所属用户
	Desc        string `json:"desc" form:"desc" gorm:"column:desc;comment:;" binding:"required"`                       //应用介绍
	Classify    string `json:"classify" form:"classify" gorm:"column:classify;comment:;" binding:"required"`           //分类
	Version     string `json:"version" form:"version" gorm:"column:version;comment:;" binding:"required"`              //应用版本
	Mode        string `json:"mode" form:"mode" gorm:"column:mode;comment:;" binding:"required"`                       //收费模式
	DevelopMode string `json:"developMode" form:"developMode" gorm:"column:develop_mode;comment:;" binding:"required"` //后续迭代
	OssPath     string `json:"ossPath" form:"ossPath" gorm:"column:oss_path;comment:;"`                                //文件地址
	IndexHtml   string `json:"indexHtml" form:"indexHtml" gorm:"column:index_html;comment:网站首页;"`                      //网站首页
	Cover       string `json:"cover" form:"cover" gorm:"column:cover;comment:;"`                                       //封面地址
	Tags        string `json:"tags" form:"tags" gorm:"column:tags;comment:;"`                                          //应用标签
	Video       string `json:"video" form:"video" gorm:"column:video;comment:;"`
	TenantID    *int   `json:"tenant_id" form:"tenant_id" gorm:"column:tenant_id;comment:所属租户id;"` //所属租户id
	TenantUser  string `json:"tenant_user" form:"tenant_user" gorm:"column:tenant_user;comment:所属租户;"`
	OperateUser string `json:"operateUser" gorm:"column:operate_user;comment:更新者"` //介绍视频
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`

	Remark string `json:"remark" form:"remark" gorm:"-"`
	//OperateUser string `json:"-" gorm:"-"`
}

// TableName biz_apphub BizAppHub自定义表名 biz_apphub
func (BizAppHub) TableName() string {
	return "biz_apphub"
}

// GetUseOssPath 访问地址
func (b BizAppHub) GetUseOssPath() string {
	return fmt.Sprintf("soft/%s/%s/index.html", b.TenantUser, b.AppCode)
}

// GetUseOssPathV2 访问地址
func (b BizAppHub) GetUseOssPathV2() string {
	return fmt.Sprintf("soft/%s/%s/index", b.TenantUser, b.AppCode)
}

// GetOssStorePath 版本存储地址
func (b BizAppHub) GetOssStorePath() string {
	return fmt.Sprintf("soft/%s/%s/%s/dist/index.html", b.TenantUser, b.AppCode, b.Version)
}
