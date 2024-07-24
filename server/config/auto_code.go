package config

import (
	"path/filepath"
	"strings"
)

type Autocode struct {
	Web    string `mapstructure:"web" json:"web" yaml:"web"`
	Root   string `mapstructure:"root" json:"root" yaml:"root"`
	Server string `mapstructure:"server" json:"server" yaml:"server"`
	Module string `mapstructure:"module" json:"module" yaml:"module"`
	AiPath string `mapstructure:"ai-path" json:"ai-path" yaml:"ai-path"`
}

func (a *Autocode) WebRoot() string {
	webs := strings.Split(a.Web, "/")
	if len(webs) == 0 {
		webs = strings.Split(a.Web, "\\")
	}
	return filepath.Join(webs...)
}
