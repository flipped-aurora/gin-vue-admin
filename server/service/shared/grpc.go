package shared

import (
	"context"

	"github.com/eyotang/game-api-admin/server/service/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

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

// GRPCClient is an implementation of Game that talks over RPC.
type GRPCClient struct{ client proto.GameClient }

func (m *GRPCClient) Open(host string, port uint32) error {
	_, err := m.client.Open(context.Background(), &proto.OpenRequest{
		Host: host,
		Port: port,
	})
	return err
}

func (m *GRPCClient) Close() (string, error) {
	rsp, err := m.client.Close(context.Background(), &proto.EmptyRequest{})
	if err != nil {
		return rsp.Message, err
	}

	return rsp.Message, nil
}

func (m *GRPCClient) Request(name string, data []byte) ([]byte, error) {
	rsp, err := m.client.Request(context.Background(), &proto.GameRequest{
		Name: name,
		Data: data,
	})
	if err != nil {
		return rsp.Data, err
	}
	return rsp.Data, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Game
}

func (m *GRPCServer) Open(
	ctx context.Context,
	req *proto.OpenRequest) (*proto.EmptyResponse, error) {
	err := m.Impl.Open(req.Host, req.Port)
	return &proto.EmptyResponse{}, err
}

func (m *GRPCServer) Close(
	ctx context.Context,
	req *proto.EmptyRequest) (*proto.CloseResponse, error) {
	message, err := m.Impl.Close()
	return &proto.CloseResponse{Message: message}, err
}

func (m *GRPCServer) Request(
	ctx context.Context,
	req *proto.GameRequest) (*proto.GameResponse, error) {
	body, err := m.Impl.Request(req.Name, req.Data)
	return &proto.GameResponse{Data: body}, err
}
