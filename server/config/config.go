package config

type Server struct {
	System     System          `mapstructure:"system" json:"system" yaml:"system"`
	TencentCOS TencentCOS      `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AliyunOSS  AliyunOSS       `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs       `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	JWT        JWT             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis      Redis           `mapstructure:"redis" json:"redis" yaml:"redis"`
	Local      Local           `mapstructure:"local" json:"local" yaml:"local"`
	Excel      Excel           `mapstructure:"excel" json:"excel" yaml:"excel"`
	AutoCode   Autocode        `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	AwsS3      AwsS3           `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	Qiniu      Qiniu           `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Timer      Timer           `mapstructure:"timer" json:"timer" yaml:"timer"`
	Cors       CORS            `mapstructure:"cors" json:"cors" yaml:"cors"`
	DBList     []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Zap        Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	Email      Email           `mapstructure:"email" json:"email" yaml:"email"`
	Sqlite     Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Mysql      Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql      Mssql           `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql      Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle     Oracle          `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Captcha    Captcha         `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
