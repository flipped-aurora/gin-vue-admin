package client

import (
	"context"
	"errors"

	mcpClient "github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func NewClient(baseURL, name, version, serverName string, headers ...map[string]string) (*mcpClient.Client, error) {
	options := make([]transport.StreamableHTTPCOption, 0, 1)
	if len(headers) > 0 && len(headers[0]) > 0 {
		options = append(options, transport.WithHTTPHeaders(headers[0]))
	}

	client, err := mcpClient.NewStreamableHttpClient(baseURL, options...)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	if err := client.Start(ctx); err != nil {
		return nil, err
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    name,
		Version: version,
	}

	result, err := client.Initialize(ctx, initRequest)
	if err != nil {
		return nil, err
	}
	if result.ServerInfo.Name != serverName {
		return nil, errors.New("server name mismatch")
	}

	return client, nil
}
