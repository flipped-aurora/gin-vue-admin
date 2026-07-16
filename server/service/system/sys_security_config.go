package system

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
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
func (s *SecurityConfigService) Get(ctx context.Context) (system.SysSecurityConfig, error) {
	var cfg system.SysSecurityConfig
	// 系统尚未初始化(未走 init 向导)或连库失败时 global.GVA_DB 为 nil
	// 此时返回 config.yaml 默认配置并带错误: 调用方 Current 据此不写缓存
	// 待数据库就绪后再惰性加载真实行 同时避免对 nil 的 *gorm.DB 解引用导致 panic
	if global.GVA_DB == nil {
		return system.DefaultSecurityConfig(global.GVA_CONFIG.Captcha), errors.New("数据库未初始化")
	}
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", 1).First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cfg = system.DefaultSecurityConfig(global.GVA_CONFIG.Captcha)
		cfg.ID = 1
		if err = global.GVA_DB.WithContext(ctx).Create(&cfg).Error; err != nil {
			return cfg, err
		}
		return cfg, nil
	}
	return cfg, err
}

// Set 持久化配置 刷新内存缓存 密码过期由关变开时回填存量 NULL 用户
func (s *SecurityConfigService) Set(ctx context.Context, cfg system.SysSecurityConfig) error {
	prev, err := s.Get(ctx)
	if err != nil {
		return err
	}
	cfg.GVA_MODEL = prev.GVA_MODEL
	if err = global.GVA_DB.WithContext(ctx).Save(&cfg).Error; err != nil {
		return err
	}
	setSecurityConfigCache(cfg)
	// 密码过期由关变开 回填存量 PasswordUpdatedAt 为 NULL 的用户
	if cfg.PwdExpireEnable && !prev.PwdExpireEnable {
		now := time.Now()
		if err = global.GVA_DB.WithContext(ctx).Model(&system.SysUser{}).
			Where("password_updated_at IS NULL").
			Update("password_updated_at", now).Error; err != nil {
			return err
		}
	}
	return nil
}

// Current 返回内存缓存当前配置 未加载则惰性 Get
func (s *SecurityConfigService) Current(ctx context.Context) system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	cfg, err := s.Get(ctx)
	if err == nil {
		setSecurityConfigCache(cfg)
	}
	return cfg
}

// LoadAll 启动时加载配置入内存缓存
func (s *SecurityConfigService) LoadAll(ctx context.Context) {
	cfg, err := s.Get(ctx)
	if err != nil {
		logger.WithCtx(ctx).Mod("biz").Error("加载安全配置失败!")
		return
	}
	setSecurityConfigCache(cfg)
}

// CurrentLimit 供中间件读取限流配置 返回 enable/window/count
func (s *SecurityConfigService) CurrentLimit(ctx context.Context) (enable bool, window int, count int) {
	cfg := s.Current(ctx)
	return cfg.LimitEnable, cfg.LimitWindow, cfg.LimitCount
}
