package config

type MCP struct {
	Name        string `mapstructure:"name" json:"name" yaml:"name"`                         // MCP名称
	Version     string `mapstructure:"version" json:"version" yaml:"version"`                // MCP版本
	SSEPath     string `mapstructure:"sse_path" json:"sse_path" yaml:"sse_path"`             // SSE路径
	MessagePath string `mapstructure:"message_path" json:"message_path" yaml:"message_path"` // 消息路径
	UrlPrefix   string `mapstructure:"url_prefix" json:"url_prefix" yaml:"url_prefix"`       // URL前缀
}
