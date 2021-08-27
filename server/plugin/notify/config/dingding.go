package config

type DingDing struct {
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // Url
	Token  string `mapstructure:"token" json:"token" yaml:"token"`    // access_token
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 密钥
}
