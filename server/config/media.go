package config

type Media struct {
	ChunkDir    string `mapstructure:"chunk-dir" json:"chunk-dir" yaml:"chunk-dir"`             // 分片暂存目录
	MaxFileSize int64  `mapstructure:"max-file-size" json:"max-file-size" yaml:"max-file-size"` // 单文件最大字节,0=不限
	SessionTTL  int    `mapstructure:"session-ttl" json:"session-ttl" yaml:"session-ttl"`       // 会话过期小时数(清理用)
}
