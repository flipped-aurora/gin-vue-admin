package config

/*
* @Author <bypanghu> (bypanghu@163.com)
* @Date 2023/12/27 17:05
**/

type Minio struct {
	AccessKey string `json:"access_key" yaml:"access_key" mapstructure:"access_key" `
	Domain    string `json:"domain" yaml:"domain" mapstructure:"domain"`
	EndPoint  string `json:"end_point" yaml:"end_point" mapstructure:"end_point"`
	Path      string `json:"path" yaml:"path" mapstructure:"path"`
	Prefix    string `json:"prefix" yaml:"prefix" mapstructure:"prefix"`
	Token     string `json:"token" yaml:"token" mapstructure:"token"`
	UseSSl    bool   `json:"use_ssl" yaml:"use_ssl" mapstructure:"use_ssl"`
	SecretKey string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
}
