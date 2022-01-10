package response

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId" example:"axxbbas" comment:"随机数id"`                      // 随机数id
	PicPath       string `json:"picPath" example:"data:img/png;base64,///" comment:"base64 图像字符串"` // base64 图像字符串
	CaptchaLength int    `json:"captchaLength" example:"6"  comment:"验证码长度" `                      // 验证码长度
}
