package config

type Ldap struct {
	BindDN   string `mapstructure:"bind_dn" json:"bind_dn" yaml:"bind_dn"`
	BindPwd  string `mapstructure:"bind_pwd" json:"bind_pwd" yaml:"bind_pwd"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	TLS      bool   `mapstructure:"tls" json:"tls" yaml:"tls"` // 是否SSL   是否开启SSL
	BaseDN   string `mapstructure:"base_dn" json:"base_dn" yaml:"base_dn"`
	Filter   string `mapstructure:"filter" json:"filter" yaml:"filter"`
	FieldMap string `mapstructure:"field_map" json:"field_map" yaml:"field_map"`
}
