package config

type Cache struct {
	Mode  string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Redis Redis
}
