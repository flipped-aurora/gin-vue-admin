package config

type Timer struct {
	Start       bool     `mapstructure:"start" json:"start" yaml:"start"`                      // 是否启用
	Spec        string   `mapstructure:"spec" json:"spec" yaml:"spec"`                         // CRON表达式
	WithSeconds bool     `mapstructure:"with_seconds" json:"with_seconds" yaml:"with_seconds"` // 是否精确到秒
	Detail      []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
}

type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}
