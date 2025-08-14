package config

import (
	"fmt"
	"net"
	"net/url"
)

type Oracle struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Oracle) Dsn() string {
	dsn := fmt.Sprintf("oracle://%s:%s@%s/%s?%s", url.PathEscape(m.Username), url.PathEscape(m.Password),
		net.JoinHostPort(m.Path, m.Port), url.PathEscape(m.Dbname), m.Config)
	return dsn

}
