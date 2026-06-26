package captcha

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// defaultCaptchaExpiration 验证码答案默认有效期。
// 验证码校验通过即焚（Verify 时 clear=true 立即删除），此 TTL 仅为未使用验证码的兜底回收时间。
const defaultCaptchaExpiration = time.Second * 180

type CacheStore struct {
	Expiration time.Duration
	PreKey     string
}

// NewCacheStore 返回一个基于 GVA_CACHE 的验证码存储，前缀隔离避免与其它缓存键冲突。
func NewCacheStore() *CacheStore {
	return &CacheStore{
		Expiration: defaultCaptchaExpiration,
		PreKey:     "CAPTCHA_",
	}
}

// Set 写入验证码答案，按 Expiration 过期。GVA_CACHE 未就绪时返回错误，由上层(Generate)优雅处理。
func (cs *CacheStore) Set(id string, value string) error {
	if global.GVA_CACHE == nil {
		return errors.New("GVA_CACHE 未初始化")
	}
	global.GVA_CACHE.Set(cs.PreKey+id, value, cs.Expiration)
	return nil
}

// Get 读取验证码答案，clear=true 时读取后删除。未命中或类型不符返回空串。
func (cs *CacheStore) Get(id string, clear bool) string {
	if global.GVA_CACHE == nil {
		return ""
	}
	key := cs.PreKey + id
	v, ok := global.GVA_CACHE.Get(key)
	if !ok {
		return ""
	}
	if clear {
		global.GVA_CACHE.Delete(key)
	}
	// 内存后端原样返回存入的 string；Redis 后端 Get 也恒返回 string，两者一致。
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

// Verify 校验验证码答案，clear=true 时校验后删除。
func (cs *CacheStore) Verify(id, answer string, clear bool) bool {
	return cs.Get(id, clear) == answer
}
