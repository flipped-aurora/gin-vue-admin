package config

type Autocode struct {
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-api" json:"server-api" yaml:"server-api"`
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`
	SRequest        string `mapstructure:"server-request" json:"server-request"  yaml:"server-request"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
}
