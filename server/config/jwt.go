package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}
