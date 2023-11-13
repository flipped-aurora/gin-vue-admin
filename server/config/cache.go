package config

type Cache struct {
	Model string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Redis Redis
}
