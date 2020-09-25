package config

// 文件上传服务器minio配置
type Minio struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `mapstructure:"access-key-id" json:"accessKeyID" yaml:"access-key-id"`
	SecretAccessKey string `mapstructure:"secret-access-key" json:"secretAccessKey" yaml:"secret-access-key"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	UseSSL          bool   `mapstructure:"use-ssl" json:"useSSL" yaml:"use-ssl"`
	Path            string `mapstructure:"path" json:"path" yaml:"path"`
}
