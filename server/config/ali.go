package config

type Ali struct {
	AppId        string `mapstructure:"appId" json:"appId" yaml:"appId"`
	PrivateKey   string `mapstructure:"privateKey" json:"privateKey" yaml:"privateKey"`
	AliPublicKey string `mapstructure:"aliPublicKey" json:"aliPublicKey" yaml:"aliPublicKey"`
	PayNotifyURL string `mapstructure:"payNotifyURL" json:"payNotifyURL" yaml:"payNotifyURL"`
}
