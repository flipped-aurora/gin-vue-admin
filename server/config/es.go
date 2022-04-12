package config

type Elasticsearch struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 哪个数据库
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 服务器地址:端口
	User string `mapstructure:"user" json:"user" yaml:"user"` // 用户名
	Pass string `mapstructure:"pass" json:"pass" yaml:"pass"` // 密码
}
