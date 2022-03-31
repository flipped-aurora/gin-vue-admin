package config

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`    // 验证码宽度
	ImgHeight int `mapstructure:"img-height" json:"img-height" yaml:"img-height"` // 验证码高度
}
