package config

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"model-path" yaml:"model-path"` // 存放casbin模型的相对路径
}
