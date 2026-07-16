package service

type ServiceGroup struct {
	CliService cliService
	McpService mcService
}

var ServiceGroupApp = new(ServiceGroup)
