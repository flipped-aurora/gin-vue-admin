package mcp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/mcp/client"
	"github.com/flipped-aurora/gin-vue-admin/server/service/mcp/service"
)

type McpGroup struct {
	ServiceGroup service.McpServiceGroup
	ClientGroup  client.ClientStruct
}
