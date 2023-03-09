package config

type Ldap struct {
	Enable          bool           `mapstructure:"enable" json:"enable" yaml:"enable"` //是否启用ldap
	Host            string         `mapstructure:"host" json:"host" yaml:"host"`
	Port            int            `mapstructure:"port" json:"port" yaml:"port"`
	BaseDn          string         `mapstructure:"base-dn" json:"base-dn" yaml:"base-dn"`
	AuthFilter      string         `mapstructure:"auth-filter" json:"auth-filter" yaml:"auth-filter"` // openldap格式: (&(uid=%s))  AD格式: (&(sAMAccountName=%s)) 查询时%s使用Username
	BindUser        string         `mapstructure:"bind-user" json:"bind-user" yaml:"bind-user"`       // openldap格式: dc=example,dc=org AD格式: manange@example.org
	BindPass        string         `mapstructure:"bind-pass" json:"bind-pass" yaml:"bind-pass"`
	TLS             bool           `mapstructure:"tls" json:"tls" yaml:"tls"`
	StartTLS        bool           `mapstructure:"start-tls" json:"start-tls" yaml:"start-tls"`
	DefaultRoleIds  []uint         `mapstructure:"default-role-ids" json:"default-role-ids" yaml:"default-role-ids"` // 从ldap创建的用户初始role id列表, 关联Authority ID
	CoverAttributes bool           `mapstructure:"cover-attributes" json:"cover-attributes" yaml:"cover-attributes"` // 是否从ldap更新用户属性, 属性映射使用Attributes struct
	Attributes      LdapAttributes `mapstructure:"attributes" json:"attributes" yaml:"attributes"`
}

type LdapAttributes struct {
	NickName  string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`       // ldap attr中属性名映射到 System.User的NickName
	Phone     string `mapstructure:"phone" json:"phone" yaml:"phone"`                //ldap attr中字段名映射到 System.User的Phone
	Email     string `mapstructure:"email" json:"email" yaml:"email"`                // ldap attr中字段名映射到 System.User的Email
	HeaderImg string `mapstructure:"header-img" json:"header-img" yaml:"header-img"` // ldap attr中字段名映射到 System.User的HeaderImg
}
