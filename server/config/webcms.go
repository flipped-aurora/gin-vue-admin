package config

type Website struct {
	Siteid int    `mapstructure:"siteid" json:"siteid" yaml:"siteid"` // 站点id
	Weburl string `mapstructure:"weburl" json:"weburl" yaml:"weburl"` // 站点url
}
