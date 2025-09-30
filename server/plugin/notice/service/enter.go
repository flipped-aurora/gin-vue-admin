package service

import "github.com/doquangtan/socketio/v4"

type ServiceGroup struct {
	NotificationService
	OnlineUserService
	SocketService
}

var ServiceGroupApp = &ServiceGroup{
	SocketService: SocketService{
		connections: make(map[string]*socketio.Socket),
		userSockets: make(map[uint]string),
	},
}
