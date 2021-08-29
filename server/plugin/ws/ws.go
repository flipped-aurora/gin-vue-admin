package ws

import (
	"github.com/flipped-aurora/ws/core/biz"
	"github.com/flipped-aurora/ws/core/data"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type wsPlugin struct {
	logger               *zap.Logger                       // 日志输出对象
	manageBuf            int64                             // buffer
	registeredMsgHandler map[int32]func(biz.IMessage) bool // 消息处理
	checkMap             map[string]biz.CheckFunc          // 用户校验

	admin     biz.IManage
	adminCase *biz.AdminCase
}

func DefaultRegisteredMsgHandler(admin biz.IManage, logger *zap.Logger) map[int32]func(biz.IMessage) bool {
	return map[int32]func(msg biz.IMessage) bool{
		1: func(msg biz.IMessage) bool {
			// w.admin 里面找到注册客户端的方法
			client, ok := admin.FindClient(msg.GetTo())
			if !ok {
				logger.Info("没有找到该用户")
				return false
			}
			return client.SendMes(msg)
		},
	}
}

func DefaultCheckMap() map[string]biz.CheckFunc {
	return map[string]biz.CheckFunc{
		"gva_ws": func(c interface{}) (string, bool) {
			// 先断言是gin.content
			cc, ok := c.(*gin.Context)
			if !ok {
				return "", false
			}
			token := cc.Query("jwt")
			// 可以携带jwt
			if len(token) == 0 {
				return "", false
			}
			// 解析 jwt...

			return token, true
		},
	}
}

func (w *wsPlugin) Register(g *gin.RouterGroup) {
	// gva_ws 为身份校验函数
	g.GET("/ws", w.adminCase.HandlerWS("gva_ws", &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	}))
	g.POST("/sendMsg", w.adminCase.SendMsg("gva_ws"))

}

func (w *wsPlugin) RouterPath() string {
	return "gva_ws"
}

func GenerateWs(logger *zap.Logger, manageBuf int64, checkMap map[string]biz.CheckFunc) *wsPlugin {
	m := data.NewManage(manageBuf)
	t := data.NewTopic()
	h := data.NewHandle()
	admin := data.NewAdmin(m, t, h, logger)
	for s, checkFunc := range checkMap {
		admin.AddCheckFunc(s, checkFunc)
	}
	registeredMsgHandler := DefaultRegisteredMsgHandler(admin, logger)

	for key, handler := range registeredMsgHandler {
		admin.RegisteredMsgHandler(key, handler)
	}
	return &wsPlugin{logger: logger, manageBuf: manageBuf,
		registeredMsgHandler: registeredMsgHandler, checkMap: checkMap, admin: admin, adminCase: biz.NewAdmin(admin)}
}
