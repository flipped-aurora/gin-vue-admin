package home

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var (
	messageService = service.ServiceGroupApp.WebcmsServiceGroup.MessageService
	store          = base64Captcha.DefaultMemStore
)

// 获取公共参数
func getwebcommon(siteurl string) gin.H {
	webconfig := getwebconfig(siteurl)
	res := gin.H{}
	for k, v := range webconfig {
		res[k] = v
	}
	return res
}

// 获取getwebconfig 参数
func getwebconfig(siteurl string) map[string]any {
	res, ok := global.BlackCache.Get(siteurl)
	if !ok {
		global.GVA_LOG.Error("获取参数失败!")
		return nil
	} else {
		return res.(map[string]any)
	}
}

// 提交留言
func (h *HomeApi) PostMessage(c *gin.Context) {
	var message webcms.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证验证码
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if !oc || store.Verify(message.Captchaid, message.Captcha, true) {
		if err := messageService.CreateMessage(message); err != nil {
			global.GVA_LOG.Error("提交失败!", zap.Error(err))
			response.FailWithMessage("提交失败", c)
		} else {
			response.OkWithMessage("提交成功", c)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}

}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (h *HomeApi) GetCaptcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
