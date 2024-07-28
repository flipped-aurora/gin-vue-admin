package config

type CloudflareR2 struct {
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	BaseURL         string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	Path            string `mapstructure:"path" json:"path" yaml:"path"`
	AccountID       string `mapstructure:"account-id" json:"account-id" yaml:"account-id"`
	AccessKeyID     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	SecretAccessKey string `mapstructure:"secret-access-key" json:"secret-access-key" yaml:"secret-access-key"`
}
