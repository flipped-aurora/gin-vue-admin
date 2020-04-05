package config

type Server struct {
	Mysql   `mapstructure:"mysql"`
	Qiniu   `mapstructure:"qiniu"`
	Casbin  `mapstructure:"casbin"`
	Redis   `mapstructure:"redis"`
	System  `mapstructure:"system"`
	JWT     `mapstructure:"jwt"`
	Captcha `mapstructure:"captcha"`
	Log     `mapstructure:"log"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint"`
	Env           string `mapstructure:"env"`
	Addr          int    `mapstructure:"addr"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path"`
}

type Mysql struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Path         string `mapstructure:"path"`
	Dbname       string `mapstructure:"db-name"`
	Config       string `mapstructure:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
type Qiniu struct {
	AccessKey string `mapstructure:"access-key"`
	SecretKey string `mapstructure:"secret-key"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long"`
	ImgWidth  int `mapstructure:"img-width"`
	ImgHeight int `mapstructure:"img-height"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix"`
	LogFile bool   `mapstructure:"log-file"`
	Stdout  string `mapstructure:"stdout"`
	File    string `mapstructure:"file"`
}
