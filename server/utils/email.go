package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"gin-vue-admin/global"

	"github.com/jordan-wright/email"
)

func Email(subject string, body string) error {
	to := strings.Split(global.GVA_CONFIG.Email.EmailTo, ",")
	return send(to, subject, body)
}

func EmailTest(subject string, body string) error {
	to := []string{global.GVA_CONFIG.Email.EmailFrom}
	return send(to, subject, body)
}

func send(to []string, subject string, body string) error {
	from := global.GVA_CONFIG.Email.EmailFrom
	nickName := global.GVA_CONFIG.Email.EmailNickName
	secret := global.GVA_CONFIG.Email.EmailSecret
	host := global.GVA_CONFIG.Email.EmailHost
	port := global.GVA_CONFIG.Email.EmailPort
	isSSL := global.GVA_CONFIG.Email.EmailIsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickName != "" {
		e.From = fmt.Sprintf("%s <%s>", nickName, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
