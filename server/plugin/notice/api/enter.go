package api

type ApiGroup struct {
	NotificationApi
	OnlineUserApi
}

var ApiGroupApp = new(ApiGroup)
