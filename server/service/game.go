package service

import (
	"encoding/json"
	"os/exec"

	"github.com/eyotang/gin-vue-admin/server/model"
	"github.com/eyotang/gin-vue-admin/server/model/request"
	"github.com/eyotang/gin-vue-admin/server/service/shared"

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

func OpenConnection(paramGame *request.ParamGame, connection *request.CreateConnection) (err error, token model.ConnectionToken) {
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
		if paramGame.ID == games[idx].ProductID {
			game = &games[idx]
			break
		}
	}

	if game == nil {
		err = errors.Errorf("游戏插件未配置: %s", paramGame.ID)
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
		err = errors.Wrap(err, "New rpcClient failed!")
		return
	}

	raw, err := rpcClient.Dispense("game_grpc")
	if err != nil {
		err = errors.Wrapf(err, "rpcClient Dispense '%s' failed!", "game_grpc")
		return
	}

	t = u.String()
	gameItf := raw.(shared.Game)

	if err = gameItf.Open(connection.Host, connection.Port); err != nil {
		err = errors.Wrapf(err, "RPC call Open failed, host: %s, port: %d", connection.Host, connection.Port)
		return
	}

	saveClient(t, client, gameItf, paramGame.ID)
	token = model.ConnectionToken{Token: t}

	return
}

func CloseConnection(paramGame *request.ParamGame, header *request.HeaderRequest) (err error) {
	var (
		client *plugin.Client
		rpc    shared.Game
		id     string
		token  = header.Token
	)

	if id, err = gameID(token); err != nil {
		return
	}

	if id != paramGame.ID {
		err = errors.Errorf("token: %s, 不属于该游戏: %s", token, paramGame.ID)
		return
	}

	if rpc, err = gameClient(token); err != nil {
		return
	}

	if message, err := rpc.Close(); err != nil {
		err = errors.Wrapf(err, "rpc close failed, message: %s", message)
		return err
	}

	if client, err = rpcClient(token); err != nil {
		return
	}

	client.Kill()

	removeClient(token)
	return
}

func GameRequest(header *request.HeaderRequest, paramRequest *request.ParamRequest, request *request.GameRequest) (err error, data interface{}) {
	var (
		id    string
		rpc   shared.Game
		body  []byte
		token = header.Token
	)

	if id, err = gameID(token); err != nil {
		return
	}
	if id != paramRequest.ID {
		err = errors.Errorf("token: %s, 不属于该游戏: %s", token, paramRequest.ID)
		return
	}

	if rpc, err = gameClient(token); err != nil {
		return
	}

	if body, err = json.Marshal(request.Data); err != nil {
		err = errors.Wrapf(err, "Marshal json failed! data: %s", request.Data)
		return
	}

	if data, err = rpc.Request(paramRequest.Name, body); err != nil {
		err = errors.Wrapf(err, "rpc close failed, body: %s", data)
		return err, nil
	}

	return
}

func gameClient(token string) (rpc shared.Game, err error) {
	var (
		ok bool
		c  interface{}
	)
	if c, ok = hm.Get(tokenGame(token)); !ok {
		err = errors.Errorf("Get RPC client failed with token: %s", token)
		return
	}

	if rpc, ok = c.(shared.Game); !ok {
		err = errors.Errorf("RPC: %v can't covert", c)
		return
	}

	return
}

func rpcClient(token string) (client *plugin.Client, err error) {
	var (
		ok bool
		c  interface{}
	)
	if c, ok = hm.Get(tokenClient(token)); !ok {
		err = errors.Errorf("Get client failed with token: %s", token)
		return
	}

	if client, ok = c.(*plugin.Client); !ok {
		err = errors.Errorf("Client: %v can't covert", c)
		return
	}

	return
}

func gameID(token string) (id string, err error) {
	var (
		ok bool
		c  interface{}
	)
	if c, ok = hm.Get(tokenGameID(token)); !ok {
		err = errors.Errorf("Get client failed with token: %s", token)
		return
	}

	if id, ok = c.(string); !ok {
		err = errors.Errorf("Client: %v can't covert", c)
		return
	}

	return
}

func tokenClient(token string) string {
	return token + ":client"
}

func tokenGame(token string) string {
	return token + ":gameItf"
}

func tokenGameID(token string) string {
	return token + ":gameID"
}

func saveClient(token string, client *plugin.Client, game shared.Game, id string) {
	hm.Set(tokenClient(token), client)
	hm.Set(tokenGame(token), game)
	hm.Set(tokenGameID(token), id)
}

func removeClient(t string) {
	hm.Del(tokenClient(t))
	hm.Del(tokenGame(t))
	hm.Del(tokenGameID(t))
}
