package service

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/eyotang/game-proxy/server/global"
	"github.com/eyotang/game-proxy/server/model"
	"github.com/eyotang/game-proxy/server/model/request"
	"github.com/eyotang/game-proxy/server/service/shared"

	"github.com/cornelk/hashmap"
	"github.com/gofrs/uuid"
	"github.com/hashicorp/go-plugin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	hm *hashmap.HashMap
)

func init() {
	hm = &hashmap.HashMap{}
}

func ServedPlugin(paramGame *request.ParamGame) (err error, game *model.ProductPlugin) {
	var (
		ok    bool
		list  interface{}
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
	return
}

func OwnedToken(paramGame *request.ParamGame, header *request.HeaderRequest) (err error) {
	var (
		id    string
		token = header.Token
	)

	if id, err = gameID(token); err != nil {
		return
	}

	if id != paramGame.ID {
		err = errors.Errorf("token: %s, 不属于该游戏: %s", token, paramGame.ID)
		return
	}
	return
}

func Destroy(paramGame *request.ParamGame) (err error, result model.DestroyResult) {
	var (
		client  *plugin.Client
		rpc     shared.Game
		message string
		err2    error
	)

	tks := tokens(paramGame.ID)
	global.GVA_LOG.Info("Destroy", zap.Any("tokens", tks))
	result.Total = len(tks)

	for _, token := range tks {
		if rpc, err2 = gameClient(token); err2 != nil {
			err = errors.Wrap(err, err2.Error())
			result.Failed++
			continue
		}

		if message, err2 = rpc.Close(); err2 != nil {
			err2 = errors.Wrapf(err2, "rpc close failed, message: %s", message)
			err = errors.Wrap(err, err2.Error())
			result.Failed++
			continue
		}

		if client, err2 = rpcClient(token); err2 != nil {
			err = errors.Wrap(err, err2.Error())
			result.Failed++
			continue
		}

		client.Kill()
		removeClient(token)
		result.Success++
	}

	return
}

func OpenConnection(paramGame *request.ParamGame, connection *request.CreateConnection) (err error, token model.ConnectionToken) {
	var (
		t    string
		u    uuid.UUID
		game *model.ProductPlugin
	)

	// 获取游戏插件
	if err, game = ServedPlugin(paramGame); err != nil {
		return
	}

	if u, err = uuid.NewV4(); err != nil {
		err = errors.New("生成token失败")
		return
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", game.PluginPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
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

func CloseConnection(header *request.HeaderRequest) (err error) {
	var (
		client *plugin.Client
		rpc    shared.Game
		token  = header.Token
	)

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

func GameRequest(header *request.HeaderRequest, param *request.ParamRequest, req *shared.GameRequest) (err error, rsp *shared.GameResponse) {
	var (
		rpc   shared.Game
		data  []byte
		body  []byte
		token = header.Token
	)

	if rpc, err = gameClient(token); err != nil {
		return
	}

	global.GVA_LOG.Info("请求消息", zap.Any("request", req))
	if body, err = json.Marshal(req); err != nil {
		return
	}

	if data, err = rpc.Request(param.Name, body); err != nil {
		err = errors.Wrapf(err, "rpc request failed, name: %s", param.Name)
		return err, nil
	}

	rsp = &shared.GameResponse{}
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err = d.Decode(rsp); err != nil {
		return
	}

	global.GVA_LOG.Info("响应消息", zap.Any("response", rsp))

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
	if c, ok = hm.Get(tokenGameName(token)); !ok {
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

func tokenGameName(token string) string {
	return token + ":gameName"
}

func tokensKey(name string) string {
	return "tokens:" + name
}

func tokens(name string) (tks []string) {
	key := tokensKey(name)

	var (
		ok bool
		v  interface{}
	)
	if v, ok = hm.Get(key); !ok {
		tks = []string{}
	} else {
		if tks, ok = v.([]string); !ok {
			tks = []string{}
		}
	}
	return
}

func saveClient(token string, client *plugin.Client, game shared.Game, name string) {
	hm.Set(tokenClient(token), client)
	hm.Set(tokenGame(token), game)
	hm.Set(tokenGameName(token), name)

	tks := tokens(name)
	global.GVA_LOG.Info("saveClient", zap.Any("tks 1", tks))
	tks = append(tks, token)
	global.GVA_LOG.Info("saveClient", zap.Any("tks 2", tks))
	hm.Set(tokensKey(name), tks)
}

func removeClient(token string) {
	name := tokenGameName(token)
	tks := tokens(name)

	global.GVA_LOG.Info("removeClient", zap.Any("tks 1", tks))
	tks = removeToken(tks, token)
	hm.Set(tokensKey(name), tks)
	global.GVA_LOG.Info("removeClient", zap.Any("tks 2", tks))

	hm.Del(tokenClient(token))
	hm.Del(tokenGame(token))
	hm.Del(tokenGameName(token))
}

func removeToken(tks []string, t string) []string {
	for idx, _ := range tks {
		if t == tks[idx] {
			tks[idx] = tks[len(tks)-1]
			return tks[:len(tks)-1]
		}
	}
	return tks
}
