package config

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath       string `mapstructure:"img-path" json:"img-path" yaml:"img-path"`
	AccessKey     string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"use-https" yaml:"use-https"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"use-cdn-domains" yaml:"use-cdn-domains"`
}
