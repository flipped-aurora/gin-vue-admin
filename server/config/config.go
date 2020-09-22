package config

type Server struct {
	Mysql       Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Postgresql  Postgresql  `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	Sqlite      Sqlite      `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Sqlserver   Sqlserver   `mapstructure:"sqlserver" json:"sqlserver" yaml:"sqlserver"`
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
	ErrorToEmail bool `mapstructure:"error-to-email" json:"errorToEmail" yaml:"error-to-email"`
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

type Postgresql struct {
	Host                 string `mapstructure:"host" json:"host" yaml:"host"`
	Username             string `mapstructure:"username" json:"username" yaml:"username"`
	Password             string `mapstructure:"password" json:"password" yaml:"password"`
	Dbname               string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Port                 string `mapstructure:"port" json:"port" yaml:"port"`
	Config               string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns         int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns         int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger               bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
	PreferSimpleProtocol bool   `mapstructure:"prefer-simple-protocol" json:"preferSimpleProtocol" yaml:"prefer-simple-protocol"`
}

type Sqlite struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger       bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type Sqlserver struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger       bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type LocalUpload struct {
	Local    bool   `mapstructure:"local" json:"local" yaml:"local"`
	FilePath string `mapstructure:"file-path" json:"filePath" yaml:"file-path"`
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

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}

type Email struct {
	EmailFrom     string `mapstructure:"email-from" json:"emailFrom" yaml:"email-from"`
	EmailNickname string `mapstructure:"email-nickname" json:"emailNickname" yaml:"email-nickname"`
	EmailSecret   string `mapstructure:"email-secret" json:"emailSecret" yaml:"email-secret"`
	EmailTo       string `mapstructure:"email-to" json:"emailTo" yaml:"email-to"`
	EmailHost     string `mapstructure:"email-host" json:"emailHost" yaml:"email-host"`
	EmailPort     int    `mapstructure:"email-port" json:"emailPort" yaml:"email-port"`
	EmailIsSSL    bool   `mapstructure:"email-is-ssl" json:"emailIsSSL" yaml:"email-is-ssl"`
}
