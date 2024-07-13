package config

import (
	"path/filepath"
	"strings"
)

type Autocode struct {
	Web    string `mapstructure:"web" json:"web" yaml:"web"`
	Root   string `mapstructure:"root" json:"root" yaml:"root"`
	Server string `mapstructure:"server" json:"server" yaml:"server"`
	AiPath string `mapstructure:"ai-path" json:"ai-path" yaml:"ai-path"`
}

func (a *Autocode) WebRoot(template string, packageName string, biz string) string {
	webs := strings.Split(a.Web, "/")
	if len(webs) == 0 {
		webs = strings.Split(a.Web, "\\")
	}
	if template == "package" {
		return filepath.Join(a.Root, filepath.Join(webs...), biz, packageName)
	}
	return filepath.Join(a.Root, filepath.Join(webs...), "plugin", packageName, biz)
}

// ServerRoot 生成服务端路径
// template: 模板 package/plugin
// biz: 业务模块(api/model/service/request/response)
// packageName: 模块名/插件名(system/example/email)
func (a *Autocode) ServerRoot(template, packageName, biz string) string {
	if template == "package" {
		if biz == "api" {
			return filepath.Join(a.Root, a.Server, biz, "v1", packageName)
		}
		if biz == "request" || biz == "response" {
			return filepath.Join(a.Root, a.Server, packageName, "model", biz)
		}
		if biz == "router" {
			return filepath.Join(a.Root, a.Server, biz, packageName)
		}
		return filepath.Join(a.Root, a.Server, biz, packageName)
	}
	return filepath.Join(a.Root, a.Server, "plugin", packageName, biz)
}

// ServerEnterRoot 生成服务端路径
// template: 模板 package/plugin
// biz: 业务模块(api/model/service/request/response)
// packageName: 模块名/插件名(system/example/email)
func (a *Autocode) ServerEnterRoot(template, packageName, biz string) string {
	if template == "package" {
		if biz == "api" {
			return filepath.Join(a.Server, biz, "v1", packageName)
		}
		if biz == "request" || biz == "response" {
			return filepath.Join(a.Server, packageName, "model", biz)
		}
		return filepath.Join(a.Server, packageName, biz)
	}
	return filepath.Join(a.Server, packageName, biz)
}
