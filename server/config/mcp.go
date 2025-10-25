package config

type MCP struct {
	Name        string `mapstructure:"name" json:"name" yaml:"name"`                         // MCP名称
	Version     string `mapstructure:"version" json:"version" yaml:"version"`                // MCP版本
	SSEPath     string `mapstructure:"sse_path" json:"sse_path" yaml:"sse_path"`             // SSE路径
	MessagePath string `mapstructure:"message_path" json:"message_path" yaml:"message_path"` // 消息路径
	UrlPrefix   string `mapstructure:"url_prefix" json:"url_prefix" yaml:"url_prefix"`       // URL前缀
	Addr        int    `mapstructure:"addr" json:"addr" yaml:"addr"`                         // 独立MCP服务端口
	Separate    bool   `mapstructure:"separate" json:"separate" yaml:"separate"`             // 是否独立运行MCP服务
}
