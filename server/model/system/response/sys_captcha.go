package response

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId" example:"axxbbas"`               // 随机数id
	PicPath       string `json:"picPath" example:"data:img/png;base64,///"` // base64 图像字符串
	CaptchaLength int    `json:"captchaLength" example:"6"`                 // 验证码长度
}
