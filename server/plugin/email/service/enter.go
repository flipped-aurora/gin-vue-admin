package service

type ServiceGroup struct {
	EmailService
}

var ServiceGroupApp = new(ServiceGroup)
