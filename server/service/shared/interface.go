// Package shared contains shared data between the host and plugins.
package shared

import (
	"context"
	"net/rpc"

	"gin-vue-admin/service/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "GAME_PLUGIN",
	MagicCookieValue: "eyotang",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"game_grpc": &GameGRPCPlugin{},
	"game":      &GamePlugin{},
}

// Game is the interface that we're exposing as a plugin.
type Game interface {
	Open(host string, port uint32) error
	Close() (string, error)
	Request(name string, data []byte) ([]byte, error)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
type GamePlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Game
}

func (p *GamePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*GamePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type GameGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Game
}

func (p *GameGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterGameServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GameGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewGameClient(c)}, nil
}
