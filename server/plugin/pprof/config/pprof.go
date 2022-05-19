package config

type Pprof struct {
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	// Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
}
