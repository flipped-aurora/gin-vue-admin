package shared

import (
	"gin-vue-admin/service/proto"

	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Game that talks over RPC.
type GRPCClient struct{ client proto.GameClient }

func (m *GRPCClient) Open(host string, port uint32) (string, error) {
	rsp, err := m.client.Open(context.Background(), &proto.OpenRequest{
		Host: host,
		Port: port,
	})
	return rsp.Token, err
}

func (m *GRPCClient) Close(token string) (int32, string, error) {
	rsp, err := m.client.Close(context.Background(), &proto.CloseRequest{
		Token: token,
	})
	if err != nil {
		return rsp.Code, rsp.Message, err
	}

	return rsp.Code, rsp.Message, nil
}

func (m *GRPCClient) Request(name string, data []byte) (int32, []byte, error) {
	rsp, err := m.client.Request(context.Background(), &proto.GameRequest{
		Name: name,
		Data: data,
	})
	if err != nil {
		return rsp.Code, rsp.Data, err
	}
	return rsp.Code, rsp.Data, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Game
}

func (m *GRPCServer) Open(
	ctx context.Context,
	req *proto.OpenRequest) (*proto.OpenResponse, error) {
	token, err := m.Impl.Open(req.Host, req.Port)
	return &proto.OpenResponse{Token: token}, err
}

func (m *GRPCServer) Close(
	ctx context.Context,
	req *proto.CloseRequest) (*proto.CloseResponse, error) {
	code, message, err := m.Impl.Close(req.Token)
	return &proto.CloseResponse{Code: code, Message: message}, err
}

func (m *GRPCServer) Request(
	ctx context.Context,
	req *proto.GameRequest) (*proto.GameResponse, error) {
	code, body, err := m.Impl.Request(req.Name, req.Data)
	return &proto.GameResponse{Code: code, Data: body}, err
}
