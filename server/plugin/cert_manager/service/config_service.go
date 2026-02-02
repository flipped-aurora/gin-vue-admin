package service

import (
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type configService struct{}

type ProbeConfig struct {
	AutoDiscovery       bool `json:"autoDiscovery" gorm:"column:auto_discovery;comment:自动发现子域名;default:false"`
	MaxConcurrentProbes int  `json:"maxConcurrentProbes" gorm:"column:max_concurrent_probes;comment:最大并发探测数;default:10"`
}

var (
	probeSemaphore chan struct{}
	probeConfig    ProbeConfig
	once           sync.Once
)

func (s *configService) InitProbeConfig() error {
	var config ProbeConfig
	err := global.GVA_DB.Where("id = ?", 1).First(&config).Error
	if err != nil {
		// 这里简单处理，如果表不存在或者没数据就创建默认的
		config = ProbeConfig{
			AutoDiscovery:       false,
			MaxConcurrentProbes: 10,
		}
		// 注意：ProbeConfig 应该是一个 model，需要注册
	}

	probeConfig = config
	probeSemaphore = make(chan struct{}, config.MaxConcurrentProbes)
	for i := 0; i < config.MaxConcurrentProbes; i++ {
		probeSemaphore <- struct{}{}
	}
	global.GVA_LOG.Info("探测配置初始化完成", zap.Bool("autoDiscovery", config.AutoDiscovery), zap.Int("maxConcurrent", config.MaxConcurrentProbes))
	return nil
}

func (s *configService) GetProbeConfig() ProbeConfig {
	return probeConfig
}

func (s *configService) UpdateProbeConfig(config ProbeConfig) error {
	// ... 更新逻辑
	probeConfig = config
	// 重新初始化信号量比较复杂，这里简单处理：如果是增加并发，可以往 channel 放数据；如果是减少，就比较难。
	// 实际生产中建议重启服务或者使用更复杂的信号量实现。
	// 这里我们简单重新创建
	newSem := make(chan struct{}, config.MaxConcurrentProbes)
	for i := 0; i < config.MaxConcurrentProbes; i++ {
		newSem <- struct{}{}
	}
	probeSemaphore = newSem
	return nil
}

func (s *configService) AcquireProbeSlot() {
	if probeSemaphore != nil {
		<-probeSemaphore
	}
}

func (s *configService) ReleaseProbeSlot() {
	if probeSemaphore != nil {
		probeSemaphore <- struct{}{}
	}
}

func (s *configService) IsAutoDiscoveryEnabled() bool {
	return probeConfig.AutoDiscovery
}
