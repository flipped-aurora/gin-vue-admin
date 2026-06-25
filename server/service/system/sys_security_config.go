package system

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type SecurityConfigService struct{}

// securityConfigCache 进程内当前生效配置 热读
var securityConfigCache atomic.Value

func setSecurityConfigCache(cfg system.SysSecurityConfig) {
	securityConfigCache.Store(cfg)
}

func getSecurityConfigCache() system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	return system.SysSecurityConfig{}
}

// Get 读取单行配置 不存在则按 config.yaml 默认创建并返回
func (s *SecurityConfigService) Get() (system.SysSecurityConfig, error) {
	var cfg system.SysSecurityConfig
	err := global.GVA_DB.Where("id = ?", 1).First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cfg = system.DefaultSecurityConfig(global.GVA_CONFIG.Captcha)
		cfg.ID = 1
		if err = global.GVA_DB.Create(&cfg).Error; err != nil {
			return cfg, err
		}
		return cfg, nil
	}
	return cfg, err
}

// Set 持久化配置 刷新内存缓存 密码过期由关变开时回填存量 NULL 用户
func (s *SecurityConfigService) Set(cfg system.SysSecurityConfig) error {
	prev, err := s.Get()
	if err != nil {
		return err
	}
	cfg.ID = 1
	if err = global.GVA_DB.Save(&cfg).Error; err != nil {
		return err
	}
	setSecurityConfigCache(cfg)
	// 密码过期由关变开 回填存量 PasswordUpdatedAt 为 NULL 的用户
	if cfg.PwdExpireEnable && !prev.PwdExpireEnable {
		now := time.Now()
		if err = global.GVA_DB.Model(&system.SysUser{}).
			Where("password_updated_at IS NULL").
			Update("password_updated_at", now).Error; err != nil {
			return err
		}
	}
	return nil
}

// Current 返回内存缓存当前配置 未加载则惰性 Get
func (s *SecurityConfigService) Current() system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	cfg, err := s.Get()
	if err == nil {
		setSecurityConfigCache(cfg)
	}
	return cfg
}

// LoadAll 启动时加载配置入内存缓存
func (s *SecurityConfigService) LoadAll() {
	cfg, err := s.Get()
	if err != nil {
		global.GVA_LOG.Error("加载安全配置失败!")
		return
	}
	setSecurityConfigCache(cfg)
}

// CurrentLimit 供中间件读取限流配置 返回 enable/window/count
func (s *SecurityConfigService) CurrentLimit() (enable bool, window int, count int) {
	cfg := s.Current()
	return cfg.LimitEnable, cfg.LimitWindow, cfg.LimitCount
}
