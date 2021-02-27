package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path" `
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"`
}

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
}
type TencentCOS struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
}
