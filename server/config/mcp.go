package config

type MCP struct {
	Name            string `mapstructure:"name" json:"name" yaml:"name"`
	Version         string `mapstructure:"version" json:"version" yaml:"version"`
	Path            string `mapstructure:"path" json:"path" yaml:"path"`
	Addr            int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	BaseURL         string `mapstructure:"base_url" json:"base_url" yaml:"base_url"`
	UpstreamBaseURL string `mapstructure:"upstream_base_url" json:"upstream_base_url" yaml:"upstream_base_url"`
	AuthHeader      string `mapstructure:"auth_header" json:"auth_header" yaml:"auth_header"`
	RequestTimeout  int    `mapstructure:"request_timeout" json:"request_timeout" yaml:"request_timeout"`
}
