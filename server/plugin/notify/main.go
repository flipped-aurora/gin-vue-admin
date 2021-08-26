package notify

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/router"
	"github.com/gin-gonic/gin"
)

type ddPlugin struct {
	Secret string
	Token  string
	Url    string
}

func CreateDDPlug(Url string, Secret string, Token string) *ddPlugin {
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.Token = Token
	global.GlobalConfig.Url = Url
	return &ddPlugin{}
}

func (*ddPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRouter(group)
}

func (*ddPlugin) RouterPath() string {
	return "notify"
}
