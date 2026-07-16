package config

type App struct {
	Node  string `mapstructure:"node" json:"node" yaml:"node"`       // 节点标识
	AppID string `mapstructure:"app-id" json:"app-id" yaml:"app-id"` // 应用标识
	Env   string `mapstructure:"env" json:"env" yaml:"env"`          // 环境：dev/test/prod
}
