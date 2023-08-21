package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
}
