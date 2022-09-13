package service

type ServiceGroup struct {
	{{ .PlugName}}Service
}

var ServiceGroupApp = new(ServiceGroup)
