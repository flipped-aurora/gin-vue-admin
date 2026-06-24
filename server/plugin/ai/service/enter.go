package service

type ServiceGroup struct {
	CliService cliService
}

var ServiceGroupApp = new(ServiceGroup)
