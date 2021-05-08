package shared

import (
	"net/rpc"
)

// RPCClient is an implementation of Game that talks over RPC.
type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) Open(host string, port uint32) error {
	// We don't expect a response, so we can just use interface{}
	var resp interface{}

	// The args are just going to be a map. A struct could be better.
	return m.client.Call("Plugin.Open", map[string]interface{}{
		"host": host,
		"port": port,
	}, &resp)
}

func (m *RPCClient) Close(key string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.Close", key, &resp)
	return resp, err
}

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Game
}

func (m *RPCServer) Open(args map[string]interface{}, token *interface{}) (err error) {
	*token, err = m.Impl.Open(args["host"].(string), args["port"].(uint32))
	return
}

func (m *RPCServer) Close(token string, code *int32, message *string) (err error) {
	*code, *message, err = m.Impl.Close(token)
	return err
}
