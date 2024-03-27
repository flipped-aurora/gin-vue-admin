package webcms

// Class 结构体
type Webconfig struct {
	ID          uint64 `json:"id" gorm:"primarykey"` // 主键ID
	Logo        string `json:"logo" form:"logo" gorm:"column:logo;comment:网站logo;size:191;"`
	HotLine     string `json:"hotLine" form:"hotLine" gorm:"column:hot_line;comment:网站热线;size:191;"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号码;size:32;"`
	Wechat      string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:wechat;size:32;"`
	Email       string `json:"email" form:"email" gorm:"column:email;comment:email;size:32;"`
	Beian       string `json:"beian" form:"beian" gorm:"column:beian;comment:备案;size:191;"`
	Address     string `json:"address" form:"address" gorm:"column:address;comment:公司地址;size:191;"`
	Company     string `json:"company" form:"company" gorm:"column:company;comment:公司名称;size:191;"`
	Keywords    string `json:"keywords" form:"keywords" gorm:"column:keywords;comment:网站关键字;size:191;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:网站描述;size:191;"`
	SiteUrl     string `json:"siteUrl" form:"siteUrl" gorm:"column:site_url;comment:网站域名;size:191;"`
}

// TableName Class 表名
func (Webconfig) TableName() string {
	return "webconfig"
}
