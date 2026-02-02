package config

type CertManager struct {
	AutoDiscovery bool `mapstructure:"auto-discovery" json:"autoDiscovery" yaml:"auto-discovery"`
	Concurrency   int  `mapstructure:"concurrency" json:"concurrency" yaml:"concurrency"`
}
