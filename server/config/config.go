package config

type Server struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql"`
	Qiniu   Qiniu   `mapstructure:"qiniu" json:"qiniu"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin"`
	Redis   Redis   `mapstructure:"redis" json:"redis"`
	System  System  `mapstructure:"system" json:"system"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha"`
	Log     Log     `mapstructure:"log" json:"log"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint"`
	Env           string `mapstructure:"env" json:"env"`
	Addr          int    `mapstructure:"addr" json:"addr"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Path         string `mapstructure:"path" json:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}
type Qiniu struct {
	AccessKey string `mapstructure:"access-key" json:"accessKey"`
	SecretKey string `mapstructure:"secret-key" json:"secretKey"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile"`
	Stdout  string `mapstructure:"stdout" json:"stdout"`
	File    string `mapstructure:"file" json:"file"`
}
