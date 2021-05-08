package service

import (
	"os/exec"

	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service/shared"

	"github.com/cornelk/hashmap"
	"github.com/gofrs/uuid"
	"github.com/hashicorp/go-plugin"
	"github.com/pkg/errors"
)

var (
	hm *hashmap.HashMap
)

func init() {
	hm = &hashmap.HashMap{}
}

func OpenConnection(connection request.CreateConnection) (err error, token model.ConnectionToken) {
	var (
		ok    bool
		t     string
		u     uuid.UUID
		list  interface{}
		game  *model.ProductPlugin
		games []model.ProductPlugin
		info  = request.ProductPluginSearch{
			PageInfo: request.PageInfo{Page: 1, PageSize: 1000},
		}
	)

	if err, list, _ = GetProductPluginInfoList(info); err != nil {
		return
	}

	if list == nil {
		err = errors.New("游戏插件列表为空")
		return
	}

	if games, ok = list.([]model.ProductPlugin); !ok {
		err = errors.New("游戏插件配置错误")
		return
	}

	for idx, _ := range games {
		if connection.ID == games[idx].ProductID {
			game = &games[idx]
			break
		}
	}

	if game == nil {
		err = errors.Errorf("游戏插件未配置: %s", connection.ID)
		return
	}

	if u, err = uuid.NewV4(); err != nil {
		err = errors.New("生成token失败")
		return
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("sh", "-c", game.PluginPath),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})

	rpcClient, err := client.Client()
	if err != nil {
		return
	}

	raw, err := rpcClient.Dispense("game_grpc")
	if err != nil {
		return
	}

	t = u.String()
	gameItf := raw.(shared.Game)
	saveClient(t, client, gameItf)
	token = model.ConnectionToken{Token: t}

	return
}

func CloseConnection(connection request.CloseConnection) (err error) {
	var (
		c      interface{}
		client *plugin.Client
		ok     bool
		token  = connection.Token
	)

	if c, ok = hm.Get(token); !ok {
		err = errors.Errorf("Get client failed with token: %s", token)
		return
	}

	if client, ok = c.(*plugin.Client); !ok {
		err = errors.Errorf("Client: %v can't covert", client)
		return
	}

	client.Kill()

	removeClient(token)
	return
}

func GameRequest(request request.GameRequest) (err error, data interface{}) {
	return
}

func saveClient(token string, client *plugin.Client, game shared.Game) {
	hm.Set(token+":client", client)
	hm.Set(token+":gameItf", game)
}

func removeClient(t string) {
	hm.Del(t + ":client")
	hm.Del(t + ":gameItf")
}
