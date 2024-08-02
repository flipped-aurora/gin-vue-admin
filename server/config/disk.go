package config

type Disk struct {
	MountPoint string `mapstructure:"mount-point" json:"mount-point" yaml:"mount-point"`
}

type DiskList struct {
	Disk `yaml:",inline" mapstructure:",squash"`
}
