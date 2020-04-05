package config

type Server struct {
	Mysql   Mysql   `json:"mysql"`
	Qiniu   Qiniu   `json:"qiniu"`
	Casbin  Casbin  `json:"casbin"`
	Redis   Redis   `json:"redis"`
	System  System  `json:"system"`
	JWT     JWT     `json:"jwt"`
	Captcha Captcha `json:"captcha"`
	Log     Log     `json:"log"`
}

type System struct { // 系统配置
	UseMultipoint bool   `json:"useMultipoint"`
	Env           string `json:"env"`
	Addr          int    `json:"addr"`
}

type JWT struct { // jwt签名
	SigningKey string `json:"signingKey"`
}

type Casbin struct { //casbin配置
	ModelPath string `json:"modelPath"` // casbin model地址配置
}

type Mysql struct { // mysql admin 数据库配置
	Username     string `json:"username"`
	Password     string `json:"password"`
	Path         string `json:"path"`
	Dbname       string `json:"dbname"`
	Config       string `json:"config"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
	LogMode      bool   `json:"maxOpenConns"`
}

type Redis struct { // Redis admin 数据库配置
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
type Qiniu struct { // 七牛 密钥配置
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type Captcha struct { // 验证码配置
	KeyLong   int `json:"keyLong"`
	ImgWidth  int `json:"imgWidth"`
	ImgHeight int `json:"imgHeight"`
}

type Log struct {
	// log 打印的前缀
	Prefix string `json:"prefix"`
	// 是否显示打印log的文件具体路径
	LogFile bool `json:"logFile"`
	// 在控制台打印log的级别， "" 默认不打印
	Stdout string `json:"stdout"`
	// 在文件中打印log的级别   "" 默认不打印
	File string `json:"file"`
}
