package shared

import (
	"net/rpc"
)

// RPCClient is an implementation of Game that talks over RPC.
type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) Open(host string, port uint32) error {
	// We don't expect a response, so we can just use interface{}
	var rsp map[string]string

	// The args are just going to be a map. A struct could be better.
	err := m.client.Call("Plugin.Open", map[string]interface{}{
		"host": host,
		"port": port,
	}, &rsp)

	return err
}

func (m *RPCClient) Close() ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.Close", nil, &resp)
	return resp, err
}

func (m *RPCClient) Request(name string, data []byte) ([]byte, error) {
	var rsp map[string]interface{}
	err := m.client.Call("Plugin.Request", map[string]interface{}{
		"name": name,
		"data": data,
	}, &rsp)
	return rsp["data"].([]byte), err
}

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Game
}

func (m *RPCServer) Open(args map[string]interface{}, rsp *interface{}) (err error) {
	err = m.Impl.Open(args["host"].(string), args["port"].(uint32))
	return
}

func (m *RPCServer) Close(rsp map[string]interface{}) error {
	message, err := m.Impl.Close()
	rsp["message"] = message
	return err
}

func (m *RPCServer) Request(args map[string]interface{}, rsp map[string]interface{}) error {
	data, err := m.Impl.Request(args["name"].(string), args["data"].([]byte))
	rsp["data"] = data
	return err
}
