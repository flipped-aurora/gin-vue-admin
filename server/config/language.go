package config

type Language struct {
	Language        string `mapstructure:"language" json:"language" yaml:"language"`
	DefaultLanguage string `mapstructure:"defualt-languge" json:"defualt-languge" yaml:"defualt-languge"`
	Dir             string `mapstructure:"dir" json:"dir" yaml:"dir"`
}
