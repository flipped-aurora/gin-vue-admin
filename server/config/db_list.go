package config

import (
	"gorm.io/gorm/logger"
	"strings"
)

type DsnProvider interface {
	Dsn() string
}

// Embeded 结构体可以压平到上一层，从而保持 config 文件的结构和原来一样
// 见 playground: https://go.dev/play/p/KIcuhqEoxmY

// GeneralDB 也被 Pgsql 和 Mysql 原样使用
type GeneralDB struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 数据库前缀
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库账号
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 数据库地址
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // 数据库引擎，默认InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true表示开启
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (c GeneralDB) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

type SpecializedDB struct {
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	Disable   bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}
