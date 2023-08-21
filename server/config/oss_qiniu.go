package config

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                  // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                            // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"img-path" yaml:"img-path"`                      // CDN加速域名
	AccessKey     string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`                // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`                // 秘钥SK
	UseHTTPS      bool   `mapstructure:"use-https" json:"use-https" yaml:"use-https"`                   // 是否使用https
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"use-cdn-domains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}
