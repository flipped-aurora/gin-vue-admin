package router

type RouterGroup struct {
	NotificationRouter
	OnlineUserRouter
	SocketRouter
}

var RouterGroupApp = new(RouterGroup)
