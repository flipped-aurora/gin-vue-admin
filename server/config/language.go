package config

type Language struct {
	Language        string `mapstructure:"language" json:"language" yaml:"language"`
	DefaultLanguage string `mapstructure:"default-language" json:"default-language" yaml:"default-language"`
	Dir             string `mapstructure:"dir" json:"dir" yaml:"dir"`
}
