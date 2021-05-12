// Package shared contains shared data between the host and plugins.
package shared

import (
	"github.com/hashicorp/go-plugin"
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
}

// Game is the interface that we're exposing as a plugin.
type Game interface {
	Open(host string, port uint32) error
	Close() (string, error)
	Request(name string, data []byte) ([]byte, error)
}
