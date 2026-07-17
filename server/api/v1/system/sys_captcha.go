package system

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/captcha"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// store 验证码存储：自动跟随 GVA_CACHE 后端——开启 Redis 时多实例共享，未开启时用进程内存。
// 无需再手动切换 Redis/内存(原两套逻辑与 GVA_CACHE 重合，已统一收敛到 captcha.CacheStore)。
var store base64Captcha.Store = captcha.NewCacheStore()

type BaseApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	cfg := securityConfigService.Current(c.Request.Context())
	// 判断验证码是否开启
	openCaptcha := cfg.CaptchaOpen               // 错误 N 次后出验证码
	openCaptchaTimeOut := cfg.CaptchaTimeout      // 计数缓存超时
	key := c.ClientIP()
	v, ok := global.GVA_CACHE.Get(key)
	if !ok {
		global.GVA_CACHE.Set(key, int64(1), time.Second*time.Duration(openCaptchaTimeOut)) // int64 以匹配 GVA_CACHE.Increment 计数类型(内存后端 IncrementInt64)
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 生成默认数字 driver
	driver := base64Captcha.NewDriverDigit(cfg.ImgHeight, cfg.ImgWidth, cfg.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("验证码获取失败!")
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: cfg.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	case int64:
		i = int(v)
	case string:
		// redis 后端 Get 返回字符串
		if n, err := strconv.Atoi(v); err == nil {
			i = n
		}
	default:
		i = 0
	}
	return
}
