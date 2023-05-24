package config

type Douyin struct {
	Key    string `mapstructure:"key" json:"key" yaml:"key"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}
