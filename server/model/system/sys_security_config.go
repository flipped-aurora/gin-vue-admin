package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysSecurityConfig 安全配置 单行表 固定 id=1 启动加载入内存缓存 保存即热更新
type SysSecurityConfig struct {
	global.GVA_MODEL
	// 验证码
	CaptchaOpen    int `json:"captchaOpen" gorm:"default:0;comment:错误N次后出验证码 0=每次都要"`
	CaptchaTimeout int `json:"captchaTimeout" gorm:"default:3600;comment:防爆破计数缓存超时(秒)"`
	KeyLong        int `json:"keyLong" gorm:"default:6;comment:验证码长度"`
	ImgWidth       int `json:"imgWidth" gorm:"default:240;comment:验证码宽度"`
	ImgHeight      int `json:"imgHeight" gorm:"default:80;comment:验证码高度"`
	// 密码复杂度
	PwdMinLength      int  `json:"pwdMinLength" gorm:"default:8;comment:密码最小长度"`
	PwdRequireUpper   bool `json:"pwdRequireUpper" gorm:"default:false;comment:需大写字母"`
	PwdRequireLower   bool `json:"pwdRequireLower" gorm:"default:false;comment:需小写字母"`
	PwdRequireDigit   bool `json:"pwdRequireDigit" gorm:"default:false;comment:需数字"`
	PwdRequireSpecial bool `json:"pwdRequireSpecial" gorm:"default:false;comment:需特殊字符"`
	// 限流
	LimitEnable bool `json:"limitEnable" gorm:"default:false;comment:是否开启限流"`
	LimitWindow int  `json:"limitWindow" gorm:"default:60;comment:限流窗口(秒)"`
	LimitCount  int  `json:"limitCount" gorm:"default:30;comment:窗口内最大次数"`
	// 失败锁定
	LockEnable    bool `json:"lockEnable" gorm:"default:false;comment:是否开启失败锁定"`
	LockThreshold int  `json:"lockThreshold" gorm:"default:5;comment:失败次数阈值"`
	LockDuration  int  `json:"lockDuration" gorm:"default:30;comment:锁定时长(分钟)"`
	// 密码过期
	PwdExpireEnable bool `json:"pwdExpireEnable" gorm:"default:false;comment:是否开启密码过期"`
	PwdExpireDays   int  `json:"pwdExpireDays" gorm:"default:90;comment:密码有效天数"`
}

func (SysSecurityConfig) TableName() string {
	return "sys_security_config"
}

// CaptchaTimeoutDuration 防爆破计数缓存超时
func (c SysSecurityConfig) CaptchaTimeoutDuration() time.Duration {
	return time.Duration(c.CaptchaTimeout) * time.Second
}

// LockDurationTimeout 锁定时长
func (c SysSecurityConfig) LockDurationTimeout() time.Duration {
	return time.Duration(c.LockDuration) * time.Minute
}

// LimitWindowDuration 限流窗口
func (c SysSecurityConfig) LimitWindowDuration() time.Duration {
	return time.Duration(c.LimitWindow) * time.Second
}

// DefaultSecurityConfig 由 config.yaml 的 captcha 生成默认单行配置 调用方负责设 id=1
func DefaultSecurityConfig(captcha config.Captcha) SysSecurityConfig {
	return SysSecurityConfig{
		CaptchaOpen:       captcha.OpenCaptcha,
		CaptchaTimeout:    captcha.OpenCaptchaTimeOut,
		KeyLong:           captcha.KeyLong,
		ImgWidth:          captcha.ImgWidth,
		ImgHeight:         captcha.ImgHeight,
		PwdMinLength:      8,
		PwdRequireUpper:   false,
		PwdRequireLower:   false,
		PwdRequireDigit:   false,
		PwdRequireSpecial: false,
		LimitEnable:       false,
		LimitWindow:       60,
		LimitCount:        30,
		LockEnable:        false,
		LockThreshold:     5,
		LockDuration:      30,
		PwdExpireEnable:   false,
		PwdExpireDays:     90,
	}
}
