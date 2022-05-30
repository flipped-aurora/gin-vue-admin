package config

import (
	"path"
	"strings"
)

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-api" json:"server-api" yaml:"server-api"`
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"server-request"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
}

// Model 返回model不带%s
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *Autocode) Model() string {
	s1 := strings.Replace(a.SModel, "%s", "", -1)
	if strings.Contains(s1, "/") {
		s1 = strings.Replace(s1, "/", "", 1)
	}
	return path.Join(s1)
}
