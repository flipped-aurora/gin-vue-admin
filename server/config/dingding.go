package config

type DingDing struct {
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // 发件人  你自己要发邮件的邮箱
	Token  string `mapstructure:"token" json:"token" yaml:"token"`    // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
}
