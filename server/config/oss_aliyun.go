package config

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
	BasePath        string `mapstructure:"base-path" json:"basePath" yaml:"base-path"`
}
