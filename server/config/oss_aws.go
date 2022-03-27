package config

type AwsS3 struct {
	Bucket           string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region           string `mapstructure:"region" json:"region" yaml:"region"`
	Endpoint         string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	S3ForcePathStyle bool   `mapstructure:"s3-force-path-style" json:"s3-force-path-style" yaml:"s3-force-path-style"`
	DisableSSL       bool   `mapstructure:"disable-ssl" json:"disable-ssl" yaml:"disable-ssl"`
	SecretID         string `mapstructure:"secret-id" json:"secret-id" yaml:"secret-id"`
	SecretKey        string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	BaseURL          string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	PathPrefix       string `mapstructure:"path-prefix" json:"path-prefix" yaml:"path-prefix"`
}
