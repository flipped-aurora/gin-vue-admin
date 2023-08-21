package config

type Timer struct {
	Spec        string   `mapstructure:"spec" json:"spec" yaml:"spec"`
	Detail      []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
	Start       bool     `mapstructure:"start" json:"start" yaml:"start"`
	WithSeconds bool     `mapstructure:"with_seconds" json:"with_seconds" yaml:"with_seconds"`
}

type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}
