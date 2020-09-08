package config

type Server struct {
	Mysql       Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite      Sqlite      `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Qiniu       Qiniu       `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Casbin      Casbin      `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Redis       Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	System      System      `mapstructure:"system" json:"system" yaml:"system"`
	JWT         JWT         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha     Captcha     `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Zap         Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	LocalUpload LocalUpload `mapstructure:"localUpload" json:"localUpload" yaml:"localUpload"`
	Email       Email       `mapstructure:"email" json:"email" yaml:"email"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	NeedInitData  bool   `mapstructure:"need-init-data" json:"needInitData" yaml:"need-init-data"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type LocalUpload struct {
	Local      bool   `mapstructure:"local" json:"local" yaml:"local"`
	AvatarPath string `mapstructure:"avatar-path" json:"avatarPath" yaml:"avatar-path"`
	FilePath   string `mapstructure:"file-path" json:"filePath" yaml:"file-path"`
}

type Qiniu struct {
	AccessKey string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath   string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type Sqlite struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link_name" json:"linkName" yaml:"link_name"`
	ShowLine      bool   `mapstructure:"show_line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode_level" json:"encodeLevel" yaml:"encode_level"`
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktraceKey" yaml:"stacktrace_key"`
	LogInConsole  bool   `mapstructure:"log_in_console" json:"logInConsole" yaml:"log_in_console"`
}

type Email struct {
	EmailFrom         string `mapstructure:"email_from" json:"emailFrom" yaml:"email_from"`
	EmailNickName         string `mapstructure:"email_nick_name" json:"emailNickName" yaml:"email_nick_name"`
	EmailSecret         string `mapstructure:"email_secret" json:"emailSecret" yaml:"email_secret"`
	EmailTo         string `mapstructure:"email_to" json:"emailTo" yaml:"email_to"`
	EmailHost         string `mapstructure:"email_host" json:"emailHost" yaml:"email_host"`
	EmailPort         int `mapstructure:"email_port" json:"emailPort" yaml:"email_port"`
	EmailIsSSL         bool `mapstructure:"email_isSSL" json:"emailIsSSL" yaml:"email_isSSL"`
}
