package utils

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type Option struct {
	BindDN  string `mapstructure:"bind_dn" json:"bind_dn" yaml:"bind_dn"`
	BindPwd string `mapstructure:"bind_pwd" json:"bind_pwd" yaml:"bind_pwd"`
	Host    string `mapstructure:"host" json:"host" yaml:"host"`
	TLS     bool   `mapstructure:"tls" json:"tls" yaml:"tls"` // 是否TLS
	BaseDN  string `mapstructure:"base_dn" json:"base_dn" yaml:"base_dn"`
	Filter  string `mapstructure:"filter" json:"filter" yaml:"filter"`
}

type Ldap struct {
	option *Option
	conn   *ldap.Conn
	Error  error
}

func NewLdap(option *Option) *Ldap {
	l := &Ldap{option: option}
	return l
}

func (l *Ldap) Connect() (err error) {
	var conn *ldap.Conn
	if l.option.TLS {
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		conn, err = ldap.DialTLS("tcp", l.option.Host, tlsConf)
	} else {
		conn, err = ldap.Dial("tcp", l.option.Host)
	}
	if err != nil {
		return err
	}
	conn.SetTimeout(5 * time.Second)
	l.conn = conn
	return
}

func (l *Ldap) Bind(username, password string) *Ldap {
	if l.Error != nil {
		return l
	}
	l.Error = l.conn.Bind(username, password)
	return l
}

func (l *Ldap) Search(req *ldap.SearchRequest, resp *ldap.SearchResult) *Ldap {
	if l.Error != nil {
		return l
	}
	l.Bind(l.option.BindDN, l.option.BindPwd)
	if l.Error != nil {
		return l
	}
	resp, l.Error = l.conn.Search(req)
	return l
}

func (l *Ldap) SearchFunc(req *ldap.SearchRequest, fn func(*ldap.SearchResult) error) *Ldap {
	if l.Error != nil {
		return l
	}
	l.Bind(l.option.BindDN, l.option.BindPwd)
	if l.Error != nil {
		return l
	}
	resp, err := l.conn.Search(req)
	if err == nil && fn != nil {
		l.Error = fn(resp)
	} else {
		l.Error = err
	}
	return l
}

func (l *Ldap) Close() {
	l.conn.Close()
}

func (l *Ldap) BuildSearchRequest(username string) *ldap.SearchRequest {
	filter := l.option.Filter
	if strings.Contains(filter, "(cn=%s)") {
		filter = fmt.Sprintf(filter, ldap.EscapeFilter(username))
	}
	return ldap.NewSearchRequest(
		l.option.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		filter,
		[]string{},
		nil)
}
