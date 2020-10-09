package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	From     string `mapstructure:"from" json:"from" yaml:"from"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`
}
